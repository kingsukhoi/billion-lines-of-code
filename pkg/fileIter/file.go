package fileIter

import (
	"bufio"
	"iter"
	"os"
)

type FileInfo struct {
	FilePath string
	file     os.File
}

func (f *FileInfo) Open() error {
	fileHandler, err := os.Open(f.FilePath)
	if err != nil {
		return err
	}
	f.file = *fileHandler
	return nil
}

func GetNumLines(filePath string) (int, error) {
	rtnMe := 0
	osFile, err := os.Open(filePath)
	if err != nil {
		return rtnMe, err
	}
	scanner := bufio.NewScanner(osFile)
	for scanner.Scan() {
		rtnMe++
	}
	return rtnMe, osFile.Close()

}

func (f *FileInfo) All() iter.Seq2[int, string] {
	scanner := bufio.NewScanner(&f.file)
	lineNum := 0
	return func(yield func(int, string) bool) {
		for scanner.Scan() {
			moreData := yield(lineNum, scanner.Text())
			lineNum++
			if !moreData {
				return
			}
		}
	}
}

func (f *FileInfo) Close() error {
	return f.file.Close()
}
