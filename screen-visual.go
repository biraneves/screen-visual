package screen_visual

import "fmt"

// Draw lines on screen with a line break at the end.
func Lineln(char string, size int) {

	line := ""
	for i := 0; i < size; i++ {
		line += char
	}
	fmt.Println(line)

}

// Draw lines on screen without a line break at the end
func Line(char string, size int) {

	line := ""
	for i := 0; i < size; i++ {
		line += char
	}
	fmt.Print(line)

}