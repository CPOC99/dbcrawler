package trace

import (
	"cpoc/dbcrawler/config"
	"fmt"
	"github.com/DataDog/gopsutil/process"
	"log"
	"strings"

	"github.com/DataDog/gopsutil/net"
)

const (
	ExecPath = "ExecPath"
)

type DBProcess struct {
	name          string
	Pid           int32
	Ppid          int32
	Port          int
	ExecPath      string
	ExecQueryPath string
	LocalAddr     string
}

func SearchPidList(config *config.Config) *[]DBProcess {

	var strNoSpace string

	pids, err := process.AllProcesses()
	dbList := make([]DBProcess, 0)

	if err != nil {
		log.Println("Error fetching processes:", err)

		return nil
	}

	// is possible goroutine ?
	for _, filledPid := range pids {

		exe := filledPid.Exe
		for key, value := range config.StructToMap() {

			//check  key is exist
			if execPath, exist := value.(map[string]interface{})[ExecPath]; exist {

				strNoSpace = strings.ReplaceAll(execPath.(string), " ", "")

				if filledPid.Pid == filledPid.Ppid && strNoSpace != "" && strings.Contains(exe, strNoSpace) {

					connections, _ := net.ConnectionsPid("tcp", filledPid.Pid)

					for _, conn := range connections {
						// conn에 대한 작업 수행
						fmt.Println("Connection:", conn)
					}

					dbProcess := DBProcess{
						name: key,
						Pid:  filledPid.Pid,
						Ppid: filledPid.Ppid,
					}

					dbList = append(dbList, dbProcess)

				}
			}
		}
	}
	return &dbList
}
