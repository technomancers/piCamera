// +build !linux !arm !pi

package piCamera

import (
	"image"
	"image/draw"

	"bytes"

	"image/jpeg"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

//Start raspivid in the background.
//Also logs the PID to os.stdOut.
//
//This is the function that is built seperatly depending on if you are building for the Raspberry Pi or not.
func (pc *PiCamera) Start() error {
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	fg, bg := image.White, image.Black
	rgba := image.NewRGBA(image.Rect(0, 0, pc.args.Width, pc.args.Height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(20)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)
	str := "You are not running running on a pi."
	base := int(float64(c.PointToFixed(20)>>6) / 2)
	pt := freetype.Pt((pc.args.Width/2)-(len(str)*5), (pc.args.Height/2)-base)
	_, err = c.DrawString(str, pt)
	if err != nil {
		return err
	}
	buff := new(bytes.Buffer)
	err = jpeg.Encode(buff, rgba, &jpeg.Options{Quality: 70})
	if err != nil {
		return err
	}
	pc.latestImg = buff.Bytes()
	return nil
}
