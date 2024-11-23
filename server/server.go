package main

import (
    "net/http"
    "os/exec"
    "log"
)

func streamVideo(w http.ResponseWriter, r *http.Request){
    cmd := exec.Command("ffmpeg",
    "-re",
    "-i", "chungus.mp4",        // Input video file
    "-c:v", "libx264",        // Encode video using H.264
    "-preset", "ultrafast",   // Encoding speed/quality trade-off
    "-f", "mpegts",           // Output format
    "pipe:1")

    cmd.Stdout = w
    cmd.Stderr = log.Writer() // Log FFmpeg errors
    w.Header().Set("Content-Type", "video/mp2t") // Set MIME type for video

    if err := cmd.Run(); err != nil {
        log.Println("Error running FFmpeg:", err)
    }
}

func main(){
    http.HandleFunc("/stream", streamVideo)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
