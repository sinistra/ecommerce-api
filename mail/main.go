package main

import (
	"bytes"
	"context"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/sinistra/ecommerce-api/domain"
)

var mgDomain string
var mgKey string

type MgMessage struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}

func main() {
	godotenv.Load("../.env")
	var ok bool
	mgDomain, ok = os.LookupEnv("MAILGUN_DOMAIN")
	if !ok {
		log.Fatal("cannot read domain")
	}
	mgKey, ok = os.LookupEnv("MAILGUN_KEY")
	if !ok {
		log.Fatal("cannot read apiKey")
	}

	uuid := "e8ca9c69-1359-4169-ace8-b1c762336dbb"
	user := domain.User{
		FirstName: "Paul",
		LastName:  "Taylor",
		Email:     "paul@nms.com.au",
		UUID:      &uuid,
	}

	body, err := parseTemplate("../templates/verify_email.html", user)
	if err != nil {
		log.Fatal(err)
	}

	mgMessage := MgMessage{
		Sender:  "Excited User <mailgun@thenotabene.com>",
		Subject: "Welcome to the Nota Bene",

		Body:      body,
		Recipient: user.Email,
	}

	SendMgMessage(mgMessage)
}

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Excited User <mailgun@thenotabene.com>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"paul@nms.com.au",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

func SendMgMessage(mm MgMessage) {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(mgDomain, mgKey)
	text := ""

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(mm.Sender, mm.Subject, text, mm.Recipient)
	// set the text to empty string and html to below
	message.SetHtml(string(mm.Body))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)
}

func parseTemplate(fileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
