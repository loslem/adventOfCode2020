package adv11

import (
	"strings"
)

func main() {
	sm := storeInput(puzzleInput)
	sm = sm.CheckSeatsAgain()
	checkSumOld := 0
	checkSumNew := 0

	for {
		sm = sm.CheckSeatsAgain()
		checkSumNew = 0
		for y, jj := range sm {
			for x := range jj {
				if sm[y][x] == "#" {
					checkSumNew++
				}
			}

		}
		if checkSumNew == checkSumOld {
				break
		} else {
			checkSumOld = checkSumNew
		}
	}
	println(checkSumNew)
}

func (g *SeatGrid) CheckSeatsAgain() *SeatGrid {
	ssg := &SeatGrid{}

	for y, jj := range g {
		for x := range jj {
			if g[y][x] == "L" {
				if g.checkTakenSeats(y, x) == 0 {
					ssg[y][x] = "#"
				} else {
					ssg[y][x] = "L"
				}
			}

			if g[y][x] == "#" {
				if g.checkTakenSeats(y, x) > 3 {
					ssg[y][x] = "L"
				} else {
					ssg[y][x] = "#"
				}
			}

			if g[y][x] == "." {
				ssg[y][x] = "."
			}
			//println(fmt.Sprintf("x:%v, y:%v, val: %v",oo,x,ssg[oo][x]))
		}
	}

	return ssg
}

type SeatGrid [98][99]string

const TestInput = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
const puzzleInput = "LLLLLLLLL.L.LLL.L.LLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLL.LLLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLL.LLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nL..LL.L.L....L..L.L....L.....L.L.L......L..L..L....L.LL.......L.....L....L..L................L....\nLLLLLLLLLLLLLLL.LLL.LLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL..LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLL.LLLLLLLLL.LLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLL.LLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLL\nLL.LLLLLLLLLLLL.LLL.LLL.LLLLLLLL.LLLLLLLL.LLLLLLL..LLLL.LLLLLLL.LLLL.LLLLLL.LLLLL.LLLL.LL.LLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLL.L.LLLLLLL.LLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLL.LLLL.LLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLL.LLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLLLLL.LL.LLLLLLLLLLLLLLLL\n.LLLL..L.L..L....L.LLL...LLL..L.L....L.L.L...L.....L...L.....L...L..L...LL...L..LLL..L...L....L..L\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLL.LLLL..LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL..LLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\n.....LL...LL..L...L..L..LL.L..LLL...L..L.LLL...LLL..L.....L...L.....L........LLLL..LLL.....L....L.\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLL.LLLL.LLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLLLLLLLL.LL\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLL.LL.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLL.L\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLL.LLLLLLLLL.LLLLLLL.LLLL.LLL.LL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.L.LLLLL.LLLLLLLL\n....L..LL.L.L......L..LL.......LL......L..L...L..L.........LL....LL.LLL..LLLL..L..L.......L....L..\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLL.LLL.L.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLL.LLLLLLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL\n.LL..LLLLL.LL.L...L...L...L....LLLL...LLLL.LL...LL...L.LL...LL..LL.L..L......L....L.L.L.L...LL...L\nLLLLLLLLL.LLLLLLLLLLLLL..LLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLLLLL.LL.LLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLL.L.LLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLL..LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LL.LL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLLLL.LLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLL.LLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLLLLLLLL.LLLLL.L.LLLLLLLL\n....L.........LL...L.L..LL.L.L..L......L.L.L...L....L.......LL.LL....L.L...L....L...L....L...L....\nL.LLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLL.LL..LLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLL..LLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLL.L.LLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLL.LL.LLLLLLLLLLLLL.LLLL.LLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLL..L.L........L...L..L.LLLLLLL..LLLL..L...L..L.L..LLL....L..L.LLL..L...L..L.......L.LLLLLL.LL.L.\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LL.LLLL.LLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LL.LL.LLLLLLL.LLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLL.LLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nL...L.....LL....L.L.LLLLL.....L....LL..L.....L...L.L.LL.......LL...LL............L.L..L..LLL..LL..\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLL.LL\nLL.LLLLLLLLLLLL.L.LLL.LLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LL.LLLL.LLLLLL.LLLLL.LLLLLLL.L.LLLLLL\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLL.LLLL\nLLLLLLLLL.LLLLL.L.LLLLL.LLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLLLLLLLLLL.LLLLLLLLLLLLLL\nLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLL..LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLL...L.........L..LLL........L.L..L....L.L.L....LLLL.LL.L....LL....L.LL.LLL..L..L....LL.L......LLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLL.L.LLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\n...L........L....L.L.LL.LL.L.L.....L.L..L.....LL....L..L.L.L..L..L.L...L.........L.LL.LLLLL..L....\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL..LLLL.LLLLLLL.LLLLLLLLLLL.LLL.L.LLLLLLLLLLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLL.L.LLLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLL.LL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLL..LLLLLLL\nLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL\n..LLLL.L..LL.L......LL.L...L.L..L....LLL.........L....L.......LL.LLL.L.L.....LL.L...L.LLL....L....\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.L.LLL.LLLLLLL.LLLL.LLLLLL.LLLL..LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLL.LLLLLLLLL.L.LLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLL.LLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLLLLLLL.LL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLL.LLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLL.LLLLLLL.LLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLL...LLLLL.....LL..L.....L....L......L.L........LL......LL.LL..L..LLL.L...L..L..L.L.....LLLL...LLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLLLLLLLLL.LLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLL.LLLLL.LLL.LLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\n.LL..L.....LLLL.LL..L.LL....L.LL...L..LL......L..LL...LL..L.L.LLLL...LL.L..L..LLL.L.L.....L.LL.LL.\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLL.LL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLLLLLLLL.LLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLL.LLLL.LLLLL.LLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLL.LLLLL.LL\nLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLL\nLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLL.LLLLLLL.LLLL..LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLL\nLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLL.LLL.L.LLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLL.LLLLLLLL\n"

func (g *SeatGrid) checkASeat(y, x int) int {
	if x < 0 || x > 99 {
		return 0
	}

	if y < 0 || y > 98 {
		return 0
	}

	seat := g[y][x]

	if seat == "." {
		return 0
	}

	if seat == "#" {
		return 1
	}
	if seat == "L" {
		return 0
	} else {
		return 0
	}
}

func (g *SeatGrid) checkTakenSeats(y, x int) int {
	checkSum := 0
	checkSum += g.checkASeat(y-1, x-1)
	checkSum += g.checkASeat(y-1, x)
	checkSum += g.checkASeat(y-1, x+1)
	checkSum += g.checkASeat(y+1, x)
	checkSum += g.checkASeat(y+1, x-1)
	checkSum += g.checkASeat(y+1, x+1)
	checkSum += g.checkASeat(y, x+1)
	checkSum += g.checkASeat(y, x-1)
	return checkSum
}

func storeInput(s string) *SeatGrid {
	var sm = &SeatGrid{}
	clean := strings.ReplaceAll(s, "\n", " ")
	split := strings.Fields(clean)

	for i, j := range split {

		for y := 0; y < len(j); y++ {
			tmp := string([]rune(j)[y])

			sm[i][y] = tmp

		}
	}

	return sm
}
