package g

import (
	"log"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logFile, err := os.Create("report.log")
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error!")
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
