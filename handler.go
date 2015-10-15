package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetDeviceByIPHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipaddr := vars["ipaddr"]

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	device, err := GetDeviceByIP(ipaddr)

	if err != nil {
		http.Error(w, "IP Address is not allocated", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(device); err != nil {
		fmt.Println("Json Encoding Error")
		panic(err)
	}
}

func AllocateIPHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var device Device
	err = json.Unmarshal(body, &device)

	if device.Name == "" {
		http.Error(w, "Invalid DeviceName", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}

	if !IsValidIPAddress(device.IPAddress) {
		http.Error(w, "Invalid IPAddress, it should 1.2.<0-255>.<0-255>", http.StatusBadRequest)
		return
	}

	_, err = GetDeviceByIP(device.IPAddress)
	if err == nil {
		http.Error(w, "IPAddress Already Assigned use different one", http.StatusBadRequest)
		return
	}

	err = WriteDevice(device)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IsValidIPAddress(ipaddr string) (valid bool) {
	valid, err := regexp.MatchString("1.2.[0-9][0-9]*.[0-9][0-9]*", ipaddr)
	if err != nil || !valid {
		return
	}
	substr := strings.Split(ipaddr, ".")

	value, err := strconv.Atoi(substr[2])

	if value > 255 {
		valid = false
		return
	}

	value, err = strconv.Atoi(substr[3])
	if value > 255 {
		valid = false
		return
	}

	valid = true
	return
}

func AddHandlers() {
	h := mux.NewRouter()
	h.HandleFunc("/ipalloc/view/{ipaddr:1.2.[0-9][0-9]*.[0-9][0-9]*}", GetDeviceByIPHandler)
	h.HandleFunc("/ipalloc/add/", AllocateIPHandler)
	http.ListenAndServe(":8080", h)
}
