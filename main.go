package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/antchfx/jsonquery"
	session_notifications "github.com/brunoqc/go-windows-session-notifications"
)

func stop_software(doc *jsonquery.Node) {
	log.Println("屏幕锁定，开始关闭程序")
	for _, n := range jsonquery.Find(doc, "software/*") {
		c := exec.Command("taskkill.exe", "/f", "/im", n.InnerText())
		err := c.Start()
		if err != nil {
			log.Panicln(err)
		}
		log.Println(n.InnerText(), "关闭成功")
	}
}

func check_time(doc *jsonquery.Node) bool {
	timeValue := jsonquery.Find(doc, "time/*")
	timeNowValue := fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute())
	return timeNowValue > timeValue[0].InnerText() && timeNowValue < timeValue[1].InnerText()
}

func main() {
	config, err := os.Open("./config.json")
	if err != nil {
		log.Panicln(err)
	}
	defer config.Close()
	doc, err := jsonquery.Parse(config)
	if err != nil {
		log.Panicln(err)
	}

	quit := make(chan int)
	chanMessages := make(chan session_notifications.Message, 100)
	chanClose := make(chan int)

	go func() {
		for {
			select {
			case m := <-chanMessages:
				switch m.UMsg {
				case session_notifications.WM_WTSSESSION_CHANGE:
					switch m.Param {
					case session_notifications.WTS_SESSION_LOCK:
						status := check_time(doc)
						if !status {
							stop_software(doc)
						} else {
							log.Println("屏幕锁定，但时间不满足")
						}
					case session_notifications.WTS_SESSION_UNLOCK:
						log.Println("屏幕解锁")
					}
				case session_notifications.WM_QUERYENDSESSION:
					log.Println("注销或关机")
				}
				close(m.ChanOk)
			}
		}
	}()

	session_notifications.Subscribe(chanMessages, chanClose)
	log.Println("程序已初始化")
	<-quit
}
