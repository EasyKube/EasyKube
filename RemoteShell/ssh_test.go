package RemoteShell

import "testing"

func Test_SSH(t *testing.T) {
	var s = NewSSHSession()
	config := NewSessionConfig()
	config.Host = "192.168.0.236"
	config.User = "user"
	config.Password = "ljlkkk"
	err := s.Open(config)
	if err != nil {
		t.Error(err)
	}
	s.Run("ls -l")
	if err != nil {
		t.Error(err)
	}

	s.DownFile("/home/user/main.go", "i:/main.txt")
	s.UpFile("i:/test.exe", "/home/user/test.exe")
	s.Close()

}
