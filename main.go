package main

import (
	"context"
	"fmt"
	"log"
	"mail-service/mail"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok v1")
	})

	r.POST("/v1/smtp/mail", mail.SendEmailHandler(mail.SendEmail()))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + viper.GetString("app.port"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
