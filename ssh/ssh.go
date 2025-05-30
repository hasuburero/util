package ssh

import (
	"errors"
	"io"
	"os"
)

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func ()

func GetSSHConfig(user, passkey, keypath string) (*ssh.ClientConfig, error) {
	var auth []ssh.AuthMethod = []ssh.AuthMethod{}
	if keypath != "" {
		fd, err := os.Open(keypath)
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		key, err := io.ReadAll(fd)
		if err != nil {
			return nil, err
		}

		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}

		auth = append(auth, ssh.PublicKeys(signer))
	}

	if passkey != "" {
		auth = append(auth, ssh.Password(passkey))
	}

	if len(auth) == 0 {
		return nil, errors.New("No valid auth method\n")
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return config, nil
}
