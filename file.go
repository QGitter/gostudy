package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

//写入文件通过io包
func WriteFileByIo(filename string, content string) int {
	file, err := createFile(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	return n
}

//写入文件通过file指针
func WriteFileByFile(filename string, content string) int {
	file, err := createFile(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n, e := file.WriteString(content)
	if e != nil {
		panic(err)
	}
	return n
}

//写入文件通过bufio包
func WriteFileByBufio(filename string, content string) int {
	file, err := createFile(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	i, e := writer.WriteString(content)
	if e != nil {
		panic(e)
	}
	return i
}

//写入文件通过ioutil包
func WriteFileByIoutil(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0666) //不能追加数据
}

//读文件通过ioutil包
func ReadFileByIoutil(filename string) (string, error) {
	if !checkFileIsExist(filename) {
		return "", os.ErrNotExist
	}
	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(e)
	}
	return string(bytes), nil
}

//读文件通过os 和ioutil包
func ReadFileByOs(filename string) (string, error) {
	if !checkFileIsExist(filename) {
		return "", os.ErrNotExist
	}
	file, e := os.OpenFile(filename, os.O_RDWR, 0644)
	if e != nil {
		return "", os.ErrNotExist
	}
	defer file.Close()
	bytes, i := ioutil.ReadAll(file)
	if i != nil {
		return "", os.ErrNotExist
	}
	return string(bytes), nil
}

//读文件通过bufio os包
func ReadFileByBufio(filename string) (string, error) {
	var s string
	if !checkFileIsExist(filename) {
		return "", os.ErrNotExist
	}
	file, e := os.Open(filename)
	if e != nil {
		return "", errors.New("file open failure")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		sr, ir := reader.ReadString('\n')
		if ir == io.EOF {
			return s, nil
		}
		s += sr
	}
}

//读文件通过file包
func ReadFileByFile(filename string) (string, error) {
	var bufAppend []byte
	if !checkFileIsExist(filename) {
		return "", os.ErrNotExist
	}
	file, e := os.Open(filename)
	if e != nil {
		return "", errors.New("file open failure")
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF && n == 0 {
			return string(bufAppend), nil
		}
		bufAppend = append(bufAppend, buf[:n]...)
	}
}

//创建文件指针
func createFile(filename string) (*os.File, error) {
	var file *os.File
	var e error
	if checkFileIsExist(filename) {
		file, e = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		file, e = os.Create(filename)
	}
	return file, e
}

//判断文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, e := os.Stat(filename); os.IsNotExist(e) {
		exist = false
	}
	return exist
}
