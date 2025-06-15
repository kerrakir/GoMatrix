package main

import (
	"fmt"
	"math/rand"
	/*"os/exec"
	"runtime"*/
	"time"
)

const (
	width  = 80
	height = 40
	minLen = 6
	maxLen = 16
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%&")

/*func clearScreen() {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "cls").Run()
	} else {
		fmt.Print("\033[2J")
	}
}*/

func drawChar(x, y int, r rune, style string) {
	if y >= 0 && y < height {
		fmt.Printf("\033[%d;%dH%s%c", y, x, style, r)
	}
}

func clearChar(x, y int) {
	if y >= 0 && y < height {
		fmt.Printf("\033[%d;%dH ", y, x)
	}
}

func rainColumn(x int) {
	for {
		length := rand.Intn(maxLen-minLen) + minLen
		startY := rand.Intn(height/2) * -1
		speed := time.Duration(rand.Intn(20)+20) * time.Millisecond

		for y := startY; y < height+length; y++ {
			if y-length >= 0 {
				clearChar(x, y-length)
			}

			for i := 1; i < length; i++ {
				yy := y - i
				drawChar(x, yy, charset[rand.Intn(len(charset))], "\033[2;32m")
			}

			drawChar(x, y, charset[rand.Intn(len(charset))], "\033[1;37m")

			time.Sleep(speed)
		}

		time.Sleep(time.Duration(rand.Intn(1000)+200) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//clearScreen()
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	for x := 1; x < width; x += 2 {
		go rainColumn(x)
		time.Sleep(10 * time.Millisecond)
	}

	select {}
}
