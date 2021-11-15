package repository

type ICacheRepos interface {
	// GetTransContentByHashID 通过hashID查询之前的翻译结果
	GetTransContentByHashID(hashID string) (text string, err error)

	// SetTransContent 设置翻译结果到缓存中
	SetTransContent(hashID string, content string) error
}
