package integrate

import (
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-homedir"
	"github.com/tddey01/switfs-auth/auth"
	"github.com/tddey01/switfs-auth/config"
	"github.com/tddey01/switfs-auth/log"
)

func setup(t *testing.T) (server *httptest.Server, dir string, token string) {
	tempDir := t.TempDir()
	log.Infof("create storage temp dir: %s", tempDir)

	cnf := config.DefaultConfig()
	dir, err := homedir.Expand(tempDir)
	if err != nil {
		t.Fatalf("could not expand repo location error:%s", err)
	} else {
		log.Infof("venus repo: %s", dir)
	}
	gin.SetMode(gin.DebugMode)
	dataPath := path.Join(dir, "data")

	app, err := auth.NewOAuthApp(dataPath, cnf.DB)
	if err != nil {
		t.Fatalf("Failed to init sophon-auth: %s", err)
	}
	token, err = app.GetDefaultAdminToken()
	if err != nil {
		t.Fatalf("Failed to get default admin token: %s", err)
	}

	router := auth.InitRouter(app)
	srv := httptest.NewServer(router)
	return srv, tempDir, token
}

func shutdown(t *testing.T, tempDir string) {
	log.Infof("shutdown, remove dir %s", tempDir)
	err := os.RemoveAll(tempDir)
	if err != nil {
		t.Fatal(err)
	}
}
