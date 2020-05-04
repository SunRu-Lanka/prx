package dailyUsers

import (
	"encoding/json"
	"fmt"
	"github.com/sudesh35139/prx/config"
	"net/http"
)

func NumberOfUsersr(w http.ResponseWriter,r *http.Request)  {
	if r.Method !="GET"{
		fmt.Println("error number 1")
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	totUser,err := TotDailyUser(r)
	if err!= nil{
		fmt.Println("error number 2")
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	toUserJson,err := json.Marshal(totUser)
	if err != nil{
			fmt.Println("unable to convert json",err)
	}
	fmt.Println("convert to json",toUserJson)
	w.Write(toUserJson)
	//fmt.Println("total users",totUser)
	//config.TPL.ExecuteTemplate(w,"dailyUser.html",totUser)

}
func DailyUsersV(w http.ResponseWriter,req * http.Request)  {
	if req.Method != "GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return

	}
	config.TPL.ExecuteTemplate(w,"dailyUserTot.html",nil)

}
func DailyUserProcess(w http.ResponseWriter,req * http.Request)  {
	if req.Method != "POST"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	_, err := TotDailyUser(req)

	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return

	}

	config.TPL.ExecuteTemplate(w,"dailyUserTot.html",nil)


}

