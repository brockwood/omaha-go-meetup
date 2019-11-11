package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

type Command struct {
	Interpreter *interp.Interpreter
	Description string
}

var (
	cmds map[string]*Command
)

func init() {
	files, err := ioutil.ReadDir(`./cmds`)
	if err != nil {
		fmt.Println("Error reading commands:", err)
		return
	}
	cmds = make(map[string]*Command)
	for _, file := range files {
		fileData, err := ioutil.ReadFile(filepath.Join(`./cmds`, file.Name()))
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
			return
		}
		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols)
		_, err = i.Eval(string(fileData))
		if err != nil {
			panic(err)
		}
		commandName, err := i.Eval(`script.Command`)
		if err != nil {
			panic(err)
		}
		commandDescription, err := i.Eval(`script.Description`)
		c := Command{
			Interpreter: i,
			Description: commandDescription.String(),
		}
		cmds[commandName.String()] = &c
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: logs <command>")
		fmt.Println("Available commands:")
		for commandName, command := range cmds {
			fmt.Printf("\t%s\t- %s\n", commandName, command.Description)
		}
		os.Exit(1)
	}
	cmd := cmds[os.Args[1]]
	if cmd == nil {
		fmt.Printf("The command '%s' is not available.\n", os.Args[1])
		os.Exit(1)
	}
	cmdFuncInterface, err := cmd.Interpreter.Eval("script.Cmd")
	if err != nil {
		panic(err)
	}
	cmdFunc := cmdFuncInterface.Interface().(func([]time.Time, []string))
	logTimes, logLines := generateDummyData()
	cmdFunc(logTimes, logLines)
}

func generateDummyData() ([]time.Time, []string) {
	times := make([]time.Time, 0, 86400)
	logLines := make([]string, 0, 86400)
	startTime := time.Now().Add(-time.Hour * 24)
	for seconds := 0; seconds < 86400; seconds++ {
		newTime := startTime.Add(time.Second * time.Duration(seconds))
		times = append(times, newTime)
		n, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			panic(err)
		}
		logLines = append(logLines, fmt.Sprintf("%d", n))
	}

	return times, logLines
}
