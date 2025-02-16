package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
)

func StartStreaming(inputFile, rtmpURL string, loopVideo bool, preset, resolution, fps, bitrate string) (*exec.Cmd, error) {
	// Konversi resolusi ke format ffmpeg (misalnya, "720p" -> "1280:720")
	var scale string
	switch resolution {
	case "480p":
		scale = "854:480"
	case "720p":
		scale = "1280:720"
	case "1080p":
		scale = "1920:1080"
	default:
		scale = "1280:720" // Default: 720p
	}

	// Buat perintah ffmpeg
	args := []string{
		"-i", inputFile,
		"-vf", fmt.Sprintf("scale=%s", scale),
		"-r", fps,
		"-c:v", "libx264",
		"-preset", preset,
		"-b:v", bitrate,
		"-bufsize", "4000k",
		"-pix_fmt", "yuv420p",
		"-g", "50",
		"-c:a", "aac",
		"-b:a", "128k",
		"-ar", "44100",
		"-f", "flv", rtmpURL,
	}

	// Tambahkan opsi looping jika dipilih
	if loopVideo {
		args = append([]string{"-stream_loop", "-1"}, args...)
	}

	// Buat perintah ffmpeg
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Jalankan perintah ffmpeg
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
