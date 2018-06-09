package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
)

type Route struct {
	gateway string
	ifc     string
}

func main() {
	var cmd *exec.Cmd
	var dg Route
	var routeError error

	os := runtime.GOOS

	switch os {
	case "linux":
		cmd = exec.Command("netstat", "-rn")
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		dg, routeError = ParseDGLinux(output)

	case "darwin":
		cmd = exec.Command("route", "get", "default")
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		dg, routeError = ParseDGDarwin(output)
	}

	if routeError != nil {
		log.Fatal(routeError)
	}

	ic, _ := net.InterfaceByName(dg.ifc)
	addrs, _ := ic.Addrs()

	hosts, err := GetHosts(addrs[0].String())
	if err != nil {
		log.Fatal(err)
	}
	for _, host := range hosts {
		fmt.Printf("%s\n", host)

	}

}
