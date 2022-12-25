package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Connection struct {
	Index                uint
	StartByte            uint64
	EndByte              uint64
	ByteLength           uint64
	DownloadedBytesCount chan uint64
	TmpFileName          string
}

func setupDir() (string, string) {
	tmpPath := "./tmp"
	resultPath := "./result"

	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		if err := os.Mkdir("./tmp", 0744); err != nil {
			log.Fatal("Failed to create tmp dir", err)
		}
	}

	if _, err := os.Stat(resultPath); os.IsNotExist(err) {
		if err := os.Mkdir("./result", 0744); err != nil {
			log.Fatal("Failed to create result dir", err)
		}
	}

	return tmpPath, resultPath
}

func main() {
	setupDB()

	e := echo.New()

	e.Use(middleware.Recover())

	TaskRoutes(e)

	e.Logger.Fatal(e.Start(":4321"))

	// if len(os.Args) <= 1 {
	// 	log.Fatal("1 argument required")
	// }

	// url := os.Args[1]
	// if url == "" {
	// 	log.Fatal("url is required")
	// }

	// tmpPath, resultPath := setupDir()

	// maxConnection := uint(4)

	// fm, err := GetFileMetadata(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if !fm.AcceptRanges {
	// 	maxConnection = 1
	// }

	// connections := []Connection{}
	// cases := []reflect.SelectCase{}
	// progressBars := []*mpb.Bar{}
	// byteLenPerConn := fm.ContentLength / uint64(maxConnection)
	// contentSizeRemaining := fm.ContentLength

	// var wg sync.WaitGroup
	// p := mpb.New(mpb.WithWidth(40))

	// for i := uint(0); i < maxConnection; i++ {
	// 	startByte := uint64(i) * byteLenPerConn
	// 	endByte := startByte + byteLenPerConn - 1
	// 	if i == maxConnection-1 {
	// 		endByte = startByte + contentSizeRemaining
	// 	}

	// 	contentSizeRemaining -= byteLenPerConn

	// 	connection := Connection{
	// 		Index:                i,
	// 		StartByte:            startByte,
	// 		EndByte:              endByte,
	// 		ByteLength:           endByte - startByte + 1,
	// 		DownloadedBytesCount: make(chan uint64),
	// 		TmpFileName:          GenerateTmpFileName(tmpPath),
	// 	}

	// 	connections = append(connections, connection)
	// 	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(connection.DownloadedBytesCount)})

	// 	wg.Add(1)
	// 	go func(i uint) {
	// 		defer wg.Done()

	// 		if err := Download(url, &connections[i]); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}(i)

	// 	name := fmt.Sprintf("#%d", i+1)
	// 	bar := p.New(int64(connections[i].ByteLength),
	// 		mpb.BarStyle().Lbound("|").Filler("█").Tip("").Padding("░").Rbound("|"),
	// 		mpb.PrependDecorators(
	// 			decor.Name(name),
	// 		),
	// 		mpb.AppendDecorators(
	// 			decor.AverageSpeed(decor.UnitKiB, "% .2f"),
	// 			decor.Name(" | "),
	// 			decor.OnComplete(
	// 				decor.AverageETA(decor.ET_STYLE_GO),
	// 				"done",
	// 			),
	// 			decor.Name(" | "),
	// 			decor.Percentage(decor.WC{W: 4}),
	// 		),
	// 	)

	// 	progressBars = append(progressBars, bar)
	// }

	// wg.Wait()

	// HandleGracefulExit(connections)

	// remaining := len(cases)

	// for remaining > 0 {
	// 	chosen, val, ok := reflect.Select(cases)
	// 	if !ok {
	// 		// The chosen channel has been closed, so zero out the channel to disable the case
	// 		cases[chosen].Chan = reflect.ValueOf(nil)
	// 		remaining -= 1
	// 		continue
	// 	}

	// 	bar := progressBars[chosen]
	// 	bar.SetCurrent(int64(val.Uint()))
	// }

	// time.Sleep(1 * time.Second)

	// fOut, err := os.Create(filepath.Join(resultPath, fm.FileName))
	// if err != nil {
	// 	log.Fatal("Failed to create file: ", err)
	// }
	// defer fOut.Close()

	// // merge file & cleanup tmp
	// for _, conn := range connections {
	// 	fIn, err := os.Open(conn.TmpFileName)
	// 	if err != nil {
	// 		log.Fatal("Failed to open tmp file: ", err)
	// 	}
	// 	defer fIn.Close()

	// 	buf := make([]byte, 1024*1024) // 1 MiB
	// 	for {
	// 		n, err := fIn.Read(buf)
	// 		if err != nil && err != io.EOF {
	// 			log.Fatalf("Failed to read file: %s\nError: %s", conn.TmpFileName, err.Error())
	// 		}

	// 		if n == 0 {
	// 			break
	// 		}

	// 		fOut.Write(buf[:n])
	// 	}

	// 	os.Remove(conn.TmpFileName)
	// }
}
