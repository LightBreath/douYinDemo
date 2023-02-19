package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host      string
	Port      int32
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml: "parse_toml"`
	Loc       string
}

type Ftp struct {
	host string
	port int32
}

type Path struct {
}

var Info Config

type Config struct {
	mysql Mysql `toml:"mysql"`
	ftp   Ftp   `toml:"ftp"`
	path  Path  `toml:"path"`
}

func ToString() string {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.mysql.Username, Info.mysql.Password, Info.mysql.Host, Info.mysql.Port, Info.mysql.Database,
		Info.mysql.Charset, Info.mysql.ParseTime, Info.mysql.Loc)
	log.Println(str)
	return str
}

func init() {
	_, err := toml.DecodeFile("/home/malchinee/douYinDemo/config/conf.toml", &Info)
	if err != nil {
		panic(err)
	}
}
