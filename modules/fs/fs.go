package fs

import (
	"os"
)

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func OpenFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

func FilePutContent(name string, value string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.WriteString(value)
		fo.Close()
	}()
	return nil
}

func Remove(file string) error {
	return os.Remove(file)
}
