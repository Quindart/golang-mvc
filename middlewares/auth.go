package middlewares

import (
	"example/golang-mvc/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	parts := strings.Fields(c.GetHeader("Authorization"))
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "A Bearer token is required."})
		return
	}
	userID, err := utils.ValidateToken(parts[1])
	if err != nil {
		// Abort ngăn các handler tiếp theo chạy khi token không hợp lệ.
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token."})
		return
	}

	// Lưu ID đã xác thực để handler dùng khi tạo event hoặc kiểm tra chủ sở hữu.
	c.Set("userId", userID)
	c.Next()
}
