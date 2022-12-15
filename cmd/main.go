package main


import (
	// "fmt"
	// "log"
	// "fmt"
	cutils "utube-downloader/common"
	video "utube-downloader/video"
)

func main() {

	// read config
	// fmt.Println(cutils.GetAppBasePath())
	config:=cutils.GetShareConfig()
	// config.RootPath = cutils.GetAppBasePath()
	// log.Println(config.VideoId)

	video.Download(config)
	// video.PrintHello()
}



// package main
/* 
import (
	// "crypto/rand"
	// "io"
	// "io/ioutil"

	pb "github.com/cheggaaa/pb/v3"
	"os"
	"time"
)

func main() {
	// var limit int64 = 1024 * 1024 * 500

	// we will copy 500 MiB from /dev/rand to /dev/null
	// reader := io.LimitReader(rand.Reader, limit)
	// writer := ioutil.Discard

	var limit int=1000
	// start new bar
	bar:=pb.New(limit)

	// refresh info every second (default 200ms)
	bar.SetRefreshRate(time.Second)

	// force set io.Writer, by default it's os.Stderr
	bar.SetWriter(os.Stdout)

	// bar will format numbers as bytes (B, KiB, MiB, etc)
	bar.Set(pb.Bytes, true)

	// bar use SI bytes prefix names (B, kB) instead of IEC (B, KiB)
	bar.Set(pb.SIBytesPrefix, true)

	// set custom bar template
	// bar.SetTemplateString(myTemplate)

	// check for error after template set
	if err := bar.Err(); err != nil {
		return
	}
	bar.Start()

	// create proxy reader
	// barReader := bar.NewProxyReader(reader)

	// copy from proxy reader
	// io.Copy(writer, barReader)
	for ind:=int(1); ind <= limit; ind++ {
		time.Sleep(1 * time.Second)
		bar.Add(100)
	}  
	// finish bar
	bar.Finish()
} */