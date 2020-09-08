package L

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	Info *log.Logger
	Warn *log.Logger
	Erro *log.Logger
)

type Log_para struct {
	Log_path   string        // 日志路径
	File_size  int64         //文件大于file_size 会被切分
	Check_time time.Duration //检查时间
	File_num   int           //保留的文件数量
	New_name   string        //最新的日志文件名称
}

func Start(log_p *Log_para) {
	var info *(os.File)
	var now_name string
	var change bool = false
	go func() {
		for {
			//查找当前日志文件名称
			files, err := ioutil.ReadDir(log_p.Log_path)

			if err != nil {
				log.Println("日志路径不存在")
				log.Println(err.Error())
				return
			} else {
				if len(files) > 0 {
					now_name = files[len(files)-1].Name()
				}
			}

			if now_name == "" {
				log_p.New_name = log_p.Log_path + "/" + fmt.Sprintf("%d", time.Now().Unix()) + ".log"
			} else {
				log_p.New_name = log_p.Log_path + "/" + now_name
			}
			// 检查文件是否存在
			fi, err := os.Stat(log_p.New_name)
			if err != nil {
				info, err = os.OpenFile(log_p.New_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
				if err != nil {
					log.Println("创建日志文件失败：", err)
				}
				change = true
			} else {
				if fi.Size() > log_p.File_size {
					info.Close()
					log_p.New_name = log_p.Log_path + "/" + fmt.Sprintf("%d", time.Now().Unix()) + ".log"
					info, err = os.OpenFile(log_p.New_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
					if err != nil {
						log.Println("创建日志文件失败：", err)
					}
					log.Println("创建日志:", log_p.New_name)
					change = true
				}
				if info == nil {
					info, err = os.OpenFile(log_p.New_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
					if err != nil {
						log.Println("打开日志文件失败：", err)
					}
					change = true
				}
			}
			if change == true {
				Info = log.New(info, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
				Warn = log.New(info, "Warn:", log.Ldate|log.Ltime|log.Lshortfile)
				Erro = log.New(info, "Erro:", log.Ldate|log.Ltime|log.Lshortfile)

				//清理多余的日志文件
				files, err = ioutil.ReadDir(log_p.Log_path)
				if err != nil {
					log.Println(err.Error())
				} else {
					if len(files) > log_p.File_num {
						end := len(files) - log_p.File_num
						for i := 0; i < end; i++ {
							os.Remove(log_p.Log_path + "/" + files[i].Name())
						}
					}
				}

				change = false
			}
			time.Sleep(log_p.Check_time * time.Second)
		}

	}()
	for {
		if Erro != nil {
			log.Println("当前日志:", log_p.New_name)
			break
		}
	}

}
