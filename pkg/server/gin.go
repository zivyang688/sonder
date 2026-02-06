package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ziv/sonder/pkg/app"
)

var (
	_ app.IApp = (*GinApp)(nil)
)

type GinApp struct {
	app.BaseApp
	engine *gin.Engine
}

func NewGinApp() *GinApp {
	return &GinApp{
		BaseApp: app.BaseApp{
			Id:      "gin",
			Name:    "gin",
			Version: "1.0.0",
		},
		engine: gin.Default(),
	}
}

func (g *GinApp) Start() error {
	return g.engine.Run()
}

func (g *GinApp) Stop() error {
	return nil
}
