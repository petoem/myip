package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/petoem/myip"
)

func main() {
	v6 := flag.Bool("v6", false, "discover IPv6 address")
	newline := flag.Bool("n", false, "print newline after the output")
	stun := flag.String("stun", "", "server to use for discovery (e.g. stun:your.server.host:3478)")
	liststunservers := flag.Bool("list-stun-servers", false, "list baked in public stun server")
	version := flag.Bool("version", false, "display version information")
	verbose := flag.Bool("verbose", false, "verbose output")
	flag.Parse()

	if !*verbose {
		log.SetFlags(0)
		log.SetOutput(io.Discard)
	}

	defer func() {
		if *newline {
			fmt.Print("\n")
		}
	}()

	if *version {
		fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), myip.Version())
		fmt.Printf("\n%s", myip.License())
		return
	}

	if *liststunservers {
		servers := myip.GetStunServers()
		for i, server := range servers {
			fmt.Print(server)
			if i < len(servers)-1 {
				fmt.Print("\n")
			}
		}
		return
	}

	if *stun == "" {
		servers := myip.GetStunServers()
		*stun = servers[rand.Intn(len(servers))]
	}

	if ipaddress, err := myip.DiscoverIP(*stun, *v6); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Print(ipaddress)
	}
}
