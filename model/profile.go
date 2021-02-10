package model

type Profile struct{
	Nickname string
	Des string  // 内心独白
	Info []string   // 个人资料
	Hobby map[string]string // 兴趣爱好，暂定
	Selection []string  // 择偶条件
}
