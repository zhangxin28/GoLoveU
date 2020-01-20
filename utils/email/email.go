package email

import (
	"bytes"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"goloveu/utils"
	"html/template"
	"net"
	"net/smtp"
)

var emailTemplate = `
<div style="background-color:white;border-top:2px solid #12ADDB;box-shadow:0 1px 3px #AAAAAA;line-height:180%;padding:0 15px 12px;width:500px;margin:50px auto;color:#555555;font-family:'Century Gothic','Trebuchet MS','Hiragino Sans GB',微软雅黑,'Microsoft Yahei',Tahoma,Helvetica,Arial,'SimSun',sans-serif;font-size:12px;">
    <h2 style="border-bottom:1px solid #DDD;font-size:14px;font-weight:normal;padding:13px 0 10px 8px;">
        <span style="color: #12ADDB;font-weight:bold;">
            {{.Title}}
        </span>
    </h2>
    <div style="padding:0 12px 0 12px; margin-top:18px;">
        {{if .Content}}
		<p>
            {{.Content}}
        </p>
		{{end}}
		{{if .QuoteContent}}
		<div style="background-color: #f5f5f5;padding: 10px 15px;margin:18px 0;word-wrap:break-word;">
            {{.QuoteContent}}
        </div>
		{{end}}
       
		{{if .Url}}
        <p>
            <a style="text-decoration:none; color:#12addb" href="{{.Url}}" target="_blank" rel="noopener">点击查看详情</a>
        </p>
		{{end}}
    </div>
</div>
`

// BuildEmailTemplate builds the default email template
func BuildEmailTemplate(title, content, quoteContent, url string) (string, error) {
	tpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		utils.LogError(err)
		return "", err
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, map[string]interface{}{
		"Title":        title,
		"Content":      content,
		"QuoteContent": quoteContent,
		"Url":          url,
	})
	if err != nil {
		utils.LogError(err)
		return "", err
	}
	return b.String(), nil
}

// SendEmail performs to send email
func SendEmail(to, from, identity, username, password, port, host, subject, html string, ssl bool) {
	addr := net.JoinHostPort(host, port)
	auth := smtp.PlainAuth(identity, username, password, host)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	e := email.NewEmail()
	e.From = username
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(html)

	if ssl {
		if err := e.SendWithTLS(addr, auth, tlsConfig); err != nil {
			utils.LogError("发送邮件异常", err)
		} else {
			if err := e.Send(addr, auth); err != nil {
				utils.LogError("发送邮件异常", err)
			}
		}
	}
}
