package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type FileMetadata struct {
	ContentLength uint64
	AcceptRanges  bool
	FileName      string
}

func GetFileMetadata(url string) (res FileMetadata, err error) {
	resp, err := http.Head(url)
	if err != nil {
		return res, err
	}

	headerContentLength := resp.Header.Get("Content-Length")
	res.ContentLength, err = strconv.ParseUint(headerContentLength, 10, 0)
	if err != nil {
		return res, err
	}

	headerAcceptRanges := resp.Header.Get("Accept-Ranges")
	res.AcceptRanges = headerAcceptRanges == "bytes"

	urlParts := strings.Split(url, "/")
	res.FileName = urlParts[len(urlParts)-1]

	headerContentDisposition := resp.Header.Get("Content-Disposition")
	if strings.HasPrefix(headerContentDisposition, "attachment") {
		r, err := regexp.Compile(`filename.+"(.+)"`)
		if err != nil {
			return res, nil
		}

		extracted := r.FindStringSubmatch(headerContentDisposition)
		if len(extracted) > 1 {
			res.FileName = extracted[1]
		}
	}

	return res, err
}

func Download(url string, connection *Connection) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error to create HTTP request: ", err.Error())
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", connection.StartByte, connection.EndByte))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error to do HTTP request: ", err.Error())
	}

	f, err := os.Create(connection.TmpFileName)
	if err != nil {
		log.Fatal("Failed to create file: ", err)
	}

	currentDownloaded := uint64(0)

	go func() {
		defer f.Close()
		defer close(connection.DownloadedBytesCount)

		buf := make([]byte, 1024*1024) // 1 MiB

		for {
			n, err := resp.Body.Read(buf)
			if err != nil && err != io.EOF {
				log.Fatal("Error reading HTTP response: ", err.Error())
			}

			if n == 0 {
				break
			}

			currentDownloaded += uint64(n)
			connection.DownloadedBytesCount <- currentDownloaded
			f.Write(buf[:n])
		}
	}()

	return nil
}

func HandleGracefulExit(connections []Connection) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-c
		log.Println("cleaning up tmp")
		for _, conn := range connections {
			os.Remove(conn.TmpFileName)
		}
		os.Exit(1)
	}()
}

func GenerateTmpFileName(prefix string) string {
	ns := time.Now().UnixNano()
	nsStr := []byte(fmt.Sprintf("%d", ns))

	rndBytes := make([]byte, 4)
	rand.Read(rndBytes)

	hash := md5.Sum(append(nsStr, rndBytes...))
	hashStr := hex.EncodeToString(hash[:])

	return filepath.Join(prefix, hashStr)
}
