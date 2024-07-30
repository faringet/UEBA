package config

type Config struct {
	Logger       Logger       `mapstructure:"LOGGER"`
	ScanningOpts ScanningOpts `mapstructure:"SCANNING_OPTS"`
	LocalURL     string       `mapstructure:"THIS_APP_URL"`
}

type Logger struct {
	Production  string `mapstructure:"PRODUCTION"`
	Development string `mapstructure:"DEVELOPMENT"`
}

type ScanningOpts struct {
	Path   string `mapstructure:"PATH"`
	Format string `mapstructure:"FORMAT"`
}
