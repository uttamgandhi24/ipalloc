package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Device struct {
	Name      string `bson:"Name"`
	IPAddress string `bson:"IPAddress"`
}

const kIPMapFileName string = "data/ip_map"

func GetDeviceByIP(ipaddr string) (device Device, err error) {
	lines := ReadFile(kIPMapFileName)
	devices := GetAllDevices(lines)
	_, found := devices[ipaddr]
	if !found {
		return device, errors.New("Not Found")
	}
	device.Name = devices[ipaddr]
	device.IPAddress = ipaddr
	return device, nil
}

func WriteDevice(device Device) (err error) {
	deviceStr := "1.2.0.0/16," + device.IPAddress + "," + device.Name + "\n"
	AppendFile(kIPMapFileName, deviceStr)
	return nil
}

func AppendFile(filename string, line string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("File Open Error")
		return
	}

	_, err = f.WriteString(line)
	if err != nil {
		fmt.Println("Write Error")
		return
	}

	f.Close()
}

func ReadFile(filename string) (lines []string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("File Open Error")
		return
	}
	reader := bufio.NewReader(file)

	for i := 0; ; i++ {
		bytes, _ := reader.ReadBytes('\n')
		n := len(bytes)
		if n <= 0 {
			break
		}
		lines = append(lines, string(bytes[:n-1]))
	}
	return
}

func GetAllDevices(lines []string) (devices map[string]string) {
	devices = make(map[string]string)
	for i := 0; i < len(lines); i++ {
		strlist := strings.Split(lines[i], ",")
		key := strings.Trim(strlist[1], " ")
		value := strings.Trim(strlist[2], " ")
		devices[key] = value
	}
	return
}
