package util

import (
	"time"
	"sync"
	"fmt"
)

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
		return true
	} else {
		userNow.count++
		return false
	}
}

func UpdateUserStartTime(ip string, t time.Time) {
	val, ok := userMap.Load(ip)
    if !ok {
        fmt.Println("Not able to update user start time, IP not found: " + ip)
        return
    }

	u := val.(*user)
    u.Lock()
    defer u.Unlock()
    u.startTime = t
}

func UpdateUserEndTime(ip string, t time.Time) {
	val, ok := userMap.Load(ip)
    if !ok {
        fmt.Println("Not able to update user start time, IP not found: " + ip)
        return
    }

	u := val.(*user)
    u.Lock()
    defer u.Unlock()
    u.endTime = t
}

func IsHacker10Ms(ip string) bool {
	val, ok := userMap.Load(ip)
    if !ok {
        fmt.Println("Not able to update user start time, IP not found: " + ip)
        return true
    }

	u := val.(*user)
	duration := u.endTime.Sub(u.startTime)
	if duration < 10*time.Millisecond {
		fmt.Println("Hacker detected")
		return true
	} else {
		return false
	}
}