package main

import (
	"fmt"
)

// HSL represents the Hue, Saturation, and Lightness values of a color
type HSL struct {
	Hue        int
	Saturation int
	Lightness  int
}

// GenerateSimilarColors generates a map of HSL colors similar to the provided color
func GenerateSimilarColors(color HSL, numColors int, hueRange, saturationBase, lightnessBase int) map[int]string {
	const maxColors = 11 // Maximum number of colors allowed
    const defaultSaturation = 0
    const defaultLightness = 0
	if numColors > maxColors {
		numColors = maxColors
	}
	similarColors := make(map[int]string)

	// Fibonacci sequence generation
	fib := fibonacci(numColors)

	// Generate similar colors by adjusting the HSL components within the specified ranges
	for i := 0; i < numColors; i++ {
		key := i * 100
		if key == 0 {
			key = 50
		} else if key == 1000 {
			key = 950
		}

		// Adjusting saturation and lightness based on Fibonacci sequence
		saturation := saturationBase * fib[i]
		lightness := lightnessBase * fib[i]

		newHue := clamp(color.Hue+(i-numColors/2)*hueRange, 0, 360)
		newSaturation := clamp(color.Saturation+saturation, 0, 100)
		newLightness := clamp(color.Lightness+lightness, 0, 100)

		similarColors[key] = fmt.Sprintf("hsl(%d, %d%%, %d%%)", newHue, newSaturation, newLightness)
	}

	return similarColors
}

// Fibonacci sequence generator
func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0], fib[1] = 1, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// clamp clamps the value between min and max
func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func main() {
	// Example usage
	color := HSL{180, 50, 70}
	numColors := 11
	hueRange := 30
	saturationBase := 10 // Base saturation change
	lightnessBase := 10  // Base lightness change

	similarColors := GenerateSimilarColors(color, numColors, hueRange, saturationBase, lightnessBase)
	fmt.Println("Similar Colors:")
	for key, value := range similarColors {
		fmt.Printf("%d: \"%s\",\n", key, value)
	}
}
