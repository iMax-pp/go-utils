package utils

import (
	"net/smtp"
)

type Mailer struct {
	Server string
	Port   string
	From   string
	To     string
}

func NewMailer(server, port, from, to string) *Mailer {
	return &Mailer{Server: server, Port: port, From: from, To: to}
}

func NewMailerFromConfig(f string) (*Mailer, error) {
	props := make(map[string]string)
	err := LoadConfig(f, props)
	if err != nil {
		return nil, err
	}

	return NewMailer(props["smtp.server"], props["smtp.port"], props["mail.from"], props["mail.to"]), nil
}

func (mail *Mailer) SendMail(msg string) error {
	client, err := smtp.Dial(mail.Server + ":" + mail.Port)
	if err != nil {
		return err
	}
	defer client.Close()

	client.Mail(mail.From)
	client.Rcpt(mail.To)
	wc, err := client.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	if _, err = wc.Write([]byte(msg)); err != nil {
		return err
	}

	return nil
}
