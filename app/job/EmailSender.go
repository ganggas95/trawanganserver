package job

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

func Send(email, isi string) error {
	from := mail.Address{"subhannizar25", "subhannizar25@gmail.com"}
	to := mail.Address{"ijangnizar09", email}
	body := isi
	subject := "this is the subject line"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// setup the remote smtpserver & auth info
	smtpserver := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", "subhannizar25@gmail.com", "K3yf4t0n", "smtp.gmail.com")

	// setup a map for the headers
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = subject
	header["Mime"] = mime

	// setup the message
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// create the smtp connection
	c, err := smtp.Dial(smtpserver)
	if err != nil {
		return err
	}

	// set some TLS options, so we can make sure a non-verified cert won't stop us sending
	host, _, _ := net.SplitHostPort(smtpserver)
	tlc := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	if err = c.StartTLS(tlc); err != nil {
		return err
	}

	// auth stuff
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}
	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	c.Quit()
	return nil
}
