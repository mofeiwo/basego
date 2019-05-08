package exutil

import (
	"io/ioutil"
	"os"
	"io"
	"fmt"
	"bufio"
	"strings"
	"github.com/pkg/errors"
)

// Return all content at once
func ReadAllContent(filename string) string {
	buf, _ := ioutil.ReadFile(filename)
	return string(buf)
}

// Write all content at once
func WriteContent(filename, content string) {
	var d1 = []byte(content)
	ioutil.WriteFile(filename, d1, 0666) //写入文件(字节数组)
}

// Check if the file exists
func WriteFileByRow(filename, message string) {
	var f *os.File
	var err1 error
	if CheckFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777) //打开文件
	} else {
		f, err1 = os.Create(filename) //创建文件
	}
	CheckErr(err1)
	_, err1 = io.WriteString(f, fmt.Sprintf("%s\n",message)) //写入文件(字符串)
	CheckErr(err1)
}


// Check if the file exists
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// Return bytes
func ReadContentReturnBytes(filename string) []byte {
	buf, err := ioutil.ReadFile(filename)
	CheckErr(err)
	return buf
}

// Return row
func GetFileContentByRow(filename string) ([]string,error) {

	if !CheckFileIsExist(filename) {
		CheckErr(errors.New("file does not exist"))
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil,err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result []string
	for {
		line,_, err := buf.ReadLine()
		lineinfo := strings.TrimSpace(string(line))
		if err != nil {
			if err == io.EOF {   //读取结束，会报EOF
				return result,nil
			}
			return nil,err
		}
		result = append(result,lineinfo)
	}
	return result,nil
}