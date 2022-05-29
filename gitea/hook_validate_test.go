// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hashers are based on https://github.com/go-gitea/gitea/blob/0dfc2e55ea258d2b1a3cd86e2b6f27a481e495ff/services/webhook/deliver.go#L105-L116

func TestVerifyWebhookSignature(t *testing.T) {
	secret := "s3cr3t"
	payload := []byte(`{"foo": "bar", "baz": true}`)

	hasher := hmac.New(sha256.New, []byte(secret))
	hasher.Write(payload)
	sig := hex.EncodeToString(hasher.Sum(nil))

	tt := []struct {
		Name    string
		Secret  string
		Payload string
		Succeed bool
	}{
		{
			Name:    "Correct secret and payload",
			Secret:  "s3cr3t",
			Payload: `{"foo": "bar", "baz": true}`,
			Succeed: true,
		},
		{
			Name:    "Correct secret bad payload",
			Secret:  "s3cr3t",
			Payload: "{}",
			Succeed: false,
		},
		{
			Name:    "Incorrect secret good payload",
			Secret:  "secret",
			Payload: `{"foo": "bar", "baz": true}`,
			Succeed: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			ok, err := VerifyWebhookSignature(tc.Secret, sig, []byte(tc.Payload))
			assert.NoError(t, err, "verification should not error")
			assert.True(t, ok == tc.Succeed, "verification should be %t", tc.Succeed)
		})
	}
}

func TestVerifyWebhookSignatureHandler(t *testing.T) {
	secret := "s3cr3t"
	payload := []byte(`{"foo": "bar", "baz": true}`)

	hasher := hmac.New(sha256.New, []byte(secret))
	hasher.Write(payload)
	sig := hex.EncodeToString(hasher.Sum(nil))

	tt := []struct {
		Name      string
		Secret    string
		Payload   string
		Signature string
		Status    int
	}{
		{
			Name:      "Correct secret and payload",
			Secret:    "s3cr3t",
			Payload:   `{"foo": "bar", "baz": true}`,
			Signature: sig,
			Status:    http.StatusOK,
		},
		{
			Name:      "Correct secret bad payload",
			Secret:    "s3cr3t",
			Payload:   "{}",
			Signature: sig,
			Status:    http.StatusUnauthorized,
		},
		{
			Name:      "Incorrect secret good payload",
			Secret:    "secret",
			Payload:   `{"foo": "bar", "baz": true}`,
			Signature: sig,
			Status:    http.StatusUnauthorized,
		},
		{
			Name:   "No signature",
			Status: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			server := httptest.NewServer(VerifyWebhookSignatureMiddleware(tc.Secret)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				_, _ = w.Write(nil)
			})))
			defer server.Close()

			req, err := http.NewRequest(http.MethodPost, server.URL, strings.NewReader(tc.Payload))
			assert.NoError(t, err, "should create request")

			if tc.Signature != "" {
				req.Header.Set("X-Gitea-Signature", tc.Signature)
			}

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err, "request should be delivered")
			assert.True(t, resp.StatusCode == tc.Status, "status should be %d, but got %d", tc.Status, resp.StatusCode)
		})
	}
}
