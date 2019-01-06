package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type uuidConfig struct {
	upper bool
	num   int
}

var config uuidConfig
var wg sync.WaitGroup

func init() {
	flag.BoolVar(&config.upper, "upper", false, "Generate a UPPER case UUID, (defalut is LOWER) ")
	flag.IntVar(&config.num, "n", 1, "Set UUID number")
}

func main() {
	flag.Parse()

	// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/16.4.md
	// - 切片、映射和通道，使用make
	// - 数组、结构体和所有的值类型，使用new
	// var ch1 chan string
	ch1 := make(chan string)
	go printer(ch1)

	wg.Add(1)
	for i := 0; i < config.num; i++ {
		UUIDGen(ch1)
	}

	close(ch1)
	wg.Wait()
}

// UUIDGen return a uuid
func UUIDGen(ch chan string) {
	id := uuid.Must(uuid.NewRandom())
	s := id.String()
	if config.upper {
		s = strings.ToUpper(s)
	}

	ch <- s
}

func printer(ch chan string) {
	for s := range ch {
		fmt.Println(s)
	}
	wg.Done()
}
