package exdt

import "testing"

func TestFTPUploadFile(t *testing.T) {
	FTPUploadFile("server","use","pwd","localfile","remotedir/","178527.mp4")
}
