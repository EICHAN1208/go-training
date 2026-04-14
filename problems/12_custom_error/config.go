package config

import (
	"fmt"
)

type FileError struct {
	FileName string
	Reason   string
}

func (e *FileError) Error() string {
	if e.FileName == "" {
		return fmt.Sprintf("error: (empty filename), reason: %s", e.Reason)
	}
	return fmt.Sprintf("error: %s, reason: %s", e.FileName, e.Reason)
}

func ReadConfig(filename string) (string, error) {
	switch filename {
	case "":
		return "", &FileError{FileName: "", Reason: "filename is empty"}
	case "config.json":
		return "debug=true", nil
	default:
		return "", &FileError{FileName: filename, Reason: "file not found"}
	}
}

// === 実行側 ===
//
// 正常系:
// v, err := ReadConfig("config.json")
// fmt.Println(v, err) => debug=true, <nil>
//
// 異常系:
// _, err := ReadConfig("missing.txt")
// var fe *FileError
// if errors.As(err, &fe) {
//   fmt.Println("ファイル名:", fe.FileName)
//   fmt.Println("理由:", fe.Reason)
// }
