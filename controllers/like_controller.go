package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := articleLikesKey(articleID)
	// Update source of truth first, then invalidate caches.
	id, _ := strconv.Atoi(articleID)
	if err := global.Db.Model(&models.Article{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + 1")).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.RedisDB.Del(likeKey).Err(); err != nil {
		log.Printf("redis删除失败: %v", err)
	}
	if err := global.RedisDB.Del(articleListCacheKey).Err(); err != nil {
		log.Printf("redis删除失败: %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked artical"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := articleLikesKey(articleID)
	likes, err := global.RedisDB.Get(likeKey).Int()
	if err == redis.Nil {
		var artical models.Article
		if err := global.Db.First(&artical, articleID).Error; err == nil {
			likes = artical.Likes
			//将数据同步到redis中
			if err := global.RedisDB.Set(likeKey, likes, time.Minute*10).Err(); err != nil {
				log.Printf("redis设置失败:%v", err)
			}
		} else {
			likes = 0
		}
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
