package config

type Config struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"app"`
	Mysql struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"databasename"`
	} `yaml:"mysql"`
}
