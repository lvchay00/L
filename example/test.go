package main

import (
	"time"

	"github.com/lvchay00/L"
)

func main() {
	//启动log文件 监控分割线程
	L.Start(&(L.Log_para{
		Log_path:   "./log",                //日志路径
		File_size:  10 * 1024 * 1024,       //log 文件超过File_size 会分割为新的文件
		Check_time: time.Duration(10 * 60), //10分钟检查一次
		File_num:   5,                      //日志文件保留数量
	}))

	//测试函数
	go func() {
		for {

			L.Info.Println("info")
			L.Erro.Println("erro")
			L.Warn.Println("warn")
			time.Sleep(1 * time.Second)

		}
	}()

	for {
		time.Sleep(100 * time.Hour)
	}

}
