package problems

import (
	"fmt"
	"github.com/wladimir/go-cake/server/problems/common"
	"log"
)

func Run() {
	max, err := common.Max(2, 3, 1)
	if err != nil {
		fmt.Printf("whoa %v", "what")
	} else {
		log.Printf("it's %v", max)
	}
}
