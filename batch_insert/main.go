package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
	"time"
)

func main() {
	var dsn = "root:123456@tcp(127.0.0.1:3306)/test?checkConnLiveness=false&maxAllowedPacket=0"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db.Debug().Table("test_1").Where("id >= ?", 1).Delete(0)

	for i := 0; i < 100000; i++ {
		var test = &tblTest1{A: "", B: "", C: "", Data: ""}
		setData(test)
		if err := db.Table("test_1").Create(test).Error; err != nil {
			panic(err)
		}
		var nt tblTest1
		db.Model(&tblTest1{}).Find(&nt)
		log.Println(nt)
	}
}

func setData(tbl *tblTest1) {
	t := time.Now()
	tbl.A = strconv.Itoa(t.Hour())
	tbl.B = strconv.Itoa(t.Minute())
	tbl.C = strconv.Itoa(t.Nanosecond())
	tbl.Data = t.String()
}

// table: test_1
type tblTest1 struct {
	ID   uint
	A    string `gorm:"column:a"`
	B    string `gorm:"column:b"`
	C    string `gorm:"column:c"`
	Data string `gorm:"column:data"`
}

func (*tblTest1) TableName() string {
	return "test_1"
}
