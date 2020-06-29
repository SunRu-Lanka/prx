package authenticatPac

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sudesh35139/prx/config"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type user struct{
	Id int8
	UserName string
	Password string
	FirstName string
	LastName string
	Role string

}
type Session struct{
	Un string
	LastActivity time.Time
}
var Dbusers = map[string]user{}//user Id,user
var DbSessions = map[string]Session{}//session Id, session
var DbSessionCleaned time.Time
//const sessionLength int = 30


func SignUser(r *http.Request) (user,error) {

	u := user{}
	u.UserName =r.FormValue("username")
	u.Password = r.FormValue("password")
	u.FirstName = r.FormValue("firstname")
	u.LastName = r.FormValue("lastname")
	u.Role = r.FormValue("role")

	if u.UserName==""|| u.Password=="" || u.FirstName==""{
		return u,errors.New("406 bad request,fields cannot be null")
			}


	//checking user name already taken
			row := config.DB.QueryRow("SELECT username FROM signup WHERE username =?",u.UserName)
			err := row.Scan(&u.UserName)
			switch {
			case err == sql.ErrNoRows:
				bs,err1 := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)

				if err1 != nil{
					return u,errors.New("500 internal server error "+err.Error())
				}

				_ , er := config.DB.Exec("insert into signup (username,password,firstname,lastname,role) values(?,?,?,?,?)",u.UserName,bs,u.FirstName,u.LastName,u.Role)
				if er != nil{

					return u,errors.New("500 internal server error  unable to insert sign up user table" +er.Error())
				}
			case err != nil:
				return  u,errors.New("Server error Unabe to create user Account."+err.Error())
			}
			fmt.Println(err)

			return u,err
}

func AllUsersDesc()([]user,error){
	rows,err := config.DB.Query(" select username,firstname,lastname,role from signup order by id desc")

	if err != nil{

		return nil,err
	}
	defer  rows.Close()
	users := make([]user,0)
	for  rows.Next(){
		use := user{}
		err := rows.Scan(&use.UserName,&use.FirstName,&use.LastName,&use.Role)
		if err != nil{

			return nil,err
		}
		users=append(users,use)

	}
	if err = rows.Err();err != nil{

		return nil,err
	}
	return users,err
}
//page logn process
func LoginProess(req *http.Request)(user,error,){
	u := user{}

	userName := req.FormValue("username")
	password := req.FormValue("password")

	if userName ==""{
		return  u,errors.New("bad request")

	}

	row := config.DB.QueryRow("select username,password,role from signup where username= ?",userName)
	err := row.Scan(&u.UserName,&u.Password,&u.Role)

	//empty result
	if err != nil{
		fmt.Println("sql error on the result",err)
		return u,err
		}
		//comparing password with db and form values
		errPassword := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
		if errPassword  != nil{
			fmt.Println("Password dose not match")
			return u,errPassword
		}
		Dbusers[u.UserName]=u
		fmt.Println(u)
		return u,err
}

func OneUser(req * http.Request)(user,error) {
	u:= user{}
	userName := req.FormValue("username")
	if userName == ""{
		return u,errors.New("400 Bad request")
	}

	row := config.DB.QueryRow("select id,username,firstname,lastname,role from signup where username= ?",userName)
	err := row.Scan(&u.Id,&u.UserName,&u.FirstName,&u.LastName,&u.Role)
	if err != nil{
		return u,err
	}

	return u,nil
}
func UpdateUser(req *http.Request)(user,error){
u := user{}
fmt.Println("test2")
  un :=req.FormValue("id")
u.UserName = req.FormValue("username")
u.Password = req.FormValue("password")
u.FirstName =req.FormValue("firstname")
u.LastName = req.FormValue("lastname")
u.Role = req.FormValue("role")
if  u.Password == ""  {
	return u,errors.New("bad request 400. username password field can't be empty")
	}
	//convert id value string to int val
	unid,errn := strconv.ParseInt(un,0,8)
	if errn !=nil{
		return u,errors.New("unable to user id string value to int")
	}
	u.Id =int8(unid)

	//password convert to byte

	bs,err1 := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	fmt.Println("password ",bs)
	if err1 != nil{

		return u,errors.New("500 internal server error "+err1.Error())

	}


	_,err := config.DB.Exec("update signup set id=?,username=?,password =?,firstname=?,lastname=?,role=? where username=?",u.Id,u.UserName,bs,u.FirstName,u.LastName,u.Role,u.UserName)
	if err != nil {
		fmt.Println("error is query",err)

		return u,err

	}

	return u,nil
}
func DeleteUser(req *http.Request) error {

	DeleteUserName:= req.FormValue("username")


	if DeleteUserName == ""{
		return errors.New("bad request.No user name")
	}
	_,err :=config.DB.Exec("delete from signup where username=?",DeleteUserName)

	if err != nil{
			return errors.New("delete query dosent work")
	}
	return nil

}
//check username avelabilety on signup page
func checkUsernameDb(req *http.Request)(CheckStatus bool){
	u := user{}
	bs,eer :=ioutil.ReadAll(req.Body)
	if eer != nil{
		fmt.Println(eer)
	}
	userName := string(bs)
	row := config.DB.QueryRow("SELECT username FROM signup WHERE username = ?",userName)
	err := row.Scan(&u.UserName)
	if err == sql.ErrNoRows{
		CheckStatus = false
		return
	}else {
		CheckStatus = true
	}
	if err!=nil{
		fmt.Println("check row is empty",err)
	}
	return CheckStatus

}