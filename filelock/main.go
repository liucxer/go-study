package main

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateTraceID(length int64) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result string
	for i := 0; i < int(length); i++ {
		result += string(charset[rand.Intn(len(charset))])
	}

	return result
}

func WriteFile(path string, sync string) error {
	var (
		err error
		tmpFile *os.File
		bts []byte
	)

	rand.Seed(time.Now().UnixNano())
	bts = []byte(GenerateTraceID(int64(rand.Intn(100))+ int64(10)) )
	if sync == "SYNC" {
		tmpFile, err = os.OpenFile(path, os.O_SYNC|os.O_RDWR, os.ModePerm)
	} else {
		tmpFile, err = os.OpenFile(path, os.O_RDWR, os.ModePerm)
	}

	if err != nil {
		logrus.Errorf("os.OpenFile err:%v,path:%s", err, path)
		return err
	}
	defer func() {
		_ = tmpFile.Close()
	}()

	_, err = tmpFile.Write(bts)
	if err != nil {
		logrus.Errorf("tmpFile.Write err:%v,path:%s", err, path)
		return err
	}
	logrus.Infof("write succcess bts:%s", bts)
	return err
}

func ReadFile(path string, sync string) error {
	var (
		err error
		tmpFile *os.File
		bts []byte
	)

	if sync == "SYNC" {
		tmpFile, err = os.OpenFile(path, os.O_SYNC|os.O_RDWR, os.ModePerm)
	} else {
		tmpFile, err = os.OpenFile(path, os.O_RDWR, os.ModePerm)
	}
	if err != nil {
		logrus.Errorf("os.OpenFile err:%v,path:%s", err, path)
		return err
	}
	defer func() {
		_ = tmpFile.Close()
	}()

	bts = make([]byte,100)
	_, err = tmpFile.Read(bts)
	if err != nil {
		logrus.Errorf("tmpFile.Read err:%v,path:%s", err, path)
		return err
	}

	logrus.Infof("read succcess bts:%s", bts)
	return err
}

func main() {
	filePath := os.Args[1]
	threadNumStr := os.Args[2]
	sync := os.Args[3]

	threadNum, err := strconv.Atoi(threadNumStr)
	if err != nil {
		logrus.Errorf("strconv.Atoi err:%v", err)
		return
	}

	for i :=0; i < threadNum; i++ {
		go func() {
			for {
				rand.Seed(time.Now().UnixNano())
				randNum := rand.Intn(100)
				time.Sleep(time.Duration(randNum) * time.Millisecond)
				ReadFile(filePath, sync)
			}
		}()
		go func() {
			for {
				rand.Seed(time.Now().UnixNano())
				randNum := rand.Intn(100)
				time.Sleep(time.Duration(randNum) * time.Millisecond)
				WriteFile(filePath, sync)
			}
		}()
	}

	for {
		time.Sleep(time.Second)
	}
}
