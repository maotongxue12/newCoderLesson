package stackqueue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPetType(t *testing.T) {
	Convey("test the catDogQueue function", t, func() {
		cdQue := new(catDotQueue)
		cdQue.add(pet{petType: "dog"})
		p, err := cdQue.pollDog()
		So(p, ShouldResemble, pet{petType: "dog", count: 0})
		So(err, ShouldBeNil)
		cdQue.add(pet{petType: "cat"})
		cdQue.add(pet{petType: "dog"})
		cdQue.add(pet{petType: "dog"})
		p, err = cdQue.pollAll()
		So(p, ShouldResemble, pet{petType: "cat", count: 1})
		p, err = cdQue.pollAll()
		So(p, ShouldResemble, pet{petType: "dog", count: 2})
		p, err = cdQue.pollCat()
		So(err, ShouldNotBeNil)
		p, err = cdQue.pollDog()
		So(err, ShouldBeNil)
	})
}
