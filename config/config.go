package config

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

var Cfg = &Config{}

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

func (c *Config) Load() (err error) {
	log.Println("Загрузка конфига")
	var data []byte
	if data, err = ioutil.ReadFile("config.yml"); err != nil {
		return
	}
	return yaml.Unmarshal(data, &c)
}

// LogToFile - Включить логирование в файл ./logs/filename.log
func LogToFile() (err error) {
	// Создаем директорию логов, если ее нет
	if _, err = os.Stat("logs"); os.IsNotExist(err) {
		if err = os.MkdirAll("logs", 0777); err != nil {
			return
		}
	}

	// В качестве имени файла ставим текущее время
	name := time.Now().Format("02-01-2006_15-04-05") + ".log"
	var file io.Writer
	if file, err = os.Create(path.Join("logs", name)); err != nil {
		return
	}

	// Выключаем вывод стандартного логера в созданный файл и терминал
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	return
}