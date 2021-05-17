package utils

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
	"time"
)

var GlobalSessions *session.Manager
var Session session.Store

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig)

	go GlobalSessions.GC()
}

//存储类型

//更多存储类型有待扩展
const (
	Version           = "1.0"
	StoreLocal string = "local"
	StoreOss   string = "oss"
)

var (
	BasePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	StoreType   = beego.AppConfig.String("store_type") //存储类型
)

func SendMail(name, mobile, toUser, toUsername, title, subject string, id int) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = "qq570850096@163.com"
	// 收件人(可以有多个)
	e.To = []string{toUser}
	// 邮件主题
	e.Subject = subject
	// 解析html模板
	t, err := template.ParseFiles("views/email-template.html")
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
		Id           int
	}{
		FromUserName: name,
		ToUserName:   toUsername,
		TimeDate:     time.Now().Format("2006/01/02"),
		Message:      "尊敬的" + toUsername + "您好, 我求职贵司的" + title + "职位,这里是我的简历。",
		Id:           id,
	})
	// html形式的消息
	e.HTML = body.Bytes()
	// 以路径将文件作为附件添加到邮件中
	fname := "tmp/" + name + "-" + mobile + ".pdf"
	e.AttachFile(fname)
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	err = e.Send("smtp.163.com:25", smtp.PlainAuth("", "qq570850096@163.com", "VWXYGAVRMNVZYJUG", "smtp.163.com"))
	fmt.Println(err)
	return err
}

func SendMailToSchool(name, toUser, toUsername, subject string) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = "qq570850096@163.com"
	// 收件人(可以有多个)
	e.To = []string{toUser}
	// 邮件主题
	e.Subject = subject
	// 解析html模板
	t, err := template.ParseFiles("views/email-template.html")
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
		Id           int
	}{
		FromUserName: name,
		ToUserName:   toUsername,
		TimeDate:     time.Now().Format("2006/01/02"),
		Message:      "尊敬的" + toUsername + "您好, 这里是本月的入职情况统计表格，请查收。",
	})
	// html形式的消息
	e.HTML = body.Bytes()
	// 以路径将文件作为附件添加到邮件中
	fname := "tmp/recruit.xlsx"
	e.AttachFile(fname)
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	err = e.Send("smtp.163.com:25", smtp.PlainAuth("", "qq570850096@163.com", "VWXYGAVRMNVZYJUG", "smtp.163.com"))
	fmt.Println(err)
	return err
}
