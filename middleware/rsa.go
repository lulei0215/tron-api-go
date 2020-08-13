package middleware



import (
"encoding/base64"
"ethereum/config"
"ethereum/handles"
"net/http"

"github.com/gin-gonic/gin"
)

// Cipher 接受密文
type Cipher struct {
	Data string `json:"data"`
}

// Rsa 中间件
func Rsa() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _data Cipher
		if err := c.ShouldBind(&_data); err != nil {
			handles.Error(http.StatusBadRequest, "获取data出错", c)
			return
		}
		// 密文
		ciphertext, err := base64.StdEncoding.DecodeString(_data.Data)
		if err != nil {
			handles.Error(http.StatusBadRequest, "base64解密出错", c)
			return
		}
		data, err := config.RsaDecrypt(ciphertext)
		if err != nil {
			handles.Error(http.StatusBadRequest, "RSA解密出错", c)
			return
		}
		c.Set("data", data)
		c.Next()
	}
}
