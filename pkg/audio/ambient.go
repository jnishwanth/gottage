package audio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

// SoundType defines cozy ambient sound options.
type SoundType int

const (
	SoundRain SoundType = iota
	SoundFire
	SoundWind
	SoundClock
)

// AmbientManager handles background procedural audio playback.
type AmbientManager struct {
	mu           sync.Mutex
	RainActive   bool
	FireActive   bool
	WindActive   bool
	ClockActive  bool
	Volume       float64 // 0.0 to 1.0
	tempDir      string
	rainCmd      *exec.Cmd
	fireCmd      *exec.Cmd
	windCmd      *exec.Cmd
	playerBinary string
	rng          *rand.Rand
}

// NewAmbientManager initializes the ambient sound engine.
func NewAmbientManager() *AmbientManager {
	tmpDir, _ := os.MkdirTemp("", "gottage_audio")
	am := &AmbientManager{
		RainActive:  false,
		FireActive:  false,
		WindActive:  false,
		ClockActive: false,
		Volume:      0.5,
		tempDir:     tmpDir,
		rng:         rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// Detect audio player binary
	if path, err := exec.LookPath("afplay"); err == nil {
		am.playerBinary = path
	} else if path, err := exec.LookPath("mpv"); err == nil {
		am.playerBinary = path
	} else if path, err := exec.LookPath("ffplay"); err == nil {
		am.playerBinary = path
	} else if path, err := exec.LookPath("aplay"); err == nil {
		am.playerBinary = path
	}

	// Generate procedural ambient WAV files in temp dir
	am.generateAmbientFiles()
	return am
}

func (am *AmbientManager) generateAmbientFiles() {
	if am.tempDir == "" {
		return
	}

	// Rain WAV (4 seconds smooth low-pass filtered loop)
	var prevRain float64
	rainWav := am.createProceduralWAV(4.0, func(t float64) float64 {
		raw := (am.rng.Float64()*2.0 - 1.0) * 0.08
		prevRain = prevRain*0.85 + raw*0.15 // Low pass filter for soft rain sound
		patter := (math.Sin(t*12.0)*0.3 + 0.7)
		return prevRain * patter
	})
	_ = os.WriteFile(filepath.Join(am.tempDir, "rain.wav"), rainWav, 0644)

	// Fireplace Crackle WAV (4 seconds warm rumble loop)
	var prevFire float64
	fireWav := am.createProceduralWAV(4.0, func(t float64) float64 {
		rumble := math.Sin(t*90.0)*0.04 + (am.rng.Float64()*2.0-1.0)*0.02
		prevFire = prevFire*0.90 + rumble*0.10
		return prevFire
	})
	_ = os.WriteFile(filepath.Join(am.tempDir, "fire.wav"), fireWav, 0644)

	// Wind Breeze WAV (4 seconds loop)
	windWav := am.createProceduralWAV(4.0, func(t float64) float64 {
		mod := math.Sin(t * 1.2)
		freq := 140.0 + mod*40.0
		val := math.Sin(t * 2.0 * math.Pi * freq)
		return val * 0.05 * (math.Sin(t*0.6)*0.5 + 0.5)
	})
	_ = os.WriteFile(filepath.Join(am.tempDir, "wind.wav"), windWav, 0644)
}

func (am *AmbientManager) createProceduralWAV(duration float64, sampleFunc func(t float64) float64) []byte {
	sampleRate := 22050
	numSamples := int(duration * float64(sampleRate))
	headerSize := 44
	dataSize := numSamples * 2 // 16-bit mono

	buf := new(bytes.Buffer)

	// WAV Header
	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, uint32(headerSize+dataSize-8))
	buf.WriteString("WAVE")
	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(16)) // Subchunk1Size (16 for PCM)
	binary.Write(buf, binary.LittleEndian, uint16(1))  // AudioFormat (1 for PCM)
	binary.Write(buf, binary.LittleEndian, uint16(1))  // NumChannels (1 mono)
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate*2)) // ByteRate
	binary.Write(buf, binary.LittleEndian, uint16(2))           // BlockAlign
	binary.Write(buf, binary.LittleEndian, uint16(16))          // BitsPerSample
	buf.WriteString("data")
	binary.Write(buf, binary.LittleEndian, uint32(dataSize))

	// Generate samples
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		val := sampleFunc(t)
		if val > 1.0 {
			val = 1.0
		} else if val < -1.0 {
			val = -1.0
		}
		sample16 := int16(val * 32767.0)
		binary.Write(buf, binary.LittleEndian, sample16)
	}

	return buf.Bytes()
}

// StartAmbientLoop launches background looped audio playback.
func (am *AmbientManager) StartAmbientLoop() {
	if am.playerBinary == "" {
		return
	}

	go am.runLoop("rain.wav", &am.RainActive, &am.rainCmd)
	go am.runLoop("fire.wav", &am.FireActive, &am.fireCmd)
	go am.runLoop("wind.wav", &am.WindActive, &am.windCmd)
}

func (am *AmbientManager) runLoop(filename string, activeFlag *bool, cmdVar **exec.Cmd) {
	filePath := filepath.Join(am.tempDir, filename)

	for {
		if *activeFlag && am.playerBinary != "" {
			var cmd *exec.Cmd
			if filepath.Base(am.playerBinary) == "afplay" {
				cmd = exec.Command(am.playerBinary, filePath)
			} else if filepath.Base(am.playerBinary) == "mpv" {
				cmd = exec.Command(am.playerBinary, "--loop", "--no-terminal", filePath)
			} else if filepath.Base(am.playerBinary) == "ffplay" {
				cmd = exec.Command(am.playerBinary, "-nodisp", "-loop", "0", filePath)
			}

			if cmd != nil {
				*cmdVar = cmd
				_ = cmd.Run()
			}
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ToggleRain toggles rain audio.
func (am *AmbientManager) ToggleRain() bool {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.RainActive = !am.RainActive
	if !am.RainActive && am.rainCmd != nil && am.rainCmd.Process != nil {
		_ = am.rainCmd.Process.Kill()
	}
	return am.RainActive
}

// ToggleFire toggles fireplace audio.
func (am *AmbientManager) ToggleFire() bool {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.FireActive = !am.FireActive
	if !am.FireActive && am.fireCmd != nil && am.fireCmd.Process != nil {
		_ = am.fireCmd.Process.Kill()
	}
	return am.FireActive
}

// ToggleWind toggles wind audio.
func (am *AmbientManager) ToggleWind() bool {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.WindActive = !am.WindActive
	if !am.WindActive && am.windCmd != nil && am.windCmd.Process != nil {
		_ = am.windCmd.Process.Kill()
	}
	return am.WindActive
}

// Cleanup stops players and removes temporary audio files.
func (am *AmbientManager) Cleanup() {
	if am.rainCmd != nil && am.rainCmd.Process != nil {
		_ = am.rainCmd.Process.Kill()
	}
	if am.fireCmd != nil && am.fireCmd.Process != nil {
		_ = am.fireCmd.Process.Kill()
	}
	if am.windCmd != nil && am.windCmd.Process != nil {
		_ = am.windCmd.Process.Kill()
	}
	if am.tempDir != "" {
		_ = os.RemoveAll(am.tempDir)
	}
}

// StatusSummary returns ambient sound toggles status string.
func (am *AmbientManager) StatusSummary() string {
	rainStr := "Rain: OFF"
	if am.RainActive {
		rainStr = "Rain: ON 🌧️"
	}
	fireStr := "Fire: OFF"
	if am.FireActive {
		fireStr = "Fire: ON 🔥"
	}
	windStr := "Wind: OFF"
	if am.WindActive {
		windStr = "Wind: ON 🍃"
	}

	playerStr := "Procedural Synth"
	if am.playerBinary != "" {
		playerStr = filepath.Base(am.playerBinary)
	}
	return fmt.Sprintf("%s | %s | %s [%s]", rainStr, fireStr, windStr, playerStr)
}
