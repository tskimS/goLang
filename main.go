package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println("hello")
	saveModuleInfo()
}

func saveModuleInfo() {
	pathList := []string{
		`c:\windows\softcamp`,
		// `\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\SCUtility_vs2019\SCUtility_vs2019`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild\00.SDS_Common`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild\10.SC_Add`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild\40.DSDateSetKeySetup`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild\D1.DS4.0_DSOnly_Default`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild\D2.DS4.0_Add_Feature`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS6.0_VS2019(ZTCA_AIP)\InstallSheild\DS27.SCMIP`,
	}

	for i, path := range pathList {
		findFileInfoList, err := getFileList(path)
		if nil == err {
			csvFilePath := fmt.Sprintf("[%d]file_info_list.csv", i)
			os.Remove(csvFilePath)
			err = saveToCSV(findFileInfoList, csvFilePath)
			findFileInfoList = nil
			if err != nil {
				log.Fatalf("Error saving to CSV: %v", err)
			}

			fmt.Printf("File info list saved to %s\n", csvFilePath)
		}
	}
}
