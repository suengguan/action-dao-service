package service

import (
	"fmt"
	"model"

	"dao-service/action-dao-service/dao"

	"github.com/astaxie/beego"
)

type ActionService struct {
}

func (this *ActionService) Create(action *model.Action) error {
	var err error
	var actionDao = dao.NewActionDao()

	err = actionDao.Create(action)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create action failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ActionService) GetAll(userId int64) ([]*model.Action, error) {
	var err error
	var actionDao = dao.NewActionDao()
	var actions []*model.Action

	// get actions
	beego.Debug("->get actions")
	actions, err = actionDao.GetByUserId(userId)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	for _, a := range actions {
		a.User = nil
	}

	beego.Debug("result:", actions)

	return actions, err
}
