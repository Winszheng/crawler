package model

type Profile struct {
	Nickname   string
	Content    string // 内心独白
	BasicInfo  []string
	DetailInfo []string
	Selection  []string // 择偶条件
}
