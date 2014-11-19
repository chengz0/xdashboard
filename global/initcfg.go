package global

import (
	"errors"
	"github.com/astaxie/beego/config"
	"github.com/ulricqin/goutils/filetool"
	"log"
)

var (
	Config       config.ConfigContainer
	HostIni      map[string]string
	SystemIni    map[string]string
	ContainerIni map[string]string
	ProcessIni   map[string]string
)

func InitGlobal(path string) error {
	if !filetool.IsExist(path) {
		log.Printf("Configuration file[%s] is nonexistent", path)
		return errors.New("No config!")
	}
	Config, err := config.NewConfig("ini", path)
	if err != nil {
		log.Printf("Error bad ini: %s", err.Error())
		return err
	}

	// read config from .ini
	HostIni, err = Config.GetSection("hosts")
	if err != nil {
		log.Printf("Error no hosts section: %s", err.Error())
	}
	SystemIni, err = Config.GetSection("system")
	if err != nil {
		log.Printf("Error no system section: %s", err.Error())
	}
	ContainerIni, err = Config.GetSection("container")
	if err != nil {
		log.Printf("Error no container section: %s", err.Error())
	}
	ProcessIni, err = Config.GetSection("process")
	if err != nil {
		log.Printf("Error no process section: %s", err.Error())
	}
	log.Println(HostIni)
	log.Println(ContainerIni)
	return nil
}
