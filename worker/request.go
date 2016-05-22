package worker

import (
	"bufio"
	"errors"
	"io"
	"log"

	"blackbox/constants"
)

func ParseRequest(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	if scanner.Text() == constants.KEY_REQUEST_HEADER {
		scanner.Scan()
		key := scanner.Text()
		if len(key) > 0 {
			return key, nil
		}
	} else {
		err := errors.New("message header error.")
		log.Println(err, scanner.Text())
		return "", err
	}
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
		return "", err
	}
	return "", nil
}
