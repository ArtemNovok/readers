package readers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFileIncMap(fileName string, size int64) (map[rune]int, error) {
	mp := make(map[rune]int)
	var offset int64 = 0
	for {
		n, err := ReadFileWithOffSetMap(fileName, offset, size, mp)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println(err)
				return mp, nil
			}
			return nil, err
		}
		if int64(n) < size {
			offset += int64(n)
		} else {
			offset += size
		}
	}
}

func ReadFileWithOffSetMap(fileName string, offset int64, size int64, mp map[rune]int) (int, error) {
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
	buf = nil
	return n, nil
}

func ReadFileWithOffSet(fileName string, offset int64, size int64) (int, []byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()
	f.Seek(offset, 0)
	r := bufio.NewReader(f)
	buf := make([]byte, size)
	n, err := r.Read(buf)
	if err != nil {
		if n == 0 || err == io.EOF {
			return 0, nil, fmt.Errorf("whole file is reded %w", err)
		}
		return 0, nil, err
	}
	buf = buf[:n]
	return n, buf, nil
}

func ReadFileInc(filename string, size int64) error {
	var offset int64 = 0
	var buf bytes.Buffer
	for {
		n, bytes, err := ReadFileWithOffSet(filename, offset, size)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println(fmt.Sprintf("%s", buf.Bytes()))
				log.Println(err)
				return nil
			}
			return err
		}
		buf.Write(bytes)
		if n < int(size) {
			offset += int64(n)
		} else {
			offset += size
		}

	}
}
