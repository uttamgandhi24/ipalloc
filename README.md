This is a webservice using following
 - golang's net/http capabilities
 - gorilla mux

To use this app following are pre-requisites
 - go should be installed from here 'https://golang.org/dl/'
 - get gorilla mux using 'go get github.com/gorilla/mux'
 - data directory should exist in current directory from where service is to be run
 - there should be a file named 'ip_map' in the 'data' directory

The ip_map file has following format
 ip_block, ip_address, device_name
e.g.

  1.2.0.0/16,1.2.3.4,device1
  1.2.0.0/16,1.2.3.5,device2
  1.2.0.0/16,1.2.3.6,device3

This service can be used client making REST calls
to this.

Steps to build and run
1. go install ipalloc
2. run the ipalloc binary, ./ipalloc

Supported REST APIs
The root is http://localhost:8080/

Supported APIs are

1. GET request
-----------
 "/ipalloc/view/1.2.<0-255>.<0-255>"

    -- this returns the device associated with this IP Address
    the result is returned as json text

    e.g. Request-> curl http://localhost:8080/ipalloc/view/1.2.3.4
         Response-> {"Name":"device1","IPAddress":"1.2.3.4"}

2. POST request
---------------
 "/ipalloc/add/"

  -- this allocates IP address to given device
  e.g. Request-> curl -d '{"Name":"device8","IPAddress":"1.2.3.25"}' http://localhost:8080/ipalloc/add/

  will add this device's name and IPAddress in data store

TODO
-------------------------
1. IP block e.g 1.2.0.0/16 is not read from the file, the value is hardcoded.
2. Unit tests need to be written
3. Some errors are ignored, all errors should be handled
