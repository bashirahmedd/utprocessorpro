package common

import (
	"log"
	"os"
	"path/filepath"
)

func GetAppBasePath() string{
	currdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}	
	return filepath.Dir(currdir)
}