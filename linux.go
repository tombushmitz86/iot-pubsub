package main

import (
	"errors"
	"strings"
)

// Row represent a route table route for serializing purposes
type Row struct {
	Destination string
	Gateway     string
	Genmask     string
	Flags       string
	MSS         string
	Window      string
	irtt        string
	Iface       string
}

func newRow(input []string) *Row {
	return &Row{
		Destination: input[0],
		Gateway:     input[1],
		Genmask:     input[2],
		Flags:       input[3],
		MSS:         input[4],
		Window:      input[5],
		irtt:        input[6],
		Iface:       input[7],
	}
}

// ParseDGLinux parses deafult gateway Route from output stdout from
// 'netstat -rn' command, which output the columns:
//[Destination, Gateway, Genmask, Flags, MSS Window, irtt, Iface]
func ParseDGLinux(output []byte) (Route, error) {
	var route Route

	rows := strings.Split(string(output), "\n")
	rows = rows[2:]
	for _, row := range rows {
		rowRoute := newRow(strings.Fields(row))
		//DG route
		if rowRoute.Destination == "0.0.0.0" {
			return Route{
				gateway: rowRoute.Gateway,
				ifc:     rowRoute.Iface,
			}, nil
		}
	}
	return route, errors.New("Could not find route")
}

// Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
// 0.0.0.0         10.0.0.138      0.0.0.0         UG        0 0          0 enp1s0
// 10.0.0.0        0.0.0.0         255.255.255.0   U         0 0          0 enp1s0
// 169.254.0.0     0.0.0.0         255.255.0.0     U         0 0          0 docker_gwbridge
// 172.17.0.0      0.0.0.0         255.255.0.0     U         0 0          0 docker0
// 172.18.0.0      0.0.0.0         255.255.0.0     U         0 0          0 br-7e1463b82617
// 172.19.0.0      0.0.0.0         255.255.0.0     U         0 0          0 docker_gwbridge
// 172.20.0.0      0.0.0.0         255.255.0.0     U         0 0          0 br-ef0091c8bfa3
// 192.168.86.0    0.0.0.0         255.255.255.0   U         0 0          0 vmnet8
// 192.168.232.0   0.0.0.0         255.255.255.0   U         0 0          0 vmnet1
