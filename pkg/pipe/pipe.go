package pipe

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var NoPipe = errors.New("no data is being piped")

func Read() ([]byte, error) {
	// Check if there is data available on stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped to stdin
		var data []byte
		reader := bufio.NewReader(os.Stdin)

		for {
			chunk, err := reader.ReadBytes('\n')
			if err != nil && err == io.EOF {
				data = append(data, chunk...)
				break
			}
			data = append(data, chunk...)
		}

		return data, nil
	}

	// No data is being piped
	return nil, NoPipe
}
