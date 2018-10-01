//Copyright (c) 2017, Technomancers. All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

/*Package piCamera is a simple wrapper for raspivid.

There is a non-RaspberryPi version that is used for local development.
This can become handy when the IDE does not know how to handle certain features.
*/
package piCamera

import (
	"context"
	"io"
	"os/exec"

	"errors"
	"os"
	"sync"
)

//nolint: varcheck, unused
var jpgMagic = []byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46} //This is covered here https://asecuritysite.com/forensics/jpeg

//PiCamera creates a way for code to be able to pull images from the camera live.
//You must start eh raspivid explicitly first but once started, PiCamera will have the latest image available to view.
//
//PiCamera is thread safe so many calls to GetFrame() will not break.
//Be careful as the more calls to GetFrame() the slower GetFrame() may become due to all the read locks.
type PiCamera struct {
	rwMutext  *sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	command   *exec.Cmd
	stdOut    io.ReadCloser
	latestImg []byte
	args      *RaspividArgs
}

//New creates an instance of PiCamera.
//Width and Height are for the image size.
//ctx is the parent context. If nil a background context will be created.
//
//This creates the command raspivid with the appropriate settings.
//The stdErr of the command is redirected to os.Stderr so that one may see why the command may have failed.
func New(parentCtx context.Context, args *RaspividArgs) (*PiCamera, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	if parentCtx == nil {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithCancel(parentCtx)
	}
	cmd := createCommand(ctx, args)
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, err
	}
	cmd.Stderr = os.Stderr
	return &PiCamera{
		ctx:      ctx,
		cancel:   cancel,
		command:  cmd,
		stdOut:   stdOut,
		rwMutext: new(sync.RWMutex),
		args:     args,
	}, nil
}

//GetFrame returns the latest frame from raspivid.
//If there is no frame available it will throw an error.
func (pc *PiCamera) GetFrame() ([]byte, error) {
	if pc.ctx.Err() != nil {
		return nil, pc.ctx.Err()
	}
	pc.rwMutext.RLock()
	defer pc.rwMutext.RUnlock()
	if pc.latestImg == nil {
		return nil, errors.New("Latest Image is empty")
	}
	return pc.latestImg, nil
}

//Stop the raspivid command.
//Safely stop all the commands and routines with this.
func (pc *PiCamera) Stop() {
	if pc.ctx.Err() == nil {
		pc.cancel()
	}
}
