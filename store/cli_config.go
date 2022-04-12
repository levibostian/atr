package store

type cliConfig struct {
	Debug bool
}

var CliConfig = cliConfig{false}

func SetCliConfig(debug bool) {
	CliConfig = cliConfig{debug}
}
