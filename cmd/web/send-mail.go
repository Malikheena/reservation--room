package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/malikheena/heena2-main/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}

	}()
}
func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10  *time.Second

	Client, err := server.Connect()
	if err !=nil{
		errorLog.Println(err)
	}
	email :=mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template ==""{
	email.SetBody(mail.TextHTML, m.Content)
	}else{
		data, err:= ioutil.ReadFile(fmt.Sprintf(".email-templates/%s",m.Template))
		if err != nil{
			app.Errorlog.Println(err)
		}
		mailTempalte := string(data)
		msgToSend := strings.Replace(mailTempalte, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}
	err = email.Send(Client)
	if err != nil{
		log.Println(err)
	}else{
		log.Println("Email sent")
	}
	}





