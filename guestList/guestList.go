package guestlist

import "github.com/Pallinder/go-randomdata"

func generate() [10][2]string {
	guestlist := [10][2]string{}
	for index, _ := range guestlist {
		guestlist[index] = roomGuests(index)
	}
	return guestlist
}
func roomGuests(roomID int) [2]string {
	guests := [2]string{}
	guests[0] = randomdata.SillyName()
	guests[1] = randomdata.SillyName()
	return guests
}
