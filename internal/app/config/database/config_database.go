package database

import "fmt"

type connectionData struct {
	Host     string
	Schema   string
	UserName string
	Password string
	Dialect  string
}

func GetConnectionLocal() connectionData {
	return connectionData{
		Host:     "localhost:3306",
		Schema:   "dulceria",
		UserName: "root",
		Password: "sandra123-",
		Dialect:  "mysql",
	}
}

func (c connectionData) GetUrl() (url string) {
	url = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.UserName, c.Password, c.Host, c.Schema)
	return
}
