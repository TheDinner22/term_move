package main

import (
	"github.com/gdamore/tcell"
);

func main() {
    s, err := tcell.NewScreen();
    if err != nil {
        panic(err);
    }

    err = s.Init();
    if err != nil {
        panic(err)
    }

    // event loop
    for {
        s.SetContent(2, 2, 'i', nil, tcell.StyleDefault)
        s.Show()
        ev := s.PollEvent();
        switch ev := ev.(type){
            case *tcell.EventKey:
                if ev.Key() == tcell.KeyUp {
                    s.Fini();
                    return
                }
        }

    }
}
