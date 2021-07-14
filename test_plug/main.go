/**
* @file main.go
* @brief 主程序入口
* @author DongChaofeng <s654632396@hotmail.com>
* @version 0.0.1
* @date 2021-07-14
 */

package main

import (
	"fmt"
	"os"
	"plugin"
)

// Plugins 全局插件变量
var Plugins map[string]*plugin.Plugin

/**
* @brief init 程序初始化
*
* @return
 */
func init() {
	fmt.Println("main initialize..")
	fmt.Println("init Plugins Mapper...")
	// 初始化我们的全局插件变量mapper
	Plugins = make(map[string]*plugin.Plugin)
	// 调用加载插件的函数
	if err := loadPlugs(); err != nil {
		fmt.Errorf("[loading plugins] %s\n", err)
		os.Exit(1)
	}
}

/**
* @brief loadPlugs 加载默认so插件的函数
*
* @param
*
* @return error
 */
func loadPlugs() (err error) {
	// 从本地加载动态链接库, 注册到全局插件变量里
	Plugins["plug_1"], err = plugin.Open("./plug_1.so")
	if err != nil {
		return err
	}
	fmt.Println("plugs loaded.")
	return
}

func main() {

	// 使用Lookup函数,从 plug_1 这个so库中查找 TestPlug 这个symbol
	funcTest, err := Plugins["plug_1"].Lookup("TestPlug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 然后使用断言， 断言这个symbol的类型, 这里是断言的 type func(string)
	if PfuncTestPlug, ok := funcTest.(func(string)); ok {
		// 现在可以调用函数了
		PfuncTestPlug("abc")
	}
}
