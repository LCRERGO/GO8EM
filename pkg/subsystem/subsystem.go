package subsystem

import (
	"log"

	"github.com/LCRERGO/GO8EM/pkg/subsystem/audio"
	"github.com/LCRERGO/GO8EM/pkg/subsystem/input/device"
	"github.com/LCRERGO/GO8EM/pkg/subsystem/input/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/subsystem/video"
	"github.com/veandco/go-sdl2/sdl"
)

type SubsystemController struct {
	inputDevice   *device.DeviceSubsystem
	inputKeyboard *keyboard.KeyboardSubsystem
	audio         *audio.AudioSubsystem
	video         *video.VideoSubsystem
}

func New() *SubsystemController {
	return &SubsystemController{
		inputKeyboard: keyboard.New(),
		audio:         audio.New(44100),
		video:         video.New(),
	}
}

func Destroy(controller *SubsystemController) {
	video.Destroy(controller.video)
	audio.Destroy(controller.audio)
	keyboard.Destroy(controller.inputKeyboard)
	// exit?
}

func AddROM(controller *SubsystemController, fname string) error {
	device, err := device.New(fname)
	if err != nil {
		return err
	}
	controller.inputDevice = device

	return nil
}

func RemoveROM(controller *SubsystemController) {
	if controller.inputDevice != nil {
		device.Destroy(controller.inputDevice)
	}
}

func Step(controller *SubsystemController, event sdl.Event) {
	if controller.inputDevice != nil {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Destroy(controller)
		case *sdl.KeyboardEvent:
			keySymbol := t.Keysym.Sym
			repr, err := keyboard.GetKeyRepr(controller.inputKeyboard, keySymbol)
			if err != nil {
				log.Printf("handle_key_up: %v", err)
			}

			if t.Type == sdl.KEYDOWN {
				keyboard.KeyDown(controller.inputKeyboard, repr)
			} else if t.Type == sdl.KEYUP {
				keyboard.KeyUp(controller.inputKeyboard, repr)
			}

		default:
			// NOT REACHED
			log.Fatal("step: invalid event")
		}

	}
}
