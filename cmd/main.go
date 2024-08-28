package main

import (
	"github.com/golangrustnode/gotrimreader/trimreader"
	"github.com/golangrustnode/log"
)

func main() {
	data, err := trimreader.ReadAsString("/sys/class/block/sdb/queue/hw_sector_size")
	if err != nil {
		log.Fatal(err)
	}
	log.Info(data)
	udata, err := trimreader.ReadAsUint64("/sys/class/block/sdb/queue/hw_sector_size")
	if err != nil {
		log.Fatal(err)
	}
	log.Info(udata)
}
