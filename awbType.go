package piCamera

import (
	"fmt"
)

//AWBType is for setting the Automatic White Balance setting.
type AWBType int

const (
	//AwbAuto is for automatic white balance
	AwbAuto AWBType = iota
	//AwbOff is for turning automatic white balance off
	AwbOff
	//AwbSun is for sunny mode
	AwbSun
	//AwbCloud is for cloudy mode
	AwbCloud
	//AwbShade is for shade mode
	AwbShade
	//AwbTungsten tungsten lighting mode
	AwbTungsten
	//AwbFluorescent fluorescent lighting mode
	AwbFluorescent
	//AwbIncandescent incandescent lighting mode
	AwbIncandescent
	//AwbFlash flash mode
	AwbFlash
	//AwbHorizon horizon mode
	AwbHorizon
)

//Convert takes the type and returns the string representation of that value.
//Returns true as well if it is the default value.
func (t AWBType) Convert() (string, bool) {
	switch t {
	case AwbAuto:
		return "auto", true
	case AwbOff:
		return "off", false
	case AwbSun:
		return "sun", false
	case AwbCloud:
		return "cloud", false
	case AwbShade:
		return "shade", false
	case AwbTungsten:
		return "tungsten", false
	case AwbFluorescent:
		return "fluorescent", false
	case AwbIncandescent:
		return "incandescent", false
	case AwbFlash:
		return "flash", false
	case AwbHorizon:
		return "horizon", false
	default:
		return "", true
	}
}

//AWBGains sets the blue and red gains to be applied when AWBOff is set.
type AWBGains struct {
	b float32
	r float32
}

//NewAWBGains creates gains to set to Red and Blue.
//Values are multiplied to the values so 1.5 is 150% and .5 is 50%.
func NewAWBGains(blue, red float32) *AWBGains {
	return &AWBGains{
		b: blue,
		r: red,
	}
}

//Convert takes the type and returns the string representation of that value.
func (t *AWBGains) Convert() string {
	return fmt.Sprintf("%.2f,%.2f", t.b, t.r)
}
