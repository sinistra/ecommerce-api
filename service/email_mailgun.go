package service

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"os"
	"time"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/utils"
)

var MailgunService mailgunServiceInterface = &mailgunService{}

type mailgunService struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}

type mailgunServiceInterface interface {
	SendVerificationEmail(user domain.User) error
	SendHtmlMessage() error
}

func (s mailgunService) SendVerificationEmail(user domain.User) error {
	s.Sender = "Excited User <mailgun@thenotabene.com>"
	s.Subject = "Welcome to the Nota Bene"
	body, err := utils.ParseTemplate("./templates/verify_email.html", user)
	if err != nil {
		log.Fatal(err)
	}
	s.Body = body
	s.Recipient = user.Email

	err = s.SendHtmlMessage()
	if err != nil {
		return err
	}
	return nil
}

func (s mailgunService) SendHtmlMessage() error {
	var ok bool
	mgDomain, ok := os.LookupEnv("MAILGUN_DOMAIN")
	if !ok {
		log.Fatal("cannot read MAILGUN_DOMAIN")
	}
	mgKey, ok := os.LookupEnv("MAILGUN_KEY")
	if !ok {
		log.Fatal("cannot read MAILGUN_KEY")
	}

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(mgDomain, mgKey)
	text := ""

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(s.Sender, s.Subject, text, s.Recipient)
	// set the text to empty string and html to below
	message.SetHtml(string(s.Body))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
