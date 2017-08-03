package piCamera

//ProfileType sets the H264 profile to be used for the encoding.
type ProfileType int

const (
	//ProfileNone tells this package to use whatever the default is.
	ProfileNone ProfileType = iota
	//ProfileBaseline is for the baseline profile
	ProfileBaseline
	//ProfileMain is for the main profile
	ProfileMain
	//ProfileHigh is for a high profile
	ProfileHigh
)

//Convert takes the type and returns the string representation of that value.
//Returns true as well if it is the default value.
func (t ProfileType) Convert() (string, bool) {
	switch t {
	case ProfileBaseline:
		return "baseline", false
	case ProfileMain:
		return "main", false
	case ProfileHigh:
		return "high", false
	case ProfileNone:
		fallthrough
	default:
		return "", true
	}
}
