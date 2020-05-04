package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
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

func Test_usersService_AddUser(t *testing.T) {
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

func Test_usersService_GetUserByUUID(t *testing.T) {
	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{}
			got, err := s.GetUserByUUID(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_RemoveUser(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{}
			got, err := s.RemoveUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_UpdatePassword(t *testing.T) {
	type args struct {
		User domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{}
			got, err := s.UpdatePassword(tt.args.User)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdatePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_UpdateUser(t *testing.T) {
	type args struct {
		User domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usersService{}
			got, err := s.UpdateUser(tt.args.User)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
