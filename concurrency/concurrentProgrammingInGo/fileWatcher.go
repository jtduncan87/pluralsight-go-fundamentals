package concurrentProgrammingInGo

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const watchedPath = "./watchedPath"

//FileWatcherMain copy files from resources to watchedPath to see this work
func FileWatcherMain() {
	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1)
		for _, file := range files {
			filePath := watchedPath + "/" + file.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					champ := new(championships)
					champ.Number, _ = strconv.Atoi(r[0])
					champ.Year, _ = strconv.Atoi(r[1])
					fmt.Printf("Received championship: %d, year %d\n", champ.Number, champ.Year)
				}
			}(string(data))
		}
	}
}

type championships struct {
	Number int
	Year   int
}
