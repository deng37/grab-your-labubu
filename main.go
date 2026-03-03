package main

import (
	"encoding/json" // JSON
	"fmt"
	"time"
	"net/http"
	"os"
	"sync"           // Mutex dan WaitGroup (For 100 Bots)
	"github.com/deng37/grab-your-labubu/engine"
	"github.com/deng37/grab-your-labubu/model"
	"github.com/deng37/grab-your-labubu/util"
)

const ( WinnerSeparator = ", " )
const ( UserId = 99999 )
const ( NoOfStock = 10 )
const ( NoOfPodium = 3 )
const ( DefaultPort = "8080" )
const ( EmptyString = "" )

var warStartTime time.Time
var podium chan int = make(chan int, NoOfPodium)

func main() {
	// Init
	store := &model.LabubuStore{
		StockName: "Labubu Tasty Macarons",
		Count: NoOfStock,
	}
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// API /grab - Grab Labubu
	http.HandleFunc("/grab", func(w http.ResponseWriter, r *http.Request) {
		util.UpdateHeaderJson(w)

		userIp := util.GetUserIP(r)
		fmt.Println("IP address: ", userIp)
		util.UpdateUserEndTime(userIp, time.Now())
		if util.IsUserOverLimit(userIp) || util.IsHacker10Ms(userIp) {	// Prevent unusual click and hacker
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		success, message := engine.GrabItem(store)

		if success {
			updatePodium(UserId)
			fmt.Fprintf(w, `{"status": "success", "message": "%s"}`, message)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"status": "failed", "message": "%s"}`, message)
		}
	})

	// API / - Serve HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// API /war-start - Serve HTML
	http.HandleFunc("/war-start", func(w http.ResponseWriter, r *http.Request) {
		userIp := util.GetUserIP(r)
		util.UpdateUserStartTime(userIp, time.Now())
		w.WriteHeader(http.StatusNoContent)
	})

	// API /war - Bot Coming to Arena
	http.HandleFunc("/war", func(w http.ResponseWriter, r *http.Request) {
		util.UpdateHeaderJson(w)

		successCount := 0
		var muWar sync.Mutex // mutex for scoring
		var wg sync.WaitGroup

		// Release 100 bots
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				ok, _ := engine.GrabItem(store)	// Bot fighting for Labubu via engine.GrabItem
				if ok {
					updatePodium(id)
					muWar.Lock()
					successCount++
					muWar.Unlock()
					fmt.Printf("Bot #%d got Labubu!\n", id)
				}
			}(i)
		}

		wg.Wait() // Wait until all Bots done attacking

		json.NewEncoder(w).Encode(map[string]interface{}{
			"bots_captured": successCount,
			"remaining":     store.Count,
			"podium_winner": getWinners(),
		})
	})

	// API /reset - Stock Reset
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		store.Lock()
		defer store.Unlock()

		store.Count = 10
		w.WriteHeader(http.StatusNoContent)
	})

	port := os.Getenv("PORT");
	if port == EmptyString {
		port = DefaultPort
	}
	port = ":" + port;
	fmt.Println("🚀 Labubu server running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func getWinners() string {
	var winners string

	PodiumLoop:
	for i := 0; i < 3; i++ {
		select {
			case winner := <-podium:
				if (winner != UserId) {	// Winner is a bot
					switch i {
						case 0: winners = fmt.Sprintf("🥇 BOT %d", winner)
						case 1: winners += fmt.Sprintf("%s🥈 BOT %d", WinnerSeparator, winner)
						case 2: winners += fmt.Sprintf("%s🥉 BOT %d", WinnerSeparator, winner)
					}
				} else {	// Winner is the user
					switch i {
						case 0: winners = fmt.Sprintf("🥇 YOU")
						case 1: winners += fmt.Sprintf("%s🥈 YOU", WinnerSeparator)
						case 2: winners += fmt.Sprintf("%s🥉 YOU", WinnerSeparator)
					}
				}
			case <-time.After(100 * time.Millisecond):
				break PodiumLoop
		}
	}

	if winners == EmptyString {
		return "No Winner :("
	} else {
		return winners
	}
}

func updatePodium(id int) {
	select {
		case podium <-id:
		default:
			// Do nothing since podium full already
	}
}