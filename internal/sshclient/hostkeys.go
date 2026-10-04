package sshclient

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

type HostKeyVerifier func() (ssh.HostKeyCallback, error)

// StaticHostKey returns an ssh.HostKeyCallback which accepts only the provided public key.
// StaticHostKey expects a base64 encoded public key which is parsed using ssh.ParsePublicKey.
// StaticHostKey uses ssh.FixedHostKey with the decoded public key.
// Use ssh.FixedHostKey instead of StaticHostKey if an ssh.PublicKey is already available.
func StaticHostKey(publicKey string) HostKeyVerifier {
	return func() (ssh.HostKeyCallback, error) {
		pk, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey))
		if err != nil {
			return nil, err
		}

		return ssh.FixedHostKey(pk), nil
	}
}

// StaticHostKeyAlgorithms returns a ConnectOption which restricts the host key algorithms to the type of the provided public key.
// StaticHostKeyAlgorithms expects the same public key as StaticHostKey.
// Use StaticHostKeyAlgorithms together with StaticHostKey because a host with multiple host keys otherwise presents the key type preferred by the client.
func StaticHostKeyAlgorithms(publicKey string) ConnectOption {
	return func(c *connectArgs) error {
		pk, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey))
		if err != nil {
			return err
		}

		if pk.Type() == ssh.KeyAlgoRSA {
			// A rsa key is used with multiple signature algorithms
			c.clientConfig.HostKeyAlgorithms = []string{ssh.KeyAlgoRSASHA512, ssh.KeyAlgoRSASHA256, ssh.KeyAlgoRSA}
			return nil
		}

		c.clientConfig.HostKeyAlgorithms = []string{pk.Type()}
		return nil
	}
}

// unconfiguredHostKey returns an ssh.HostKeyCallback which always fails and reminds to configure host key validation.
func unconfiguredHostKey() ssh.HostKeyCallback {
	return func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return fmt.Errorf("sshclient: host key validation is not configured")
	}
}
