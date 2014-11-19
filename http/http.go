package http

import (
	"encoding/json"
	"github.com/chengz0/xdashboard/global"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	martini_m *martini.ClassicMartini
	timeout   = time.Duration(2) * time.Second
)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func StartMartini() {
	// http client
	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	// martini
	martini_m = martini.Classic()
	martini_m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{{
			"nl2br":      nl2br,
			"htmlquote":  htmlQuote,
			"str2html":   str2html,
			"dateformat": dateFormat,
		}},
	}))

	martini_m.Get("/", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		ret, err := json.Marshal(global.HostIni)
		if err != nil {
			log.Printf("Error hostini: %s", err.Error())
			return RenderErrDto(err.Error())
		}
		log.Println(string(ret))
		return RenderDataDto(string(ret))
	})

	// //hosts
	// martini_m.Get("/state", func(w http.ResponseWriter) string {
	// 	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	// 	ret, err := json.Marshal(global.HostIni)
	// 	if err != nil {
	// 		log.Printf("Error hostini: %s", err.Error())
	// 		return err.Error()
	// 	}
	// 	log.Println(string(ret))
	// 	return string(ret)
	// })

	// system state
	martini_m.Get("/system", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		url := "http://172.16.110.134:3000/system"
		ret := XRequest(&client, url)
		log.Println(ret)
		return ret
	})

	http.ListenAndServe(":4000", martini_m)
}

type Dto struct {
	Succ bool
	Msg  string
	Data interface{}
}

func ErrDto(message string) Dto {
	return Dto{Succ: false, Msg: message}
}

func DataDto(d interface{}) Dto {
	return Dto{Succ: true, Msg: "", Data: d}
}

func RenderErrDto(message string) string {
	dto := ErrDto(message)
	bs, err := json.Marshal(dto)
	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

func RenderDataDto(d interface{}) string {
	dto := DataDto(d)
	bs, err := json.Marshal(dto)

	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}
