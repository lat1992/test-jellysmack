package main

import (
	"bufio"
	"flag"
	"fmt"
	"jellysmack-test/internal/service"
	"os"
	"strconv"
	"strings"
	"sync"
)

type dependence struct {
	mowers []service.MowerService
}

func main() {
	var input string
	flag.StringVar(&input, "f", "./input.txt", "path of input file")
	flag.Parse()
	dep, err := initialisation(input)
	if err != nil {
		return
	}
	run(dep)
}

func run(dep dependence) {
	var wg sync.WaitGroup
	for i := range dep.mowers {
		wg.Add(1)
		j := i
		go func() {
			fmt.Println(dep.mowers[j].Run())
			wg.Done()
		}()
	}
	wg.Wait()
}

func initialisation(input string) (dependence, error) {
	readFile, err := os.Open(input)
	if err != nil {
		return dependence{}, err
	}
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	var lines []string

	fs.Scan()
	lim := fs.Text()
	for fs.Scan() {
		lines = append(lines, fs.Text())
	}
	if err = readFile.Close(); err != nil {
		return dependence{}, err
	}

	if len(lines)%2 != 0 {
		return dependence{}, fmt.Errorf("the instructions are wrong")
	}

	limX, limY, err := getLimit(lim)
	if err != nil {
		return dependence{}, err
	}

	yard, err := service.NewYardService(limX, limY)
	if err != nil {
		return dependence{}, err
	}

	var mowers []service.MowerService
	for i := 0; i < len(lines); i += 2 {
		posX, posY, posO, err := getPosition(lines[i])
		if err != nil {
			return dependence{}, err
		}
		m, err := service.NewMowerService(posX, posY, posO, lines[i+1], yard)
		if err != nil {
			return dependence{}, err
		}
		mowers = append(mowers, m)
	}

	return dependence{
		mowers: mowers,
	}, nil
}

func getLimit(input string) (int, int, error) {
	limit := strings.Split(input, " ")

	if len(limit) != 2 {
		return 0, 0, fmt.Errorf("the limit is wrong")
	}

	x, err := strconv.Atoi(limit[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(limit[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func getPosition(input string) (int, int, string, error) {
	pos := strings.Split(input, " ")

	if len(pos) != 3 {
		return 0, 0, "", fmt.Errorf("the position is wrong")
	}

	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return 0, 0, "", err
	}
	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return 0, 0, "", err
	}

	return x, y, pos[2], nil
}
