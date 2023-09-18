package config

import (
	"fmt"

	"github.com/shplume/csv-import/util/config"
)

func main() {
	name := config.Getstring("name")
	fmt.Println("name:", name)
}
