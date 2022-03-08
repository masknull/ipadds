/*
	Mask
*/
package main

import (
	"bufio"
	"fmt"
	"ipadds"
	"os"
	"strings"
	"time"
)

func main() {
	begin := time.Now()
	_, err := os.Stat("db")
	if os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "当前目录不存在db文件"))
		return
	}
	region, _ := ip2region.New("db")
	defer region.Close()
	if len(os.Args) == 1 {
		fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "参数错误 你应该想要-h"))
		return
	}
	data := os.Args[1]
	gostr := os.Args[0]

	commands := strings.Fields(string(data))

	commands = append(commands, "")
	switch commands[0] {
	case "--help":
		print("	" + gostr + " 1.1.1.1")
		fmt.Println(fmt.Sprintf("\u001B[0;32m%s\u001B[0m", " 单个IP"))
		print("	-f [ipfile]")
		fmt.Println(fmt.Sprintf("\u001B[0;32m%s\u001B[0m", " 读本地文件一行一个IP"))
		print("	-o [write file]")
		fmt.Println(fmt.Sprintf("\u001B[0;32m%s\u001B[0m", " 写入本地文件地址，默认输出控制台"))
		print("	原作者项目地址：https://gitee.com/lionsoul/ip2region\n")
		print("	修改：Maks\n")
		return
	case "-h":
		print("	" + gostr + " 1.1.1.1")
		fmt.Println(fmt.Sprintf("\u001B[0;32m%s\u001B[0m", " 单个IP"))
		print("	-f [ipfile]")
		fmt.Println(fmt.Sprintf("\u001B[0;32m%s\u001B[0m", " 读本地文件一行一个IP"))
		print("	-o [write file]")
		fmt.Println(fmt.Sprintf("\u001B[0;32m%s\u001B[0m", " 写入本地文件地址，默认输出控制台"))
		print("	原作者项目地址：https://gitee.com/lionsoul/ip2region\n")
		print("	修改：Maks\n")
		return
	case "-f":
		if len(os.Args) < 3 {
			fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "参数错误 你应该想要-h"))
			return
		}
		file := os.Args[2]
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "本地文件不存在："+file))
			return
		}
		ipdata := ""
		errred := ""
		i := 0
		f, err := os.Open(file)
		s := bufio.NewScanner(f)

		if len(os.Args) == 3 {
			errred = "\x1b[0;32m"
			print(errred)
			for s.Scan() {
				ip := ip2region.IpInfo{}
				ipi := s.Text()
				ip, err = region.MemorySearch(ipi)
				i++
				if err != nil {
					errred = "\x1b[0;31m%s\x1b[0m"
					ipdata = err.Error()
					break
				} else {
					fmt.Println(ipi + " " + ip.String())
				}
			}
			err = s.Err()
			if err != nil {
				fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "读取本地IP文件运行错误："+err.Error()))
			}
			print("\x1b[0m")
			print("数据：")
			print(i)
			print("行 耗时：" + time.Since(begin).String())
			return
		} else if len(os.Args) == 5 && os.Args[3] == "-o" {
			open, _ := os.Create(os.Args[4])
			for s.Scan() {
				ip := ip2region.IpInfo{}
				ipi := s.Text()
				ip, err = region.MemorySearch(ipi)
				i++
				if err != nil {
					errred = "\x1b[0;31m%s\x1b[0m"
					ipdata = err.Error()
					break
				} else {
					errred = "\x1b[0;32m%s"
					ipdata = ipi + " " + ip.String() + "\n"
					open.WriteString(ipdata)
				}
			}
			open.Close()
			err = s.Err()
			if err != nil {
				fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "读取本地IP文件运行错误："+err.Error()))
			}
			print("\x1b[0m")
			print("数据：")
			print(i)
			print("行 耗时：" + time.Since(begin).String())
			return
		}
	default:
		err = nil
	}
	ip := ip2region.IpInfo{}
	ip, err = region.MemorySearch(commands[0])
	if err != nil {
		fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("\x1b[0;32m%s", ip.String()))
	}
}
