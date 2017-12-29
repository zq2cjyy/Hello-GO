package dbopt

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Sleep struct {
	Id        int
	Begintime string
	Endtime   string
}

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1985@/babycare", 30)
	orm.RegisterModel(new(Sleep))
}

func OrmQuery() {
	o := orm.NewOrm()
	u := Sleep{Id: 1}
	o.Read(&u)
	fmt.Println(u)
}

func OrmInsert() {
	s := Sleep{}
	s.Begintime = time.Now().Format("2006-01-02 15:04:05")
	s.Endtime = time.Now().Add(1 * time.Hour).Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	id, err := o.Insert(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}
