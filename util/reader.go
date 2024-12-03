package util

import (
	"bufio"
	"os"
)

func Read(path string, ch chan<- string) error {
	file, err := os.Open(path)
	if err != nil {
		close(ch)
		return nil
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	close(ch)

	if err := scanner.Err(); err != nil {
		return nil
	}

	return nil
}
