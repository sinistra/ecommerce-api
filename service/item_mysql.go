package service

import (
	"log"

	"github.com/sinistra/ecommerce-api/domain"
	"github.com/sinistra/ecommerce-api/driver"
)

const (
	queryInsertItem    = "INSERT INTO items(code, title, description, seller, image, price, qty_avail, qty_sold, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryGetItem       = "SELECT * FROM items WHERE id=?;"
	queryGetItemByCode = "SELECT * FROM items WHERE code=?;"
	queryGetItems      = "SELECT * FROM items"
	queryUpdateItem    = "UPDATE items SET code=?, title=?, description=?, seller=?, image=?, price=?, qty_avail=?, qty_sold=?, status=? WHERE id=?;"
	queryDeleteItem    = "DELETE FROM items WHERE id=?;"
	queryTruncateItems = "TRUNCATE items"
)

// type ItemService interface to itemsService{}
var ItemsService itemsServiceInterface = &itemsService{}

type itemsService struct{}

type itemsServiceInterface interface {
	GetItems(map[string]string) ([]domain.Item, error)
	GetItem(id int) (domain.Item, error)
	GetItemByCode(code string) (domain.Item, error)
	AddItem(Item domain.Item) (int, error)
	UpdateItem(Item domain.Item) (int64, error)
	RemoveItem(id int) (int64, error)
}

func (s itemsService) GetItems(keys map[string]string) ([]domain.Item, error) {
	db := driver.ConnectDB()
	defer db.Close()
	var items []domain.Item

	sql := queryGetItems
	count := 0
	for index, key := range keys {
		if len(key) > 0 {
			if count > 0 {
				sql = sql + " AND"
			} else {
				sql = sql + " WHERE"
			}
			sql = sql + " " + index + "='" + key + "'"
			count++
		}
	}

	sql = sql + " ORDER BY id ASC"
	err := db.Select(&items, sql)

	if err != nil {
		return []domain.Item{}, err
	}

	return items, nil
}

func (s itemsService) GetItem(id int) (domain.Item, error) {
	db := driver.ConnectDB()
	defer db.Close()
	var item domain.Item

	err := db.Get(&item, queryGetItem, id)

	return item, err
}

func (s itemsService) GetItemByCode(code string) (domain.Item, error) {
	db := driver.ConnectDB()
	defer db.Close()
	var item domain.Item

	err := db.Get(&item, queryGetItemByCode, code)

	return item, err
}

func (s itemsService) AddItem(Item domain.Item) (int, error) {
	db := driver.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(queryInsertItem)
	if err != nil {
		log.Println(err)
	}
	res, err := stmt.Exec(Item.Code, Item.Title, Item.Description, Item.Seller, Item.Image, Item.Price,
		Item.AvailableQuantity, Item.SoldQuantity, Item.Status)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Printf("Id = %d, affected = %d\n", lastId, rowCnt)

	return int(lastId), nil
}

func (s itemsService) UpdateItem(Item domain.Item) (int64, error) {
	db := driver.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(queryUpdateItem)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(Item.Code, Item.Title, Item.Description, Item.Seller, Item.Image, Item.Price, Item.AvailableQuantity,
		Item.SoldQuantity, Item.Status, Item.Id)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return rowCnt, err
	}

	return rowCnt, nil
}

func (s itemsService) RemoveItem(id int) (int64, error) {
	db := driver.ConnectDB()
	defer db.Close()

	result, err := db.Exec(queryDeleteItem, id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

func TruncateItemTable() error {
	db := driver.ConnectDB()
	defer db.Close()

	_, err := db.Exec(queryTruncateItems)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
