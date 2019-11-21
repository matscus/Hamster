package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	MultiHosts struct {
		Hosts []string `yaml:"hosts"`
		Ports []string `yaml:"ports"`
	} `yaml:"Multi_hosts"`
	SingleHost []struct {
		Host  string   `yaml:"host"`
		Ports []string `yaml:"ports"`
	} `yaml:"Single_host"`
}
type result struct {
	host   string
	port   string
	status string
}

func main() {
	fmt.Println("Start")
	file, err := ioutil.ReadFile("config.yml")
	config := config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	d := net.Dialer{Timeout: 1 * time.Second}
	res := make([]result, 0, len(config.MultiHosts.Hosts)+len(config.MultiHosts.Ports))
	for _, host := range config.MultiHosts.Hosts {
		for _, port := range config.MultiHosts.Ports {
			_, err = d.Dial("tcp", host+":"+port)
			if err != nil {
				ok := os.IsTimeout(err)
				if ok {
					res = append(res, result{host: host, port: port, status: "Fail"})
				} else {
					res = append(res, result{host: host, port: port, status: "Pass"})
				}
			} else {
				res = append(res, result{host: host, port: port, status: "Pass"})
			}
		}
	}
	for _, host := range config.SingleHost {
		for _, port := range host.Ports {
			_, err = d.Dial("tcp", host.Host+":"+port)
			if err != nil {
				ok := os.IsTimeout(err)
				if ok {
					res = append(res, result{host: host.Host, port: port, status: "Fail"})
				} else {
					res = append(res, result{host: host.Host, port: port, status: "Pass"})
				}
			} else {
				res = append(res, result{host: host.Host, port: port, status: "Pass"})
			}
		}
	}
	for _, v := range res {
		fmt.Printf("Host: %s, port: %s, status %s\n", v.host, v.port, v.status)
	}
	fmt.Println("complited")
}
