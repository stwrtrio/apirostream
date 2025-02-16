# Apirostream

<div align="center">
   <img src="https://github.com/stwrtrio/apirostream/blob/main/Assets/images/apirostream-logo.png" alt="" border="0" width="200">
</div>   
<br>
<br>

Apirostream adalah aplikasi live streaming yang memungkinkan pengguna melakukan streaming ke YouTube dengan fitur looping video, konfigurasi bitrate, resolusi, dan FPS, serta dukungan untuk platform Windows, macOS, dan Linux.

## Fitur
- Opsi looping video selama live streaming.
- Stop streaming otomatis setelah durasi tertentu.
- Konfigurasi bitrate, resolusi, dan FPS.
- Pilihan preset encoding untuk mengoptimalkan performa.

## Cara Penggunaan

### Instal Dependensi ffmpeg
Aplikasi ini memerlukan `ffmpeg` untuk melakukan streaming. Pastikan `ffmpeg` sudah terinstall di sistem Anda.

#### Panduan instalasi ffmpeg
Anda bisa mengikuti panduan instalasi pada website berikut: [https://www.hostinger.com/tutorials/how-to-install-ffmpeg](https://www.hostinger.com/tutorials/how-to-install-ffmpeg).<br>
sesuaikan dengan OS yang anda gunakan.

### Instal aplikasi

1. **Download Aplikasi**:
   - Download executable aplikasi untuk platform Anda (Windows, macOS, atau Linux) dari [Releases](https://github.com/stwrtrio/apirostream/releases).

2. **Jalankan Aplikasi**:
   - Buka terminal atau command prompt.
   - Navigasikan ke folder tempat aplikasi berada.
   - Jalankan aplikasi dengan perintah:
     ```bash
     ./apirostream
     ```
   - (Untuk Windows, jalankan `apirostream.exe`).

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

7. **Pilih Opsi Timer**:
   - Pilih apakah timer akan diaktifkan (`Ya` atau `Tidak`).
   - Jika timer diaktifkan, masukkan durasi streaming dalam menit.

8. **Mulai Streaming**:
   - Aplikasi akan memulai streaming ke YouTube dengan opsi looping video jika dipilih.
   - Jika timer diaktifkan, streaming akan berhenti secara otomatis setelah durasi yang ditentukan.

9. **Stop Streaming**:
    - Jika timer tidak diaktifkan, streaming bisa dihentikan secara manual dengan menekan `Ctrl+C`.

## Dependensi
- Golang
- ffmpeg

## Kontribusi
Silakan buat issue atau pull request jika Anda ingin berkontribusi.

