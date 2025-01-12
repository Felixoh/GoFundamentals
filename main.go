package main

import "fmt"

func main() {
	//mes := [4]string{"kurt", "janis", "amy", "mochamas"}

	//iterate over the array
	//"for i := 0; i < len(names); i++ {
	//fmt.Println(names[i])
	//}

	//use of the range keyword:
	//

	//use of map and initialisations
	var users map[string]string

	//initialising the array

	users = make(map[string]string)

	users["kurt@gmail.com"] = "mochama"
	users["devfeloh1@gmail.com"] = "janisa"
	users["jamii@2024"] = "jafrice"

	for key, value := range users {
		fmt.Printf("%s %s'\n", key, value)
	}

}
