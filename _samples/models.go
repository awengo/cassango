package main

import (
	"log"

	"github.com/awengo/cassango"
)

type Users struct {
	Email string `cassango:"column:email;primary_key"`
	Name  string
}

func (Users) TableName() string {
	return "users"
}

func main() {
	c := cassango.Cluster{
		Hosts:    []string{"cassandra"},
		Keyspace: "test",
	}

	db, err := c.Open()
	if err != nil {
		log.Fatal(err)
	}

	db.Debug = true

	u := &Users{
		Email: "1@1.com",
		Name:  "test",
	}
	//us := &[]Users{}
	//if err := db.DropTable(u).CreateTable(u).Error; err != nil {
	//	log.Fatal(err)
	//}

	db.Create(u)

	if err := db.Excute(`select * from users;`).Error; err != nil {
		log.Fatal(err)
	}
}
