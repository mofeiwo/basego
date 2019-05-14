package conf

import (
	"os"
	"bufio"
	"strings"
	"io"
)

type ConfigIni struct {
	filepath string                         //your ini file path directory+file
	conflist []map[string]map[string]string //configuration information slice
}

func LoadConfig(filepath string) *ConfigIni {
	c := new(ConfigIni)
	c.filepath = filepath
	return c
}

func (c *ConfigIni) GetValue(mark, name string) string {
	c.ReadList()
	conf := c.ReadList()
	for _, v := range conf {
		for key, value := range v {
			if key == mark {
				return value[name]
			}
		}
	}
	return "no value"
}

//List all the configuration file
func (c *ConfigIni) ReadList() []map[string]map[string]string {

	file, err := os.Open(c.filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#": //增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1: len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1: len(line)])
			data[section][strings.TrimSpace(line[0:i])] = value
			if c.uniquappend(section) == true {
				c.conflist = append(c.conflist, data)
			}
		}

	}

	return c.conflist
}

func (c *ConfigIni) uniquappend(conf string) bool {
	for _, v := range c.conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
