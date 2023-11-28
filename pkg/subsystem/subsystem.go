// SDL controller subsystem.
package subsystem

import (
	"fmt"
	"os"
	"time"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction"
	"github.com/LCRERGO/GO8EM/pkg/chip8/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/LCRERGO/GO8EM/pkg/config"
	"github.com/LCRERGO/GO8EM/pkg/constants"
	"github.com/LCRERGO/GO8EM/pkg/subsystem/audio"
	"github.com/LCRERGO/GO8EM/pkg/subsystem/input/device"
	keyboard_map "github.com/LCRERGO/GO8EM/pkg/subsystem/input/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/subsystem/video"
	"github.com/veandco/go-sdl2/sdl"
)

type SubsystemController struct {
	inputDevice *device.DeviceSubsystem
	audio       *audio.AudioSubsystem
	video       *video.VideoSubsystem

	chip8 *chip8.Chip8
}

func New(chip8 *chip8.Chip8) *SubsystemController {
	return &SubsystemController{
		audio: audio.New(44100),
		video: video.New(),

		chip8: chip8,
	}
}

func Destroy(controller *SubsystemController) {
	video.Destroy(controller.video)
	audio.Destroy(controller.audio)
}

func AddROM(controller *SubsystemController, fname string) error {
	device, err := device.New(fname)
	if err != nil {
		return err
	}
	controller.inputDevice = device
	loadROM2Memory(controller)

	return nil
}

func RemoveROM(controller *SubsystemController) {
	if controller.inputDevice != nil {
		device.Destroy(controller.inputDevice)
	}
}

func WaitForKeyPress(controller *SubsystemController) func() (int, error) {
	return func() (int, error) {
		if controller != nil {
			var event sdl.Event

			for event = sdl.WaitEvent(); event != nil; event = sdl.WaitEvent() {
				switch t := event.(type) {
				case *sdl.KeyboardEvent:
					if t.Type == sdl.KEYDOWN {
						keySymbol := t.Keysym.Sym
						repr, err := keyboard_map.GetKeyRepr(keySymbol)
						if err != nil {
							config.GetLogger(config.GetInstance()).
								Printf("wait_for_key_press: %v", err)
						}
						return repr, nil
					}

				default:
					// NOOP
					continue
				}
			}
		}

		return 0xFF, fmt.Errorf("wait_for_key_press: impossible branch")
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

		opcode := chip8.FetchOpcode(controller.chip8)
		instr, args := instruction.Decode(opcode)
		instr.ExecFunc(controller.chip8, args)
	}
}

func loadROM2Memory(controller *SubsystemController) {
	if controller.inputDevice != nil {
		rom := make([]byte, constants.ROMMaxSize)
		_, err := controller.inputDevice.File.Read(rom)
		if err != nil {
			config.GetLogger(config.GetInstance()).
				Fatalf("load_rom_2_memory: %v", err)
		}

		memory.LoadROM(controller.chip8.Memory, rom, controller.inputDevice.Size)
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
			repr, err := keyboard_map.GetKeyRepr(keySymbol)
			if err != nil {
				config.GetLogger(config.GetInstance()).
					Printf("handle_events: %v", err)
			}

			if t.Type == sdl.KEYDOWN {
				keyboard.SetKeyDown(controller.chip8.Keyboard, repr)
			} else if t.Type == sdl.KEYUP {
				keyboard.SetKeyUp(controller.chip8.Keyboard, repr)
			}
		default:
			// NOOP
			continue
		}
	}
}

func handleDelayTimer(controller *SubsystemController) {
	for controller.chip8.Registers.DT > 0 {
		time.Sleep(constants.ClockFrequency * time.Microsecond)
		register.DecDT(controller.chip8.Registers)
	}
}

func handleSoundTimer(controller *SubsystemController) {
	if controller.chip8.Registers.ST > 0 {
		go audio.Beep(controller.audio,
			12500,
			100*int(controller.chip8.Registers.ST))
		register.ResetST(controller.chip8.Registers)
	}
}
