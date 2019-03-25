package a

// single import
import "time"

// multiple import
import (
	"fmt"
	"log"
)

var now = time.Now()

func hoge(t time.Time) bool {
	/* multiline
	   comment */
	log.Println("hoge is called")
	fmt.Printf("It is %s now\n", t)
	return true
}

func main() {
	_ = hoge(now)
}
