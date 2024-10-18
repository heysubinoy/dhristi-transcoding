package ffmpeg

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"
)

func RunFFmpegCommand(inputFile string) error {
	// Parse the input URL
	parsedURL, err := url.Parse(inputFile)
	if err != nil {
		return fmt.Errorf("error parsing URL: %v", err)
	}

	// Extract the last segment of the URL path
	urlPath := parsedURL.Path
	segments := strings.Split(urlPath, "/")
	if len(segments) == 0 {
		return fmt.Errorf("no segments found in the URL path")
	}

	// Get the last segment to use as part of the output filename
	secretThing := segments[len(segments)-1]
	fmt.Println("Secret Thing:", secretThing)
	dirPath := fmt.Sprintf("./s3/stream/%s", secretThing)
	fmt.Println("Dir Path:", dirPath)
	cmd := exec.Command("mkdir", "-p", dirPath)
	cmd.Run()

	output := fmt.Sprintf("./s3/stream/%s/output.m3u8", secretThing)

	// Construct the FFmpeg command
	cmd = exec.Command("ffmpeg",
		"-i", inputFile,
		"-c:v", "libx264",
		"-c:a", "aac",
		"-strict", "experimental",
		"-b:a", "128k",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-f", "hls",
		output,
	)

	// Redirect FFmpeg output to log
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	// Run the command and handle errors
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ffmpeg: %v", err)
	}

	return nil
}
