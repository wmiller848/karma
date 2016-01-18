// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3.1 and OpenGL 4.1 core forward-compatible profile.
package main

import (
	"fmt"
	_ "image/png"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/wmiller848/Karma/game"
	"github.com/wmiller848/Karma/renderer"
)

// 	// 2560 x 1600
const windowWidth = 2560 / 2
const windowHeight = 1600 / 2

func main() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	r := renderer.CreateRenderer(windowWidth, windowHeight)
	r.SetCamera(windowWidth, windowHeight)
	// r.Render()
	g := game.CreateGame(r)
	g.LoadLevel()
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		g.Pause()
	}()
	g.Play()
	fmt.Println(g)
	// fmt.Println("%+v", r)
}
