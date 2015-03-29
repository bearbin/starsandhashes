package main

type player bool

var gamestate *miniboard

func init() {
	gamestate = &miniboard{
		contents: [XY][XY]miniboard{
			contents: [XY][XY]singlecell{},
		},
	}
}

func main() {
	println("LOL")
}
