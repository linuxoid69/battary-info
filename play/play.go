package play

import (
	"fmt"
	"github.com/cocoonlife/goalsa"
	"github.com/cryptix/wav"
	"io"
	"log"
	"os"
)

// PlaySound play sound
func PlaySound() {

	samplerate := 25000
	player, err := alsa.NewPlaybackDevice("default", 1, alsa.FormatS16LE, samplerate, alsa.BufferParams{})
	if err != nil {
		log.Println("Can't create playback device")
	}
	defer player.Close()

	if _, err := os.Stat("/tmp/sound/battery_is_low.wav"); err != nil {
		RestoreAsset("/tmp/", "sound/battery_is_low.wav")
	}

	soundfile, err := os.Open("/tmp/sound/battery_is_low.wav")

	fi, _ := soundfile.Stat()

	if err != nil {
		fmt.Printf("Error open file: %s", err)
	}

	wavReader, err := wav.NewReader(soundfile, fi.Size())
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	if wavReader == nil {
		fmt.Sprint("nil wav reader")
	}

	for {
		s, err := wavReader.ReadSampleEvery(2, 0)

		var cvert []int16

		for _, b := range s {
			cvert = append(cvert, int16(b))
		}

		if cvert != nil {
			// play!
			player.Write(cvert)
		}

		cvert = []int16{}

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Errorf("WAV Decode: %s", err)
		}
	}
	return
}
