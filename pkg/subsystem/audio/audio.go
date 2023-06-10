package audio

// #include <stdlib.h>
// #include <stdint.h>
// #include <math.h>
// typedef unsigned char Uint8;
// typedef uint32_t Uint32;
// void SineWave(void *userdata, Uint8 *stream, int len);
//
// typedef struct UserData {
// 		int frequency;
// 		int sampleRate;
//		float dPhase;
// } UserData;
//
// __attribute__((weak))
// UserData* NewUserData(Uint32 frequency, Uint32 sampleRate) {
//    UserData* userdata = malloc(sizeof(UserData));
//    userdata->frequency = frequency;
//    userdata->sampleRate = sampleRate;
//    userdata->dPhase = 2 * M_PI * frequency / sampleRate;
//
//    return userdata;
// }
//
// __attribute__((weak))
// void DestroyUserData(UserData* userdata) {
//    if (userdata)
//        free(userdata);
// }
import "C"
import (
	"log"
	"math"
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type AudioSubsystem struct {
	SampleRate int
}

func New(sampleRate int) *AudioSubsystem {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Fatal("new_audio_subsystem:", err)
	}

	return &AudioSubsystem{
		SampleRate: sampleRate,
	}
}

func Beep(audio AudioSubsystem, frequency, duration int) {
	userData := C.NewUserData(C.Uint32(frequency), C.Uint32(audio.SampleRate))
	defer C.DestroyUserData(userData)

	spec := sdl.AudioSpec{
		Freq:     int32(audio.SampleRate),
		Format:   sdl.AUDIO_U8,
		Channels: 2,
		Samples:  512,
		Callback: sdl.AudioCallback(C.SineWave),
		UserData: unsafe.Pointer(userData),
	}

	if err := sdl.OpenAudio(&spec, nil); err != nil {
		log.Fatal("beep_audio_subsystem:", err)
	}

	sdl.PauseAudio(false)
	sdl.Delay(uint32(duration))
}

func Destroy(audio *AudioSubsystem) {
	sdl.CloseAudio()
}

//export SineWave
func SineWave(userdata unsafe.Pointer, stream *C.Uint8, len C.int) {
	n := int(len)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	data := (*C.UserData)(userdata)

	var phase float64
	for i := 0; i < n; i += 2 {
		phase += float64(data.dPhase)
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128)
		buf[i] = sample
		buf[i+1] = sample
	}
}
