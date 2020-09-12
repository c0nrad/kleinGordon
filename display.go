package main

import (
	"fmt"
	"math"
	"os"
)

var (
	ColorBlack = "\033[30m"
	ColorWhite = "\033[37m"

	ColorRed     = "\033[31m" // 625 - 740
	ColorYellow  = "\033[33m" // 565 - 590
	ColorGreen   = "\033[32m" // 520 - 565
	ColorCyan    = "\033[36m" // 500 - 520
	ColorBlue    = "\033[34m" // 435 - 500
	ColorMagenta = "\033[35m" // 380 - 435

	ColorReset = "\033[0m"

	Clear = "\033[2J"

	Corner         = "+"
	HorizontalLine = "-"
	VerticalLine   = "|"
	Box            = "▉"

	SpectralWidth        = 25
	SpectralHeightOffset = 1
	SpectralWidthOffset  = 5
	MaxLevels            = 6
)

func DrawSpring(s *Spring, X, Y int, scale float64) {
	Jump(X, Y)
	fmt.Print("-")
	length := int(math.Abs(math.Round(s.X * scale)))
	dir := s.X / math.Abs(s.X)
	for x := 1; x < length; x++ {
		Jump(X, Y-(int(dir)*x))
		fmt.Print("|")
	}
	// for x := int(math.Round(s.X * scale)); x < 0; x++ {
	// 	Jump(X, Y-(x))
	// 	fmt.Print("|")
	// }
	Jump(X, Y-int(math.Round(s.X*scale)))
	fmt.Print("*")
}

func DrawSpringSystem(system *SpringSystem, scale float64) {
	fmt.Println(Clear)
	Y := 10
	for i, s := range system.Springs {
		DrawSpring(s, i+2, Y, scale)
	}
	DrawStats(system)
	Jump(1, Y+int(scale)+5)
}

func DrawStats(system *SpringSystem) {
	Jump(2, 18)
	fmt.Print("Connected Mass-Springs Simulation")
	Jump(3, 19)
	fmt.Printf("Count       = %d", len(system.Springs))
	Jump(3, 20)
	fmt.Printf("Spring_K    = %3.1f", system.Springs[0].K)
	Jump(3, 21)
	fmt.Printf("Spring_M    = %3.1f", system.Springs[0].M)
	Jump(3, 22)
	fmt.Printf("Band_J      = %3.1f", system.J)

	Jump(25, 19)
	fmt.Printf("Kinetic Energy =   %3.1f", system.KE())
	Jump(25, 20)
	fmt.Printf("Potential Energy = %3.1f", system.PE())
	Jump(25, 22)
	fmt.Print("F_i = (-kx) + (j*(xprev-x)) + (j*(xnext-x))")

	// fmt.Printf("Total Energy =     %2.1f", system.KE()+system.PE())

	// fmt.Print("Each spring is connected to the spring next to it by factor J")

	//F_i = (-kx) + (j*(xprev-x)) + (j*(xnext-x))
}

func Jump(x, y int) {
	os.Stdout.WriteString(fmt.Sprintf("\033[%d;%df", y, x))
}
