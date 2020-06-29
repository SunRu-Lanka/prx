package authenticatPac

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)
const sessionLength int = 300

func GetUser(w http.ResponseWriter,req *http.Request)user {
	//get cookie
	c, err := req.Cookie("proximity")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge =sessionLength
	http.SetCookie(w,c)
	//if the user exits already ,get user
	var U user
	if s,ok := DbSessions[c.Value];ok{
		s.LastActivity = time.Now()
		DbSessions[c.Value]= s
		U = Dbusers[s.Un]
	}
	return U

}
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("proximity")
	if err != nil {
		return false
	}
	s, ok := DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		DbSessions[c.Value] = s
	}
	_, ok = Dbusers[s.Un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
func cleanSessions()  {
	for k,v := range DbSessions{
		if time.Now().Sub(v.LastActivity)>(time.Second *30){
			delete(DbSessions,k)
		}
	}
	DbSessionCleaned = time.Now()


}
func showSessions() {
	fmt.Println("********")
	for k, v := range DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
