package hash_map_algorithm

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRandomPool(t *testing.T) {
	Convey("test the random pool", t, func() {
		pool := newRandomPool()
		pool.insert("mao")
		pool.insert("wen")
		pool.insert("qing")
		fmt.Println(pool.getRandom())
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())
		pool.delete("wen")
		fmt.Println()
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())
		time.Sleep(1*time.Microsecond)
		fmt.Println(pool.getRandom())

	})
}
