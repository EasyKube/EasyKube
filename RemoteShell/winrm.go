package RemoteShell

/**
服务端配置，设置为认证方式为Basic,不用加密传输，这时可以直接使用用户名和密码登录，传输为明文

    winrm quickconfig
    y
    winrm set winrm/config/service/Auth '@{Basic="true"}'
    winrm set winrm/config/service '@{AllowUnencrypted="true"}'
    winrm set winrm/config/winrs '@{MaxMemoryPerShellMB="1024"}'

**/

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/masterzen/winrm"
)

type WimRmSession struct {
	host     string
	user     string
	password string
	port     int
	client   *winrm.Client
	config   *Config
}

func NewWinRmSession() *WimRmSession {
	return &WimRmSession{}
}

func (this *WimRmSession) Open(config *SessionConfig) error {
	this.host = config.Host
	this.user = config.User
	this.password = config.Password
	this.port = 5985
	if config.Port > 0 {
		this.port = config.Port
	}

	this.config = &Config{}
	this.config.Auth.Password = this.password
	this.config.Auth.User = this.user
	this.config.Https = false
	this.config.Insecure = true
	this.config.MaxOperationsPerShell = 15
	this.config.OperationTimeout = 60

	endpoint := winrm.NewEndpoint(this.host, this.port, false, true, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, this.user, this.password)
	if err != nil {
		return err
	}
	this.client = client
	return nil
}

func (this *WimRmSession) Close() error {
	if this.client != nil {
		this.client = nil
	}
	this.client = nil
	return nil
}

func (this *WimRmSession) Run(cmd string) error {
	out1, out2, out3, err := this.client.RunWithString(cmd, "")
	fmt.Printf("RESULT:  %s:\n%s %s,%d", this.host, out1, out2, out3)
	return err
}

func (this *WimRmSession) List(remotePath string) ([]FileItem, error) {
	return fetchList(this.client, winPath(remotePath))
}

func (this *WimRmSession) Write(toPath string, src io.Reader) error {
	return doCopy(this.client, this.config, src, winPath(toPath))
}

func (this *WimRmSession) UpFile(src string, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
		//return fmt.Errorf("Couldn't read file %s: %v", fromPath, err)
	}

	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
		//return fmt.Errorf("Couldn't stat file %s: %v", fromPath, err)
	}

	if !fi.IsDir() {
		return this.Write(dst, f)
	} else {
		fw := fileWalker{
			client:  this.client,
			config:  this.config,
			toDir:   src,
			fromDir: dst,
		}
		return filepath.Walk(src, fw.copyFile)
	}
}

func (this *WimRmSession) DownFile(src string, dst string) error {
	return nil
}
