package common

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/spf13/viper"
)

type ConfigVideo struct {
	VideoId     string `yaml:"video.in_video_list"`
	NextIterId  string `yaml:"video.next_iteration_file"`
	BackupId    string `yaml:"video.log_id"`
	LogPath     string `yaml:"video.log_path"`
	VideoDlPath string `yaml:"video.video_dl_path"`
}

type configParam struct {
	VideoDlPath string
	Counter     int64
}

func GetShareConfig() *ConfigVideo {

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")  // path required for debugging
	viper.SetConfigName("config.yaml") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Config file not found: %w", err))
		} else {
			panic(fmt.Errorf("Config file error: %w", err))
		}
	}
	//fmt.Println(viper.Get("analytics"))
	//fmt.Println(viper.Get("keywords"))

	var vcfg ConfigVideo
	vcfg.VideoDlPath = viper.GetString("video.video_dl_path")
	vcfg.VideoId = viper.GetString("video.in_video_list")
	vcfg.NextIterId = viper.GetString("video.next_iteration_file")
	vcfg.BackupId = viper.GetString("video.log_id")
	vcfg.LogPath = viper.GetString("video.log_path")

	// err1 := cleanenv.ReadConfig("../common/config.yaml", &cfg)
	// if err1 != nil {
	// 	log.Fatal("config.yaml read failed")
	// }

	//update placeholder
	val := configParam{vcfg.VideoDlPath, time.Now().UnixMilli()}
	processConfig(&vcfg.BackupId, val)
	processConfig(&vcfg.NextIterId, val)
	processConfig(&vcfg.VideoId, val)
	processConfig(&vcfg.LogPath, val)
	return &vcfg
}

func processConfig(tpl *string, val configParam) {

	var result bytes.Buffer
	tmpl, err := template.New("test").Parse(*tpl)
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
		// panic(err)
	}
	err = tmpl.Execute(&result, val)
	if err != nil {
		log.Fatal("Error executing template: ", err)
		// panic(err)
	}
	*tpl = result.String()
}
