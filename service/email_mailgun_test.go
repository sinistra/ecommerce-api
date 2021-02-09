package service

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"

	"github.com/sinistra/ecommerce-api/domain"
)

func Test_mailgunService_SendHtmlMessage(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	s := mailgunService{
		Sender:    "Excited User <mailgun@thenotabene.com>",
		Subject:   "Test email",
		Body:      "Test email",
		Recipient: "paul@nms.com.au",
	}

	err = s.SendHtmlMessage()
	assert.Nil(t, err)
}

func Test_mailgunService_SendVerificationEmail(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	user := domain.User{}
	user.Email = "paul@nms.com.au"
	user.FirstName = "Test"

	s := mailgunService{
		Sender:    "Excited User <mailgun@thenotabene.com>",
		Subject:   "Test Subject",
		Body:      "Test Body",
		Recipient: "",
	}

	err = s.SendVerificationEmail(user)
	assert.Nil(t, err)
}
