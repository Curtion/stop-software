package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/antchfx/jsonquery"
)

func stop_software() {
	config, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer config.Close()
	doc, err := jsonquery.Parse(config)
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range jsonquery.Find(doc, "software/*") {
		fmt.Println(n.InnerText())
		c := exec.Command("taskkill.exe", "/f", "/im", n.InnerText())
		err = c.Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {

}
