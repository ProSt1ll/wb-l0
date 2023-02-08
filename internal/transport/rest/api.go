package rest

import (
	"github.com/ProSt1ll/wb-l0/internal/database"
	"github.com/ProSt1ll/wb-l0/internal/nats"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	g  *gin.Engine
	db database.Database
}

func New(db database.Database) Rest {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	return Rest{
		g:  router,
		db: db,
	}
}

func (r *Rest) Run() error {

	//middlewares settings
	r.g.Use(gin.Logger())
	r.g.Use(gin.ErrorLogger())
	r.g.Use(gin.Recovery())
	r.g.Any("/", r.HandlerHome)
	r.g.Any("/order/:uid", r.HandlerOrderByUID)
	r.g.LoadHTMLFiles("index.html")

	go nats.Subscribe(r.db)

	return r.g.Run("localhost:" + "8080")
}
