package piCamera

//MeteringType is used for setting the Metering Mode.
type MeteringType int

const (
	//MeterNone tell this package to use whatever the default is
	MeterNone MeteringType = iota
	//MeterAverage average the whole frame for metering
	MeterAverage
	//MeterSpot use spot metering
	MeterSpot
	//MeterBacklit will assume a backlit image
	MeterBacklit
	//MeterMatrix use matrix metering
	MeterMatrix
)

//Convert takes the type and returns the string representation of that value.
//Returns true as well if it is the default value.
func (t MeteringType) Convert() (string, bool) {
	switch t {
	case MeterAverage:
		return "average", false
	case MeterSpot:
		return "spot", false
	case MeterBacklit:
		return "backlit", false
	case MeterMatrix:
		return "matrix", false
	case MeterNone:
		fallthrough
	default:
		return "", true
	}
}
