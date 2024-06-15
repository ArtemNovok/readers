package readers

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFileInc(fileName string) (map[rune]int, error) {
	mp := make(map[rune]int)
	var offset int64 = 0
	for {
		n, err := ReadFileWithOffSet(fileName, offset, 8, mp)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println(err)
				return mp, nil
			}
			return nil, err
		}
		offset += int64(n)
	}
}

func ReadFileWithOffSet(fileName string, offset int64, size int64, mp map[rune]int) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	_, err = f.Seek(offset, 0)
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(f)
	buf := make([]byte, size)
	n, err := r.Read(buf)
	if err != nil {
		if n == 0 || err == io.EOF {
			return 0, fmt.Errorf("whole file is reded %w", err)
		}
		return 0, err
	}
	buf = buf[:n]
	str := fmt.Sprintf("%s", buf)
	for _, char := range str {
		mp[char]++
	}
	log.Println(fmt.Sprintf("%s", buf))
	return n, nil
}
