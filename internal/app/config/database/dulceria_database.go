package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Client struct{}

func (c Client) Connect() (db *gorm.DB, err error) {
    data := GetConnectionLocal()
    connectionString := data.GetUrl()

    db, err = gorm.Open(
        mysql.Open(connectionString),
        &gorm.Config{},
    )

    if err != nil {
        panic(err.Error())
    }

    return db, nil
}