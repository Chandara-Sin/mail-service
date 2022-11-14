package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok v1")
	})

	r.Run(":" + viper.GetString("app.port"))
}

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err) // Handle errors reading the config file
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
