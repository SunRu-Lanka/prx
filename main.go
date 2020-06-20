package main

import (
	"github.com/sudesh35139/prx/authenticatPac"
	"github.com/sudesh35139/prx/config"
	"github.com/sudesh35139/prx/dailyUsers"
	"github.com/sudesh35139/prx/dwellTime"
	"github.com/sudesh35139/prx/repeitUsers"
	"net/http"
)
type User struct {
	UserName string
	Password []byte
	FirstName string
	LastName string
	CompanyName string

}

func main()  {

	http.HandleFunc("/",index)
	http.Handle("/favicon.icon",http.NotFoundHandler())
	http.HandleFunc("/dailyUsers",dailyUsers.NumberOfUsersr)
	http.HandleFunc("/dailyUsersV",dailyUsers.DailyUsersV)
	http.HandleFunc("/dailyUsersV/process",dailyUsers.DailyUserProcess)
	http.HandleFunc("/avdDTimeV",dwellTime.DwTView)
	http.HandleFunc("/avgDTime",dwellTime.AvgDwTime)
	http.HandleFunc("/rpUsers", repeitUsers.NumberOfRpUsers)
	http.HandleFunc("/rpUsersV",repeitUsers.RpUsersVp)
	http.HandleFunc("/usersDetail",usersDetails)
	http.HandleFunc("/signup",authenticatPac.Signup)
	http.HandleFunc("/authenticatePac/signup/process",authenticatPac.SignupProcess)
	http.HandleFunc("/alreadytaken",authenticatPac.UserAlreadyTaken)
	http.HandleFunc("/userinfo",authenticatPac.AllUsersDetails)
	http.HandleFunc("/authenticate/update",authenticatPac.Update)
	http.HandleFunc("/authenticate/update/updateprocess",authenticatPac.UpdateProcess)
	http.HandleFunc("/authenticate/delete/deleterprocess",authenticatPac.DeleteProcess)
	http.HandleFunc("/login",authenticatPac.Login)
	http.HandleFunc("/logout",authenticatPac.Authorized1(authenticatPac.Logout))
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	//http.Handle("/public/",http.StripPrefix("/public",http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":4000", nil)

}


func index(w http.ResponseWriter, req *http.Request){
	config.TPL.ExecuteTemplate(w,"index.html",nil)


}
func usersDetails(w http.ResponseWriter,req * http.Request){
	Us := authenticatPac.GetUser(w,req)
	if !authenticatPac.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if !((Us.Role=="user")||(Us.Role=="admin")){

		http.Error(w,"You are not in appropriate account",http.StatusForbidden)
		return
	}

	config.TPL.ExecuteTemplate(w,"usersDetails.html",nil)
}
