package e

// @Author KHighness
// @Update 2022-02-13

var CodeMsgDic = map[int]string{
	SUCCESS                    :  "ok",
	ERROR                      :  "error",
	InvalidParams              :  "参数错误",

	ErrorExistUser             :  "用户已存在",
	ErrorNotExistUser          :  "用户不存在",
	ErrorFailEncryption        :  "加密失败",
	ErrorNotCompare            :  "不匹配",

	ErrorAuthCheckTokenFail    :  "Token签权失败",
	ErrorAuthCheckTokenExpire  :  "Token超时",
	ErrorAuthToken             :  "Token生成失败",
	ErrorAuth                  :  "Token错误",

	ErrorDatabase              :  "数据库操作出错,请重试",
}

// 根据状态码获取错误信息
func GetMsg(code int) string {
	msg, ok := CodeMsgDic[code]
	if ok {
		return msg
	}
	return CodeMsgDic[ERROR]
}