// author: liushihao
// email: liushihao1993@hotmail.com
// 如果你想要实现什么样的文本处理功能，欢迎发邮件给我

/*
# vigo
一款准强大的文本处理工具（旨在替换linux一些命令工具）
提示：命令行要传入\n参数 请加\进行转义，即输入\\n， 同理\\s \\d
Usage of vigo: vigo [option] <file>
options:
  -mn
        合并连续多个换行为一个
  -n string
        起始行,默认全文搜索, e.g. 10:18
  -p string
        匹配模式
  -r    替换
  -repl string
        匹配模式
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

var (
	isReplaced bool // isReplaced replace替换  默认是搜索
	// i_         bool   // i_ interactive 交互式样，r为true时生效
	p       string // p pattern
	line    string
	repl_   string
	isMerge bool
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage of vigo: vigo [option] <file>\noptions:")
		flag.PrintDefaults()
	}
	flag.StringVar(&line, "n", "", "起始行,默认全文搜索, e.g. 10:18")
	flag.StringVar(&p, "p", "", "匹配的字符串(正则表达式)")
	flag.StringVar(&repl_, "repl", "", "替换的字符串")
	flag.BoolVar(&isReplaced, "r", false, "是否进行替换重写文件")
	flag.BoolVar(&isMerge, "mn", false, "合并连续多个换行为一个")
	flag.Parse()
	repl_ = strings.ReplaceAll(repl_, `\n`, "\n")
	if len(flag.Args()) != 1 {
		exitWithFlagUsage("")
	}
	if isReplaced {
		fmt.Printf("search: %q, replace: %q\n", p, repl_)
	}
}
func main() {
	filename := flag.Args()[0]
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	if isMerge {
		write(filename, mergeMultiN(b))
		fmt.Println("合并多个换行成功！")
		os.Exit(0)
	}
	startLine, endLine, err := getLineRange()
	if err != nil {
		exitWithFlagUsage(err.Error())
	}
	result := find(b, p, startLine, endLine)
	if result != "" {
		fmt.Println(result)
	}
	if !isReplaced {
		return
	}
	if repl_ == "" {
		confirmYesOrNo("您当前输入的替换字符串为空,请输入y确认继续，否则退出> ")
	}
	replStr := replace(b, p, repl_)
	write(filename, replStr)
}

// exitWithFlagUsage 退出并输出打印使用帮助信息.
func exitWithFlagUsage(message string) {
	if message != "" {
		fmt.Println(message)
	}
	flag.Usage()
	os.Exit(1)
}

// find 查找文本内容.
func find(content []byte, pattern string, startLine, endLine int64) string {
	text := bytesToStr(content)
	rows := strings.Split(text, "\n")
	result := make([]string, 0, len(rows))
	var all bool
	if startLine == 0 && endLine == 0 {
		all = true
	}
	var appendMap = map[int]bool{}
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("正则表达式有误, pattern: %q error: %s\n", pattern, err.Error())
		os.Exit(1)
	}
	if endLine < 0 {
		endLine = int64(len(rows))
	}
	for n, row := range rows {
		if !all {
			if int64(n) < startLine || int64(n) > endLine {
				continue
			}
		}
		if reg.MatchString(row) {
			if appendMap[n] {
				continue
			}
			if !appendMap[n-1] && n > 0 {
				result = append(result, fmt.Sprintf("%s %s", blueStr(strconv.Itoa(n-1)+":"), reg.ReplaceAllStringFunc(rows[n-1], redStr)))
				appendMap[n-1] = true
			}
			result = append(result, fmt.Sprintf("%s %s", blueStr(strconv.Itoa(n)+":"), reg.ReplaceAllStringFunc(row, redStr)))
			appendMap[n] = true
			if !appendMap[n+1] && n < len(rows)-1 {
				result = append(result, fmt.Sprintf("%s %s", blueStr(strconv.Itoa(n+1)+":"), reg.ReplaceAllStringFunc(rows[n+1], redStr)))
				appendMap[n+1] = true
			}
		}

	}

	return strings.Join(result, "\n")
}

// replace 获取替换为的文本内容
func replace(content []byte, pattern string, repl string) string {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("正则表达式有误, pattern: %q error: %s\n", pattern, err.Error())
		os.Exit(1)
	}
	text := bytesToStr(content)
	if strings.Contains(pattern, ".*") && !strings.Contains(pattern, ".*?") {
		confirmYesOrNo("警告：当前为非贪婪模式匹配: " + pattern + "，请输入y确认继续，否则退出>")
	}
	return reg.ReplaceAllString(text, repl)
}

// mergeMultiN 合并多行空行
func mergeMultiN(content []byte) string {
	return replace(content, "\n+", "\n")
}

// bytesToStr 零拷贝字节切片转换字符串
func bytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type color int

const (
	blue = iota
	red
)

// redStr  显示红色.
func redStr(s string) string {
	return colorS(s, red)
}

// blueStr 显示蓝色.
func blueStr(s string) string {
	return colorS(s, blue)
}

// colorS 彩色显示. 不支持windows平台.
func colorS(s string, c color) string {
	if runtime.GOOS == "windows" {
		return s
	}
	switch c {
	case blue:
		return fmt.Sprintf("\033[0;34m%s\033[0m", s)
	case red:
		return fmt.Sprintf("\033[0;31m%s\033[0m", s)
	default:
		return s
	}
}

// getLineRange 获取命令行参数输入的起始行号.
func getLineRange() (startLine, endLine int64, err error) {
	if line == "" {
		return
	}
	lineS := strings.Split(line, ":")
	err = errors.New("n 参数输入错误")
	if len(lineS) != 2 {
		return
	}
	startLine, err = strconv.ParseInt(lineS[0], 10, 64)
	if lineS[0] != "" && err != nil {
		return
	}
	endLine, err = strconv.ParseInt(lineS[1], 10, 64)
	if lineS[1] != "" && err != nil {
		return
	}
	if endLine == 0 {
		endLine = -1
	}
	if startLine > endLine && endLine != -1 {
		err = errors.New("n 参数输入错误, 结束行号不能大于起始行号")
		return
	}
	err = nil
	return
}

func promptEmptyRepl(repl_ string) {
	if repl_ != "" {
		return
	}

}

func write(filename, text string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("文件覆写失败 error: " + err.Error())
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		fmt.Println("文件覆写失败 error: " + err.Error())
		os.Exit(1)
	}
}

// confirmYesOrNo 交互式确认.
func confirmYesOrNo(prompt string) {
	fmt.Print(prompt)
	var confirm string
	_, err := fmt.Scan(&confirm)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(0)
	}
	switch confirm {
	case "y", "Y", "yes", "YES":
	default:
		os.Exit(0)
	}
}
