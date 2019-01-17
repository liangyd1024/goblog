//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/20

package constant

const (
	//草稿状态
	BOWEN_STATUS_INIT = "INIT"
	//已发布状态
	BOWEN_STATUS_PUBLISH = "PUBLISH"
	//回收站状态
	BOWEN_STATUS_DISCARD = "DISCARD"

	//博文类型-原创
	BOWEN_TYPE_ORIGIN = "ORIGIN"
	//博文类型-转载
	BOWEN_TYPE_ORIGINAL = "ORIGINAL"

	//博文内容-富文本
	BOWEN_CONTENT_RICH = "RICH"
	//博文内容-markdown
	BOWEN_CONTENT_MD = "MD"
)

var constMap = make(map[string]string, 10)

func init() {
	constMap[BOWEN_STATUS_INIT] = "草稿"
	constMap[BOWEN_STATUS_PUBLISH] = "发布"
	constMap[BOWEN_STATUS_DISCARD] = "回收站"
}

func GetValue(key string) string{
	return constMap[key]
}
