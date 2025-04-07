package routes

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishna102001/simplify-url/database"
)

func generateshorturl(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	// log.Println(hasher)
	shorturl := hex.EncodeToString(hasher.Sum(nil))
	// log.Println(shorturl)
	return shorturl[:8]
}

func CreateSimplifyUrl(c *gin.Context) {
	var data struct {
		Original_url string `json:"original_url"`
	}
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(400, gin.H{"success": true, "message": "Invalid request body"})
		return
	}
	// log.Println(data.Original_url)
	shorturl := generateshorturl(data.Original_url)
	var url database.Shorturl
	url.Original_url = data.Original_url
	url.Created_at = time.Now()
	url.Short_url = shorturl
	if err := database.DB.Model(&database.Shorturl{}).Create(&url).Error; err != nil {
		log.Println("failed to save the data")
		c.JSON(400, gin.H{"success": false, "message": "Failed to save data"})
	}
	c.JSON(200, gin.H{"short_url": shorturl})
}

func GetSimplifyUrl(c *gin.Context) {
	var short_url string = c.Param("id")
	var original_url struct {
		Original_url string
	}
	if err := database.DB.Model(&database.Shorturl{}).Select("original_url").Where("short_url = ?", short_url).Find(&original_url).Error; err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Failed to get original url"})
	}

	log.Println(original_url.Original_url)

	c.Redirect(302, original_url.Original_url)
}
