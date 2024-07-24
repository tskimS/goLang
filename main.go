package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println("hello")
	saveModuleInfo()
}

func saveModuleInfo() {
	pathList := []string{
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\DS_vs2019\DS_vs2019\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\DS_vs2019\DS_vs2019\x64\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\DS6.0_VS2019(ZTCA_AIP)\DS6.0_VS2019(ZTCA_AIP)\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\DS6.0_VS2019(ZTCA_AIP)\DS6.0_VS2019(ZTCA_AIP)\x64\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\SCSDK_vs2019\SCSDK_vs2019\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\SCSDK_vs2019\SCSDK_vs2019\x64\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\SCShield\SCShield\Release`,
		`\\tfsbuild.softcamp.co.kr\TFSBuild_new\2024\SCShield\SCShield\x64\Release`,
		`c:\\windows\\softcamp`,
	}

	var totalFileList []string
	for _, path := range pathList {
		findFileInfoList, err := getFileList(path)
		if nil == err {
			totalFileList = append(totalFileList, findFileInfoList...)
		}
	}

	csvFilePath := "file_info_list.csv"
	err := saveToCSV(totalFileList, csvFilePath)
	if err != nil {
		log.Fatalf("Error saving to CSV: %v", err)
	}

	fmt.Printf("File info list saved to %s\n", csvFilePath)

}
