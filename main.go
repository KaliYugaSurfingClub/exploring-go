package main

import (
	"fmt"
	"vk_old/second"
	//for init only
	//_ "vk_old/second"
)

func init() {
	fmt.Println("init main")
}

func main() {
	second.UseEntities()
}
