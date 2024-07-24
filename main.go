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
		`c:\\windows\\softcamp`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS_vs2019\InstallSheild`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\Package\Product\DS6.0_VS2019(ZTCA_AIP)\InstallSheild`,
	}

	var totalFileList []string
	for i, path := range pathList {
		findFileInfoList, err := getFileList(path)
		if nil == err {
			totalFileList = append(totalFileList, findFileInfoList...)
		}

		csvFilePath := fmt.Sprintf("[%d]file_info_list.csv", i)
		os.Remove(csvFilePath)
		err = saveToCSV(totalFileList, csvFilePath)
		if err != nil {
			log.Fatalf("Error saving to CSV: %v", err)
		}

		fmt.Printf("File info list saved to %s\n", csvFilePath)

	}

}
