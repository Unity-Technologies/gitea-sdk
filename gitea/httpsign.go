// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-fed/httpsig"
	"golang.org/x/crypto/ssh"
)

type HTTPSign struct {
	ssh.Signer
}

// NewHTTPSigner returns a new HTTPSigner
// For now this only works with a ssh agent.
// It will try to find a valid certificate in the loaded keys and use it for signing.
func NewHTTPSign() *HTTPSign {
	agent, err := getAgent()
	if err != nil {
		return nil
	}

	signers, err := agent.Signers()
	if err != nil {
		return nil
	}

	if len(signers) == 0 {
		return nil
	}

	signer := findCertSigner(signers)
	if signer == nil {
		return nil
	}

	return &HTTPSign{
		Signer: signer,
	}
}

// SignRequest signs a HTTP request
func (c *Client) SignRequest(r *http.Request) error {
	var contents []byte = nil

	headersToSign := []string{httpsig.RequestTarget, "(created)", "(expires)"}

	// add our certificate to the headers to sign
	pubkey, _ := ssh.ParsePublicKey(c.httpsigner.Signer.PublicKey().Marshal())
	if cert, ok := pubkey.(*ssh.Certificate); ok {
		certString := base64.RawStdEncoding.EncodeToString(cert.Marshal())
		r.Header.Add("x-ssh-certificate", certString)

		headersToSign = append(headersToSign, "x-ssh-certificate")
	}

	// if we have a body, the Digest header will be added and we'll include this also in
	// our signature.
	if r.Body != nil {
		body, err := r.GetBody()
		if err != nil {
			return fmt.Errorf("getBody() failed: %s", err)
		}

		contents, err = ioutil.ReadAll(body)
		if err != nil {
			return fmt.Errorf("failed reading body: %s", err)
		}

		headersToSign = append(headersToSign, "Digest")
	}

	// create a signer for the request and headers, the signature will be valid for 10 seconds
	signer, _, err := httpsig.NewSSHSigner(c.httpsigner.Signer, httpsig.DigestSha512, headersToSign, httpsig.Signature, 10)
	if err != nil {
		return fmt.Errorf("httpsig.NewSSHSigner failed: %s", err)
	}

	// sign the request, the keyid is irrelevant, we don't use it
	err = signer.SignRequest("gitea", r, contents)
	if err != nil {
		return fmt.Errorf("httpsig.Signrequest failed: %s", err)
	}

	return nil
}

// findCertSigner returns the Signer containing a valid certificate
func findCertSigner(sshsigners []ssh.Signer) ssh.Signer {
	for _, s := range sshsigners {
		// Check if the key is a certificate
		if !strings.Contains(s.PublicKey().Type(), "cert-v01@openssh.com") {
			continue
		}

		// convert the ssh.Signer to a ssh.Certificate
		mpubkey, _ := ssh.ParsePublicKey(s.PublicKey().Marshal())
		cryptopub := mpubkey.(crypto.PublicKey)
		cert := cryptopub.(*ssh.Certificate)
		t := time.Unix(int64(cert.ValidBefore), 0)

		// make sure the certificate is at least 10 seconds valid
		if time.Until(t) <= time.Second*10 {
			continue
		}

		return s
	}

	return nil
}
