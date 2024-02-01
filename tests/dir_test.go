package tests

import (
	"fmt"
	"palworld-chan/pkg/utility/utils"
	"testing"
)

func Test_Dir(t *testing.T) {
	folderPath := "D:\\test"

	// 获取目录下的全部文件名及创建日期
	filesAndCreationDates, err := utils.GetBackUpFiles(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%v\n", filesAndCreationDates)
}

func Test_Find_Ini_File(t *testing.T) {
	directoryPath := "F:/pal/steamapps/common/PalServer/Pal/Saved"

	palWorldSettingsPath, err := utils.FindPalWorldSettingsIni(directoryPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("PalWorldSettings.ini found at: %s\n", palWorldSettingsPath)
}
