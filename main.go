package main

import (
	"github.com/sudesh35139/prx/config"
	"github.com/sudesh35139/prx/dailyUsers"
	"github.com/sudesh35139/prx/dwellTime"
	"github.com/sudesh35139/prx/repeitUsers"
	"net/http"
)

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
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	//http.Handle("/public/",http.StripPrefix("/public",http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":4000", nil)

}


func index(w http.ResponseWriter, req *http.Request){
	config.TPL.ExecuteTemplate(w,"index.html",nil)


}


