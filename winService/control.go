package winService

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/svc"
)

func Usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func CommandParser(serviceName string, desc string) {

	cmd := strings.ToLower(os.Args[1])

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if cmd == "debug" {
		runService(serviceName, true)
		return
	}
	if !isIntSess {
		runService(serviceName, false)
		return
	}
	// var err error
	if len(os.Args) < 2 {
		Usage("no command specified")
	}

	switch cmd {
	// case "debug":
	// 	runService(serviceName, true)
	// 	return
	case "install":
		err = installService(serviceName, desc)
	case "remove":
		err = removeService(serviceName)
	case "start":
		err = startService(serviceName)
	case "stop":
		err = controlService(serviceName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(serviceName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(serviceName, svc.Continue, svc.Running)
	default:
		Usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, serviceName, err)
	}
}
