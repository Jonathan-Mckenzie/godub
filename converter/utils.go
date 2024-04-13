package converter

import (
	"os/exec"
)

const (
	encoderProgram = "ffmpeg"
)

// GetEncoderName returns the name of the encoder program, just ffmpeg for now.
func GetEncoderName() string {
	return encoderProgram
}

// IsCommandAvailable returns true if the known encoder program is available from the system path.
func IsCommandAvailable() bool {
	// run a harmless "-version" to check if ffmpeg is available.
	// (Note: some systems don't have the program 'which' installed on path, 'which ffmpeg' will fail.)
	cmd := exec.Command(encoderProgram, "-version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
