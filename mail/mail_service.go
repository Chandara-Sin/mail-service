package mail

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"mime"
	"net/mail"
	"net/smtp"

	"github.com/spf13/viper"
)

func SendEmail() func(Mail) error {
	return func(ml Mail) error {
		to := mail.Address{Name: "dome", Address: ml.To[0]}
		from := mail.Address{Name: viper.GetString("smtp.from"), Address: viper.GetString("smtp.sender")}
		addr := fmt.Sprintf("%v:%v", viper.GetString("smtp.host"), viper.GetString("smtp.port"))
		host := fmt.Sprintf("%v", viper.GetString("smtp.host"))

		auth := smtp.PlainAuth("", viper.GetString("smtp.user"), viper.GetString("smtp.password"), host)
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}
		connection, err := tls.Dial("tcp", addr, tlsConfig)
		if err != nil {
			fmt.Println(err)
		}
		smtpClient, err := smtp.NewClient(connection, host)
		if err != nil {
			fmt.Println(err)
		}
		if err = smtpClient.Auth(auth); err != nil {
			fmt.Println(err)
		}
		if err = smtpClient.Mail(from.Address); err != nil {
			fmt.Println(err)
		}
		if err = smtpClient.Rcpt(to.Address); err != nil {
			fmt.Println(err)
		}
		writer, err := smtpClient.Data()
		if err != nil {
			fmt.Println(err)
		}
		_, err = writer.Write([]byte(ml.setDefaultTemplate(to, from)))
		if err != nil {
			fmt.Println(err)
		}
		err = writer.Close()
		if err != nil {
			fmt.Println(err)
		}
		err = smtpClient.Quit()

		return err
	}
}

func (ml *Mail) setDefaultTemplate(to, from mail.Address) string {
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = mime.QEncoding.Encode("UTF-8", ml.Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(ml.Template))
	return message
}
