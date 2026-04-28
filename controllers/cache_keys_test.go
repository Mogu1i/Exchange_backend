package controllers

import "testing"

func TestArticleListCacheKey(t *testing.T) {
	if articleListCacheKey != "article" {
		t.Fatalf("articleListCacheKey mismatch: got %q", articleListCacheKey)
	}
}

func TestArticleLikesKey(t *testing.T) {
	got := articleLikesKey("123")
	want := "article:123:likes"
	if got != want {
		t.Fatalf("articleLikesKey mismatch: got %q want %q", got, want)
	}
}
