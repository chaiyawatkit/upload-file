package testhelper

import (
	"bytes"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"net/http/httptest"
)

func SetupMockDB() *gorm.DB {
	gomocket.Catcher.Register()

	db, err := gorm.Open(gomocket.DriverName, "")
	if err != nil {
		log.Fatalf("error mocking up the database %s", err)
	}

	db.LogMode(true)

	return db
}

func MakeStubContext(method string, url string, params string) (c *gin.Context) {
	const MIMEJSON = "application/json"

	body := bytes.NewBufferString(params)

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request, _ = http.NewRequest(method, url, body)
	context.Request.Header.Add("Content-Type", MIMEJSON)

	return context
}
