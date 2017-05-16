package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	_ "dao-service/action-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	base_url = "http://localhost:8080/v1/action"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:corex123@tcp(117.122.198.151:3306)/PME?charset=utf8")

	o := orm.NewOrm()
	o.Using("PME")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM USER_T WHERE ID = ?", 1).Values(&maps)

	if err != nil {
		fmt.Println("get user failed!", err)
		return
	}

	if num == 0 {
		// create user
		_, err := o.Raw("insert into USER_T(ID) values(1)").Exec()
		if err != nil {
			fmt.Println("insert user failed!", err)
			return
		}
	} else if num == 1 {
		// user is existed, nothing todo
	} else {
		// error
		fmt.Println("get user failed!", err, num)
		return
	}
}

func Test_Create(t *testing.T) {
	var action model.Action
	action.Id = 0
	action.Time = time.Now().Unix()
	action.SessionId = "sessionId"
	action.DevType = 1
	action.Type = model.LOG_TYPE_INFO
	action.Content = "this is action test!"
	var user model.User
	user.Id = 1
	action.User = &user

	// post create action
	requestData, err := json.Marshal(&action)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err := http.Post(base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))
}

func Test_GetAll(t *testing.T) {
	res, err := http.Get(base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))
}
