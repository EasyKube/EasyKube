package RemoteShell

// We use a forked version of net/http and crypto/tls
// because the standard libs do now support renegotiation

import "testing"

func Test_WinRM(t *testing.T) {
	s := NewWinRmSession()
	config := NewSessionConfig()
	config.Host = "192.168.0.239"
	config.User = "administrator"
	config.Password = "ljlkkk"

	err := s.Open(config)
	if err != nil {
		panic(err)
	}
	err = s.Run("ipconfig /all")
	println(err)
}
