package mail

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

func SendEmail() func(Mail) error {
	return func(ml Mail) error {
		addr := fmt.Sprintf("%v:%v", viper.GetString("smtp.host"), viper.GetString("smtp.port"))
		host := fmt.Sprintf("%v", viper.GetString("smtp.host"))
		auth := smtp.PlainAuth("", viper.GetString("smtp.user"), viper.GetString("smtp.password"), host)
		err := smtp.SendMail(addr, auth, viper.GetString("smtp.sender"), ml.To, []byte(ml.setDefaultTemplate()))
		return err
	}
}

func (ml *Mail) setDefaultTemplate() string {
	temp := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	temp += fmt.Sprintf("From: %s\r\n", ml.From)
	temp += fmt.Sprintf("To: %s%s\r\n", ml.To[0], ";")
	temp += fmt.Sprintf("Subject: %s\r\n", ml.Subject)
	temp += fmt.Sprintf("\r\n%s\r\n", ml.Template)
	return temp
}
