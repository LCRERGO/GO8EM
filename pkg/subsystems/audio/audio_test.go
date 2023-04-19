package audio

import (
	"os"
	"testing"
)

func TestAudio(t *testing.T) {
	if os.Getenv("TEST_ALL") != "" ||
		os.Getenv("TEST_AUDIO") != "" {
		audio := NewAudioSubsystem(44100)
		Beep(*audio, 660, 800)
		Destroy(audio)
	}
}
