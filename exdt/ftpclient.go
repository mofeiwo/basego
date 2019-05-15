package exdt

import (
	ftp2 "github.com/jlaffaye/ftp"
	"fmt"
	"os"
)

func FTPUploadFile(ftpserver, ftpuser, ftppwd, localFile, remoteSavePath, saveName string) {
	ftp, err := ftp2.Connect(ftpserver)

	if err != nil {
		fmt.Println("connect error:", err)
		panic(err)
	}

	err = ftp.Login(ftpuser, ftppwd)

	if err != nil {
		fmt.Println("login error:",err)
		panic(err)
	}

	ftp.ChangeDir(remoteSavePath)

	file , err := os.Open(localFile)

	if err != nil {
		fmt.Println("open local file error:",err)
	}

	defer file.Close()

	err = ftp.Stor(saveName, file)

	if err != nil {
		fmt.Println("store error:",err)
	}

	ftp.Logout()

	ftp.Quit()

	fmt.Println("success upload file:", localFile)
}