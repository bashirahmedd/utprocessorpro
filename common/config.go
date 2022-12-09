package common

import (
	"bytes"
	"text/template"
	"time"

	"github.com/ilyakaznacheev/cleanenv"

	// "github.com/spf13/viper"
	"log"
)

type ConfigDownloader struct {
	VideoId     string `yaml:"in_video_list"`
	NextIterId  string `yaml:"next_iteration_file"`
	BackupId    string `yaml:"log_id"`
	LogPath    string `yaml:"log_path"`
	VideoDlPath string `yaml:"video_dl_path"`
}

type configVideo struct {
	VideoDlPath string
	Counter     int64
}

func GetShareConfig() *ConfigDownloader {

	var cfg ConfigDownloader
	err1 := cleanenv.ReadConfig("../common/config.yaml", &cfg)
	if err1 != nil {
		log.Fatal("config.yaml read failed")
	}

	//update placeholder
	val := configVideo{cfg.VideoDlPath, time.Now().UnixMilli()}
	processConfig(&cfg.BackupId, val)
	processConfig(&cfg.NextIterId, val)
	processConfig(&cfg.VideoId, val)
	processConfig(&cfg.LogPath, val)
	return &cfg
}

func processConfig(tpl *string, val configVideo) {
	
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
