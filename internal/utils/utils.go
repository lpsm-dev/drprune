package utils

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ci-monk/drprune/internal/constants"
	"github.com/sirupsen/logrus"
)

// IsEmpty function - check if a string is empty.
func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// IsDirExists function - check fi a directory exist in te system.
func IsDirExists(path string) bool {
	result, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return result.IsDir()
}

// IsFileExists function - check fi a file exist in te system.
func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// MakeDirIfNotExist create a new directory if they not exist.
func MakeDirIfNotExist(dir string) {
	fullDir, _ := filepath.Abs(filepath.Dir(dir))
	mode := os.FileMode(0775)
	if !IsDirExists(fullDir) {
		err := os.MkdirAll(fullDir, mode)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

// CreateLogFile function
func CreateLogFile(logdir, logfile string) string {
	pid := strconv.Itoa(os.Getpid())
	return filepath.Join(logdir,
		logfile+
			".pid"+
			pid+
			"."+
			time.Now().Format(constants.DefaultTimestampFormat)+
			".log",
	)
}

// EncodeParam function to encode a string a return a url format
func EncodeParam(s string) string {
	return url.QueryEscape(s)
}

// DecodeParam function to decode a url format to string
func DecodeParam(s string) string {
	in, _ := url.QueryUnescape(s)
	return in
}
