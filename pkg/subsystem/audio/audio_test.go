package audio_test

import (
	"os"
	"testing"
	"time"

	"github.com/LCRERGO/GO8EM/pkg/subsystem/audio"
)

func TestAudio(t *testing.T) {
	if os.Getenv("TEST_ALL") != "" ||
		os.Getenv("TEST_AUDIO") != "" {
		audioSubsystem := audio.New(44100)
		audio.Start(audioSubsystem)
		time.Sleep(800 * time.Millisecond)
		audio.Stop(audioSubsystem)
		audio.Destroy(audioSubsystem)
	}
}
