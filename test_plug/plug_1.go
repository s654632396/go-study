/**
* @file plug_1.go
* @brief 这是一个插件入口文件
* @author DongChaofeng <s654632396@hotmail.com>
* @version 0.0.1
* @date 2021-07-14
 */
package main

import "fmt"

/**
* @brief TestPlug 插件的测试函数,通过加载该源码编译后的动态so文件来调用
*
* @param string 参数
*
* @return
 */
func TestPlug(test string) {
	fmt.Println("Hello, World, this is plug 1!")
	fmt.Println("ok, dump test parameter: " + test)
}
