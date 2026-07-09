package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Move struct {
	Color int
	Col   int
}

type Board struct {
	BoardString string
}

type ConnectGame int

func (t *ConnectGame) Move(args *Move, reply *int) error {
	log.Println("Move received:", args.Color, args.Col)

	if args.Color == lastColorMoved { // check to see if same player/color is going again
		return fmt.Errorf("Turn Order Violation")
	}

	if args.Col >= 0 && args.Col < 7 {
		for row := range gameBoard { // scan through each row
			cell := gameBoard[row][args.Col] //gets each cell in from that column
			if cell == 1 || cell == 2 {      //planning to make white = 1 and black = 2
				if row == 0 {
					return fmt.Errorf("column number %d is full", args.Col) // if the column is full you can't place
				}
				gameBoard[row-1][args.Col] = args.Color + 1 // the color +1 so that we don't have any zeros, since that is what is already filling the board
				break
			}
			if 0 == gameBoard[len(gameBoard)-1][args.Col] { //for if we reach last row in column with no filled cell
				gameBoard[len(gameBoard)-1][args.Col] = args.Color + 1
				break
			}
		}
	} else {
		return fmt.Errorf("not a valid column dude")
	}
	time.Sleep(10 * time.Millisecond)
	lastColorMoved = args.Color
	return nil
}

func (t *ConnectGame) Get(args *int, reply *Board) error {
	reply.BoardString = ""        //we will build from this
	for rows := range gameBoard { // nested loop to scan through each cell
		for cols := range gameBoard[rows] {
			if gameBoard[rows][cols] == 0 { //empty spot
				reply.BoardString += "."
			} else if gameBoard[rows][cols] == 1 { //has a white checker
				reply.BoardString += "W"
			} else if gameBoard[rows][cols] == 2 { // has a black checker
				reply.BoardString += "B"
			}
		}
		reply.BoardString += "\n" // makes sure the printed board isnt just printed in a line
	}
	return nil
}

var gameBoard [][]int
var lastColorMoved int = -1 // -1 means no moves yet

func main() {
	rows := 6
	cols := 7
	gameBoard = make([][]int, rows) // Allocates the outer slice

	for i := range gameBoard {
		gameBoard[i] = make([]int, cols) // Allocates each inner row
	}
	cg := new(ConnectGame)
	rpc.Register(cg)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Println("Serving on PORT 1234")
	http.Serve(l, nil)
}
