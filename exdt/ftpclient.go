package exdt

import (
	"fmt"
	ftp2 "github.com/jlaffaye/ftp"
	"os"
)

func FTPUploadFile(ftpserver, ftpuser, ftppwd, localFile, remoteSavePath, saveName string) error {
	ftp, err := ftp2.Connect(ftpserver)

	if err != nil {
		fmt.Println("connect error:", err)
		return err
	}

	err = ftp.Login(ftpuser, ftppwd)

	if err != nil {
		fmt.Println("login error:", err)
		//panic(err)
		return err
	}

	ftp.ChangeDir(remoteSavePath)

	file, err := os.Open(localFile)

	if err != nil {
		fmt.Println("open local file error:", err)
		return err
	}

	defer file.Close()

	err = ftp.Stor(saveName, file)

	if err != nil {
		fmt.Println("store error:", err)
		//panic(err)
		return err
	}

	ftp.Logout()

	ftp.Quit()

	fmt.Println("success upload file:", localFile)
	return nil
}
