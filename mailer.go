// Copyright (c) 2014 Maxime SIMON. All rights reserved.

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

func NewMailerFromConfig(filename string) (*Mailer, error) {
	props, err := LoadConfig(filename)
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
