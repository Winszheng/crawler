package model

type Profile struct {
	Id string
	Url string
	Nickname string
	Content string
	BasicInfo  []string
	DetailInfo []string
	Selection  []string // 择偶条件
}
