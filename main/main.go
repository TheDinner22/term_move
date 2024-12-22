package main

import (
	"github.com/gdamore/tcell"
)

func clamp(v int, lower int, upper int) int {
    if (lower > upper) {
        panic("bad clamp call")
    }

    if v >= lower && v <= upper {
        return v
    } else if v < lower {
        return lower
    } else {
        return upper
    }

}

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	err = s.Init()
	if err != nil {
		panic(err)
	}

	defer s.Fini()

	// set screen to be all .'s
	x, y := s.Size()
	for i := range x {
		for j := range y {
			s.SetContent(i, j, '.', nil, tcell.StyleDefault)
		}
	}
	s.Show()

    current_x := 0
    current_y := 0

	// spawn character in top left
	s.SetContent(0, 0, '/', nil, tcell.StyleDefault)

	// event loop
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				return
			}

			dx := 0
			dy := 0
			if ev.Key() == tcell.KeyUp {
				dy = -1
			}

			if ev.Key() == tcell.KeyDown {
				dy = 1
			}

			if ev.Key() == tcell.KeyLeft {
				dx = -1
			}

			if ev.Key() == tcell.KeyRight {
				dx = 1
			}

            s.SetContent(current_x, current_y, '.', nil, tcell.StyleDefault)
            current_x += dx
            current_y += dy
            current_x = clamp(current_x, 0, x-1)
            current_y = clamp(current_y, 0, y-1)
            s.SetContent(current_x, current_y, '/', nil, tcell.StyleDefault)
            s.Show()
		}
		s.Show()
	}
}
