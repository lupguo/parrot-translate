package entity

// TransEntity 翻译实体
type TransEntity struct {
	HashID     string // 原始翻译内容hash值，如果一致直接返回，不做http请求，减少调用量
	Text       string // 目标text
	Additional string // 附带信息，暂无用
}
