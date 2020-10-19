package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var path string
var out string
var protoPath string
var protoc string
var tsplugin string

func main() {
	var cmd string
	var script string
	flag.StringVar(&cmd, "cmd", "install", "gen : install protoc-gen-go; gen : gen proto file;")
	flag.StringVar(&script, "sc", "golang", "golang/javascript/typescript")
	flag.StringVar(&path, "p", "", "protoc path")
	flag.StringVar(&out, "o", "", "proto out path")
	flag.Parse()

	ostat, err := os.Stat(out)
	if (err != nil && os.IsNotExist(err)) || !ostat.IsDir() {
		os.MkdirAll(out, os.ModePerm)
	}

	protoPath = path + "prtobuf/proto/"

	switch runtime.GOOS {
	case "windows":
		protoc = path + "/protoc/protoc_3_13_0_win64.exe"
		tsplugin = path + "/node_modules/.bin/protoc-gen-ts_proto.cmd"
	case "darwin":
		protoc = path + "/protoc/protoc_3_13_0_osx_x86_64"
		tsplugin = path + "/node_modules/.bin/protoc-gen-ts_proto"
	}

	switch cmd {
	case "install":
		installGo()
	case "gen":
		gen(script)
	}
}

func installGo() {
	cmd := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("安装失败")
		fmt.Println(err)
	} else {
		fmt.Println("安装成功")
	}
}

func gen(script string) {
	switch script {
	case "golang":
		gen4golang()
	case "javascript":
		gen4javascript()
	case "typescript":
		gen4typescript()
	}
}

func gen4golang() {
	walkProto(protoPath, func(file string) {
		cmd := exec.Command(protoc, "--proto_path="+protoPath, "--go_out="+out, file)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("生成失败")
			fmt.Println(err.Error())
		}
	})
}

func gen4javascript() {
	walkProto(protoPath, func(file string) {
		cmd := exec.Command(protoc, "--proto_path="+protoPath, "--js_out=import_style=commonjs,binary:"+out, file)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("生成失败")
			fmt.Println(err.Error())
		}
	})
}

func gen4typescript() {
	walkProto(protoPath, func(file string) {
		cmd := exec.Command(protoc, "--proto_path="+protoPath, "--plugin=protoc-gen-ts_proto="+tsplugin, "--ts_proto_out="+out, file)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("生成失败")
			fmt.Println(err.Error())
		}
	})
}

func walkProto(path string, cb func(file string)) {
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			if file.IsDir() {
				walkProto(path+file.Name()+"/", cb)
				continue
			}
			if strings.HasSuffix(file.Name(), ".proto") {
				cb(path + file.Name())
			}
		}
	}
}
