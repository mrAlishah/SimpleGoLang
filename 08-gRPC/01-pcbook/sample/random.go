package sample

import (
	"math/rand"
	"pcbook/pb"
	"time"

	"github.com/google/uuid"
)

func init() {
	// https://golang.cafe/blog/golang-random-number-generator.html
	rand.Seed(time.Now().UnixNano())
}

//---------------------------------------  Random Boolean
func randomBool() bool {
	return rand.Intn(2) == 1
}

//---------------------------------------  Random enum
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

//---------------------------------------  Random String from set
func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

//---------------------------------------  Random google.UUID
func randomID() string {
	return uuid.New().String()
}

//---------------------------------------  Random int between min and max
func randomInt(min, max int) int {
	//rand.Intn() function will return an integer from 0 to max - min. So if we add min to it, we will get a value from min to max
	return min + rand.Int()%(max-min+1)
	//rand.Intn(max - min + 1) + min
}

//---------------------------------------  Random float64 between min and max
func randomFloat64(min, max float64) float64 {
	//rand.Float64() function will return a random float between 0 and 1. So we will multiply it with (max - min) to get a value between 0 and max - min
	return min + rand.Float64()*(max-min)
}

//---------------------------------------  Random float64 between min and max
func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

//---------------------------------------
func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

//---------------------------------------
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

//---------------------------------------
func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

//---------------------------------------
func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
