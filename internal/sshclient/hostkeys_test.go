package sshclient

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	"testing"
)

func TestStaticHostKeyAlgorithms(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Desc      string
		PublicKey string
		Expect    []string
		ExpectErr bool
	}

	ed25519PublicKey, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	tcs := []testCase{
		{
			Desc:      "ed25519",
			PublicKey: testAuthorizedKey(t, ed25519PublicKey),
			Expect:    []string{ssh.KeyAlgoED25519},
		},
		{
			Desc:      "rsa",
			PublicKey: testAuthorizedKey(t, &rsaPrivateKey.PublicKey),
			Expect:    []string{ssh.KeyAlgoRSASHA512, ssh.KeyAlgoRSASHA256, ssh.KeyAlgoRSA},
		},
		{
			Desc:      "invalid public key",
			PublicKey: "invalid",
			ExpectErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Desc, func(t *testing.T) {
			args, err := processConnectOpts([]ConnectOption{StaticHostKeyAlgorithms(tc.PublicKey)})

			if tc.ExpectErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.Expect, args.clientConfig.HostKeyAlgorithms)
		})
	}
}

func testAuthorizedKey(t *testing.T, key interface{}) string {
	pk, err := ssh.NewPublicKey(key)
	require.NoError(t, err)

	return string(ssh.MarshalAuthorizedKey(pk))
}
