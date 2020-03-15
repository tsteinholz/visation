// desktop/main.go -- Visation
// Copyright (C) 2020 Thomas Steinholz
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
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Plaform-Specific Code for the Desktop Enviornment!
// - Creates, Manages and Destroys desktop resoucres.

package main

import (
	"log"
	"runtime"
	"time"

	as "github.com/vulkan-go/asche"
	"github.com/vulkan-go/demos/vulkancube"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
	"github.com/xlab/closer"
)

func init() {
	runtime.LockOSThread()
	log.SetFlags(log.Lshortfile)
}

type Application struct {
	*vulkancube.SpinningCube
	debugEnabled bool
	windowHandle *glfw.Window
}

func (a *Application) VulkanSurface(instance vk.Instance) (surface vk.Surface) {
	surfPtr, err := a.windowHandle.CreateWindowSurface(instance, nil)
	if err != nil {
		log.Println(err)
		return vk.NullSurface
	}
	return vk.SurfaceFromPointer(surfPtr)
}

func (a *Application) VulkanAppName() string {
	return "VulkanCube"
}

func (a *Application) VulkanLayers() []string {
	return []string{
		// "VK_LAYER_GOOGLE_threading",
		// "VK_LAYER_LUNARG_parameter_validation",
		// "VK_LAYER_LUNARG_object_tracker",
		// "VK_LAYER_LUNARG_core_validation",
		// "VK_LAYER_LUNARG_api_dump",
		// "VK_LAYER_LUNARG_swapchain",
		// "VK_LAYER_GOOGLE_unique_objects",
	}
}

func (a *Application) VulkanDebug() bool {
	return false // a.debugEnabled
}

func (a *Application) VulkanDeviceExtensions() []string {
	return []string{
		"VK_KHR_swapchain",
	}
}

func (a *Application) VulkanSwapchainDimensions() *as.SwapchainDimensions {
	return &as.SwapchainDimensions{
		Width: 500, Height: 500, Format: vk.FormatB8g8r8a8Unorm,
	}
}

func (a *Application) VulkanInstanceExtensions() []string {
	extensions := a.windowHandle.GetRequiredInstanceExtensions()
	if a.debugEnabled {
		extensions = append(extensions, "VK_EXT_debug_report")
	}
	return extensions
}

func NewApplication(debugEnabled bool) *Application {
	return &Application{
		SpinningCube: vulkancube.NewSpinningCube(1.0),

		debugEnabled: debugEnabled,
	}
}

func main() {
	orPanic(glfw.Init())
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())
	orPanic(vk.Init())
	defer closer.Close()

	app := NewApplication(true)
	reqDim := app.VulkanSwapchainDimensions()
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(int(reqDim.Width), int(reqDim.Height), "Visation", nil, nil)
	orPanic(err)
	app.windowHandle = window

	// creates a new platform, also initializes Vulkan context in the app
	platform, err := as.NewPlatform(app)
	orPanic(err)

	dim := app.Context().SwapchainDimensions()
	log.Printf("Initialized %s with %+v swapchain", app.VulkanAppName(), dim)

	// some sync logic
	doneC := make(chan struct{}, 2)
	exitC := make(chan struct{}, 2)
	defer closer.Bind(func() {
		exitC <- struct{}{}
		<-doneC
		log.Println("Bye!")
	})

	fpsDelay := time.Second / 60
	fpsTicker := time.NewTicker(fpsDelay)
	for {
		select {
		case <-exitC:
			app.Destroy()
			platform.Destroy()
			window.Destroy()
			glfw.Terminate()
			fpsTicker.Stop()
			doneC <- struct{}{}
			return
		case <-fpsTicker.C:
			if window.ShouldClose() {
				exitC <- struct{}{}
				continue
			}
			glfw.PollEvents()
			app.NextFrame()

			imageIdx, outdated, err := app.Context().AcquireNextImage()
			orPanic(err)
			if outdated {
				imageIdx, _, err = app.Context().AcquireNextImage()
				orPanic(err)
			}
			_, err = app.Context().PresentImage(imageIdx)
			orPanic(err)
		}
	}
}

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case vk.Result:
		if err := vk.Error(v); err != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}
