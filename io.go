package transitivedep

import (
	"bufio"
	"log"
	"strings"
	"os"
	"sort"
)

type Input interface {
	ReadInput(td Dependency)
}

type Output interface {
	WriteOutput(td Dependency)
}

type FileIO struct {
	Filename  string
}

func (f *FileIO) ReadInput(td Dependency) {
	input, err := os.Open(f.Filename)

	if err != nil {
		log.Fatal("Cannot open the file!!!")
		return
	}
	defer input.Close()
	
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		td.AddDirect(line[0], line[1:])	
	}
}

func (f *FileIO) WriteOutput(td Dependency) {
	output, err := os.Create(f.Filename)

	if err != nil {
		log.Fatal("Cannot open the file to write it!!!")
		return
	}
	defer output.Close()
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	keys := []string{}
	for key := range td.GetDependency() {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		line := key+"   "
		for _, dep := range td.DependencyFor(key) {
			line += dep + " "
		}
		writer.WriteString(line+"\n")
	}
}

type TransitiveDependencyApp struct {
	input      Input
	output     Output
	dependency Dependency
}

func NewTransDependencyApp(input Input, output Output, dependency Dependency) TransitiveDependencyApp {
	return TransitiveDependencyApp{input: input, output: output, dependency: dependency}
}

func (t TransitiveDependencyApp) Analyze() {
	t.input.ReadInput(t.dependency)
	t.output.WriteOutput(t.dependency)
}
