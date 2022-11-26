package controller

import (
	"fmt"

	"github.com/uekiGityuto/othello-go/app/model"
)

type Controller struct {
	turn  model.Color
	board *model.Board
}

func New(turn model.Color) (*Controller, error) {
	board, err := model.NewBoard()
	if err != nil {
		return nil, fmt.Errorf("failed to create board: %w", err)
	}
	return &Controller{
		turn:  turn,
		board: board,
	}, nil
}

func (c *Controller) ChangeTurn() {
	c.turn = c.turn.GetReversed()
}

func (c *Controller) Start() {
	fmt.Println("石を置きたい場所を「列番号,行番号」の形式で入力して下さい。例）左上隅の場合: 0,0")
	fmt.Println("石を置きたい場所を「列番号,行番号」の形式で入力して下さい。例）左上隅の場合: 0,0")
	fmt.Println("パスをしたい時は「pass」と入力して下さい。")

	c.board.Draw()
}
