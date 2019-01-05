package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Config struct {
	lower bool
	upper bool
	num   int
}

var config Config

func init() {
	flag.BoolVar(&config.lower, "lower", true, "Generate UUID by lower case (default) ")
	flag.BoolVar(&config.upper, "upper", false, "Generate UUID by upper case")
	flag.IntVar(&config.num, "n", 1, "How many uuid to be generated")
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
