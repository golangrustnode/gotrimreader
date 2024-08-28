package trimreader

import (
	"github.com/golangrustnode/log"
	"io/ioutil"
	"strings"
)

func ReadAsString(file_path string) (string, error) {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Info("Error reading file: ", file_path, " Error: ", err)
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
