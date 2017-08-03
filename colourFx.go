package piCamera

import (
	"fmt"
)

//ColourEffect is used fo setting which color effect one wants.
type ColourEffect struct {
	u int
	v int
}

//NewColourEffect creates a colour effect of the U and Y channel of the image.
//Values should be in the range from (0, 255).
func NewColourEffect(u, v int) *ColourEffect {
	return &ColourEffect{
		u: u,
		v: v,
	}
}

//Convert takes the type and returns the string representation of that value.
func (t *ColourEffect) Convert() string {
	return fmt.Sprintf("%d:%d", t.u, t.v)
}
