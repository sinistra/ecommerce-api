package service

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"

	"github.com/sinistra/ecommerce-api/domain"
)

func init() {
	err := godotenv.Load("../test.env")
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestTruncateItemTable(t *testing.T) {
	err := TruncateItemTable()
	assert.Nil(t, err)
}

func Test_itemsService_AddItem(t *testing.T) {
	err := TruncateItemTable()
	assert.Nil(t, err)

	item := domain.Item{}
	item.Code = "product-code"
	item.Description = "Product Description"
	item.Status = "Active"
	item.Seller = 99
	item.AvailableQuantity = 9
	item.SoldQuantity = 10
	item.Price = 99.99
	item.Picture = "path to picture"

	s := itemsService{}
	got, err := s.AddItem(item)
	assert.Nil(t, err)
	assert.Equal(t, got, 1)

	log.Println("start item 2")

	got, err = s.AddItem(item)
	if err != nil {
		log.Println(err)
	}
	assert.EqualError(t, err, "Error 1062: Duplicate entry 'product-code' for key 'items_code_uindex'")
	assert.Equal(t, got, 0)

	log.Println("start item 3")

	item.Code = "product-code2"
	item.Description = "2nd Product Description"
	item.Status = "Active"
	item.Seller = 11
	item.AvailableQuantity = 12
	item.SoldQuantity = 55
	item.Price = 0.78
	item.Picture = "path to picture 2"

	got, err = s.AddItem(item)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, got, 3)
	assert.Nil(t, err)

}

func Test_itemsService_GetItem(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := itemsService{}
			got, err := s.GetItem(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsService_GetItems(t *testing.T) {
	type args struct {
		keys map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := itemsService{}
			got, err := s.GetItems(tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItems() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsService_RemoveItem(t *testing.T) {
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
			s := itemsService{}
			got, err := s.RemoveItem(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsService_UpdateItem(t *testing.T) {
	type args struct {
		Item domain.Item
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
			s := itemsService{}
			got, err := s.UpdateItem(tt.args.Item)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsService_GetItemByCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := itemsService{}
			got, err := s.GetItemByCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemByCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
