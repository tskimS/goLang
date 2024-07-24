package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

func getFileVersion(filePath string) (string, error) {
	infoSize, err := windows.GetFileVersionInfoSize(filePath, nil)
	if err != nil {
		return "", err
	}

	info := make([]byte, infoSize)
	err = windows.GetFileVersionInfo(filePath, 0, infoSize, unsafe.Pointer(&info[0]))
	if err != nil {
		return "", err
	}

	var fixedInfo *windows.VS_FIXEDFILEINFO
	var length uint32
	err = windows.VerQueryValue(unsafe.Pointer(&info[0]), `\`, unsafe.Pointer(&fixedInfo), &length)
	if err != nil {
		return "", err
	}

	major := fixedInfo.FileVersionMS >> 16
	minor := fixedInfo.FileVersionMS & 0xFFFF
	build := fixedInfo.FileVersionLS >> 16
	revision := fixedInfo.FileVersionLS & 0xFFFF

	return fmt.Sprintf("%d.%d.%d.%d", major, minor, build, revision), nil
}

func isExtensionMatch(filename string, extList []string) bool {
	filename = strings.ToLower(filename)
	ext := filepath.Ext(filename)
	for _, _ext := range extList {
		if ext == _ext {
			return true
		}
	}

	return false
}

func getFileList(searchPath string) ([]string, error) {
	var fileInfoList []string
	var supportExt = []string{".dll", ".exe", ".ocx", ".sys", ".png", ".bmp", ".ico", ".ini", ".rc", ".json", ".html"}

	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && isExtensionMatch(path, supportExt) {
			version, err := getFileVersion(path)
			if err == nil {
				fileInfoList = append(fileInfoList, fmt.Sprintf("%s\t%s\t%s\t%d", info.Name(), version, filepath.Dir(path), info.Size()))
			} else {
				fileInfoList = append(fileInfoList, fmt.Sprintf("%s\t%s\t%s\t%d", info.Name(), "none", filepath.Dir(path), info.Size()))
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("[Error] walking the path %v: %v", searchPath, err)
	}

	return fileInfoList, nil
}

// saveToCSV saves the provided fileInfoList to a CSV file.
func saveToCSV(fileInfoList []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"File Name", "Version", "Directory", "byte size"})

	// Write data
	for _, fileInfo := range fileInfoList {
		fields := strings.Split(fileInfo, "\t")
		if err := writer.Write(fields); err != nil {
			return fmt.Errorf("could not write to csv: %v", err)
		}
	}

	return nil
}
