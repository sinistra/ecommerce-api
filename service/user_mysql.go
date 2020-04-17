package service

import (
    "log"

    "github.com/sinistra/ecommerce-api/domain"
    "github.com/sinistra/ecommerce-api/driver"
)

const (
    queryInsertUser             = "INSERT INTO users(first_name, last_name, email, status, password) VALUES(?, ?, ?, ?, ?);"
    queryGetUser                = "SELECT * FROM users WHERE id=?;"
    queryGetUserByEmail         = "SELECT * FROM users WHERE email=?"
    queryGetUsers               = "SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users"
    queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
    queryUpdateUserPassword     = "UPDATE users SET password=? WHERE id=?;"
    queryDeleteUser             = "DELETE FROM users WHERE id=?;"
)

// type UserService struct{}
var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
    GetUsers(map[string][]string) ([]domain.User, error)
    GetUser(id int) (domain.User, error)
    GetUserByEmail(email string) (domain.User, error)
    AddUser(User domain.User) (int, error)
    UpdateUser(User domain.User) (int64, error)
    UpdatePassword(User domain.User) (int64, error)
    RemoveUser(id int) (int64, error)
}

func (s usersService) GetUsers(keys map[string][]string) ([]domain.User, error) {
    db := driver.ConnectDB()
    defer db.Close()
    var users []domain.User

    sql := queryGetUsers
    if len(keys) > 0 {
        count := 0
        sql = sql + " WHERE"
        for index, key := range keys {
            if count > 0 {
                sql = sql + " AND"
            }
            sql = sql + " " + index + "='" + key[0] + "'"
            count++
        }
    }

    sql = sql + " ORDER BY id ASC"
    err := db.Select(&users, sql)

    if err != nil {
        return []domain.User{}, err
    }

    return users, nil
}

func (s usersService) GetUser(id int) (domain.User, error) {
    db := driver.ConnectDB()
    defer db.Close()
    var user domain.User

    err := db.Get(&user, queryGetUser, id)

    return user, err
}

func (s usersService) GetUserByEmail(email string) (domain.User, error) {
    db := driver.ConnectDB()
    defer db.Close()
    var user domain.User

    err := db.Get(&user, queryGetUserByEmail, email)

    return user, err
}

func (s usersService) AddUser(User domain.User) (int, error) {
    db := driver.ConnectDB()
    defer db.Close()

    stmt, err := db.Prepare(queryInsertUser)
    if err != nil {
        log.Println(err)
    }
    res, err := stmt.Exec(User.FirstName, User.LastName, User.Email, User.Status, User.Password)
    if err != nil {
        log.Println(err)
    }
    lastId, err := res.LastInsertId()
    if err != nil {
        log.Println(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Println(err)
    }
    log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

    return int(lastId), nil
}

func (s usersService) UpdateUser(User domain.User) (int64, error) {
    db := driver.ConnectDB()
    defer db.Close()

    stmt, err := db.Prepare(queryUpdateUser)

    if err != nil {
        log.Println(err)
    }
    res, err := stmt.Exec(User.FirstName, User.LastName, User.Email, User.Status, User.Password, User.Id)

    if err != nil {
        log.Println(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Println(err)
    }

    return rowCnt, nil
}

func (s usersService) UpdatePassword(User domain.User) (int64, error) {
    db := driver.ConnectDB()
    defer db.Close()

    stmt, err := db.Prepare(queryUpdateUserPassword)

    if err != nil {
        log.Println(err)
    }
    res, err := stmt.Exec(User.Password, User.Id)

    if err != nil {
        log.Println(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Println(err)
    }

    return rowCnt, nil
}

func (s usersService) RemoveUser(id int) (int64, error) {
    db := driver.ConnectDB()
    defer db.Close()

    result, err := db.Exec(queryDeleteUser, id)
    if err != nil {
        return 0, err
    }

    rowsDeleted, err := result.RowsAffected()

    if err != nil {
        return 0, err
    }

    return rowsDeleted, nil
}
