package piCamera

import (
	"fmt"
)

//RegionOfIntrest is used to set the cameras area to be used as the source.
type RegionOfIntrest struct {
	tlX float32
	tlY float32
	w   float32
	h   float32
}

//NewROI creates a new Region of Intrest.
//tlx and tly are the top left x and y of ROI.
//w and h are the width and height of the ROI.
//All points should be normalized from 0.0 - 1.0.
func NewROI(tlx, tly, w, h float32) *RegionOfIntrest {
	return &RegionOfIntrest{
		tlX: tlx,
		tlY: tly,
		w:   w,
		h:   h,
	}
}

//Convert takes the type and returns the string representation of that value.
func (t *RegionOfIntrest) Convert() string {
	return fmt.Sprintf("%.2f,%.2f,%.2f,%.2f", t.tlX, t.tlY, t.w, t.h)
}
