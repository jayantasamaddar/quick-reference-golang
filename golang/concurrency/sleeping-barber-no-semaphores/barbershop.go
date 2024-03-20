package slb

import (
	"fmt"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

type BarberShop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		// Set a variable
		isSleeping := false
		fmt.Println(std.PrintC(std.Yellow, barber, "goes to the waiting room to check for clients..."))

		for {
			// If there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				fmt.Println(std.PrintC(std.Yellow, fmt.Sprintf("There is nothing to do, so %s, takes a nap.", barber)))
				isSleeping = true
			}

			// If the channel is still open, we have a real client
			client, ok := <-shop.ClientsChan
			if ok {
				// Wake the barber up if he's sleeping
				if isSleeping {
					fmt.Println(std.PrintC(std.Yellow, fmt.Sprintf("%s wakes the barber, %s up.", client, barber)))
					isSleeping = false
				}
				// cut hair
				shop.CutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				shop.SendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) CutHair(barber, client string) {
	fmt.Println(std.PrintC(std.Green, fmt.Sprintf("Barber %s is cutting %s's hair.", barber, client)))
	time.Sleep(shop.HaircutDuration)
	fmt.Println(std.PrintC(std.Green, fmt.Sprintf("Barber %s has finished cutting %s's hair.", barber, client)))
}

func (shop *BarberShop) SendBarberHome(barber string) {
	fmt.Println(std.PrintC(std.Cyan, barber, "is going home!"))
	// Cannot take anymore clients
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) CloseShopForTheDay() {
	fmt.Println(std.PrintC(std.Blue, "Closing shop for the day!"))

	// Cannot accept anymore clients
	close(shop.ClientsChan)

	// Close shop
	shop.Open = false

	// Block until every single barber is done
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)

	fmt.Println(std.PrintC(std.Green, "The Barbershop is now closed for the day and everybody has gone home!"))
}

func (shop *BarberShop) AddClient(client string) {
	fmt.Println(std.PrintC(std.Green, fmt.Sprintf("Client %s has arrived!", client)))

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			fmt.Println(std.PrintC(std.Yellow, fmt.Sprintf("%s takes a seat in the waiting room.", client)))
		default:
			fmt.Println(std.PrintC(std.Red, fmt.Sprintf("The waiting room is full, so %s leaves!", client)))
		}
	} else {
		fmt.Println(std.PrintC(std.Red, fmt.Sprintf("The shop is already closed, so %s leaves!", client)))
	}
}
