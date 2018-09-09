package handler

import (
	"fmt"
	"g"
	"log"
	"strings"
	"time"

	"github.com/go-gomail/gomail"
)

func Sendmail(isok bool) {
	m := gomail.NewMessage()
	m.SetHeader("From", g.Config().From)
	tos := g.Config().Tos
	toslice := strings.Split(tos, ",")
	m.SetHeader("To", toslice...)
	subject := fmt.Sprintf("平台巡检报告-%v", time.Now().In(time.Local))
	m.SetHeader("Subject", subject)
	if !isok {
		m.SetBody("text/html", "本次巡检有异常，请查看巡检结果！")
	} else {
		m.SetBody("text/html", "本次巡检无异常！")
	}

	m.Attach(g.Config().ExportExecelPath)

	d := gomail.NewDialer(g.Config().MailServer, g.Config().MailServerPort, g.Config().User, g.Config().Passwd)

	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail failed,err:", err.Error())
		return
	}
	log.Print("send mail success")
}
