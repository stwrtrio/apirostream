package main

import (
	"fmt"
	"log"

	"github.com/stwrtrio/loopsphere/internal/ffmpeg"
	"github.com/stwrtrio/loopsphere/internal/utils"
)

func main() {
	fmt.Println("LoopSphere: Aplikasi Live Streaming dengan Looping Video")

	// Input RTMP URL (dengan default value)
	var rtmpURL string
	fmt.Print("Masukkan RTMP URL (default: rtmp://a.rtmp.youtube.com/live2): ")
	fmt.Scanln(&rtmpURL)
	if rtmpURL == "" {
		rtmpURL = "rtmp://a.rtmp.youtube.com/live2" // Default value
	}

	// Input stream key
	var streamKey string
	fmt.Print("Masukkan Stream Key: ")
	fmt.Scanln(&streamKey)

	// Gabungkan RTMP URL dan stream key
	fullRTMPURL := fmt.Sprintf("%s/%s", rtmpURL, streamKey)

	// Input file video
	var inputFile string
	fmt.Print("Masukkan path ke file video (misalnya, input.mp4): ")
	fmt.Scanln(&inputFile)

	// Pilih apakah video akan di-looping
	loopOptions := []string{"Ya", "Tidak"}
	loopChoice := utils.SelectOption("Apakah video akan di-looping?", loopOptions)
	loopVideo := loopChoice == "Ya"

	// Pilih preset encoding
	presetOptions := []string{"ultrafast", "superfast", "veryfast", "faster", "fast"}
	preset := utils.SelectOption("Pilih preset encoding:", presetOptions)

	// Pilih resolusi
	resolutionOptions := []string{"480p", "720p", "1080p"}
	resolution := utils.SelectOption("Pilih resolusi (default: 720p):", resolutionOptions)

	// Pilih FPS
	fpsOptions := []string{"24", "30", "60"}
	fps := utils.SelectOption("Pilih FPS (default: 60):", fpsOptions)
	if fps == "" {
		fps = "60" // Default value
	}

	// Input bitrate
	var bitrate string
	fmt.Print("Masukkan bitrate (misalnya, 2000k): ")
	fmt.Scanln(&bitrate)
	if bitrate == "" {
		bitrate = "2000k" // Default value
	}

	// Mulai streaming
	cmd, err := ffmpeg.StartStreaming(inputFile, fullRTMPURL, loopVideo, preset, resolution, fps, bitrate)
	if err != nil {
		log.Fatalf("Gagal memulai streaming: %v", err)
	}

	fmt.Println("Streaming berhasil dimulai! Tekan Ctrl+C untuk menghentikan.")

	// Tunggu hingga proses streaming selesai (misalnya, dihentikan manual)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Proses streaming berakhir dengan error: %v", err)
	}
}
