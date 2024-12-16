package service

type Leaderboard struct {
	Listen       string `yaml:"listen" mapstructure:"listen"`
	Database     string `yaml:"database" mapstructure:"database"`
	ReadTimeout  int64  `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout int64  `yaml:"write_timeout" mapstructure:"write_timeout"`
}

type Module struct {
	Leaderboard Leaderboard `yaml:"leaderboard" mapstructure:"leaderboard"`
}

type ModuleConfig struct {
	Name   string `yaml:"name" mapstructure:"name"`
	Owner  string `yaml:"owner" mapstructure:"owner"`
	Module Module `yaml:"module" mapstructure:"module"`
}
