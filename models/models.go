package models
import (
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
	"time"
)
type AuthUser struct {
	Id          int
	First       string
	Last        string
	Email       string
	Password    string
	Reg_key     string
	Reg_date    time.Time `orm:"auto_now_add;type(datetime)"`
}



func init() {
	orm.RegisterModel(new(AuthUser))
	//orm.RegisterModelWithPrefix("snaphy_", new(User))
	//// orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	//orm.RegisterDriver("postgres", orm.DRPostgres)
	//orm.RegisterDataBase("default", "postgres", "postgres://robins:myfunzone2030@localhost/robins?sslmode=disable")
	//name := "default"
	//force := false
	//verbose := true
	//orm.Debug = true
	//err := orm.RunSyncdb(name, force, verbose)
	//if err != nil {
	//	fmt.Println(err)
	//}

}



