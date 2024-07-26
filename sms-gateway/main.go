package main

import "sms-gateway/routers"

func main() {
	r := routers.InitRouter()
	r.Run()
}