package reader

import (
	"io/ioutil"
	"log"
	"strings"
)

/*
    Ini File Format (example.ini):
        key=value
        key=value
        ...
*/
func ReadIni(path string) (content map[string]string, ok bool) {
	log.Printf("Trying to read from file %s\n", path)
	ini, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("ioutil.ReadFile: Error reading %s (%v)\n", path, err)
		ok = false
		return content, ok
	}
	content = make(map[string]string)
	var tmp []string
	for _, line := range strings.Split(string(ini), "\n") {
		if len(line) == 0 {
			continue
		}
		tmp = strings.Split(line, "=")
		content[strings.Trim(tmp[0], " ")] = strings.Trim(tmp[1], " ")
	}
	ok = true
	return content, ok
}
