package usecase

import (
	// "context"
	"bytes"
	"fmt"
	"html/template"
	"log"
	r "regate-core/domain/repository"
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type mailUcase struct{
	gomailAuth     *gomail.Dialer
}

type Chart struct {
	Code string
}


func NewUseCase(gomailAuth *gomail.Dialer) r.MailUseCase{
	return &mailUcase{
		gomailAuth: gomailAuth,
	}
}
func(u *mailUcase)SendResetPasswordEmail(email string,token string){
	path := fmt.Sprintf(`%semail-templates/reset_password.html`,viper.GetString("path"))
	t, _ := template.ParseFiles(path)
	url := viper.GetString("admin_web_url") + "/auth/password-reset/" + token
	urlReset := viper.GetString("admin_web_url") + "/auth/forgot-password"
	var body bytes.Buffer
		data := struct {
			Url string
			UrlReset string
		}{
			Url: url,
			UrlReset: urlReset,
		}
		t.Execute(&body, data)
		u.SendEmail([]string{email},&body,"Credenciales de usuario")
}

func(u *mailUcase)SendEmailUserAdmin(email string,password string){
	path := fmt.Sprintf(`%semail-templates/user_admin_account.html`,viper.GetString("path"))
	t, err := template.ParseFiles(path)
	if err != nil{
		log.Println(err)
	}
	url := viper.GetString("admin_web_url") + "/auth/login"
	var body bytes.Buffer
		data := struct {
			Url string
			Username string
			Email string
			Password string
			Content string
			EmpresaName string
		}{
			Url: url,
			Username:"Jorge",
			Email:email,
			Password:password,
			EmpresaName:"Complejo deportivo las palmas",
		}
		log.Println("1---------",password)
		t.Execute(&body, data)
		log.Println("2---------")
		err = u.SendEmail([]string{email},&body,"Credenciales de usuario")
		if err != nil {
			log.Println(err)
		}
		log.Println("3---------")
}

func(u *mailUcase)SendEmailVerification(email string,otp int,username string){
	path := fmt.Sprintf(`%semail-templates/email-verification.html`,viper.GetString("path"))
	log.Println(otp)
	t, err := template.ParseFiles(path)
	if err != nil{
		log.Println(err)
	}
	otpString := strconv.Itoa(otp)
	var otpCode []Chart
	for _,char := range otpString{
		t := Chart{}
		t.Code = string(char)
		otpCode = append(otpCode, t)
	}
	var body bytes.Buffer
		data := struct {
			OtpCode []Chart 
			Username string
		}{
			OtpCode: otpCode,
			Username:username,
		}
		t.Execute(&body, data)
		log.Println("2---------")
		err = u.SendEmail([]string{email},&body,"Email Verificacion")
		if err != nil {
			log.Println(err)
		}
		log.Println("3---------")
}



func(u *mailUcase)SendEmail(emails []string, body *bytes.Buffer,subject string)(err error){
		m := gomail.NewMessage()	
		m.SetHeader("From", viper.GetString("mail.user"))
		m.SetHeader("To", emails...)
		// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body.String())
		// m.Attach("/home/Alex/lolcat.jpg")
		// d := gomail.NewDialer("mail.teclu.com", 25, "jmiranda@teclu.com", "jmiranda2022")
		if err = u.gomailAuth.DialAndSend(m); err != nil {
			return 
		}
		return
}
