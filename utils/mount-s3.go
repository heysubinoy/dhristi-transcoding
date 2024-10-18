package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// MountS3 mounts the specified S3 bucket to the given mount point using s3fs.
func MountS3(bucketName, mountPoint string) error {
	// Create a buffer to capture standard output and standard error
	var outBuf, errBuf bytes.Buffer

	// Create the command with stdout and stderr redirection
	cmd := exec.Command("s3fs", bucketName, mountPoint, "-o", "passwd_file=.passwd-s3fs")
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	// Run the command
	err := cmd.Run()
	if err != nil {
		// Log the error along with the captured stderr output
		return fmt.Errorf("failed to mount S3 bucket: %v. Output: %s. Error: %s", err, outBuf.String(), errBuf.String())
	}

	// Log the standard output
	log.Println("Mount S3 Output:", outBuf.String())
	return nil
}
