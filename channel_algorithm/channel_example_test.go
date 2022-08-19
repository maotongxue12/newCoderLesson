package channel_algorithm

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConnectRequestWeb(t *testing.T) {
	Convey("test the connectRequestWeb function", t, func() {
		connectRequestWeb(100, 1000, "www.baidu.com")
	})
}

func TestPrintTotal(t *testing.T) {
	Convey("test print cat dog fish ...", t, func() {
		printTotal(100)
	})
}
