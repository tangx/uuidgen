package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type uuidConfig struct {
	lower bool
	upper bool
	num   int
}

var config uuidConfig

func init() {
	flag.BoolVar(&config.lower, "lower", true, "Generate a lower case UUID (default)")
	flag.BoolVar(&config.upper, "upper", false, "Generate a upper case UUID ")
	flag.IntVar(&config.num, "n", 1, "Set UUID number")
}

func main() {
	flag.Parse()

	// multi gen
	for i := 0; i < config.num; i++ {
		id := UUIDGen()
		func() {
			if config.upper {
				fmt.Println(strings.ToUpper(id))
				return
			}
			// if config.lower {
			fmt.Println(strings.ToLower(id))
			return
			// }
		}()
	}

}

// UUIDGen return a uuid
func UUIDGen() string {
	id := uuid.Must(uuid.NewRandom())
	return id.String()
}
