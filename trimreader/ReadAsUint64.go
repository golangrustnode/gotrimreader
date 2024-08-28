package trimreader

import (
	"github.com/golangrustnode/log"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadAsUint64(file_path string) (uint64, error) {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Info("Error reading file: ", file_path, " Error: ", err)
		return 0, err
	}
	data_str := strings.TrimSpace(string(data))
	return strconv.ParseUint(data_str, 10, 64)
}
