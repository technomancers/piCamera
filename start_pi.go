// +build linux,arm,pi

package piCamera

import (
	"bytes"
	"log"
)

//Start raspivid in the background.
//Also logs the PID to os.stdOut.
//
//This is the function that is built seperatly depending on if you are building for the Raspberry Pi or not.
func (pc *PiCamera) Start() error {
	err := pc.command.Start()
	go pc.updateLatest()
	log.Printf("PiCamera running PID: %d", pc.command.Process.Pid)
	return err
}

func (pc *PiCamera) updateLatest() {
	readBuff := make([]byte, 4096)         //Buffer of the currently read bytes (4 kilobytes)
	var work = make([]byte, len(jpgMagic)) //This is the currently working bytes in process data. Must be outside function as it can carry over calls
	var buffer = new(bytes.Buffer)         //The new image buffer. The one currently being processed
	for {
		select {
		case <-pc.ctx.Done():
			break
		default:
			n, err := pc.stdOut.Read(readBuff)
			if err != nil {
				log.Printf("Reading from raspivid stdOut error: %v", err)
				break
			}
			start := 0 //This is where the image starts if this data splits images
			//Read through the data
			for i := 0; i < n; i++ {
				//add byte to working bytes
				work = append(work[1:], readBuff[i])
				//If we are at the start of a new image
				if bytes.Compare(work, jpgMagic) == 0 {
					buffer.Write(readBuff[start:i]) //write what is left of the old image
					if buffer.Len() > 0 {
						end := buffer.Len() - len(jpgMagic) + 1 //figure out where the end of the previous image was
						image := buffer.Bytes()[:end]
						cpyImage := make([]byte, len(image))
						copy(cpyImage, image)
						rest := buffer.Bytes()[end:]
						//write the image to the latest
						pc.rwMutext.Lock()
						pc.latestImg = cpyImage
						pc.rwMutext.Unlock()
						buffer.Reset()     //Clear out the buffer
						buffer.Write(rest) //Include the partial image that was left back into the buffer
						start = i          //make sure to change I so that the rest of the readBuffer is cleared correctly
					}
				}
			}
			buffer.Write(readBuff[start:n]) //write to the buffer readBuffer depending on where the start was
		}
	}
}
