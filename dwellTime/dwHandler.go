package dwellTime

import (
	"encoding/json"
	"fmt"
	"github.com/sudesh35139/prx/config"
	"net/http"
)

func DwTView(w http.ResponseWriter,req * http.Request){
	if req.Method != "GET"{
		http.Error(w,http.StatusText(405),http.StatusNotAcceptable)
		return
	}
	config.TPL.ExecuteTemplate(w,"avgDwellTime.html",nil)

}

func AvgDwTime(w http.ResponseWriter,req * http.Request) {
	if req.Method != "POST"{
		http.Error(w,http.StatusText(405),http.StatusNotAcceptable)
		return
	}
	totDwtime,err := DwTimeProcess(req)
	if err!= nil{
		fmt.Println("error number 2")
		http.Error(w,http.StatusText(500),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	toDwtimeJson,err := json.Marshal(totDwtime)
	if err != nil{
		fmt.Println("unable to convert json",err)
	}
	fmt.Println("convert to json",toDwtimeJson)
	w.Write(toDwtimeJson)


}
