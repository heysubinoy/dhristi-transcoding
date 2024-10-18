package main

import (
	ffmpeg "dhristi-transcoding/ffmpeg"
	"dhristi-transcoding/utils"
	"log"
	"os"
)

func main() {
	rtmpURL := os.Getenv("RTMP_URL")
	if rtmpURL == "" {
		log.Fatal("RTMP_URL environment variable is required")
	}

	bucketName := "dhristi-bucket"
	mountPoint := "./s3"

	if err := utils.MountS3(bucketName, mountPoint); err != nil {
		log.Fatalf("Failed to mount S3 bucket: %v", err)
	}

	if err := ffmpeg.RunFFmpegCommand(rtmpURL); err != nil {
		log.Fatalf("Failed to run FFmpeg command: %v", err)
	}

	log.Println("Transcoding completed successfully.")
}
