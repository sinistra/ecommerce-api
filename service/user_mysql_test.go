package service

import (
	"github.com/joho/godotenv"
	"github.com/satori/go.uuid"
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
	assert.Nil(t, err)
}

func Test_usersService_AddUser(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")

	s := usersService{}
	got, err := s.AddUser(user)
	assert.Nil(t, err)
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
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")

	user1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user1, 1)

	user.FirstName = "Second"
	user.Email = "user2@test.com"

	user2, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user2, 2)

	keys := make(map[string][]string)
	got, err := s.GetUsers(keys)
	assert.Nil(t, err)
	// spew.Dump(got)
	assert.Len(t, got, 2)
	assert.IsType(t, []domain.User{}, got)
}

func Test_usersService_GetUser(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")

	user1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user1, 1)

	got, err := s.GetUser(1)
	assert.Nil(t, err)
	assert.Equal(t, got.FirstName, "First")
	assert.Equal(t, got.LastName, "Last")
	assert.Equal(t, got.Email, "user1@test.com")
}

func Test_usersService_GetUserByEmail(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")

	user1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user1, 1)

	got, err := s.GetUserByEmail("user1@test.com")
	assert.Nil(t, err)
	assert.Equal(t, got.Id, 1)
}

func Test_usersService_GetUserByUUID(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")
	uuidString := uuid.NewV4().String()
	user.UUID = &uuidString

	record1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, record1, 1)

	got, err := s.GetUser(record1)
	assert.Nil(t, err)
	assert.IsType(t, domain.User{}, got)

	user1, err := s.GetUserByUUID(got.UUID)
	assert.Nil(t, err)
	assert.IsType(t, domain.User{}, user1)
	assert.Equal(t, got, user1)

}

func Test_usersService_RemoveUser(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")
	uuidString := uuid.NewV4().String()
	user.UUID = &uuidString

	record1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, record1, 1)

	got, err := s.RemoveUser(record1)
	assert.Nil(t, err)
	// spew.Dump(got)
	assert.Equal(t, got, int64(1))
}

func Test_usersService_UpdatePassword(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")
	uuidString := uuid.NewV4().String()
	user.UUID = &uuidString

	record1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, record1, 1)

	user1, err := s.GetUser(record1)
	assert.Nil(t, err)
	assert.IsType(t, domain.User{}, user1)

	user.Id = record1
	user.Password = utils.EncryptPassword("password2")

	record2, err := s.UpdatePassword(user)
	assert.Nil(t, err)
	user2, err := s.GetUser(int(record2))
	assert.Nil(t, err)
	assert.IsType(t, domain.User{}, user2)

	// spew.Dump(user2)
	assert.NotEqual(t, user1.Password, user2.Password)
	assert.Equal(t, user1.Id, user2.Id)
}

func Test_usersService_UpdateUser(t *testing.T) {
	err := TruncateUserTable()
	assert.Nil(t, err)

	s := usersService{}

	user := domain.User{}
	user.FirstName = "First"
	user.LastName = "Last"
	user.Email = "user1@test.com"
	user.Password = utils.EncryptPassword("password")
	uuidString := uuid.NewV4().String()
	user.UUID = &uuidString

	record1, err := s.AddUser(user)
	assert.Nil(t, err)
	assert.Equal(t, record1, 1)

	user1, err := s.GetUser(record1)
	assert.Nil(t, err)
	assert.IsType(t, domain.User{}, user1)

	user1.FirstName = "Second"
	user1.Email = "user2@test.com"

	record2, err := s.UpdateUser(user1)
	assert.Nil(t, err)
	user2, err := s.GetUser(int(record2))
	assert.Nil(t, err)
	assert.IsType(t, domain.User{}, user2)
	assert.Equal(t, user1.Password, user2.Password)
	assert.Equal(t, user1.FirstName, user2.FirstName)
	assert.NotEqual(t, user.FirstName, user2.FirstName)
}
