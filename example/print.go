package main

import (
	"fmt"

	"github.com/MingxuanGame/enkanetwork-go/api"
)

func main() {
	uids := []uint32{}
	done := make(chan bool)
	for _, uid := range uids {
		go func(uid uint32) {
			data, err := api.GetInfo(uid)
			if err != nil {
				fmt.Println(err)
				done <- true
				return
			}
			fmt.Printf("%+v\n", data)
			done <- true
		}(uid)
	}
	for n := len(uids); n > 0; n-- {
		<-done
	}
	close(done)
}
