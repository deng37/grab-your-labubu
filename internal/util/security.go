package util

import (
	"time"
	"sync"
	"fmt"
)

const ( HackerDuration = 10.0 )

var userMap sync.Map
type user struct {
	sync.Mutex
	lastReset time.Time
	count int
	startTime time.Time
	endTime time.Time
}

func IsUserOverLimit(ip string) bool {
	now := time.Now()
	val, _ := userMap.LoadOrStore(ip, &user{lastReset: now, count: 0})
	userNow, userNowOk := val.(*user)

	if !userNowOk {	// Not user
		return true
	}

	userNow.Lock()
	defer userNow.Unlock()

	if now.Sub(userNow.lastReset) > time.Second {	// Reset counter every 1s
		userNow.count = 1
		userNow.lastReset = now
		return false
	}

	if userNow.count >= 5 {	// Over limit
		fmt.Println("User is over limit")
		return true
	} else {
		userNow.count++
		return false
	}
}

func UpdateUserStartTime(ip string, t time.Time) {
	now := time.Now()
	val, _ := userMap.LoadOrStore(ip, &user{lastReset: now, count: 0})

	u := val.(*user)
    u.Lock()
    defer u.Unlock()
    u.startTime = t
}

func UpdateUserEndTime(ip string, t time.Time) {
	now := time.Now()
	val, _ := userMap.LoadOrStore(ip, &user{lastReset: now, count: 0})

	u := val.(*user)
    u.Lock()
    defer u.Unlock()
    u.endTime = t
}

func GetUserDuration(ip string) float64 {
	val, ok := userMap.Load(ip)
    if !ok {
        fmt.Println("Not able to update user start time, IP not found: " + ip)
        return 999
    }

	u := val.(*user)
	duration := u.endTime.Sub(u.startTime)
	durationMS := float64(duration) / float64(time.Millisecond)
	return durationMS
}

func IsHacker10Ms(duration float64) bool {
	if duration < HackerDuration {
		fmt.Println("Hacker detected")
		return true
	} else {
		return false
	}
}