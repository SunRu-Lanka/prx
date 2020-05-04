package dailyUsers

import (
	"fmt"
	"github.com/sudesh35139/prx/config"
	"net/http"
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
	FromDate time.Time
	ToDate time.Time
	//MinCapturedTime time.Time `json:"min_captured_time"`
	//MaxCapturedTime time.Time `json:"max_captured_time"`

}

func TotDailyUser(r *http.Request)([]DailyUse,error) {
	dUser :=DailyUse{}
	//taken value from daily user html form(select date range from date and to date.
	//  fromDateStart := r.FormValue("fromDate")
	// toDateEnd := r.FormValue("toDate")
	//if fromDateStart ==""||toDateEnd ==""{
	//	fmt.Println("error  4")
	//	return nil,errors.New("FromDate or ToDate cannot be null")
	//}
	////Chang Date format Patten
	//DFrom,err :=time.Parse("2006-01-02",fromDateStart)
	//if err!= nil{
	//	return nil,errors.New("406 not acceptable.wrong format")
	//}
	//dUser.FromDate = DFrom
	//DTo,err :=time.Parse("2006-01-02",toDateEnd)
	//if err!= nil{
	//	return nil,errors.New("406 not acceptable.wrong format")
	//}
	//dUser.ToDate = DTo

	//number of people  count from given time periods.
	rows,err := config.DB.Query("SELECT date(capturedtime)as 'CapturedTime', count(*)as'NumOfUser'FROM proximity.`6c-3b-6b-68-e4-f3`where capturedtime between '2020-01-10' and '2020-02-11'group by date(capturedtime)order by capturedtime")
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
		//fmt.Println("daily users value",dailyUsers)
	}
	if err = rows.Err();err != nil{
		panic(err)
	}

return dailyUsers,nil
	
}
