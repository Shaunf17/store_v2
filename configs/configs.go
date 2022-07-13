package configs

type ServerConfigs struct {
	Host string
	Port string
	Root string
}

type LoggerConfigs struct {
	On             bool
	WriteToFile    bool
	WriteToConsole bool
	Filename       string
}
