package audio_test

import (
	"os"
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/subsystem/audio"
)

func TestAudio(t *testing.T) {
	if os.Getenv("TEST_ALL") != "" ||
		os.Getenv("TEST_AUDIO") != "" {
		audioSubsystem := audio.New(44100)
		audio.Beep(*audioSubsystem, 660, 800)
		audio.Destroy(audioSubsystem)
	}
}
