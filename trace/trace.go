package trace

import (
	"cpoc/dbcrawler/config"
	"log"
	"strings"

	"github.com/shirou/gopsutil/process"
)

const (
	ExecPath = "ExecPath"
)

type DBProcess struct {
	name          string
	Pid           int
	Ppid          int
	Port          int
	ExecPath      string
	ExecQueryPath string
	LocalAddr     string
}

// later needs goroutine
// if expect slice enough big ... use pointer
func SearchPidList(config *config.Config) *[]DBProcess {

	var strNoSpace string

	pids, err := process.Processes()
	dbList := make([]DBProcess, 0)

	if err != nil {
		log.Println("Error fetching processes:", err)

		return nil
	}

	// is possible goroutine ?
	for _, pid := range pids {

		exe, _ := pid.Exe()

		for key, value := range config.StructToMap() {

			//check  key is exist
			if execPath, exist := value.(map[string]interface{})[ExecPath]; exist {

				strNoSpace = strings.ReplaceAll(execPath.(string), " ", "")

				if strNoSpace != "" && strings.Contains(exe, strNoSpace) {

					dbProcess := DBProcess{
						name: key,
					}

					dbList = append(dbList, dbProcess)

				}
			}
		}
	}
	return &dbList
}
