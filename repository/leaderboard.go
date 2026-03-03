package repository

import (
	"time"
	"github.com/deng37/grab-your-labubu/util"
)

type LeaderboardEntry struct {
	IP         string
	UserName   string
	DurationMS float64
	UpdatedAt  time.Time
}

func UpsertWinner(ip string, duration float64) error {
	query := `
		INSERT INTO leaderboard (ip_address, duration_ms)
		VALUES (?, ?)
		ON CONFLICT(ip_address) DO UPDATE SET
			duration_ms = CASE
				WHEN excluded.duration_ms < leaderboard.duration_ms THEN excluded.duration_ms
				ELSE leaderboard.duration_ms
			END,
			updated_at = CURRENT_TIMESTAMP;
	`
	_, err := util.DB.Exec(query, ip, duration)
	return err
}

func GetTopWinners(limit int) ([]LeaderboardEntry, error) {
	query := `SELECT ip_address, duration_ms FROM leaderboard
			  ORDER BY duration_ms ASC LIMIT ?`

	rows, err := util.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []LeaderboardEntry
	for rows.Next() {
		var e LeaderboardEntry
		if err := rows.Scan(&e.IP, &e.UserName, &e.DurationMS); err == nil {
			list = append(list, e)
		}
	}
	return list, nil
}