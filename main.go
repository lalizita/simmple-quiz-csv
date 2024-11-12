package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Print("Seja bem vindo(a) ao quiz")
	fmt.Print("\033[33;1m Escreva seu nome:\033[0m\n")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	g.Name = name

	fmt.Printf("Vamos ao jogo %s \n", g.Name)
}

func (g *GameState) ProccessCSV() {
	f, err := os.Open("quiz-go.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		panic("Error reading csv")
	}

	for _, record := range records {
		question := Question{
			Text:    record[0],
			Options: record[1:5],
			Answer:  toInt(record[5]),
		}

		g.Questions = append(g.Questions, question)
	}
}

func (g *GameState) Run() {
	for i, q := range g.Questions {
		fmt.Printf("\033[33;1m %d. %s\033[0m\n", i+1, q.Text)

		for j, opt := range q.Options {
			fmt.Printf("[%d] %s\n", j+1, opt)
		}

		fmt.Println("Digite a alternativa:")

		reader := bufio.NewReader(os.Stdin)
		read, _ := reader.ReadString('\n')
		answer, err := strconv.Atoi(read[:len(read)-1])
		if err != nil {
			panic(err)
		}

		if answer == q.Answer {
			fmt.Println("Parabéns você acertou!")
			g.Points += 10
		} else {
			fmt.Println("Ops! Errou! Vamos para a próxima...")
			fmt.Println("------------------------------------")
		}
	}
}

func main() {
	game := &GameState{Points: 0}
	go game.ProccessCSV()

	game.Init()
	game.Run()
	fmt.Printf("Fim de jogo, você fez %d pontos", game.Points)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
