package video

import (
	"fmt"
	"io"
	// "io"
	"encoding/json"
	"log"
	"os"
	"time"

	cutils "utube-downloader/common"

	youtube "github.com/kkdai/youtube/v2"
)

type youtubeVideo struct {
	Id          string `json:"id"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int32  `json:"duration"`
	ViewCount   int32  `json:"view_count"`
	Uploader    string `json:"uploader"`
	// Seller struct {
	//     Id int `json:"id"`
	//     Name string `json:"name"`
	//     CountryCode string `json:"country_code"`
	// } `json:"seller"`
}

func Download(cfg *cutils.ConfigVideo) {

	if status := validateinputfile(cfg); status {

		// backup video ids
		if n, err := copyFile(cfg.VideoId, cfg.BackupId); err == nil && n > 0 {
			log.Println("Input id list is backed up.")
			log.Printf("Backed-up to : %s", cfg.BackupId)

			counter := time.Now().UnixMilli()
			var vids []youtubeVideo
			if jsonIn, err := os.ReadFile(cfg.VideoId); err == nil {
				if err := json.Unmarshal(jsonIn, &vids); err == nil {
					/*
						var objMap []map[string]interface{}
						if err := json.Unmarshal(jsonIn, &objMap); err == nil {
							fmt.Println(objMap)
						}
					*/
					for _, v := range vids {
						youtubeClient(cfg, &v)
					}
				}
			}
			fmt.Println(counter)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Println("Initial state is invalid, please check:")
		log.Printf("Either, the file is empty:  %s", cfg.VideoId)
		log.Printf("or the file is not empty: %s", cfg.NextIterId)
		os.Exit(1)
	}
}

func youtubeClient(cfg *cutils.ConfigVideo, yt *youtubeVideo) {
	
	client := youtube.Client{}
	video, err := client.GetVideo(yt.Url)
	if err != nil {
		log.Println(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	ind := getFormatIndex(formats)
	stream, _, err := client.GetStream(video, &formats[ind])
	if err != nil {
		log.Println(err)
	}

	file, err := os.Create(cfg.VideoDlPath+"_"+"NEW GO"+"_"+yt.Url+".mp4")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		log.Println(err)
	}
}

func getFormatIndex(list youtube.FormatList) int {

	for ind, f := range list {
		h := f.Height
		if h > 480 {
			continue
		}
		return ind
	}
	return 0
}

/*
baseUrl='https://www.youtube.com/watch?v='
target='/home/naji/Downloads/temp/ytdown/'
*/

func validateinputfile(cfg *cutils.ConfigVideo) bool {

	//check if exist
	exists(cfg.VideoDlPath)
	exists(cfg.VideoId)
	exists(cfg.NextIterId)
	exists(cfg.LogPath)

	//check if contents are valid
	vsize := fileSize(cfg.VideoId)
	nsize := fileSize(cfg.NextIterId)

	if vsize > 0 && nsize == 0 {
		return true
	}
	return false
}

func exists(path string) (bool, error) {

	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		log.Fatalln("Path Not Found:", path)
	}
	return true, nil
}

func fileSize(fname string) int64 {

	fInfo, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	fsize := fInfo.Size()
	fmt.Printf("%s size is %d bytes\n", fname, fsize)
	return fsize
}

// func copyFile(SrcF string, backupF string) bool {

// 	return true
// }

func copyFile(src, dst string) (n int64, err error) {
	r, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	w, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	return w.ReadFrom(r)
}

/*
func copyFileContents(src, dst string) (err error) {

    in, err := os.Open(src);
    if err != nil {
        return
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}
*/
