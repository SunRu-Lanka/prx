package repeitUsers

import (
	"encoding/json"
	"fmt"
	"github.com/sudesh35139/prx/config"
	"net/http"
)

func NumberOfRpUsers(w http.ResponseWriter,req * http.Request)  {
	if req.Method != "POST"{
		http.Error(w,http.StatusText(405),http.StatusNotAcceptable)
		return
	}
	torRpUsers,err := RpUsersProcess(req)
	if err!= nil{
		fmt.Println("error number 2")
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type","application/json")
	torRpUsersJson,err :=json.Marshal(torRpUsers)
	if err!= nil{
		fmt.Println("unable to convert json",err)
	}
	w.Write(torRpUsersJson)

}
func RpUsersVp(w http.ResponseWriter,req * http.Request)  {
	if req.Method != "GET"{
		http.Error(w,http.StatusText(405),http.StatusNotAcceptable)
		return
	}
	config.TPL.ExecuteTemplate(w,"rpUsersS.html",nil)

}
