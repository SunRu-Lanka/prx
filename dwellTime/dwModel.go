package dwellTime

import (
	"fmt"
	"github.com/sudesh35139/prx/config"
	"io/ioutil"
	"net/http"
	"strings"
)

type DwellTimeSt struct {
    CpaturedDate string `json:"cpatured_date"`
	Hours int
	minutes float32
	CapturedTimeDw float32 `json:"captured_time_dw"`
}

func DwTimeProcess(r *http.Request) ([]DwellTimeSt, error) {

	req ,err :=ioutil.ReadAll(r.Body)

	if err !=nil{
		fmt.Println("nill data",err)
	}
	reqs := string(req)
	fmt.Println("check value",reqs)
	if len (reqs) == 0{
		fmt.Println("empty string")

	}
	spiletDate := strings.Split(reqs,"&")
	fromDateStartV:= spiletDate[0]
	fromdateCha:= fromDateStartV[8:]
	fmt.Println("Spilt date ###",fromdateCha)
	toDateEndV := spiletDate[1]
	toDateCha := toDateEndV[6:]

	fmt.Println("spilt End date",toDateCha)

	dwTime := DwellTimeSt{}
	rows,err := config.DB.Query("SELECT date(capturedtime) as CpaturedDate,hour(timediff(max(capturedtime), min(capturedtime))) as hours , minute(timediff(max(capturedtime), min(capturedtime)))/60 as minutes,(hour(timediff(max(capturedtime), min(capturedtime))) + minute(timediff(max(capturedtime), min(capturedtime)))/60) as CapturedTimeDw	FROM proximity.`6c-3b-6b-68-e4-f3`where capturedtime between '"+fromdateCha+"' and '"+toDateCha+"' group by date(capturedtime)")
	if err != nil{

		fmt.Println("Empty row from number of daily users",err)
	}
	defer  rows.Close()
	dwellTimeSts := make([]DwellTimeSt,0)

	for  rows.Next(){

		err := rows.Scan(&dwTime.CpaturedDate,&dwTime.Hours,&dwTime.minutes,&dwTime.CapturedTimeDw)
		if err!=nil{
			//panic(err)
			fmt.Println("test1")
			fmt.Println("Empty Result from total users",err)

		}
		dwellTimeSts=append(dwellTimeSts,dwTime)
		fmt.Println("daily users value",dwellTimeSts)
	}
	if err = rows.Err();err != nil{
		panic(err)
	}

	return dwellTimeSts,nil
}