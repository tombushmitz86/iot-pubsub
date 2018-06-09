package main

import (
	"errors"
	"strings"
)

// ParseDGDarwin parses deafult gateway Route from output stdout from
// 'route get default' command, which output the columns:
//[Destination, Gateway, Genmask, Flags, MSS Window, irtt, Iface]
func ParseDGDarwin(output []byte) (Route, error) {
	var route Route
	var gateway, ifc string

	rows := strings.Split(string(output), "\n")
	for _, row := range rows {
		s := strings.Trim(row, " ")
		if strings.Contains(s, "gateway") {
			gateway = strings.Split(s, ":")[1]
		} else if strings.Contains(s, "interface") {
			ifc = strings.Split(s, ":")[1]
		}
	}
	if gateway == "" || ifc == "" {
		return route, errors.New("Could not get Route object")
	}
	return route, nil
}

// route to: default
// destination: default
//        mask: default

//     gateway: 172.20.10.1
//   interface: en0
//       flags: <UP,GATEWAY,DONE,STATIC,PRCLONING>
//  recvpipe  sendpipe  ssthresh  rtt,msec    rttvar  hopcount      mtu     expire
//    0         0         0         0         0         0      1500         0
