package video

import (
	"fmt"
	"io"
	// "reflect"
	"strconv"

	// "io"
	"encoding/json"
	"log"
	"os"
	"time"

	pb "github.com/cheggaaa/pb/v3"
	youtube "github.com/kkdai/youtube/v2"
	cutils "utube-downloader/common"
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


// failedVids = make([]youtubeVideo)

func Download(cfg *cutils.ConfigVideo) {

	failedVids := []youtubeVideo{}
	if status := validateinputfile(cfg); status {

		// backup video ids
		if n, err := copyFile(cfg.VideoId, cfg.BackupId); err == nil && n > 0 {
			log.Println("Input id list is backed up.")
			log.Printf("Backed-up to : %s", cfg.BackupId)

			var vids []youtubeVideo
			if jsonIn, err := os.ReadFile(cfg.VideoId); err == nil {
				if err := json.Unmarshal(jsonIn, &vids); err == nil {
					/*
						var objMap []map[string]interface{}
						if err := json.Unmarshal(jsonIn, &objMap); err == nil {
							fmt.Println(objMap)
						}
					*/
					fmt.Printf("Starting download of %d tasks\n", len(vids))
					fmt.Println("-----------------------------")
					for _, v := range vids {						
						if !youtubeClient(cfg, &v) {
							failedVids = append(failedVids, v)
						}
						fmt.Println("-----------------------------")
					}
				}
			}
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

func youtubeClient(cfg *cutils.ConfigVideo, yt *youtubeVideo) bool{

	counter := time.Now().UnixMilli()
	client := youtube.Client{}
	video, err := client.GetVideo(yt.Url)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Target file id: %s", yt.Url)

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	ind := getFormatIndex(formats)
	stream, size, err := client.GetStream(video, &formats[ind])
	if err != nil {
		log.Println(err)
	}
	log.Printf("Target video size : %s", cutils.SizeReadable(int(size), 1))
	defer stream.Close()

	fname := cfg.VideoDlPath + strconv.Itoa(int(counter)) + "_" + yt.Title + "_" + yt.Url + ".mp4"
	partialFname := fname + ".part"
	log.Printf("Destination: %s", fname)
	file, err := os.Create(partialFname)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// _, err = io.Copy(file, stream)
 	if _, err = writeBufferedStream(file, stream, nil, int(size)); err!=nil{
		log.Println(err)
		return false
	} else {
		fi, err := os.Stat(fname)
		if err != nil {
			log.Fatalln(err)
		}
		sz := fi.Size()
		if size != sz { // incomplete dowload
			log.Printf("Failed: %s Stream: %d Download: %d ", yt.Url, size, sz)
			return false
		} else {
			log.Printf("Success: %s ", yt.Url)
			e := os.Rename(partialFname, fname)
			if e != nil {
				log.Fatal(e)
			}
			return true
		}
	}
}

func writeBufferedStream(dst io.Writer, src io.Reader, buf []byte, szdl int) (written int64, err error) {

	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	var totBytes int64 = 0
	bar := initPBar(szdl)
	bar.Start()
	for {
		nr, er := src.Read(buf)
		totBytes += int64(nr)
		bar.Add(nr)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					//ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				//err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	bar.Finish()
	return written, err
}

func initPBar(sz int) *pb.ProgressBar {

	bar := pb.New(sz)
	bar.SetRefreshRate(time.Second)
	bar.SetWriter(os.Stdout)
	bar.Set(pb.Bytes, true)
	bar.Set(pb.SIBytesPrefix, true)
	bar.SetTemplateString(string(pb.Full))

	if err := bar.Err(); err != nil {
		return nil
	}
	return bar
}

func getFormatIndex(list youtube.FormatList) int {

	for ind, f := range list {
		h := f.Height
		if h > 480 {
			continue
		}
		/*
					//prints the selected format
			 		func(fr youtube.Format) {
						v := reflect.ValueOf(fr)
						typeOfS := v.Type()
						fields := []string{"InitRange", "IndexRange", "Cipher", "URL", "AudioChannels", "ProjectionType"}

						for i := 0; i < v.NumField(); i++ {

							isThere := func(s []string, str string) bool {
								for _, v := range s {
									if v == str {
										return true
									}
								}
								return false
							}(fields, typeOfS.Field(i).Name)
							if isThere || v.Field(i).Interface() == nil {
								continue
							} else {
								fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
							}
						}

					}(f) */

		return ind
	}
	return 0
}

func validateinputfile(cfg *cutils.ConfigVideo) bool {

	//check if exist
	exists(cfg.VideoDlPath)
	exists(cfg.VideoId)
	exists(cfg.NextIterId)
	exists(cfg.LogPath)

	//check if contents are valid
	fmt.Println("Input File: ")
	vsize := fileSize(cfg.VideoId)
	fmt.Println("Previously Failed IDs File: ")
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
	fmt.Printf("\t %s size is %d bytes\n", fname, fsize)
	return fsize
}

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
