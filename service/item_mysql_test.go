package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
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

	// log.Println("start item 2")

	got, err = s.AddItem(item)
	if err != nil {
		log.Println(err)
	}
	assert.EqualError(t, err, "Error 1062: Duplicate entry 'product-code' for key 'items_code_uindex'")
	assert.Equal(t, got, 0)

	// log.Println("start item 3")

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
	recordID := 1
	got, err := ItemsService.GetItem(recordID)

	assert.Nil(t, err)
	assert.IsType(t, domain.Item{}, got)
	// spew.Dump(got)
	assert.Equal(t, got.Description, "Product Description")
	assert.Equal(t, got.SoldQuantity, 10)
}

func Test_itemsService_GetItems(t *testing.T) {
	keys := make(map[string][]string)
	got, err := ItemsService.GetItems(keys)

	assert.Nil(t, err)
	// spew.Dump(got)
	assert.Equal(t, len(got), 2)
	assert.IsType(t, got, []domain.Item{})

}

func Test_itemsService_GetItemByCode(t *testing.T) {
	code := "product-code"
	got, err := ItemsService.GetItemByCode(code)
	assert.Nil(t, err)
	// spew.Dump(got)
	assert.IsType(t, domain.Item{}, got)
	assert.Equal(t, got.Status, "Active")
}

func Test_itemsService_UpdateItem(t *testing.T) {
	recordID := 1
	item, err := ItemsService.GetItem(recordID)
	assert.Nil(t, err)
	assert.IsType(t, domain.Item{}, item)

	item.Description = "Updated Description"
	got, err := ItemsService.UpdateItem(item)
	assert.Nil(t, err)
	assert.Equal(t, got, int64(1))
	spew.Dump(got)
	item2, err := ItemsService.GetItem(recordID)
	assert.Nil(t, err)
	assert.IsType(t, domain.Item{}, item2)
	assert.Equal(t, item.Code, item2.Code)
	assert.Equal(t, item.Description, item2.Description)
	assert.Equal(t, item.Status, item2.Status)
}

func Test_itemsService_RemoveItem(t *testing.T) {
	recordID := 1
	got, err := ItemsService.RemoveItem(recordID)
	assert.Nil(t, err)
	assert.Equal(t, got, int64(1))

	item, err := ItemsService.GetItem(recordID)
	// spew.Dump(item)
	assert.Equal(t, err.Error(), "sql: no rows in result set")
	assert.IsType(t, domain.Item{}, item)

}
