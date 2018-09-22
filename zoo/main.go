package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName()) /* AppNameはmainパッケージに定義された関数のためパッケージ指定不要*/

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
