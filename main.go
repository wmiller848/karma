// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3.1 and OpenGL 4.1 core forward-compatible profile.
package main

import (
	_ "image/png"
	"runtime"

	"github.com/wmiller848/Karma/renderer"
)

const windowWidth = 800
const windowHeight = 600

func main() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	r := renderer.CreateRenderer(windowWidth, windowHeight)
	r.Render()
	// fmt.Println("%+v", r)
}
