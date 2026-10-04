package sshclient

import "golang.org/x/crypto/ssh"

func Password(password string) AuthMethod {
	return func() ([]ssh.AuthMethod, error) {
		// The error of the last attempted method is returned when the authentication fails.
		// Keyboard interactive is attempted first because a server without a keyboard interactive device rejects
		// the method with an error which hides the authentication failure.
		return []ssh.AuthMethod{
			ssh.KeyboardInteractive(sshPasswordKeyboardInteractive(password)),
			ssh.Password(password),
		}, nil
	}
}

// sshPasswordKeyboardInteractive is an implementation of ssh.KeyboardInteractiveChallenge that simply sends
// back the password for all questions.
// go://github.com/hashicorp/terraform@v1.0.4/internal/communicator/ssh.PasswordKeyboardInteractive
func sshPasswordKeyboardInteractive(password string) ssh.KeyboardInteractiveChallenge {
	return func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		// Just send the password back for all questions
		answers := make([]string, len(questions))
		for i := range answers {
			answers[i] = password
		}

		return answers, nil
	}
}
