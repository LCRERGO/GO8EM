// SDL audio subsystem.
package audio

// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include <stdint.h>
// #include <math.h>
// typedef unsigned char Uint8;
// typedef uint32_t Uint32;
//
// typedef struct UserData {
//     int playing;
//     double dPhase;
//     double phase;
// } UserData;
//
// void SineWave(void *userdata, Uint8 *stream, int len) {
//     UserData* data = (UserData*)userdata;
//     if (!data || !data->playing) {
//         for (int i = 0; i < len; i++) {
//             stream[i] = 128;
//         }
//         if (data) data->phase = 0.0;
//         return;
//     }
//     for (int i = 0; i < len; i++) {
//         data->phase += data->dPhase;
//         double sample = sin(data->phase);
//         stream[i] = (Uint8)((sample + 1.0) * 127.5);
//     }
// }
//
// UserData* NewUserData(int sampleRate) {
//     UserData* data = (UserData*)malloc(sizeof(UserData));
//     data->playing = 0;
//     data->dPhase = 2.0 * M_PI * 440.0 / sampleRate;
//     data->phase = 0.0;
//     return data;
// }
//
// void DestroyUserData(UserData* data) {
//     if (data) free(data);
// }
//
// void SetPlaying(UserData* data, int playing) {
//     if (data) data->playing = playing;
// }
import "C"
import (
	"unsafe"

	"github.com/LCRERGO/GO8EM/pkg/config"
	"github.com/veandco/go-sdl2/sdl"
)

type AudioSubsystem struct {
	SampleRate int
	DeviceID   sdl.AudioDeviceID
	userData   *C.UserData
}

func New(sampleRate int) *AudioSubsystem {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		config.GetLogger(config.GetInstance()).
			Fatal("new_audio_subsystem:", err)
	}

	userData := C.NewUserData(C.int(sampleRate))

	audio := &AudioSubsystem{
		SampleRate: sampleRate,
		DeviceID:   0,
		userData:   userData,
	}

	spec := sdl.AudioSpec{
		Freq:     int32(sampleRate),
		Format:   sdl.AUDIO_U8,
		Channels: 1,
		Samples:  512,
		Callback: sdl.AudioCallback(C.SineWave),
		UserData: unsafe.Pointer(userData),
	}

	deviceID, err := sdl.OpenAudioDevice(
		"",
		false,
		&spec,
		nil,
		sdl.AUDIO_ALLOW_ANY_CHANGE,
	)
	if err != nil {
		config.GetLogger(config.GetInstance()).
			Fatal("new_audio_subsystem:", err)
	}

	audio.DeviceID = deviceID
	return audio
}

func Destroy(audio *AudioSubsystem) {
	if audio != nil {
		if audio.DeviceID > 0 {
			sdl.CloseAudioDevice(audio.DeviceID)
		}
		C.DestroyUserData(audio.userData)
	}
}

func Start(audio *AudioSubsystem) {
	if audio != nil && audio.userData != nil {
		C.SetPlaying(audio.userData, 1)
		sdl.PauseAudioDevice(audio.DeviceID, false)
	}
}

func Stop(audio *AudioSubsystem) {
	if audio != nil && audio.userData != nil {
		sdl.PauseAudioDevice(audio.DeviceID, true)
		C.SetPlaying(audio.userData, 0)
	}
}
