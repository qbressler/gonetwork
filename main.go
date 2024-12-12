package main

// This application gets basic network info
// Quintin Bressler Dec 10 2024
//
import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), nil

}

func sendPing() bool {
	out, _ := exec.Command("ping", "google.com", "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		return false
	} else {
		return true
	}

}

func main() {

	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("Please provide a domain name")
		fmt.Println("gonetwork google.com")
		return
	}

	pingDomain := os.Args[1]
	fmt.Println("Gathering network info....")

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Something went horribly wrong %v", err)
	}

	for _, iface := range interfaces {
		fmt.Println("================================================")
		fmt.Printf("Name: %s\n", iface.Name)
		fmt.Printf("Hardware Address: %s\n", iface.HardwareAddr)

		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		} else {
			for _, arr := range addrs {
				fmt.Printf("IP Address: %s\n", arr.String())
			}
		}

	}
	fmt.Println("================================================")

	fmt.Printf("Sending packets of data for ping to %s...hold up, bro\n", pingDomain)
	if sendPing() {
		fmt.Println("Internet is working as expected!")
		publicIP, err := getPublicIP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Your public IP Address is %s\n", publicIP)
	}
}
