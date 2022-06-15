// Copyright 2022 The Gitea Authors. All rights reserved.
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

// HTTPSign contains the signer used for signing requests
type HTTPSign struct {
	ssh.Signer
	cert bool
}

// HTTPSignConfig contains the configuration for creating a HTTPSign
type HTTPSignConfig struct {
	fingerprint string
	principal   string
	pubkey      bool
	cert        bool
}

// NewHTTPSignWithPubkey can be used to create a HTTPSign with a public key
// if no fingerprint is specified it returns the first public key found
func NewHTTPSignWithPubkey(fingerprint string) *HTTPSign {
	return newHTTPSign(&HTTPSignConfig{
		fingerprint: fingerprint,
		pubkey:      true,
	})
}

// NewHTTPSignWithCert can be used to create a HTTPSign with a certificate
// if no principal is specified it returns the first certificate found
func NewHTTPSignWithCert(principal string) *HTTPSign {
	return newHTTPSign(&HTTPSignConfig{
		principal: principal,
		cert:      true,
	})
}

// NewHTTPSign returns a new HTTPSign
// For now this only works with a ssh agent.
// Depending on the configuration it will either use a certificate or a public key
func newHTTPSign(config *HTTPSignConfig) *HTTPSign {
	agent, err := GetAgent()
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

	var signer ssh.Signer

	if config.cert {
		signer = findCertSigner(signers, config.principal)
		if signer == nil {
			return nil
		}
	}

	if config.pubkey {
		signer = findPubkeySigner(signers, config.fingerprint)
		if signer == nil {
			return nil
		}
	}

	return &HTTPSign{
		Signer: signer,
		cert:   config.cert,
	}
}

// SignRequest signs a HTTP request
func (c *Client) SignRequest(r *http.Request) error {
	var contents []byte

	headersToSign := []string{httpsig.RequestTarget, "(created)", "(expires)"}

	if c.httpsigner.cert {
		// add our certificate to the headers to sign
		pubkey, _ := ssh.ParsePublicKey(c.httpsigner.Signer.PublicKey().Marshal())
		if cert, ok := pubkey.(*ssh.Certificate); ok {
			certString := base64.RawStdEncoding.EncodeToString(cert.Marshal())
			r.Header.Add("x-ssh-certificate", certString)

			headersToSign = append(headersToSign, "x-ssh-certificate")
		}
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

	// sign the request, use the fingerprint if we don't have a certificate
	keyID := "gitea"
	if !c.httpsigner.cert {
		keyID = ssh.FingerprintSHA256(c.httpsigner.Signer.PublicKey())
	}

	err = signer.SignRequest(keyID, r, contents)
	if err != nil {
		return fmt.Errorf("httpsig.Signrequest failed: %s", err)
	}

	return nil
}

// findCertSigner returns the Signer containing a valid certificate
// if no principal is specified it returns the first certificate found
func findCertSigner(sshsigners []ssh.Signer, principal string) ssh.Signer {
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

		if principal == "" {
			return s
		}

		for _, p := range cert.ValidPrincipals {
			if p == principal {
				return s
			}
		}
	}

	return nil
}

// findPubkeySigner returns the Signer containing a valid public key
// if no fingerprint is specified it returns the first public key found
func findPubkeySigner(sshsigners []ssh.Signer, fingerprint string) ssh.Signer {
	for _, s := range sshsigners {
		// Check if the key is a certificate
		if strings.Contains(s.PublicKey().Type(), "cert-v01@openssh.com") {
			continue
		}

		if fingerprint == "" {
			return s
		}

		if strings.TrimSpace(string(ssh.MarshalAuthorizedKey(s.PublicKey()))) == fingerprint {
			return s
		}

		if ssh.FingerprintSHA256(s.PublicKey()) == fingerprint {
			return s
		}
	}

	return nil
}
