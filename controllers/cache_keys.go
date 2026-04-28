package controllers

const (
	articleListCacheKey = "article"
)

func articleLikesKey(articleID string) string {
	return "article:" + articleID + ":likes"
}
