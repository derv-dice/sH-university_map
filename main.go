package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"api_template/api"
	. "api_template/config"
)

func init() {
	logFileDisabled := flag.Bool("nolog", true, "Log to file")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var err error
	if !*logFileDisabled {
		if err = LogToFile(); err != nil {
			log.Fatal(err)
		}
	}

	if err = Cfg.Load(); err != nil {
		log.Fatalf("Не удалось загрузить конфиг. Ошибка: %v", err)
	}
}

func main() {
	log.Println(fmt.Sprintf("Сервер запущен по адресу %s:%d", Cfg.Server.Addr, Cfg.Server.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", Cfg.Server.Addr, Cfg.Server.Port), api.MainHandler()))
}
