package exiftool

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func Read(path string) (map[string]interface{}, error) {

	cmd := exec.Command("exiftool", "-a", "-s", "-G", "-j", path)
	out, err := cmd.Output()
	if err != nil {
		return make(map[string]interface{}), fmt.Errorf("Unable to read exif metadata using exiftool")
	}

	result := make([]map[string]interface{}, 0)
	err = json.Unmarshal(out, &result)
	if err != nil {
		return make(map[string]interface{}), fmt.Errorf("Unable to unmarshal json output from exiftool")
	}

	if len(result) != 1 {
		return make(map[string]interface{}), fmt.Errorf("Expected single metadata array in exiftool output")
	}

	return result[0], nil

}
