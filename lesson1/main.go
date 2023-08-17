package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	getTime()
	testIota()
}

/*
Hello now()
Завести Go репозиторий на GitHub,
написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода.
*/

func getTime() {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		if _, err := os.Stderr.WriteString(err.Error()); err != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}

	fmt.Println(currentTime.Format("01.02.2006 15:04:05"))
}

/*
напиши iota для констант b, kb, mb, gb, tb
*/

/*
0001
0010
0100
1000
*/

func testIota() {
	const (
		b  = 1 << (iota * 10) // 2^0
		kb                    // 2^10
		mb                    // 2^20
		gb                    // 2^30
		tb
	)

	fmt.Println(b)
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)
	fmt.Println(tb)
}
