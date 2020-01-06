package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func EOL() string {
	var ret string
	osi := runtime.GOOS
	switch osi {
	case "linux":
		ret = "\n"
	case "windows":
		ret = "\r\n"
	case "macos":
		ret = "\r"
	default:
		ret = "\n"
	}

	return ret
}

// get started file path
func GetRunPath() (string, error) {
	var dirAbsPath string
	ex, err := os.Executable()
	if err == nil {
		dirAbsPath = filepath.Dir(ex)
		return dirAbsPath, err
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		return exReal, err
	}
	dirAbsPath = filepath.Dir(exReal)
	return dirAbsPath, err
}

func MD5sum(filePath string) (result string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	result = hex.EncodeToString(hash.Sum(nil))
	return
}
