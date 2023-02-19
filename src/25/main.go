package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(sleepVer1(5 * time.Second))
	fmt.Println(time.Now().Sub(start))
	fmt.Println(sleepVer2(2 * time.Second))
	fmt.Println(time.Now().Sub(start))
}

func sleepVer1(dur time.Duration) (val time.Time) { //  с использованием тикера
	ticker := time.Tick(dur)
	for val = range ticker { //  ждем когда в канал тикер прийдет значение и выходим из программы
		return val
	}
	return
}

func sleepVer2(dur time.Duration) time.Time { // с использования канал time.After
	return <-time.After(dur) //  ждем пока в канал прийдет значение и выходим из программы
}
