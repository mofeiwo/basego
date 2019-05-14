package exutil

import (
	"testing"
	"fmt"
)

// 接口返回的数据信息
type ResponseVideoInfo struct {
	Success bool
	Data struct{
		Sort_key int
		Task_id,
		Group_id,
		Keyword,
		Video_dir string

	}
	Code int
	Msg string
}

func TestDecodeUnMarshal(t *testing.T) {
	var res ResponseVideoInfo
	DecodeUnMarshal([]byte("{\"success\":true,\"data\":{\"sort_key\":3,\"task_id\":\"899\",\"group_id\":\"109856\",\"keyword\":\"乳腺癌复发怎么办\",\"video_dir\":\"mingyi/baidu/20190328/899_20190505_gongjianping/\"},\"code\":6000000,\"msg\":\"操作成功\"}"),&res)
	fmt.Println(res.Data)
}
