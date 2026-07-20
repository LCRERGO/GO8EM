# GO8EM

A CHIP-8 emulator written in Go, using SDL2 for graphics, audio, and input.

## Features

- Full CHIP-8 instruction set
- SDL2 video output (64×32 display, 10× scaled)
- Audio beep via SDL audio subsystem
- Hex keyboard input
- Configurable ROM loading

## Requirements

- Go >= 1.20
- SDL2 development libraries
- make

### Installing SDL2

**Debian/Ubuntu:**
```
sudo apt install libsdl2-dev
```

**macOS (Homebrew):**
```
brew install sdl2
```

**Arch Linux:**
```
sudo pacman -S sdl2
```

## Building

```
make
```

The binary is generated at `build/go8em`.

### Running tests

```
make check          # unit tests
make check/all      # full tests (includes audio/input integration tests)
```

## Usage

```
./build/go8em <path-to-rom>
```

### Controls

The CHIP-8 hex keypad (0–F) maps to the following PC keys:

| CHIP-8 | PC Key |
|--------|--------|
| 0      | 0      |
| 1      | 1      |
| 2      | 2      |
| 3      | 3      |
| 4      | 4      |
| 5      | 5      |
| 6      | 6      |
| 7      | 7      |
| 8      | 8      |
| 9      | 9      |
| A      | A      |
| B      | B      |
| C      | C      |
| D      | D      |
| E      | E      |
| F      | F      |

## Project Structure

```
cmd/              — entry point (main.go)
pkg/
├── chip8/        — core CHIP-8 emulation (CPU, memory, registers, screen)
│   ├── instruction/ — instruction decoders and executors
│   ├── keyboard/    — internal CHIP-8 key state
│   ├── memory/      — memory + sprite data
│   ├── register/    — register file (V0–VF, I, DT, ST, PC, SP)
│   ├── screen/      — display buffer
│   └── stack/       — call stack
├── config/       — global configuration and logging
├── constants/    — emulator constants
├── subsystem/    — SDL I/O (video, audio, input)
│   ├── audio/       — sound output
│   ├── video/       — display rendering
│   └── input/       — keyboard device abstraction
└── utils/        — misc utilities (LCG, logging)
```

## License

GNU General Public License v3.0. See [LICENSE](LICENSE).
