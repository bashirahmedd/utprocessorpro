package main

import (
	// "fmt"
	// "log"
	video "utube-downloader/video"
	cutils "utube-downloader/common"

)

func main() {
	// read config 
	config:=cutils.GetShareConfig()
	// config.RootPath = cutils.GetAppBasePath()
	// log.Println(config.VideoId)
	
	video.Download(config)
	// video.PrintHello()
}
