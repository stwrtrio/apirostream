# LoopSphere

LoopSphere adalah aplikasi untuk melakukan live streaming ke YouTube dengan opsi looping video.

## Fitur
- Opsi looping video selama live streaming.
- Stop streaming otomatis setelah durasi tertentu.
- Konfigurasi bitrate, resolusi, dan FPS.
- Pilihan preset encoding untuk mengoptimalkan performa.

## Cara Penggunaan
1. **Download Aplikasi**:
   - Download executable aplikasi untuk platform Anda (Windows, macOS, atau Linux) dari [Releases](https://github.com/<username-anda>/loopshpere/releases).

2. **Jalankan Aplikasi**:
   - Buka terminal atau command prompt.
   - Navigasikan ke folder tempat aplikasi berada.
   - Jalankan aplikasi dengan perintah:
     ```bash
     ./loopshpere
     ```
   - (Untuk Windows, jalankan `loopshpere.exe`).

3. **Masukkan RTMP URL dan Stream Key**:
   - Aplikasi akan meminta Anda untuk memasukkan RTMP URL (default: `rtmp://a.rtmp.youtube.com/live2`). Tekan `Enter` untuk menggunakan default.
   - Masukkan stream key yang didapatkan dari YouTube Studio.

4. **Pilih File Video**:
   - Masukkan path ke file video yang ingin Anda streaming (misalnya, `input.mp4`).

5. **Pilih Opsi Looping**:
   - Pilih apakah video akan di-looping (`Ya` atau `Tidak`).

6. **Pilih Konfigurasi Streaming**:
   - **Preset Encoding**: Pilih preset encoding dari opsi yang tersedia (`ultrafast`, `superfast`, `veryfast`, dll.).
   - **Resolusi**: Pilih resolusi dari opsi yang tersedia (`480p`, `720p`, `1080p`). Tekan `Enter` untuk menggunakan default (`720p`).
   - **FPS**: Pilih FPS dari opsi yang tersedia (`24`, `30`, `60`). Tekan `Enter` untuk menggunakan default (`60`).
   - **Bitrate**: Masukkan bitrate (misalnya, `2000k`). Tekan `Enter` untuk menggunakan default (`2000k`).

7. **Mulai Streaming**:
   - Aplikasi akan memulai streaming ke YouTube dengan opsi looping video jika dipilih.

8. **Stop Streaming**:
   - Untuk saat ini, streaming bisa dihentikan secara manual dengan menekan `Ctrl+C`. Next update akan ada fitur dimana user bisa menentukan durasi streaming.

## Dependensi
- Golang
- ffmpeg

## Kontribusi
Silakan buat issue atau pull request jika Anda ingin berkontribusi.

## Lisensi
Proyek ini dilisensikan di bawah [MIT License](LICENSE).