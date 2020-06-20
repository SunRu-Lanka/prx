package authenticatPac

import (
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/sudesh35139/prx/config"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter,req *http.Request){
	u := GetUser(w,req)
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin"{
		http.Error(w,"You are not in appropriate account",http.StatusForbidden)
		return
	}


	config.TPL.ExecuteTemplate(w,"signup.html",nil)

}
func SignupProcess(w http.ResponseWriter,req *http.Request){
	u := GetUser(w,req)
	//if !AlreadyLoggedIn(w, req) {
	//	http.Redirect(w, req, "/", http.StatusSeeOther)
	//	return
	//}
	if u.Role != "admin"{
		http.Error(w,"You are not in appropriate account",http.StatusForbidden)
		return
	}


	if req.Method != "POST"{
		http.Error(w, http.StatusText(405),http.StatusMethodNotAllowed)

	}

	_,err := SignUser(req)
	if err !=nil{
		users,err := AllUsersDesc()
		if err!=nil{
			fmt.Println("passing value to dt",err)
			return
		}
		fmt.Println(users)

		config.TPL.ExecuteTemplate(w,"usersinfor.html",users)


	}
	http.Redirect(w,req,"/alreadytaken",http.StatusSeeOther)
	//config.TPL.ExecuteTemplate(w,"alreadytaken.html",user)

}
func UserAlreadyTaken(w http.ResponseWriter,req *http.Request)  {
	u := GetUser(w,req)
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin"{
		http.Error(w,"You are not in appropriate account",http.StatusForbidden)
		return
	}
	config.TPL.ExecuteTemplate(w,"alreadytaken.html",nil)
}
func Login(w http.ResponseWriter,req *http.Request){
	if AlreadyLoggedIn(w,req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost{
		u,err := LoginProess(req)

		if err == sql.ErrNoRows{
			http.Error(w,"username or password do not  match",http.StatusForbidden)
			//http.Redirect(w,req,"/",http.StatusSeeOther)
			return

		}
		if err != nil{
			http.Error(w,"username or password do not match" ,http.StatusForbidden)
			return
			//http.Redirect(w,req,"/",http.StatusSeeOther)
		}else {
			sId,_ := uuid.NewV4()
			c := &http.Cookie{
				Name:       "proximity",
				Value:      sId.String(),

			}
			c.MaxAge = sessionLength
			http.SetCookie(w,c)
			DbSessions[c.Value] = Session{u.UserName,time.Now()}
			http.Redirect(w,req,"/usersDetail",http.StatusSeeOther)
		}

	}
	showSessions()
	config.TPL.ExecuteTemplate(w,"/",nil)

}
func Logout(w http.ResponseWriter,req * http.Request)  {
	c,_ := req.Cookie("proximity")
	//delete the session
	delete(DbSessions,c.Value)
	//remove the cookie
	c=&http.Cookie{
		Name:       "proximity",
		Value:      "",
		RawExpires: "",
		MaxAge:     -1,

	}
	//clean up dbsessions
	if time.Now().Sub(DbSessionCleaned)>(time.Second *30){
		go cleanSessions()
	}
	http.Redirect(w,req,"/",http.StatusSeeOther)


}
func AllUsersDetails(w http.ResponseWriter,req * http.Request)  {
	u := GetUser(w,req)
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/usersDetail", http.StatusSeeOther)
		return
	}
	if u.Role != "admin"{
		http.Error(w,"You are not in appropriate account",http.StatusForbidden)
		return
	}

	if req.Method != "GET"{
		http.Error(w, http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	users,err := AllUsersDesc()
	if err != nil{
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return

	}
	config.TPL.ExecuteTemplate(w,"usersinfor.html",users)

}
//authorized function for handle logout

func Authorized1(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !AlreadyLoggedIn(w, r) {
			//http.Error(w, "not logged in", http.StatusUnauthorized)
			http.Redirect(w, r, "/usersDetail", http.StatusSeeOther)
			return // don't call original handler
		}
		h.ServeHTTP(w, r)
	})
}
func Update(w http.ResponseWriter,req *http.Request){
	u :=GetUser(w,req)
	if !AlreadyLoggedIn(w, req) {
		fmt.Println("login in error1")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin"{
		http.Error(w,"You are not in appropriate account",http.StatusForbidden)
		return
	}

	if req.Method !="GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("test1")
	up,err := OneUser(req)
	switch{
	case err == sql.ErrNoRows:
		http.NotFound(w,req)
		return
	case err != nil:
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		}
		fmt.Println("update value",up)
		config.TPL.ExecuteTemplate(w,"updateuser.html",up)
	}
func UpdateProcess(w http.ResponseWriter,req *http.Request){
	u := GetUser(w,req)
	//if !AlreadyLoggedIn(w,req){
	//	fmt.Println("log worng 2")
	//	http.Redirect(w,req,"/",http.StatusSeeOther)
	//	return
	//}
	if u.Role !="admin"{
		http.Error(w,"You are not in appropriate account ",http.StatusForbidden)
	}
	if req.Method !="POST"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	_,err := UpdateUser(req)
	if err != nil {
		fmt.Println("test1")
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
	}
	http.Redirect(w,req,"/userinfo",http.StatusSeeOther)
}
func DeleteProcess(w http.ResponseWriter,req *http.Request){
	fmt.Println("test4444")
	u:= GetUser(w,req)
	//fmt.Println("delete user details",U)
	//if !AlreadyLoggedIn(w,req){
	//	fmt.Println("user allready log")
	//	http.Redirect(w,req,"/",http.StatusSeeOther)
	//	return
	//}
	if u.Role != "admin"{
		http.Error(w,"Your are not in appoprate account",http.StatusSeeOther)
	}
	if req.Method !="GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("test11")
	err := DeleteUser(req)
	if err !=nil{

		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	http.Redirect(w,req,"/userinfo",http.StatusSeeOther)

}