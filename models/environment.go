package models

type storageEnv struct {
	Connection string `mapstructure:"connection"`
	Password   string `mapstrcuture:"password"`
}

// store environment variables on runtime
type Environment struct {
	Production  bool   `mapstructure:"production"`
	GitRoot     string `mapstructure:"git_root"`
	ProjectRoot string `mapstructure:"project_root"`
	Storage     struct {
		Db    storageEnv `mapstructure:"db"`
		Cache storageEnv `mapstructure:"cache"`
	} `maspstructure:"storage"`
	URLs struct {
		PfpApi  string `mapstructure:"pfp_api"`
		Project string `mapstructure:"project"`
	} `mapstructure:"urls"`
	Email struct {
		Host     string `mapstrcuture:"host"`
		User     string `mapstrcuture:"password"`
		Password string `mapstructure:"password"`
		Port     string `mapstructure:"port"`
	} `mapstructure:"email"`
}
