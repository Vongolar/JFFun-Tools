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

var protoc string
var protoPath string
var out string

func main() {
	var cmd string
	var script string
	flag.StringVar(&cmd, "cmd", "install", "gen : install protoc-gen-go; gen : gen proto file;")
	flag.StringVar(&script, "sc", "golang", "golang/javascript")
	flag.StringVar(&protoc, "p", "", "protoc path")
	flag.StringVar(&protoPath, "pp", "", "proto files path")
	flag.StringVar(&out, "o", "", "proto out path")
	flag.Parse()

	switch runtime.GOOS {
	case "windows":
		protoc += "protoc_3_13_0_win64.exe"
	case "darwin":
		protoc += "protoc_3_13_0_osx_x86_64"
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
		cmd := exec.Command(protoc, "--proto_path="+protoPath, "--js_out="+out, file)
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
