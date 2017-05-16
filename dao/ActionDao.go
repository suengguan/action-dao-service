package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

var cfg = beego.AppConfig

type ActionDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Action
}

func NewActionDao() *ActionDao {
	d := new(ActionDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *ActionDao) Create(action *model.Action) error {
	num, err := this.m_Orm.Insert(action)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *ActionDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no action to delete")
		return err
	}

	return err
}

// update
func (this *ActionDao) Update(action *model.Action) error {
	num, err := this.m_Orm.Update(action)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no action to update")
	}

	return err
}

// find
func (this *ActionDao) GetByUserId(userId int64) ([]*model.Action, error) {
	var actions []*model.Action

	beego.Debug("userid:", userId)
	//num, err := this.m_QuerySeter.Filter("USER_ID", userId).RelatedSel().All(&actions)

	//var action model.Action
	//err := this.m_QuerySeter.Filter("ID", userId).One(&action)
	//num, err := this.m_QuerySeter.Filter("DEV_TYPE", userId).All(&actions)
	//num, err := this.m_QuerySeter.Filter("User", userId).RelatedSel().All(&actions)
	num, err := this.m_QuerySeter.Filter("USER_ID", userId).RelatedSel().All(&actions)
	// o := orm.NewOrm()
	// o.Using("PME")
	// var t *model.Action
	// num, err := o.QueryTable(t).Filter("Usaer", 1).RelatedSel().All(&actions)
	//num, err := this.m_QuerySeter.Filter("USER_ID", userId).All(&actions)
	beego.Debug(num)
	beego.Debug(err)
	//beego.Debug(action)

	if err != nil {
		beego.Debug(num, err)
		return nil, err
	}

	return actions, err
}

func (this *ActionDao) GetById(Id int64) (*model.Action, error) {
	var action model.Action

	err := this.m_QuerySeter.Filter("ID", Id).One(&action)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &action, err
}
