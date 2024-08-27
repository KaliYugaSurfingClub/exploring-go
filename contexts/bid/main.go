/*
	Аукцион состоит из нескольких участников, которые делают постоянно увеличивающиеся ставки по лоту
	Он может закончиться в следующих случаях:
	1. дана максимальная ставка
	2. сделано определённое количество ставок
	3. прошло определённое время
	4. todo the user thinks too long
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bid struct {
	PlayerID int
	Value    int
}

type Lot struct {
	sync.Mutex

	MaxTime  time.Duration
	MaxPrice int
	MaxBids  int

	LastBid    *Bid
	StartPrice int
	CurrentCnt int
}

func (b *Lot) GetLastBid() int {
	if b.LastBid == nil {
		return b.StartPrice
	}

	return b.LastBid.Value
}

func (b *Lot) SetNewBid(newBid *Bid) bool {
	b.Lock()
	defer b.Unlock()

	if newBid.Value > b.GetLastBid() {
		fmt.Printf("new bid: %+v\n", newBid)

		b.LastBid = newBid
		b.CurrentCnt++

		if b.MaxBids <= b.CurrentCnt {
			println("finish by count")
			return true
		}

		if b.MaxPrice <= newBid.Value {
			println("finish by bid")
			return true
		}
	}

	return false
}

const (
	AvgBidStep = 100
)

func makeBid(ctx context.Context, playerID int, lot *Lot, bids chan<- *Bid) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case bids <- &Bid{PlayerID: playerID, Value: lot.GetLastBid() + 2*rand.Intn(AvgBidStep)}:
		default:
		}
	}
}

func main() {
	lot := &Lot{MaxBids: 100, MaxPrice: 10000, MaxTime: time.Second * 10}
	bids := make(chan *Bid)
	ctx, finish := context.WithTimeout(context.Background(), lot.MaxTime)

	for i := 0; i < 5; i++ {
		go makeBid(ctx, i, lot, bids)
	}

LOOP:
	for {
		select {
		case bid := <-bids:
			if lot.SetNewBid(bid) {
				finish()
				break LOOP
			}
		case <-ctx.Done():
			fmt.Println("Done:", ctx.Err())
			break LOOP
		}
	}

	fmt.Println("Auction finished by player", lot, "with price", lot.GetLastBid())
}
