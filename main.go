package main

import (
	"fmt"
	"vk_old/fourth"
	//for init only
	//_ "vk_old/second"
)

func init() {
	fmt.Println("init main")
}

func main() {
	fourth.HttpMyExampleRace()
}
