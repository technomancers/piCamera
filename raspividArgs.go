package piCamera

import (
	"context"
	"os/exec"
	"strconv"
)

//Command has a non zero default so these are the values used to make up for it
const (
	defBrightness = 50
	defMode       = -1
)

//RaspividArgs are arguments used to set camera settings for the desired output
//https://www.raspberrypi.org/documentation/raspbian/applications/camera.md
type RaspividArgs struct {
	Width         int              // width of the image
	Height        int              // height of the image
	HorizFlip     bool             // flip the image horizontally
	VertFlip      bool             // flip the camera vertically
	Sharpness     int              // change the sharpness of the camera (-100 , 100 DEF 0)
	Contrast      int              // change the contrast of the camera (-100 , 100 DEF 0)
	Brightness    int              // change the brightness of the camera (0 , 100 DEF 50)
	Saturation    int              // change the saturation of the camera (-100 , 100 DEF 0)
	ISO           int              // change the sensitivity the camera is to light (100 , 800 DEF 100)
	VideoStable   bool             // try to stableize the video
	EV            int              // Slightly under or over expose the camera (-10 , 10 DEF 0)
	ExposureMode  ExposureType     // set which mode to use for exposure
	AWB           AWBType          // set the automatic white balance mode
	ImageFx       ImgEffectType    // set the image effect
	ColourFx      *ColourEffect    // set the color effects to an image
	Metering      MeteringType     // ste the metering mode
	Rotation      int              // set the rotation of the image. (0, 90, 180, 270)
	ROI           *RegionOfIntrest // set the cameras region of intrest
	ShutterSpeed  int              // set the shutter speed in microseconds (Max 6000000)
	DRC           DRCType          // set the dynamic range compression
	AWBGains      *AWBGains        // set the AWBGains when AWB is off
	Mode          int              // set the mode of the camera by checking the documentation
	Annotate      string           // annotate the image according to the documentation
	AnnotateExtra string           // annotate the image according to the documentation
	Bitrate       int              // set the bitrate in bits per second. Max is 25000000
	FPS           int              // set the frames per second (2 , 30)
	IntraRate     int              // set number of frames before next intra frame
	Quantization  int              // set Quantization parameter
	Profile       ProfileType      // set the profile type
	InsertHeaders bool             // insert pps, sps headers to every I-Frame
}

//NewArgs returns a RaspividArgs with the default settings
func NewArgs() *RaspividArgs {
	return &RaspividArgs{
		Brightness: defBrightness,
		Mode:       defMode,
	}
}

func createCommand(ctx context.Context, args *RaspividArgs) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "raspivid", "-cd", "MJPEG", "-t", "0")
	var final []string
	if args.Width != 0 {
		final = append(final, "-w", strconv.Itoa(args.Width))
	}
	if args.Height != 0 {
		final = append(final, "-h", strconv.Itoa(args.Height))
	}
	if args.HorizFlip {
		final = append(final, "-hf")
	}
	if args.VertFlip {
		final = append(final, "-vf")
	}
	if args.Sharpness != 0 {
		final = append(final, "-sh", strconv.Itoa(args.Sharpness))
	}
	if args.Contrast != 0 {
		final = append(final, "-co", strconv.Itoa(args.Contrast))
	}
	if args.Brightness != defBrightness {
		final = append(final, "-br", strconv.Itoa(args.Brightness))
	}
	if args.Saturation != 0 {
		final = append(final, "-sa", strconv.Itoa(args.Saturation))
	}
	if args.ISO != 0 {
		final = append(final, "-ISO", strconv.Itoa(args.ISO))
	}
	if args.VideoStable {
		final = append(final, "-vs")
	}
	if args.EV != 0 {
		final = append(final, "-ev", strconv.Itoa(args.EV))
	}
	if mode, def := args.ExposureMode.Convert(); !def {
		final = append(final, "-ex", mode)
	}
	if mode, def := args.AWB.Convert(); !def {
		final = append(final, "-awb", mode)
	}
	if mode, def := args.ImageFx.Convert(); !def {
		final = append(final, "-ifx", mode)
	}
	if args.ColourFx != nil {
		final = append(final, "-cfx", args.ColourFx.Convert())
	}
	if mode, def := args.Metering.Convert(); !def {
		final = append(final, "-mm", mode)
	}
	if args.Rotation != 0 {
		final = append(final, "-rot", strconv.Itoa(args.Rotation))
	}
	if args.ROI != nil {
		final = append(final, "-roi", args.ROI.Convert())
	}
	if args.ShutterSpeed != 0 {
		final = append(final, "-ss", strconv.Itoa(args.ShutterSpeed))
	}
	if mode, def := args.DRC.Convert(); !def {
		final = append(final, "-drc", mode)
	}
	if args.AWBGains != nil {
		final = append(final, "-awbg", args.AWBGains.Convert())
	}
	if args.Mode != defMode {
		final = append(final, "-md", strconv.Itoa(args.Mode))
	}
	if args.AnnotateExtra != "" {
		final = append(final, "-ae", args.AnnotateExtra)
	}
	if args.Annotate != "" {
		final = append(final, "-a", args.Annotate)
	}
	if args.Bitrate != 0 {
		final = append(final, "-b", strconv.Itoa(args.Bitrate))
	}
	if args.FPS != 0 {
		final = append(final, "-fps", strconv.Itoa(args.FPS))
	}
	if args.IntraRate != 0 {
		final = append(final, "-g", strconv.Itoa(args.IntraRate))
	}
	if args.Quantization != 0 {
		final = append(final, "-qp", strconv.Itoa(args.Quantization))
	}
	if mode, def := args.Profile.Convert(); !def {
		final = append(final, "-pf", mode)
	}
	if args.InsertHeaders {
		final = append(final, "-ih")
	}
	final = append(final, "-o", "-")
	cmd.Args = append(cmd.Args, final...)
	return cmd
}
