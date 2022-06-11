package main

import (
	"errors"
	"fmt"
)

//宠物类
type pet struct {
	petType string
}

func (p *pet) getPetType() string {
	return p.petType
}

func (p *pet) setPetType(petType string) {
	p.petType = petType
}

//封装一个包含宠物和当前个数的类
type enterQueuePet struct{
	animals pet
	count int
}

func (e *enterQueuePet) getPet() pet {
	return e.animals
}

func (e *enterQueuePet) getCount() int {
	return e.count
}

func (e *enterQueuePet) getQueuePetType() string {
	return e.animals.getPetType()
}

//猫狗队列
type catDogQueue struct {
	catQueue []enterQueuePet
	DogQueue []enterQueuePet
	count int
}

func (c *catDogQueue) add(pets pet) error {
	if "dog" == pets.getPetType() {
		c.count++
		c.DogQueue = append(c.DogQueue, enterQueuePet{pets, c.count})
	} else if "cat" == pets.getPetType() {
		c.count++
		c.catQueue = append(c.catQueue, enterQueuePet{pets, c.count})
	} else {
		return errors.New("the pet is not dog or cat")
	}
	return nil
}

func (c *catDogQueue) pollAll() {
	if len(c.catQueue) > 0 && len(c.DogQueue) > 0 {
		if c.catQueue[0].getCount() < c.DogQueue[0].getCount() {
			c.catQueue = c.catQueue[1:len(c.catQueue)]
		} else {
			c.DogQueue = c.DogQueue[1:len(c.DogQueue)]
		}
	}else if len(c.catQueue) > 0 {
		c.catQueue = c.catQueue[1:len(c.catQueue)]
	}else if len(c.DogQueue) > 0 {
		c.DogQueue = c.DogQueue[1:len(c.DogQueue)]
	}
}

func (c *catDogQueue) pollDog() error {
	if len(c.DogQueue) > 0 {
		c.DogQueue = c.DogQueue[1:len(c.DogQueue)]
	} else {
		return errors.New("dog queue not exist")
	}
	return nil
}

func (c *catDogQueue) pollCat() error {
	if len(c.catQueue) > 1 {
		c.catQueue = c.catQueue[1:len(c.catQueue)]
	} else if len(c.catQueue) == 1 {
		c.catQueue = []enterQueuePet{}
	}else {
		return errors.New("cat queue not exist")
	}
	return nil
}

func (c *catDogQueue) isEmpty() bool {
	if len(c.catQueue) == 0 && len(c.DogQueue) == 0 {
		return true
	}
	return false
}

func (c *catDogQueue) isCatEmpty() bool {
	return len(c.catQueue) == 0
}
func(c *catDogQueue) isDogEmpty() bool {
	return len(c.DogQueue) == 0
}

func main() {
	var catDogQueue catDogQueue
	catDogQueue.add(pet{"dog"})
	catDogQueue.add(pet{"dog"})
	catDogQueue.add(pet{"cat"})
	catDogQueue.add(pet{"cat"})
	catDogQueue.add(pet{"dog"})
	fmt.Println(catDogQueue)
	catDogQueue.pollAll()
	catDogQueue.pollCat()
	fmt.Println(catDogQueue)
	catDogQueue.pollAll()
	fmt.Println(catDogQueue)
	catDogQueue.pollDog()
	fmt.Println(catDogQueue)
	catDogQueue.pollDog()
}