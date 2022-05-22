package greedy_algorithm
import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCutGoldBar(t *testing.T) {
	arr1 := []int{20, 30, 10}
	arr2 := []int{10, 40, 60}
	fmt.Println(cutGoldBar(arr1))
	fmt.Println(cutGoldBar(arr2))
}

func TestGetProgramProfits(t *testing.T) {
	costs := []int{5, 6, 9, 12}
	profit := []int{2, 3, 5, 11}
	k := 2
	m := 6
	sum := getProgramProfits(costs, profit, k, m)
	fmt.Println(sum)
}

func TestGetMedianNum(t *testing.T) {
	arr := generateRandomNArr(20, 10)
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		insert(arr[i])
		medianNum := getMedianNum()
		fmt.Println(medianNum)
	}
}

func TestLowestString(t *testing.T) {
	arr1 := []string{"jibw", "ji", "jp", "bw", "jibw"}
	fmt.Println(lowestString(arr1))
}

func TestMostPreachArrange(t *testing.T) {
	programs := make([]*preachInfo, 5)
	programs[0] = &preachInfo{8, 10}
	programs[1] = &preachInfo{9, 11}
	programs[2] = &preachInfo{10, 14}
	programs[3] = &preachInfo{11, 13}
	programs[4] = &preachInfo{13, 15}
	fmt.Println(mostPreachArrange(programs, 7))
	preachs := make([]*preachInfo, 3)
	preachs[0] = &preachInfo{1, 27}
	preachs[1] = &preachInfo{26, 31}
	preachs[2] = &preachInfo{29, 300}
	fmt.Println(mostPreachArrange(preachs, 1))
}

func generateRandomNArr(maxValue, maxLength int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength)
	arr := make([]int, length)
	for index := 0; index < length; index++ {
		arr[index] = rand.Intn(maxValue)
	}
	return arr
}