package exutil

import (
	"fmt"
	"crypto/md5"
	"os/exec"
	"bytes"
	"regexp"
)

func Md5(str interface{}) string {
	data := []byte(fmt.Sprintf("%v", str))
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}

// return top domain
func MainHost(host string) string {
	pattern := `(?msU)([a-z0-9][a-z0-9\-]+?\.(?:com|cn|net|org|info|la|cc|co|gov|sx|gz|sh|sc|faith|date|space|top|ga|gq)(?:\.cn)?)$`
	regStr := regexp.MustCompile(pattern)
	matchStr := regStr.FindStringSubmatch(host)
	if len(matchStr) == 2 {
		return matchStr[1]
	}
	return ""
}

// IsEmpty
func IsEmpty(arg interface{}) bool {
	switch arg.(type) {
	case int:
		return If(arg.(int) == 0, true, false).(bool)
	case int64:
		return If(arg.(int64) == 0, true, false).(bool)
	case string:
		return If(arg.(string) == "", true, false).(bool)
	default:
		return true
	}
}

// If : ternary operator (三元运算)
// condition:比较运算
// trueVal:运算结果为真时的值
// falseVal:运算结果为假时的值
// return: 由于不知道传入值的类型, 所有, 必须在接收结果时, 指定对应的值类型
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
