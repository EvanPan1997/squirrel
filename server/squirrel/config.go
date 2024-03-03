package squirrel

type Config struct {
	Server `ini:"server"`
}

type Server struct {
	Host           string `ini:"host" json:"host"`
	Port           int    `ini:"port" json:"port"`
	MaxConnections int    `ini:"max_connections" json:"max_connections"`
}

func loadConfig() {
	return
}
