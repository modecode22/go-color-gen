package main

import (
    "fmt"
    "math"
)

// HSLtoRGB converts HSL to RGB
func HSLtoRGB(h, s, l float64) (float64, float64, float64) {
 
	h = math.Mod(h, 360)
    s = math.Max(0, math.Min(100, s)) / 100.0
    l = math.Max(0, math.Min(100, l)) / 100.0

    c := (1 - math.Abs(2*l-1)) * s
    x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
    m := l - c/2.0

    r, g, b := 0.0, 0.0, 0.0

    if h >= 0 && h < 60 {
        r, g, b = c, x, 0
    } else if h >= 60 && h < 120 {
        r, g, b = x, c, 0
    } else if h >= 120 && h < 180 {
        r, g, b = 0, c, x
    } else if h >= 180 && h < 240 {
        r, g, b = 0, x, c
    } else if h >= 240 && h < 300 {
        r, g, b = x, 0, c
    } else if h >= 300 && h < 360 {
        r, g, b = c, 0, x
    }

    r = (r + m) * 255
    g = (g + m) * 255
    b = (b + m) * 255

    return r, g, b
}

// FormatHSLString formats HSL string
func FormatHSLString(h, s, l float64) string {
    return fmt.Sprintf("hsl(%.0f, %.0f%%, %.0f%%)", h, s, l)
}

// GenerateShades generates HSL shades
func GenerateShades(h, s, l float64) map[int]string {
    // Initialize a map called shades using make
    shades := make(map[int]string)

    // Generate shades for different lightness levels
    for i := 50; i <= 950; i += 50 {
        // Adjust the lightness level
        l += float64(i) / 1000.0
        // Ensure that the lightness level is within the range [0, 1]
        if l < 0 {
            l = 0
        } else if l > 1 {
            l = 1
        }
        // Format the HSL (Hue, Saturation, Lightness) string and store it in the shades map
        shades[i] = FormatHSLString(h, s, l*100)
    }
    // Return the generated shades map
    return shades
}

func main() {
    h := 228.0
    s := 8.0
    l := 12.0
    shades := GenerateShades(h, s, l)
    for key, value := range shades {
        fmt.Printf("%d: %s\n", key, value)
    }
}