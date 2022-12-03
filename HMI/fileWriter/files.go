package fileWriter

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

const FILE_DIR string = "files/"

type FileWriter struct {
	file         *os.File
	writer       *csv.Writer
	InputChannel chan []string
	QuitChannel  chan bool
}

func NewWriter(fileName string) (*FileWriter, error) {
	dirName := filepath.Dir(FILE_DIR + fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	f, err := os.Create(FILE_DIR + fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	w := csv.NewWriter(f)
	var file_writer FileWriter
	file_writer.file = f
	file_writer.writer = w
	file_writer.InputChannel = make(chan []string, 10)
	file_writer.QuitChannel = make(chan bool)
	return &file_writer, nil
}

func (fw *FileWriter) Record() {
	var records [][]string
	defer close(fw.InputChannel)
	defer close(fw.QuitChannel)
	defer fw.writer.Flush()
	defer fw.file.Close()
	for {
		select {
		case done := <-fw.QuitChannel:
			if done {
				return
			}
		}
		select {
		case data := <-fw.InputChannel:
			records = append(records, [][]string{data}...)
		default:
		}
		if len(records) > 10 {
			err := fw.writer.WriteAll(records)
			if err != nil {
				log.Println(err)
			} else {
				records = [][]string{}
			}
		}
	}
}
