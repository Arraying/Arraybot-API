package ratelimits

import (
	"fmt"
	"github.com/arraying/Arraybot-API/files"
	"net/http"
	"time"
)

// Ratelimit is a ratelimit object.
type Ratelimit struct {
	Requests int
	Time     int
}

// RatelimitUser is a user ratelimit object, per endpoint.
type RatelimitUser struct {
	RequestsLeft int
	Reset        int64
}

// RatelimitData is ratelimit information and data.
type RatelimitData struct {
	Ratelimit Ratelimit
	Users     map[string]*RatelimitUser
}

// Handle handles the ratelimit. Returns true if b1nzy'd.
func (ratelimit *RatelimitData) Handle(writer http.ResponseWriter, request *http.Request) bool {
	ip := request.RemoteAddr
	user, exists := ratelimit.Users[ip]
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1e6
	if !exists {
		ratelimitInfo := ratelimit.Ratelimit
		reset := millis + int64(ratelimitInfo.Time*1000)
		newUser := &RatelimitUser{
			RequestsLeft: ratelimitInfo.Requests - 1,
			Reset:        reset,
		}
		ratelimit.Users[ip] = newUser
		handleTimer(ratelimit, ip, time.Duration(ratelimitInfo.Time))
		handleWriter(writer, newUser)
		return false
	}
	if user.RequestsLeft == 0 {
		handleWriter(writer, user)
		writer.WriteHeader(http.StatusTooManyRequests)
		return true
	}
	newRequests := user.RequestsLeft - 1
	user.RequestsLeft = newRequests
	handleWriter(writer, user)
	return false
}

// Ratelimits is a map of ratelimits.
var Ratelimits = map[string]RatelimitData{
	files.APIGetLanguages: RatelimitData{
		Ratelimit: Ratelimit{
			Requests: 5,
			Time:     10,
		},
		Users: make(map[string]*RatelimitUser),
	},
	files.APIGetLanguage: RatelimitData{
		Ratelimit: Ratelimit{
			Requests: 5,
			Time:     60,
		},
		Users: make(map[string]*RatelimitUser),
	},
	files.APIPatchLanguage: RatelimitData{
		Ratelimit: Ratelimit{
			Requests: 5,
			Time:     5,
		},
		Users: make(map[string]*RatelimitUser),
	},
}

func handleWriter(writer http.ResponseWriter, user *RatelimitUser) {
	remaining := user.RequestsLeft
	reset := user.Reset
	writer.Header().Add("RateLimit-Remaining", fmt.Sprintf("%v", remaining))
	writer.Header().Add("RateLimit-Reset", fmt.Sprintf("%v", reset))
}

func handleTimer(ratelimit *RatelimitData, ip string, delay time.Duration) {
	timer := time.NewTimer(time.Second * delay)
	go func() {
		<-timer.C
		delete(ratelimit.Users, ip)
	}()
}
