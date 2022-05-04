package main

import (
	"fmt"
	"math/rand"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var ant = "🐜"
var leaf = "🍃"
var empty_space = "🟫"

var ant_count = 4
var leaf_count = 5

var moves = [4]string{"l", "r", "u", "d"}

var mapa = [10][10]string{
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space},
	{empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space, empty_space}}

var ant_positions = [][]int{}

var leaf_positions = [][]int{}

func spawn_ants(ant_count int) {
	for i := 0; i < ant_count; i++ {
		var x = r1.Intn(9)
		var y = r1.Intn(9)
		mapa[x][y] = ant
		var position = []int{x, y}
		ant_positions = append(ant_positions, position)
	}
}

func spawn_leaves(leaf_count int) {
	for i := 0; i < leaf_count; i++ {
		var x = r1.Intn(9)
		var y = r1.Intn(9)
		mapa[x][y] = leaf
		var position = []int{x, y}
		leaf_positions = append(leaf_positions, position)
	}
}

func print(to_print [10][10]string) {
	for i := 0; i < len(to_print); i++ {
		fmt.Println(to_print[i])
	}
}

func change_ant_position(positions [][]int) {
	for i := 0; i < len(positions); i++ {
		mapa[positions[i][0]][positions[i][1]] = empty_space
		var move = moves[r1.Intn(4)]
		switch move {
		case "l":
			if positions[i][0] > 0 && mapa[positions[i][0]-1][positions[i][1]] != ant {
				positions[i][0] = positions[i][0] - 1
			}
		case "r":
			if positions[i][0] < 9 && mapa[positions[i][0]+1][positions[i][1]] != ant {
				positions[i][0] = positions[i][0] + 1
			}
		case "u":
			if positions[i][1] > 0 && mapa[positions[i][0]][positions[i][1]-1] != ant {
				positions[i][1] = positions[i][1] - 1
			}
		case "d":
			if positions[i][1] < 9 && mapa[positions[i][0]][positions[i][1]+1] != ant {
				positions[i][1] = positions[i][1] + 1
			}
		}
	}
	for i := 0; i < len(positions); i++ {
		if mapa[positions[i][0]][positions[i][1]] == leaf {
			if positions[i][0] < 9 && positions[i][1] < 9 {
				mapa[positions[i][0]+1][positions[i][1]+1] = leaf
			}
			if positions[i][0] > 0 && positions[i][1] > 0 {
				mapa[positions[i][0]-1][positions[i][1]-1] = leaf
			}
		}
		mapa[positions[i][0]][positions[i][1]] = ant
	}
}
func main() {
	spawn_ants(ant_count)
	spawn_leaves(leaf_count)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range ticker.C {
			fmt.Println()
			print(mapa)
			change_ant_position(ant_positions)
			fmt.Println()
		}
	}()
	time.Sleep(1600 * time.Second)
	ticker.Stop()
}
