// A simple CHIP-8 emulator written in Golang.
// Copyright (C) 2023  Lucas Cruz dos Reis <lcr.ergo@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/handler"
	"github.com/LCRERGO/GO8EM/pkg/config"
	"github.com/LCRERGO/GO8EM/pkg/subsystem"
	sdllog "github.com/LCRERGO/GO8EM/pkg/utils/log/sdl"
)

func init() {
	globalConfig := config.New()
	config.AddLogger(globalConfig, sdllog.New())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}

	// setup emulator
	c8 := chip8.New()
	defer chip8.Destroy(c8)
	chip8.AddRandGen(c8, config.GetRandomGenerator(config.GetInstance()))
	controller := subsystem.New(c8)
	subsystem.AddROM(controller, os.Args[1])
	defer subsystem.RemoveROM(controller)

	// setup handlers
	waitForKeyPress := subsystem.WaitForKeyPress(controller)
	handlerAggregator := handler.New(waitForKeyPress)
	defer handler.Destroy(handlerAggregator)
	chip8.AddHandler(c8, handlerAggregator)

	subsystem.Run(controller)
}
