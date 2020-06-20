package repeitUsers

import (
	"fmt"
	"github.com/sudesh35139/prx/config"
	"io/ioutil"
	"net/http"
	"strings"
)

type RpUsers struct {
	PDate string `json:"p_date"`
	RpCount int32 `json:"rp_count"`
}

func RpUsersProcess(r *http.Request) ([]RpUsers, error) {
	req ,err :=ioutil.ReadAll(r.Body)

	if err !=nil{
		fmt.Println("nill data",err)
	}
	reqs := string(req)
	fmt.Println("check value******",reqs)
	if len (reqs) == 0{
		fmt.Println("empty string")

	}
	spiletDate := strings.Split(reqs,"&")
	fromDateStartV:= spiletDate[0]
	fromdateCha:= fromDateStartV[8:]
	fmt.Println("Spilt date ###OOOOOOOOOOOOOO",fromdateCha)
	toDateEndV := spiletDate[1]
	toDateCha := toDateEndV[6:]

	fmt.Println("spilt End dateOOOOOOOOOOOOO",toDateCha)



	rpTotUser := RpUsers{}

	rows,err := config.DB.Query("select pdate as'PDate',count(*) as'RpCount' from vw_repeatusers where pdate between  '"+fromdateCha+"' and '"+toDateCha+"' group by pdate")
	if err != nil{
		fmt.Println("Empty RpUsers ",err)
	}
	defer rows.Close()
	rpTotUsers := make([]RpUsers,0)
	for rows.Next(){
		err:= rows.Scan(&rpTotUser.PDate,&rpTotUser.RpCount)
		if err!= nil{
			fmt.Println("Empty rows get from database ",err)
		}
		rpTotUsers= append(rpTotUsers,rpTotUser)
		//fmt.Println("test Understand",rpTotUsers)

	}
	if err = rows.Err();err != nil{
		panic(err)
	}
	return rpTotUsers,nil


}