package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func GetAppBasePath() string{
	currdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}	
	return filepath.Dir(currdir)
}

const (	
	KB = int(1024)
	MB = int(KB * 1024)
	GB = int(MB * 1024)
	TB = int(GB * 1024)
)
func SizeReadable(length int, decimals int) (out string) {
	var unit string
	var i int
	var remainder int
	// Get whole number, and the remainder for decimals
	if length > TB {
		unit = "TB"
		i = length / TB
		remainder = length - (i * TB)
	} else if length > GB {
		unit = "GB"
		i = length / GB
		remainder = length - (i * GB)
	} else if length > MB {
		unit = "MB"
		i = length / MB
		remainder = length - (i * MB)
	} else if length > KB {
		unit = "KB"
		i = length / KB
		remainder = length - (i * KB)
	} else {
		return strconv.Itoa(length) + " B"
	}
	if decimals == 0 {
		return strconv.Itoa(i) + " " + unit
	}
	// This is to calculate missing leading zeroes
	width := 0
	if remainder > GB {
		width = 12
	} else if remainder > MB {
		width = 9
	} else if remainder > KB {
		width = 6
	} else {
		width = 3
	}
	// Insert missing leading zeroes
	remainderString := strconv.Itoa(remainder)
	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}
	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}
	return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
}