package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/utils"
)

func init() {
	err := godotenv.Load("../test.env")
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestTruncateUserTable(t *testing.T) {
	err := TruncateUserTable()
	if err != nil {
		log.Println(err)
	}
	assert.Nil(t, err)
}

func Test_usersService_AddUser1(t *testing.T) {
	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")

	s := usersService{}
	got, err := s.AddUser(user)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, got, 1)

	got, err = s.AddUser(user)
	assert.Equal(t, got, 0)
	assert.EqualError(t, err, "Error 1062: Duplicate entry 'user1@test.com' for key 'users_email_uindex'")

	user.FirstName = "Second"
	user.LastName = "Last"
	user.Email = "user2@test.com"
	user.Password = utils.EncryptPassword("password2")
	got, err = s.AddUser(user)
	assert.Equal(t, got, 3)
	assert.Nil(t, err)
}

func Test_usersService_GetUsers(t *testing.T) {
	s := usersService{}
	keys := make(map[string][]string)
	got, err := s.GetUsers(keys)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(got)
	assert.Len(t, got, 2)
	assert.Nil(t, err)
}

func Test_usersService_GetUser(t *testing.T) {
	s := usersService{}
	got, err := s.GetUser(3)
	assert.Nil(t, err)
	assert.Equal(t, got.FirstName, "Second")
	assert.Equal(t, got.LastName, "Last")
	assert.Equal(t, got.Email, "user2@test.com")

	got, err = s.GetUser(2)
	// spew.Dump(got)
	assert.Equal(t, got.Id, 0)
	assert.Error(t, err, "not found")

}

func Test_usersService_GetUserByEmail(t *testing.T) {
	s := usersService{}
	got, err := s.GetUserByEmail("user2@test.com")
	assert.Nil(t, err)
	assert.Equal(t, got.Id, 3)
}
