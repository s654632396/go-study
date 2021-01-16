package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver package
	"log"
	"sync"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/kana?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
}

type Inventory struct {
	ID  int64 `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Num int64 `gorm:"column:num;type:int"`
}

func main() {
	db.Debug()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			updateNumByID(1)
		}()
	}
	wg.Wait()

	fmt.Println("done.")
}

func updateNumByID(id int) {
	// using transaction and tx's session to update
	tx := db.Begin()
	defer func() {
		if rcv := recover(); rcv != nil {
			tx.Rollback()
			log.Println("err with rollback:", rcv)
		}
	}()

	var n1 Inventory
	if err := tx.Debug().Set("gorm:query_option", "FOR UPDATE").Find(&n1, id).Error; err != nil {
		tx.Rollback()
		log.Println("query with for update opt err:", err)
		return
	}

	n1.Num += 100
	tx.Debug().Save(&n1)

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}
	return
}
