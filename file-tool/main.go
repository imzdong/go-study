package main

import (
	"os"
	"log"
	"path/filepath"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/qiniu/iconv"
)

func main() {

	logger, err := CreateLogger("logfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		logger.Println("Flushing log buffer")
		_ = logger.Writer().(*os.File).Sync()
	}()

	var inEnCodingComboBox, outEnCodingComboBox *walk.ComboBox
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
				//Value:    Bind(&inEncoding),
				AssignTo: &inEnCodingComboBox,
				Model:    []string{"UTF-8", "GB2312", "GBK"}, // 添加更多编码格式
				OnCurrentIndexChanged: func() {
					// 处理下拉选框值变化的逻辑
					inEncoding = inEnCodingComboBox.Text()
					logger.Println(inEncoding)
				},
			},
			Label{
				Text: "输出编码：",
			},
			ComboBox{
				AssignTo: &outEnCodingComboBox,
				//Value:    Bind(outEncoding),
				Model: []string{"UTF-8", "GB2312", "GBK"}, // 添加更多编码格式
				OnCurrentIndexChanged: func() {
					// 处理下拉选框值变化的逻辑
					outEncoding = outEnCodingComboBox.Text()
					logger.Println(outEncoding)
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
						convertFileEncoding(file, inEncoding, outEncoding, logger)
					}
					var tmp walk.Form
					walk.MsgBox(tmp, "转换完成", "文件的编码转换已完成！", walk.MsgBoxIconInformation)
				},
			},
		},
	}.Run()
}
func convertFileEncoding(file, inEncoding, outEncoding string, logger *log.Logger) {
	logger.Println("file:%s, inEncoding：%s，outEncoding：%s\n", file, inEncoding, outEncoding)
	// 打开源文件
	inFile, err := os.Open(file)
	if err != nil {
		logger.Println("无法打开文件：%s，错误：%s\n", file, err.Error())
		return
	}
	defer inFile.Close() // 创建输出文件
	newFilePath,err := generateConvertedFilePath(file, logger)
	if err != nil {
		logger.Println("无法创建新文件目录：%s，错误：%s\n", file, err.Error())
		return
	}
	outFile, err := os.Create(newFilePath)
	if err != nil {
		logger.Println("无法创建输出文件：%s，错误：%s\n", file, err.Error())
		return
	}

	defer outFile.Close()

	// 创建编码转换器
	cd, err := iconv.Open(outEncoding, inEncoding)
	if err != nil {
		logger.Println("无法创建编码转换器：%s\n", err.Error())
		return
	}
	defer cd.Close()

	// 转换文件内容并写入输出文件
	buf := make([]byte, 1024)
	for {
		n, err := inFile.Read(buf)
		if err != nil && err.Error() != "EOF" {
			logger.Println("读取文件时发生错误：%s\n", err.Error())
			return
		}
		if n == 0 {
			break
		}

		inBytes := buf[:n]
		outBytes := make([]byte, len(inBytes)*2) // 可能需要调整转换后的字节长度

		_, _, err = cd.Conv(inBytes, outBytes)
		if err != nil {
			logger.Println("转换文件时发生错误：%s\n", err.Error())
			return
		}

		_, err = outFile.Write(outBytes)
		if err != nil {
			logger.Println("写入输出文件时发生错误：%s\n", err.Error())
			return
		}
	}

	logger.Println("文件 %s 的编码已成功转换并保存为 %s\n", file, newFilePath)
}

func CreateLogger(filePath string) (*log.Logger, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	logger := log.New(file, "", log.LstdFlags)
	logger.SetOutput(file)

	return logger, nil
}

func generateConvertedFilePath(filePath string, logger *log.Logger) (string,error) {
	fileName := filepath.Base(filePath)
	basePath := filepath.Dir(filePath)
	newDirPath := filepath.Join(basePath, "converted")

	// 检查目录是否已存在
	if _, err := os.Stat(newDirPath); os.IsNotExist(err) {
		// 目录不存在，创建新目录
		err := os.Mkdir(newDirPath, 0755)
		if err != nil {
			return "", err
		}
		logger.Println("Directory created:", newDirPath)
	} else {
		logger.Println("Directory already exists:", newDirPath)
	}

	newFilePath := filepath.Join(newDirPath, fileName)
	return newFilePath,nil
}
