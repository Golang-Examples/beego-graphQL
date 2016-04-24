package models
import (
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/Pallinder/go-randomdata"
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


//Connect to database..
func init() {
	//orm.RegisterModel(new(AuthUser))
	orm.RegisterModelWithPrefix("snaphy_", new(AuthUser))
	// orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RegisterDriver("postgres", orm.DRPostgres)
	database, user, password := getDatabaseCredentials()
	connString := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, database)

	orm.RegisterDataBase("default", "postgres", connString )
	name := "default"
	force := false
	verbose := true
	orm.Debug = true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	before := makeTimestamp()
	for i := 0; i< 1000000; i++ {
		insertData()
	}

	after := makeTimestamp()
	diff := after - before
	beego.Info("Data added in ", diff)


}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}


func insertData(){
	auth_user := new(AuthUser)
	auth_user.First = randomdata.FirstName(randomdata.RandomGender)
	auth_user.Last = randomdata.LastName()
	auth_user.Email = randomdata.Email()
	auth_user.Password = randomdata.Digits(6)
	auth_user.Reg_key = randomdata.Digits(16)

	o := orm.NewOrm()
	o.Insert(auth_user)
}



//Get the database name, username and password info for postgresql.
func getDatabaseCredentials() (string, string, string){
	database := beego.AppConfig.String("postgres::database")
	user := beego.AppConfig.String("postgres::user")
	password := beego.AppConfig.String("postgres::password")
	return  database, user, password
}

