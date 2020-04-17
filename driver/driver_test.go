package driver

import (
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	os.Setenv("SQL_DB", "macd_test")
	os.Setenv("SQL_HOST", "192.168.188.13")
	os.Setenv("SQL_PORT", "3306")
	os.Setenv("SQL_USER", "macdtest")
	os.Setenv("SQL_PASS", "3xtr453cur3")

	os.Setenv("MONGO_DB", "macdtest")
	os.Setenv("MONGO_HOST", "localhost")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_USER", "macdtest")
	os.Setenv("MONGO_PASS", "3xtr453cur3")
}

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TestDBConnection",
			want: "*sqlx.DB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConnectDB()
			gotType := reflect.TypeOf(got).String()
			log.Println(gotType)
			if gotType != tt.want {
				t.Errorf("ConnectDB() = %v, want %v", gotType, tt.want)
			}
		})
	}
}

func TestConnectMongo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TestMongoConnection",
			want: "*mgo.Session",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConnectMongo()
			gotType := reflect.TypeOf(got).String()
			log.Println(gotType)
			if gotType != tt.want {
				t.Errorf("ConnectDB() = %v, want %v", gotType, tt.want)
			}
		})
	}
}
