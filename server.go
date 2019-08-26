package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers   []Server
	ServersID string
}

type ServerConfig struct {
	ServerIP   string
	ServerName string
	ServerPort uint16
	PostPath   string
	GetPath    string

	EndPoint            string
	AccessKey           string
	AccessKeySecret     string
	BucketName          string
	EndPointInternal    string
	UseInternalEndPoint bool
}

type Res struct {
	Id  string
	Url string
}

var cfg ServerConfig

func loadconfig() bool {
	data, err := ioutil.ReadFile("/etc/alioss-uploader/config.json")
	if err != nil {
		return false
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return false
	}

	fmt.Println(cfg)
	return true
}

func sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

type User struct {
	Id      string
	Balance uint64
}

func main() {

	if loadconfig() {
		http.HandleFunc(cfg.PostPath, handler)
		var serverstr string = cfg.ServerIP + ":" + strconv.Itoa(int(cfg.ServerPort))
		fmt.Println(serverstr)
		var success bool
		if cfg.UseInternalEndPoint {
			success = AliCreateBlunker(cfg.EndPointInternal, cfg.AccessKey, cfg.AccessKeySecret, cfg.BucketName)
		} else {
			success = AliCreateBlunker(cfg.EndPoint, cfg.AccessKey, cfg.AccessKeySecret, cfg.BucketName)
		}

		if success {
			http.ListenAndServe(serverstr, nil)
		} else {
			fmt.Println("Create bucket fail!")

		}

	} else {
		fmt.Println("Load config error!")
	}
}
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println("method:", r.Method)
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}

	} else if r.Method == "POST" {
		fmt.Println("method:", r.Method)
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println(handler.Header)

		_id := sha1s(handler.Filename + time.Now().String())
		fmt.Println(_id)
		_filepath := "/tmp/" + _id
		f, err := os.OpenFile(_filepath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		if cfg.UseInternalEndPoint {
			AliPut(cfg.EndPointInternal, cfg.AccessKey, cfg.AccessKeySecret, cfg.BucketName, _id+path.Ext(handler.Filename), _filepath)
		} else {
			AliPut(cfg.EndPoint, cfg.AccessKey, cfg.AccessKeySecret, cfg.BucketName, _id+path.Ext(handler.Filename), _filepath)
		}

		u := "http://" + cfg.BucketName + "." + cfg.EndPoint + "/" + _id + path.Ext(handler.Filename)
		r := Res{Id: _id, Url: u}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
		//clear
		os.Remove(_filepath)
	}
}
