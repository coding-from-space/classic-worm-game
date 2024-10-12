package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
)

const WINDOW_WIDTH = 800
const WINDOW_HIGHT = 600
const BLOCK_SIZE = 20

var blue rl.Color = rl.Color{34, 36, 54, 255}
var orange = rl.Color{255, 150, 108, 255}
var purple rl.Color = rl.Color{192, 153, 255, 255}

// worm variables
var worm []rl.Vector2
var wormLength = 1
var wormHead rl.Vector2 = rl.NewVector2(float32(WINDOW_WIDTH)/2, float32(WINDOW_HIGHT)/2)
var velocity rl.Vector2 = rl.Vector2{}

// let's make a function to get a random position in the screen
// we'll need it later
func getRandomPosition() rl.Vector2 {
	randomX := float64(rand.Intn(WINDOW_WIDTH - BLOCK_SIZE))
	randomY := float64(rand.Intn(WINDOW_HIGHT - BLOCK_SIZE))

	return rl.Vector2{
		X: float32(math.Round(randomX/20.0) * 20.0),
		Y: float32(math.Round(randomY/20.0) * 20.0),
	}
}

// let's make a function to restart the game
func restart() {
	wormHead = rl.NewVector2(float32(WINDOW_WIDTH)/2, float32(WINDOW_HIGHT)/2)
	velocity = rl.Vector2{}
	worm = []rl.Vector2{}
	wormLength = 1
}

func main() {
	defer rl.CloseWindow()
	rl.SetConfigFlags(rl.FlagBorderlessWindowedMode)
	rl.SetTargetFPS(10)
	rl.SetConfigFlags(rl.FlagWindowTopmost)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HIGHT, "worm")
	rl.SetWindowPosition(1000, 280)

	// first let's spawn some food
	pizza := getRandomPosition() // this worm eats pizza

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(blue)

		// let's make that little piece of s*** move
		if rl.IsKeyPressed(rl.KeyDown) {
			velocity = rl.NewVector2(0, BLOCK_SIZE)
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			velocity = rl.NewVector2(0, -BLOCK_SIZE)
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			velocity = rl.NewVector2(BLOCK_SIZE, 0)
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			velocity = rl.NewVector2(-BLOCK_SIZE, 0)
		}

		// how does the worm move?
		// not like a normal worm
		// it grows a new head every frame where it should go
		wormHead = rl.Vector2Add(wormHead, velocity)
		worm = append(worm, wormHead)

		// because it has a new head we need to cut out a piece of its tail
		// so that it has the same length
		if len(worm) > wormLength {
			worm = worm[1:]
			// don't worry, it doesn't feel pain
			// ...because I couldn't find a way to program that, I' not that good
		}

		// now let's check if the worm is drunk

		// check if the worm hits itself
		for _, segment := range worm[:wormLength-1] {
			if segment == wormHead {
				restart()
			}
		}

		// check if the worm hits a wall
		if wormHead.X >= float32(WINDOW_WIDTH) || wormHead.X < 0 ||
			wormHead.Y >= float32(WINDOW_HIGHT) || wormHead.Y < 0 {
			restart()
		}

		// draw the ugly worm
		for _, segment := range worm {
			rl.DrawRectangleV(segment, rl.NewVector2(BLOCK_SIZE, BLOCK_SIZE), purple)
		}

		// make the little bastard eat and grow, and add more food.
		if wormHead == pizza {
			// the worm's head is not a pizza, remember, these are vectors.
			wormLength += 1
			pizza = getRandomPosition()
		}

		// draw the piaasdf. I can't type
		rl.DrawRectangleV(
			pizza,
			rl.NewVector2(BLOCK_SIZE, BLOCK_SIZE),
			orange,
		)

		// that's it. enjoy the game
		// there could be bugs, ignore or fix them yourself
		// don't be lazy
		// like me

		rl.EndDrawing()
	}
}
