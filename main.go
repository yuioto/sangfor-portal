package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkPortal(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Request unsuccessful.")
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading response failed.")
		return
	}

	if response.StatusCode == 200 && strings.Contains(string(body), "portal") {
		fmt.Println("Portal detection succeed.")
		//connect_network(user)
	} else {
		fmt.Println("Portal detection fail.")
		fmt.Println("Network is online.")
	}
}

/*
func connect_network(user User) {
	//post login request.
}
*/

func main() {
	check_url := "https://ping.archlinux.org/nm-check.txt"
	checkPortal(check_url)
}
