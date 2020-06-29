package dailyUsers

import (
	"context"
	"fmt"
	"github.com/sudesh35139/prx/config"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type DailyUse struct {
	//Id int32 `json:"id"`
	//ApMac string `json:"ap_mac"`
	//SourceMac string`json:"source_mac"`
	//DestinationMac string`json:"destination_mac"`
	CapturedTime string`json:"captured_time"`
	//SignalRate int8 `json:"signal_rate"`
	//Chanel string `json:"chanel"`
	NumOfUser int32 `json:"num_of_user"`
	//NumOfRepUser int32 `json:"num_of_rep_user"`
	//AvgDwellTime time.Time `json:"avg_dwell_time"`
	fromDate time.Time
	toDate time.Time
	//MinCapturedTime time.Time `json:"min_captured_time"`
	//MaxCapturedTime time.Time `json:"max_captured_time"`

}

var (
	ctx context.Context
)
func TotDailyUser(r *http.Request)([]DailyUse,error) {
	dUser :=DailyUse{}

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

	//"SELECT * FROM users WHERE "+timeColoumn+" BETWEEN '"+formatedFrom+"' AND '"+formatedTo+"'"

	//number of people  count from given time periods.
	rows,err := config.DB.Query("SELECT date(capturedtime)as CapturedTime, count(*)as NumOfUser FROM proximity.`6c-3b-6b-68-e4-f3`where capturedtime between '"+fromdateCha+"' and '"+toDateCha+"' group by date(capturedtime)order by capturedtime")
	if err != nil{

		fmt.Println("Empty row from number of daily users",err)
	}
	defer  rows.Close()
	dailyUsers := make([]DailyUse,0)

	for  rows.Next(){

		err := rows.Scan(&dUser.CapturedTime,&dUser.NumOfUser)
		if err!=nil{
			//panic(err)
			fmt.Println("Empty Result from total users",err)

		}
		dailyUsers=append(dailyUsers,dUser)

	}
	if err = rows.Err();err != nil{
		panic(err)
	}

return dailyUsers,nil
	
}
