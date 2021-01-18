package model

import "goexmail/config"

type EmailModle struct {
	Host     string // 邮件服务器地址
	Port     string // 端口
	User     string // 发送邮件用户账号
	Pwd      string // 授权密码
	NickName string // 发送邮件昵称
	Receiver string // 邮件接收者
	Body     string // 邮件内容
	Subject  string // 邮件主题
}


func InitEmailModel() EmailModle {
	return EmailModle{
		Host:     config.SMTP_MAIL_HOST,
		Port:     config.SMTP_MAIL_PORT,
		User:     config.SMTP_MAIL_USER,
		Pwd:      config.SMTP_MAIL_PWD,
		NickName: config.SMTP_MAIL_NICKNAME,
		Receiver: config.SMTP_MAIL_RECEIVER,
		Body:     config.SMTP_MAIL_BODY,
		Subject:  config.SMTP_MAIL_SUBJECT,
	}
}