//author: richard
package single

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var Config	map[string]interface{}

func Init() (err error){
	defer func() {
		if re := recover(); re != nil {
			err = errors.New(re.(string))
		}
	}()
	Config = make(map[string]interface{})

	initSingle(1)
	// 初始化配置
	err = initConfig()
	checker(err)
	// 初始化系统
	err = initSys()
	checker(err)
	return
}


func initSingle(cap uint) {
	go func() {
		c := make(chan os.Signal, cap)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
		for {
			s := <- c
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				os.Exit(0)
			case syscall.SIGUSR1: //kill -10 pid
				//SetLogLevel("DEBUG")  //切换日志级别DEBUG
				fmt.Println("usr1", s)
			case syscall.SIGUSR2: //kill -12 pid
				//SetLogLevel("ERROR")  //切换日志级别ERROR
				fmt.Println("usr2", s)
			default:
				fmt.Println("debug", s)
			}
		}
	}()
}

func initSys() (err error){
	fmt.Println(Config)
	var pid = os.Getpid()
	var filename = Config["pid"]
	file, err := os.OpenFile(filename.(string), os.O_WRONLY | os.O_CREATE, 0666)
	checker(err)
	defer func() {
		checker(file.Close())
	}()

	_, err = file.WriteString(strconv.Itoa(pid))
	checker(err)
	return
}

func initConfig() (err error) {
	var args = os.Args
	var length = len(args)
	var configPath string
	for i := 1; i < length; i += 2 {
		switch args[i] {
		case "-c":
			if j := i + 1; j < length {
				configPath = args[j]
			} else {
				configPath = "/etc/letaotao/letaotao.json"
			}
		default:
			continue
		}
	}
	config, err := os.Open(configPath)
	checker(err)
	defer func() {
		checker(config.Close())
	}()
	buf := make([]byte, 1024)
	n, err := config.Read(buf)
	checker(err)
	err = json.Unmarshal(buf[:n], &Config)
	checker(err)
	return
}

func checker(err error) {
	if err != nil {
		panic(err)
	}
}