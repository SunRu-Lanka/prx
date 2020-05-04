package repeitUsers

import (
	"fmt"
	"github.com/sudesh35139/prx/config"
	"net/http"
)

type RpUsers struct {
	PDate string `json:"p_date"`
	RpCount int32 `json:"rp_count"`
}

func RpUsersProcess(r *http.Request) ([]RpUsers, error) {
	rpTotUser := RpUsers{}
	rows,err := config.DB.Query("select pdate as'PDate',count(*) as'RpCount' from vw_repeatusers where pdate between '2020-01-10' and '2020-02-11' group by pdate")
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