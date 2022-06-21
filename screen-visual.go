package screen_visual

import "fmt"

// Draw lines on screen
func Line(char string, size int) {

	line := ""
	for i := 0; i < size; i++ {
		line += char
	}
	fmt.Println(line)

}
