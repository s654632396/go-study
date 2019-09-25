package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影

	// test := Test{Name: "梦开始到地方.."}
	// db.Create(&test)

	var user Test
	db.Find(&user, 2)
	fmt.Println(user)

	var users []Test
	db.Where("name = ?", "poi").Find(&users)
	fmt.Println(users)

	db.FirstOrCreate(&user, Test{Name: "23333"})
	fmt.Println(user)

}

type Test struct {
	gorm.Model
	Id   int64  `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"size:255"`
}

func (Test) TableName() string {
	return "test"
}
