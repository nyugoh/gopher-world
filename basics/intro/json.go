package intro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Weight float64 `json:"weight"`
	Address Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	Post int `json:"post"`
	PostalCode string `json:"postalCode"`
}

func parsingJson() {
	usersFile, err:= os.Open("users.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	defer usersFile.Close()
	bs, err := ioutil.ReadAll(usersFile)
	if err != nil {
		fmt.Printf("Error")
	}

	/*type Users map[string]interface{}
	var users Users
	json.Unmarshal(bs, &users)
	fmt.Println(string(bs[:]))
	*/

	var users Users

	json.Unmarshal(bs, &users)
	for i:=0;i< len(users.Users); i++ {
		user := users.Users[i]
		fmt.Printf("%s is %d yrs old, weighs %.2f and lives at %s, %d-%s\n",
			user.Name, user.Age, user.Weight, user.Address.Street, user.Address.Post, user.Address.PostalCode)
	}

}
