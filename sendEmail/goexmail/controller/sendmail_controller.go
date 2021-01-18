package controller

import (
	"fmt"
	"goexmail/model"
	"net/smtp"
	"strings"
)

func SendEmail(content, subject string)  {
	emailModel := model.InitEmailModel()
	fmt.Println("send email")
	emailModel.Body = content
	emailModel.Subject = subject
	err := sendToMail(emailModel,"html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

func sendToMail(email model.EmailModle, mailType string) error {
	auth := smtp.PlainAuth("", email.User, email.Pwd, email.Host)
	var contentType string
	if mailType == "html" {
		contentType = "Content_Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content_Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To:" + email.Receiver + "\r\nFrom: " + email.NickName + "<" +
		email.User + ">\r\nSubject: " + email.Subject + "\r\n" +
		contentType + "\r\n\r\n " + email.Body)
	sendToUser := strings.Split(email.Receiver, ";")
	addr := fmt.Sprintf("%s:%s", email.Host, email.Port)
	err := smtp.SendMail(addr, auth, email.User, sendToUser, msg)
	return err
}

