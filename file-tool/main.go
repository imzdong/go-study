package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/qiniu/iconv"
)

func main() {
	var inEncoding, outEncoding string
	var files []string
	// 创建主窗口
	MainWindow{
		Title:  "文件编码转换工具",
		Size:   Size{Width: 400, Height: 300},
		Layout: VBox{},
		Children: []Widget{
			Label{
				Text: "输入编码：",
			},
			ComboBox{
				Value:    Bind(inEncoding),
				Model:    []string{"UTF-8", "GBK", "..."}, // 添加更多编码格式
				OnCurrentIndexChanged: func() {
					// 处理下拉选框值变化的逻辑
					fmt.Println(inEncoding)
				},
			},
			Label{
				Text: "输出编码：",
			},
			ComboBox{
				Value:    Bind(outEncoding),
				Model:    []string{"UTF-8", "GBK", "..."}, // 添加更多编码格式
				OnCurrentIndexChanged: func() {
					// 处理下拉选框值变化的逻辑
					fmt.Println(outEncoding)
				},
			},
			PushButton{
				Text: "选择文件",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Title = "选择文件"
					dlg.Filter = "所有文件 (*.*)|*.*"
					//dlg.Multiselect = true
					var tmp walk.Form
					if ok, _ := dlg.ShowOpenMultiple(tmp); ok {
						files = dlg.FilePaths
					}
				},
			},
			PushButton{
				Text: "开始转换",
				OnClicked: func() {
					for _, file := range files {
						convertFileEncoding(file, inEncoding, outEncoding)
					}
					var tmp walk.Form
					walk.MsgBox(tmp, "转换完成", "文件的编码转换已完成！", walk.MsgBoxIconInformation)
				},
			},
		},
	}.Run()
}
func convertFileEncoding(file, inEncoding, outEncoding string) {
	// 打开源文件
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("无法打开文件：%s，错误：%s\n", file, err.Error())
		return
	}
	defer inFile.Close() // 创建输出文件
	outFile, err := os.Create(fmt.Sprintf("%s_converted", file))
	if err != nil {
		fmt.Printf("无法创建输出文件：%s，错误：%s\n", file, err.Error())
		return
	}
	defer outFile.Close()

	// 创建编码转换器
	cd, err := iconv.Open(outEncoding, inEncoding)
	if err != nil {
		fmt.Printf("无法创建编码转换器：%s\n", err.Error())
		return
	}
	defer cd.Close()

	// 转换文件内容并写入输出文件
	buf := make([]byte, 1024)
	for {
		n, err := inFile.Read(buf)
		if err != nil && err.Error() != "EOF" {
			fmt.Printf("读取文件时发生错误：%s\n", err.Error())
			return
		}
		if n == 0 {
			break
		}

		inBytes := buf[:n]
		outBytes := make([]byte, len(inBytes)*2) // 可能需要调整转换后的字节长度

		_, _, err = cd.Conv(inBytes, outBytes)
		if err != nil {
			fmt.Printf("转换文件时发生错误：%s\n", err.Error())
			return
		}

		_, err = outFile.Write(outBytes)
		if err != nil {
			fmt.Printf("写入输出文件时发生错误：%s\n", err.Error())
			return
		}
	}

	fmt.Printf("文件 %s 的编码已成功转换并保存为 %s_converted\n", file, file)
}
