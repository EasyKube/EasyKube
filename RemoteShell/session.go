package RemoteShell

//远程会话配置
type SessionConfig struct {
	Host     string //主机名或ip
	Port     int    //主机侦听的端口 ssh默认为22 winrm默认为5985
	User     string //登录用户名
	Password string //登录密码
}

//远程会话接口
type Session interface {
	//打开会话
	Open(config *SessionConfig) error
	//关闭会话
	Close() error
	//上传文件
	UpFile(src string, dst string) error
	//下载文件
	DownFile(src string, dst string) error
	//执行远程命令
	Run(cmd string) error
	//运行远程命令文件
	//RunFile(scriptFile string) error
}

func NewSession() Session {
	return NewWinRmSession()
}

func NewSessionConfig() *SessionConfig {
	return &SessionConfig{}
}
