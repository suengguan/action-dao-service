package test

import (
	"action-dao-service/models"
	_ "action-dao-service/routers"
	"encoding/json"
	"model"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	err := model.InitEnv()
	beego.Debug(err)
}

// TestCreate is a sample to run an endpoint test
func Test_Create(t *testing.T) {
	// err := model.InitEnv()
	// t.Log(err)

	var action models.Action
	action.Id = 0
	action.Time = time.Now().Unix()
	action.SessionId = "1234567890"
	action.DevType = 1
	action.Content = "this is action-dao-service test"
	var user models.User
	user.Id = 1
	action.User = &user

	result, err := json.Marshal(&action)
	if err == nil {
		beego.Debug(err)
	}

	r, _ := http.NewRequest("POST", "/v1/action", strings.NewReader(string(result)))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// Test_GetAll is a sample to run an endpoint test
func Test_GetAll(t *testing.T) {
	// err := model.InitEnv()
	// t.Log(err)

	r, _ := http.NewRequest("GET", "/v1/action/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
