package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host      string
	Port      int32
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_toml"`
	Loc       string
}

type Ftp struct {
	Ip   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

var Info Config

type Config struct {
	Mysql `toml:"mysql"`
	Ftp   `toml:"ftp"`
	Path  `toml:"path"`
}

func ToString() string {
	// "root:030628@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.Mysql.Username, Info.Mysql.Password, Info.Mysql.Host, Info.Mysql.Port, Info.Mysql.Database,
		Info.Mysql.Charset, Info.Mysql.ParseTime, Info.Mysql.Loc)
	log.Println(str)
	return str
}

func init() {
	_, err := toml.DecodeFile("./config/conf.toml", &Info)
	fmt.Println(Info)
	if err != nil {
		panic(err)
	}
	strings.Trim(Info.Mysql.Host, " ")
	strings.Trim(Info.Ftp.Ip, " ")
}
