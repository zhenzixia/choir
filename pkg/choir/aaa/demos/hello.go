package main

import (
	"fmt"
	"runtime"
	"time"

	"engo.io/audio"
)

func main() {
	player, err := audio.NewSimplePlayer("/Users/zixiazhen/work/src/audio/demos/assets/birds.wav")
	if err != nil {
		panic(err)
	}
	player.Play()

	fmt.Println("Volume: ", player.Volume())
	time.Sleep(time.Second * 5)

	fmt.Println("Setting Volume to: 0.2")
	player.SetVolume(0.2)

	fmt.Println("Volume: ", player.Volume())

	player.Current()
	if runtime.GOARCH != "js" {
		for {
			time.Sleep(time.Hour)
		}
	}
}
