package piCamera

//ImgEffectType is used for setting the image effect to use.
type ImgEffectType int

const (
	//ImfxNone no effect
	ImfxNone ImgEffectType = iota
	//ImfxNegative invert the image colours
	ImfxNegative
	//ImfxSolarise solarise the image
	ImfxSolarise
	//ImfxPosterise posterise the image
	ImfxPosterise
	//ImfxWhiteboard whiteboard effect
	ImfxWhiteboard
	//ImfxBlackboard blackboard effect
	ImfxBlackboard
	//ImfxSketch sketch effect
	ImfxSketch
	//ImfxDenoise denoise the image
	ImfxDenoise
	//ImfxEmboss the image
	ImfxEmboss
	//ImfxOilpaint oil paint effect
	ImfxOilpaint
	//ImfxHatch hatch sektch effect
	ImfxHatch
	//ImfxGPen graphite sketch effect
	ImfxGPen
	//ImfxPastel pastel effect
	ImfxPastel
	//ImfxWatercolour watercolour effect
	ImfxWatercolour
	//ImfxFilm film grain effect
	ImfxFilm
	//ImfxBlur blur the image
	ImfxBlur
	//ImfxSaturation colour saturate the image
	ImfxSaturation
)

//Convert takes the type and returns the string representation of that value.
//Returns true as well if it is the default value.
func (t ImgEffectType) Convert() (string, bool) {
	switch t {
	case ImfxNone:
		return "none", true
	case ImfxNegative:
		return "negative", false
	case ImfxSolarise:
		return "solarise", false
	case ImfxPosterise:
		return "posterise", false
	case ImfxWhiteboard:
		return "whiteboard", false
	case ImfxBlackboard:
		return "blackboard", false
	case ImfxSketch:
		return "sketch", false
	case ImfxDenoise:
		return "denoise", false
	case ImfxEmboss:
		return "emboss", false
	case ImfxOilpaint:
		return "oilpaint", false
	case ImfxHatch:
		return "hatch", false
	case ImfxGPen:
		return "gpen", false
	case ImfxPastel:
		return "pastel", false
	case ImfxWatercolour:
		return "watercolour", false
	case ImfxFilm:
		return "film", false
	case ImfxBlur:
		return "blur", false
	case ImfxSaturation:
		return "saturation", false
	default:
		return "", true
	}
}
