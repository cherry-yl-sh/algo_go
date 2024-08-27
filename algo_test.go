package dylan_algo

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(input, target))
}
func TestThreeSum(t *testing.T) {
	//-4, -1, -1, 0, 1, 2
	input := []int{-1, 0, 1, 2, -1, -4}
	//input := []int{0, 1, 1}
	//input := []int{0, 0, 0}
	res := threeSum(input)
	fmt.Println(res)

}
func TestSubSetS(t *testing.T) {
	input := []int{1, 2, 3}
	res := subsets(input)
	fmt.Println(res)

}
func TestLRU(t *testing.T) {
	s := "abc"
	fmt.Printf("%c\n", s[1])
}

func TestCanJump(t *testing.T) {
	input := []int{0}
	res := canJump(input)
	fmt.Println(res)
}

func TestF(t *testing.T) {
	grid := [][]byte{{'1'}}
	dfs1(grid)
	fmt.Println(grid)
}
func dfs1(grid [][]byte) {

}
