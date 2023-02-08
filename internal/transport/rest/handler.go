package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Rest) HandlerOrderByUID(c *gin.Context) {

	uid := c.Params.ByName("uid")

	r.g.LoadHTMLFiles("index.html")

	model, ok := r.db.Load(uid)
	if !ok {
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"order": "Not found",
		})
		return
	}
	js, _ := json.Marshal(model)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"order": string(js),
	})

}

func (r *Rest) HandlerHome(c *gin.Context) {
	r.g.LoadHTMLFiles("index.html")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"order": "main page",
	})
}
