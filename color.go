package main

import "strings"

var colorPalette = []string{
	`"#ff1744"`, `"#D500F9"`, `"#2979FF"`, `"#00E676"`, `"#FF9100"`, `"#8bc34a"`, `"#ffeb3b"`, `"#ff9800"`,
	`"#795548"`, `"#607d8b"`, `"#e91e63"`, `"#673ab7"`, `"#2196f3"`, `"#00bcd4"`, `"#4caf50"`, `"#cddc39"`,
	`"#ffc107"`, `"#ff5722"`, `"#9e9e9e"`, `"#9e9e9d"`}

var colorI = 0

// Reset restarts the palette iterator. Following Reset(), invoking Next() returns the first color in the palette.
func colorReset() {
	colorI = 0
}

// Next iterates through the color palette.
func colorNext() string {
	result := colorPalette[colorI]
	colorI++
	if colorI > len(colorPalette)-1 {
		colorI = 0
	}

	return result
}

func colorIndex(i int) string {
	return colorPalette[i%len(colorPalette)]
}

// FirstN returns a comma-separated string of the first n colors in the palette.
func colorFirstN(n int) string {
	k := 0
	var cs []string
	for j := 0; j < n; j++ {
		cs = append(cs, colorPalette[k])
		k++
		if k > len(colorPalette)-1 {
			k = 0
		}
	}
	return strings.Join(cs, ",")
}

func colorRepeat(i, n int) string {
	return strings.Repeat(colorIndex(i)+",", n)
}
