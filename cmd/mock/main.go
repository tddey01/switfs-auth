package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-homedir"
	"github.com/tddey01/switfs-auth/auth"
	"github.com/tddey01/switfs-auth/config"
	"github.com/tddey01/switfs-auth/log"
)

func main() {
	absoluteTmp := "~/.switfs-auth"
	dir, err := homedir.Expand(absoluteTmp)
	if err != nil {
		log.Printf("could not expand repo location error:%s", err)
	} else {
		log.Printf("venus repo: %s", dir)
	}

	gin.SetMode(gin.DebugMode)

	cnfPath := path.Join(dir, "config.toml")

	dataPath := path.Join(dir, "data")
	cnf, err := config.DecodeConfig(cnfPath)
	if err != nil {
		return
	}
	log.InitLog(cnf.Log)
	app, err := auth.NewOAuthApp(dataPath, cnf.DB)
	if err != nil {
		log.Fatalf("Failed to init sophon-auth: %s", err)
	}
	router := auth.InitRouter(app)
	server := &http.Server{
		Addr:         "127.0.0.1:8989",
		Handler:      router,
		ReadTimeout:  cnf.ReadTimeout,
		WriteTimeout: cnf.WriteTimeout,
		IdleTimeout:  cnf.IdleTimeout,
	}
	log.Infof("server start and listen on %s", cnf.Listen)
	server.ListenAndServe() //nolint
	select {}
}
