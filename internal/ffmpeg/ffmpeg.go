package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
)

func StartStreaming(inputFile, rtmpURL string) (*exec.Cmd, error) {
	cmd := exec.Command("ffmpeg", "-stream_loop", "-1", "-i", inputFile, "-c:v", "libx264", "-preset", "veryfast", "-b:v", "3000k", "-bufsize", "6000k", "-pix_fmt", "yuv420p", "-g", "50", "-c:a", "aac", "-b:a", "128k", "-ar", "44100", "-f", "flv", rtmpURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("gagal memulai streaming: %v", err)
	}

	return cmd, nil
}

func StopStreaming(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return fmt.Errorf("tidak ada proses streaming yang berjalan")
	}

	err := cmd.Process.Kill()
	if err != nil {
		return fmt.Errorf("gagal menghentikan streaming: %v", err)
	}

	return nil
}
