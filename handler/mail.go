package Accountapi

import (
	"log"
	"strings"

	"github.com/alexcesaro/mail/gomail"
)

func mail(token string) bool {

	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", "kung20706@gmail.com", "BearLab")
	msg.SetHeader("To", token)
	// msg.AddHeader("To", "kung20707@gmail.com")
	msg.SetHeader("Subject", "Hello!")
	msg.SetBody("text/plain", "變更您所使用的密碼")
	text := `
	<html>
	
	<head>
	  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	  <title>Wonderland 註冊會員</title>
	  <style type="text/css">
		body {
		  font-family: 'Arial','Open Sans',sans-serif;
		  margin: 0;
		  padding: 0;
		  min-width: 100% !important;
		}
		.content {
		  width: 100%;
		  max-width: 600px;
		}
		
		.header {
		  padding: 40px 30px 20px 30px;
		}
		
		.big-title {
		  color: #222222;
		  font-family: montserrat;
		  text-align: center;
		  font-size: 42px;
		  font-weight: 900;
		}
		
		.body {
		  padding: 32px 32px;
		  color: white;
		  font-family: montserrat, helvetica, arial, sans-serif;
		  background-color: #eeeeee;
		}
		
		.txt-drk {
		  color: #3a3a3a;
		}
		
		.body > p {
		  font-size: 14px;
		}
		
		.btn {
		  text-align: center;
		  background-color: #9dc1e1;
		  color: white;
		  border-radius: 8px;
		  font-size: 24px;
		  padding: 8px 24px;
		  margin-left: auto;
		  text-decoration: none;
		}
		
		.footer {
		  margin: 128px 32px;
		  background-color: #9dc1e1;
		  padding: 128px 32px;
		}
		
		.footer > p, .footer > * > p {
		  font-weight: 100;
		  font-size: 11px;
		  font-family: verdana;
		}
		
		a:not(.btn) {
		  color: #f1683e;
		}
		
		.footer > p > a:not(.btn), .footer > * > p > a:not(.btn) {
		  color: #f1683e;
		  color: #e6e6e6;
		}
	  </style>
	
	  <style type="text/css">
		@media only screen and (min-device-width: 601px) {
		  .content {
			width: 600px !important;
		  }
		}
	  </style>
	</head>
	
	<body>
	  <table width="100%" bgcolor="white" border="0" cellpadding="0" cellspacing="0">
		<tr>
		  <td>
			<table class="content" align="center" cellpadding="0" cellspacing="0" border="0" style="">
			  <tr>
				<td class="header" bgcolor="#ffffff" style="background-image: url('https://rescuedigital.net/images/email/image.jpg'); background-size: cover; background-position: center bottom">
				  <table width="500" align="center" border="0" cellpadding="0" cellspacing="0">
					<tr>
					  <td height="70" style="padding: 0 20px 20px 0;">
					   
						<h1 class="big-title" align=center>重設密碼</h1>
					  </td>
					</tr>
				  </table>
				</td>
				<tr class="body">
				  <td class="body">
					<h2 class="txt-drk" align=center>忘記密碼</h2>
					<p class="txt-drk"align=center >歡迎使用，點擊下方按鈕設置新的密碼</p>
				  </td>
				  </tr>
			  <tr class="body">
				<td align="center">
				  <a class="btn" href="http://192.168.0.100:8081/resetpassword2" title="Verify your email">驗證</a>
				</td>
			  </tr>
				 <tr class="body">
				  <td class="body">
					<p class="txt-drk" align=center>如果您沒有在本站註冊，請忽略本信件或通知本站客服</p>
					
				  </td>
				  </tr>
			  
			  <tr class="body footer">
				<td align="center" cellpadding=50 style="padding: 32px 24px">
				   </br></br>
				
			  </tr>
			  
			  </tr>
			</table>
		  </td>
		  </tr>
	  </table>
	</body>
	
	</html>`
	form := strings.Replace(text, "wEgyaPhGhbRfscwPWjpMHqpeHLHD7cK9", "fuckoff", -1)
	msg.AddAlternative("text/html", form)

	m := gomail.NewMailer("smtp.gmail.com", "kung20706", "kung22980991", 587)
	if err := m.Send(msg); err != nil {
		log.Println(err)
	}
	return true
}
