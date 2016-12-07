package main

/*
#cgo pkg-config: sdl2 SDL2_ttf
#include <SDL2/SDL.h>
#include <SDL2/SDL_ttf.h>

// Don't know how to access event struct fields in Go,
// event._type won't work ...
int should_quit() {
	SDL_Event event;
	if (SDL_PollEvent(&event) != 0) {
		if (event.type == SDL_QUIT) {
			return 1;
		}
	}
	return 0;
}
*/
import "C"

import (
	"log"
	"strconv"
	"time"
	"unsafe"
)

func sdl_error() {
	log.Fatalf("SDL error: %v\n", C.GoString(C.SDL_GetError()))
}

func ttf_error() {
	log.Fatalf("SDL TTF error: %v\n", C.GoString(C.TTF_GetError()))
}

func main() {
	if C.SDL_Init(C.SDL_INIT_VIDEO) < 0 {
		sdl_error()
	}
	defer C.SDL_Quit()
	if C.TTF_Init() < 0 {
		ttf_error()
	}
	defer C.TTF_Quit()
	title := C.CString("Hello SDL")
	defer C.free(unsafe.Pointer(title))
	window := C.SDL_CreateWindow(
		title,
		C.SDL_WINDOWPOS_CENTERED,
		C.SDL_WINDOWPOS_CENTERED,
		800,
		600,
		0,
	)
	if window == nil {
		sdl_error()
	}
	renderer := C.SDL_CreateRenderer(
		window,
		-1,
		0,
	)
	if renderer == nil {
		sdl_error()
	}
	defer C.SDL_DestroyRenderer(renderer)
	fontName := C.CString("font.ttf")
	defer C.free(unsafe.Pointer(fontName))
	font := C.TTF_OpenFont(fontName, 24)
	if font == nil {
		ttf_error()
	}
	color := C.struct_SDL_Color{255, 0, 0, 0}
	t := time.Now()
	var frameTimes [100]float64
	frameFirst := 1
	frameLast := 0
	frameTimeTotal := 0.0
	for quit := false; !quit; {
		frameTime := time.Since(t).Seconds()
		t = time.Now()
		frameTimes[frameLast] = frameTime
		frameTimeTotal = frameTimeTotal - frameTimes[frameFirst] + frameTimes[frameLast]
		frameFirst = (frameFirst + 1) % len(frameTimes)
		frameLast = (frameLast + 1) % len(frameTimes)
		fps := float64(len(frameTimes)) / frameTimeTotal
		text := C.CString(strconv.Itoa(int(fps)))
		C.SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255)
		C.SDL_RenderClear(renderer)
		surface := C.TTF_RenderText_Solid(font, text, color)
		if surface == nil {
			ttf_error()
		}
		C.free(unsafe.Pointer(text))
		msg := C.SDL_CreateTextureFromSurface(renderer, surface)
		if msg == nil {
			sdl_error()
		}
		rect := C.struct_SDL_Rect{250, 250, 300, 100}
		if C.SDL_RenderCopy(renderer, msg, nil, &rect) < 0 {
			sdl_error()
		}
		C.SDL_RenderPresent(renderer)
		if C.should_quit() == 1 {
			quit = true
		}
	}
}
