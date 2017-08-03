package piCamera

//DRCType is to set the dynamic range compression
type DRCType int

const (
	//DRCNone is so this package can use whatever the default is
	DRCNone DRCType = iota
	//DRCOff turns of DRC
	DRCOff
	//DRCLow compress the range slightly
	DRCLow
	//DRCMedium compress the range more
	DRCMedium
	//DRCHigh compress the range even more
	DRCHigh
)

//Convert takes the type and returns the string representation of that value.
//Returns true as well if it is the default value.
func (t DRCType) Convert() (string, bool) {
	switch t {
	case DRCOff:
		return "off", false
	case DRCLow:
		return "low", false
	case DRCMedium:
		return "medium", false
	case DRCHigh:
		return "high", false
	case DRCNone:
		fallthrough
	default:
		return "", true
	}
}
