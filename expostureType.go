package piCamera

//ExposureType is for setting the exposure mode.
type ExposureType int

const (
	//ExpNone is used to tell this package to use whatever the default is.
	ExpNone ExposureType = iota
	//ExpAuto use automatic exposure mode
	ExpAuto
	//ExpNight select setting for night shooting
	ExpNight
	//ExpBacklight select setting for backlit subject
	ExpBacklight
	//ExpSpotlight select setting for spotlit subject
	ExpSpotlight
	//ExpSports select setting for sports
	ExpSports
	//ExpSnow select setting optimised for snowy scenery
	ExpSnow
	//ExpBeach select setting optimised for beach
	ExpBeach
	//ExpVerylong select setting for long exposures
	ExpVerylong
	//ExpFixedfps constrain fps to a fixed value
	ExpFixedfps
	//ExpAntishake turns on antishake mode
	ExpAntishake
	//ExpFireworks select setting optimised for fireworks
	ExpFireworks
)

//Convert takes the type and returns the string representation of that value.
//Returns true as well if it is the default value.
func (t ExposureType) Convert() (string, bool) {
	switch t {
	case ExpAuto:
		return "auto", false
	case ExpNight:
		return "night", false
	case ExpBacklight:
		return "backlight", false
	case ExpSpotlight:
		return "spotlight", false
	case ExpSports:
		return "sports", false
	case ExpSnow:
		return "snow", false
	case ExpBeach:
		return "beach", false
	case ExpVerylong:
		return "verylong", false
	case ExpFixedfps:
		return "fixedfps", false
	case ExpAntishake:
		return "antishake", false
	case ExpFireworks:
		return "fireworks", false
	case ExpNone:
		fallthrough
	default:
		return "", true
	}
}
