package readPvLog

import (
	"os"
	"bufio"
	"strings"
	"ideaparLog/Repo"
	"encoding/json"
	"ideaparLog/model"
	"time"
	"fmt"
)

var pvRepo = new(Repo.PvRepo)
var path = "/var/space/icl/icl-main/log/pv.log"
var timeLayout = "2006-01-02 15:04:05"
var local, _ = time.LoadLocation("Asia/Shanghai")

func Run() {
	lastLineNum := findLastLineNumInDatabase()

	insertNewLogToDatabase(lastLineNum)
}
func parseOneLineLog(oneLineLog string) (model.Pv, error) {
	start := strings.Index(oneLineLog, "{")
	end := strings.Index(oneLineLog, "}")
	rawString := oneLineLog[start:end+1]
	date := oneLineLog[1:20]

	pv := model.Pv{}
	err := json.Unmarshal([]byte(rawString), &pv)
	if err != nil {
		fmt.Println("Unmarshal string error")
		return model.Pv{}, err
	}
	pv.Date, err = time.ParseInLocation(timeLayout, date, local)

	return pv, err
}

func insertNewLogToDatabase(lastLineNum int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// Splits on newlines by default.
	scanner := bufio.NewScanner(file)

	line := 1

	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {

		if line > lastLineNum {
			pv, err := parseOneLineLog(scanner.Text())
			if err != nil {
				continue
			}
			pvRepo.Create(&pv)
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
	}
}

func findLastLineNumInDatabase() int {
	lastPvRow := pvRepo.GetLastRow()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// Splits on newlines by default.
	scanner := bufio.NewScanner(file)

	line := 1

	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), lastPvRow.Date.Format(timeLayout)) {
			return line
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
	}

	return 0
}
