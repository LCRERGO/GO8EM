package subsystem

import (
	"log"
	"os"
	"time"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/constants"
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

	chip8 *chip8.Chip8
}

func New(chip8 *chip8.Chip8) *SubsystemController {
	return &SubsystemController{
		inputKeyboard: keyboard.New(),
		audio:         audio.New(44100),
		video:         video.New(),

		chip8: chip8,
	}
}

func Destroy(controller *SubsystemController) {
	video.Destroy(controller.video)
	audio.Destroy(controller.audio)
	keyboard.Destroy(controller.inputKeyboard)
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

func Run(controller *SubsystemController) {
	for controller != nil {
		step(controller)
		time.Sleep(constants.ClockFrequency * time.Microsecond)
	}
}

func step(controller *SubsystemController) {
	if controller.inputDevice != nil {
		handleEvents(controller)
		video.Render(controller.video, controller.chip8.Screen)
		handleDelayTimer(controller)
		handleSoundTimer(controller)
	}
}

func handleEvents(controller *SubsystemController) {
	var event sdl.Event

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			RemoveROM(controller)
			Destroy(controller)
			os.Exit(0)
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
			// NOOP
		}
	}
}

func handleDelayTimer(controller *SubsystemController) {
	if controller.chip8.Registers.DT > 0 {
		time.Sleep(1500 * time.Microsecond)
		register.DecDT(controller.chip8.Registers)
		log.Printf("delay: %v", controller.chip8.Registers.DT)
	}

}

func handleSoundTimer(controller *SubsystemController) {
	if controller.chip8.Registers.ST > 0 {
		register.DecST(controller.chip8.Registers)
		log.Printf("sound: %v", controller.chip8.Registers.ST)
	}
}
