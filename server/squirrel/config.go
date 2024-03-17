package squirrel

type Config struct {
	server Server `ini:"server"`
}

type Server struct {
	Post int `ini:"post"`
}

func loadConfig() {
	return
}
