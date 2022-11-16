package conf

import "github.com/spf13/viper"

func LoadConfig(path string) {
	// setup authentication configuration
	viper.SetConfigName("twitterbot")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
