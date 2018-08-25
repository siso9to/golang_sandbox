package main

import "fmt"

type Ide interface {
	Debug() string
	Run() int
}

type RubyMine struct{}

func (r *RubyMine) Debug() string {
	return "Buuuuuuuuuug!!!"
}

func (r *RubyMine) Run() int {
	return 0
}

type Goland struct{}

func (g *Goland) Debug() string {
	return "Go go go!"
}

func (g *Goland) Run() int {
	return -1
}

// Ide interfaceに無いメソッド
func (g *Goland) ExitIde() int {
	return 0
}

func main() {
	ides := []Ide{new(RubyMine), new(Goland)}

	for _, i := range ides {
		debugAndRun(i)
	}
}

func debugAndRun(ide Ide) {
	fmt.Printf("debug message: %s\n", ide.Debug())
	fmt.Printf("Process finished with exit code: %d\n", ide.Run())
}
