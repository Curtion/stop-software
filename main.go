package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	config, err := os.Open("./config.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer config.Close()
	byteValue, _ := ioutil.ReadAll(config)
	arrayStr := strings.Split(string(byteValue), "\r\n")

	for _, v := range arrayStr {
		c := exec.Command("taskkill.exe", "/f", "/im", v)
		err = c.Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}
