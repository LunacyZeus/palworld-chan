package tests

import (
	"fmt"
	"palworld-chan/pkg/utility/utils"
	"testing"
)

func Test_Dir(t *testing.T) {
	folderPath := "D:\\test"

	// 获取目录下的全部文件名及创建日期
	filesAndCreationDates, err := utils.GetFilesAndCreationDates(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印文件名及创建日期
	fmt.Println("Files and Creation Dates:")
	for file, creationDate := range filesAndCreationDates {
		fmt.Printf("%s: %s\n", file, creationDate)
	}
}
