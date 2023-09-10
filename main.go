package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Ball struct{
	// Where is the ball?
	x int
	y int
	char rune // What does is look like?
}
var ball = Ball{1,1,'#'}

func main(){
	screen, err :=  tcell.NewScreen()
	if err != nil{
		log.Fatal("%+v",err)
	}
	if screen.Init(); err != nil {
		log.Fatal("%+v",err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)
	screen.Clear()

	// this executes when main() returns
	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Event loop
	for {
		screen.SetContent(ball.x, ball.y, ball.char, nil,defStyle)
		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type){
			case *tcell.EventKey:
				// exiting the program
				if ev.Key() == tcell.KeyESC || ev.Key() == tcell.KeyCtrlC {
					return
				}
				if ev.Key() == tcell.KeyUp{
					ball.y--	
				}
				if ev.Key() == tcell.KeyDown{
					ball.y++
				}
				if ev.Key() == tcell.KeyLeft{
					ball.x--
				}
				if ev.Key() == tcell.KeyRight{
					ball.x++
				}
		} // End-Switch

		screen.Clear()

	}
	
}

func registerInput(screen tcell.Screen){

}


// EXAMPLE-function
// x1 and y1 tell: where to start
// x2 and y2 tell: how much space
func drawText(screen tcell.Screen, x1, y1, x2, y2 int, style tcell.Style){
	text := "#"
	row := y1
	col := x1
	// a rune is an alias for type32, because go uses UTF-8 encoding
	for _, r := range []rune(text){
		// setContent takes a rune
		screen.SetContent(col, row, r, nil,style)
		col++
		if col >= x2 {
			row++
			col = x1
		}

		if row > y2 {
			break
		}
		
	}

}
