package e

// @Author KHighness
// @Update 2022-02-13

const (
	SUCCESS                    = 200
	ERROR                      = 500
	InvalidParams              = 400

	ErrorExistUser             = 10001
	ErrorNotExistUser          = 10002
	ErrorFailEncryption        = 10006
	ErrorNotCompare            = 10007

	ErrorAuthCheckTokenFail    = 30001
	ErrorAuthCheckTokenExpire  = 30002
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004

	ErrorDatabase              = 40001
)
