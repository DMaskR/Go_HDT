package main

import "fmt"

func main() {
	data, err := loadAllData("IP-locations.rar", "./")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Ip[0])
}
