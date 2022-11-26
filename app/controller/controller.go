package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func (c *Controller) changeTurn() {
	c.turn = c.turn.GetReversed()
}

func (c *Controller) validate(input string) error {
	inputs := strings.Split(input, ",")
	for i, v := range inputs {
		inputs[i] = strings.TrimSpace(v)
	}
	if len(inputs) != 2 {
		return fmt.Errorf(`input string "%s" is invalid`, input)
	}
	for _, v := range inputs {
		num, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf(`input string "%s" is invalid`, input)
		}
		if num < 0 || num > 7 {
			return fmt.Errorf(`input string "%s" is invalid`, input)
		}
	}
	return nil
}

func (c *Controller) Start() {
	fmt.Println("石を置きたい場所を「列番号,行番号」の形式で入力して下さい。例）左上隅の場合: 0,0")
	fmt.Println("やめたい時は「quit」と入力して下さい。")
	fmt.Println("パスをしたい時は「pass」と入力して下さい。")

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%sの番です。\n", c.turn)
		c.board.Draw()
		fmt.Print("> ")
		sc.Scan()
		input := sc.Text()
		if input == "quit" {
			break
		}
		if input == "pass" {
			fmt.Println("パスします。")
			c.changeTurn()
			continue
		}
		if c.validate(input) != nil {
			fmt.Println("入力内容が不正です。「列番号,行番号」の形式で入力して下さい。例）左上隅の場合: 0,0")
			continue
		}

		inputs := strings.Split(input, ",")
		for i, v := range inputs {
			inputs[i] = strings.TrimSpace(v)
		}
		x, _ := strconv.Atoi(inputs[0])
		y, _ := strconv.Atoi(inputs[1])
		address := model.NewAddress(x, y)
		if err := c.board.Put(c.turn, address); err != nil {
			fmt.Println("そこには置けません。")
		}
		c.changeTurn()
	}
	c.end()
}

func (c *Controller) end() {
	whiteNum := c.board.CountWhite()
	blackNum := c.board.CountBlack()
	fmt.Printf("白: %d\n", whiteNum)
	fmt.Printf("黒: %d\n", blackNum)
	switch {
	case whiteNum > blackNum:
		fmt.Println("白の勝利です。")
	case blackNum > whiteNum:
		fmt.Println("黒の勝利です。")
	default:
		fmt.Println("引き分けです。")
	}
}
