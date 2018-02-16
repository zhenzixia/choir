package main

import (
	"fmt"
	"runtime"
	"time"

	"choir/pkg/choir"
)

func main() {
	player, err := choir.NewMySimplePlayer("")
	if err != nil {
		panic(err)
	}
	//player.Play()

	srcChl := make(chan choir.Source)
	stopChl := make(chan error)
	go player.Run(srcChl, stopChl)

	fmt.Println("Volume: ", player.Volume())
	time.Sleep(time.Second * 5)

	fmt.Println("Setting Volume to: 0.2")
	player.SetVolume(0.2)

	fmt.Println("Volume: ", player.Volume())


	//input a source
	source := choir.Source{
		Src: "/Users/zixiazhen/work/src/choir/pkg/choir/demos/assets/birds.wav",
	}
	srcChl <- source

	//player.Current()
	if runtime.GOARCH != "js" {
		for {
			time.Sleep(time.Hour)
		}
	}
}
