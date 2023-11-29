package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	path := "D:/GOPROJECT/src/Path/path.go"
	/*
		目录
	*/
	Dirpath := filepath.Dir(path)
	fmt.Println(Dirpath)
	/*
		文件名
	*/
	Basepath := filepath.Base(path)
	fmt.Println(Basepath)
	/*
		扩展名
	*/
	Extpath := filepath.Ext(path)
	fmt.Println(Extpath)
	/*
		绝对路径
	*/
	boolean := filepath.IsAbs(path)
	fmt.Println(boolean)
	Abspath, _ := filepath.Abs(path)
	fmt.Println(Abspath)

	/*相对路径*/
	RelPath, _ := filepath.Rel("D:/GOPROJECT/src", path)
	fmt.Println(RelPath)

	/* 路径分割 拼接*/
	dir, file := filepath.Split(path)
	fmt.Println(dir, "", file)
	JoinDir := filepath.Join("C:/Users/user/Desktop/", "GoLanguage.docx")
	fmt.Println(JoinDir)

	/* filepath.walk 目录遍历*/
	var files []string
	root := "D:/GOPROJECT/src"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	/* 打开文件*/
	f, err := os.Open("C:/Users/user/Desktop/学习/GoLanguage.docx")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	/* defer 关闭*/
	defer f.Close()
	/* 读取文件结构*/
	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	/* 获取文件大小*/
	filesize := fileinfo.Size()
	/* 创建文件大小内存*/
	buffer := make([]byte, filesize)
	/* 读取文件*/
	bytesread, err := f.Read(buffer)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("bytes read: ", bytesread)
	fmt.Println("bytestream to string: ", string(buffer))

}
