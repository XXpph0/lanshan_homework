package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 创建文件
	newFile, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("创建文件出错:", err)
	} else {
		defer newFile.Close()
		fmt.Println("成功创建文件:", newFile.Name())
	}

	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer file.Close()
		fmt.Println("成功打开文件:", file.Name())
	}

	// 打开文件（使用 OpenFile）
	anotherFile, err := os.OpenFile("anotherfile.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer anotherFile.Close()
		fmt.Println("成功打开文件:", anotherFile.Name())
	}

	// 写入文件
	dataToWrite := []byte("Hello, Golang!")
	bytesWritten, err := anotherFile.Write(dataToWrite)
	if err != nil {
		fmt.Println("写入文件出错:", err)
		return
	}
	fmt.Printf("成功写入 %d 字节到文件\n", bytesWritten)

	// 读取文件
	readBuffer := make([]byte, 100)
	bytesRead, err := file.Read(readBuffer)
	if err != nil {
		if err != io.EOF {
			fmt.Println("读取文件出错:", err)
			return
		}
	}
	fmt.Printf("从文件中读取 %d 字节的数据: %s\n", bytesRead, readBuffer[:bytesRead])

	// 移除文件
	_, _ = os.Create("toRemove.txt")
	err = os.Remove("toRemove.txt")
	if err != nil {
		fmt.Println("删除文件出错:", err)
	} else {
		fmt.Println("成功删除文件")
	}

	// 重命名文件
	err = os.Rename("anotherfile.txt", "renamedfile.txt")
	if err != nil {
		fmt.Println("重命名文件出错:", err)
	} else {
		fmt.Println("成功重命名文件")
	}

	// 创建目录
	err = os.Mkdir("example_dir", 0755) // 0755是目录权限，表示所有者具有读写执行权限，其他人具有读和执行权限
	if err != nil {
		fmt.Println("创建目录出错:", err)
		return
	}
	// 改变工作目录
	err = os.Chdir("example_dir")
	if err != nil {
		fmt.Println("改变工作目录出错:", err)
	} else {
		newCurrentDir, _ := os.Getwd()
		fmt.Println("新的工作目录:", newCurrentDir)
	}

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错:", err)
	} else {
		fmt.Println("当前工作目录:", currentDir)
	}
	_, _ = os.Create("dir_file")
	_ = os.Chdir("..")

	// 递归删除目录及其内容
	err = os.RemoveAll("example_dir")
	if err != nil {
		fmt.Println("删除目录出错:", err)
	} else {
		fmt.Println("成功删除目录及其内容")
	}

	// 获取环境变量
	goPath := os.Getenv("GOPATH")
	fmt.Println("GOPATH环境变量:", goPath)

	// 设置环境变量
	err = os.Setenv("MY_VARIABLE", "my_value")
	if err != nil {
		fmt.Println("设置环境变量出错:", err)
	} else {
		fmt.Print("成功设置环境变量：")
		fmt.Println(os.Getenv("MY_VARIABLE"))
	}

	// 获取命令行参数
	fmt.Println("命令行参数:", os.Args)
}
