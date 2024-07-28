package commands

import (
	"io/ioutil"
)

func Complete(line string, fileName string) {
	// lineNumber := strconv.Atoi(line)

	content, _ := ioutil.ReadFile(fileName)
	print(content)

}
