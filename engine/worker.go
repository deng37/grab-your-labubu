package engine

import (
	"github.com/deng37/grab-your-labubu/model"
)

// To return boolean whether got the Labubu or not, with the message value
func GrabItem(store *model.LabubuStore) (bool, string) {
	store.Lock()
	defer store.Unlock()

	if store.Count > 0 {
		store.Count--
		return true, "Got the Labubu!"
	}
	return false, "Sorry Labubu out of stock :("
}