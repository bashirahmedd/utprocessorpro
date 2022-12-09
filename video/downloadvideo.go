package video

import (
	"fmt"
	// "io"
	"log"
	"os"
	"time"

	cutils "utube-downloader/common"
)

func Download(cfg *cutils.ConfigDownloader) {

	if validateinputfile(cfg) {
		copyFile(cfg.VideoId, cfg.BackupId)

		counter := time.Now().UnixMilli()
		fmt.Println(counter)
	}
}

/* # validate state
if [ -s $in_video_list -a ! -s $try_again_video_list ];then
   fn_say "Initial state is good..."
   cat $in_video_list > $backup_id           # backup intial ids
   fn_say "Input id list is backed up."
   echo "backed-up to "$backup_id
else
   fn_say "Initial state is invalid, please check:"
   fn_say "Either, the file is empty:  $in_video_list"
   fn_say "or the file is not empty: $try_again_video_list"
   exit 1
fi */

func validateinputfile(cfg *cutils.ConfigDownloader) bool {

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



func copyFile(src, dst string)  {
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
	w.ReadFrom(r)
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