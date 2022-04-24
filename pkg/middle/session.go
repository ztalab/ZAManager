package middle

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Session(name string) gin.HandlerFunc {
	return sessions.Sessions(name, sessions.NewCookieStore([]byte("secret")))
}
