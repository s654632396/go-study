package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var s ServerList
	s.Servers = append(s.Servers, Server{"Shanghai_VPN", "127.0.0.1"})
	s.Servers = append(s.Servers, Server{"Wuhan_VPN", "127.0.0.2"})

	v, err := json.MarshalIndent(s, "", "\t")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(v))
}

type ServerList struct {
	Servers []Server
}

type Server struct {
	ServerName string `json:"server_name"`
	ServerIP   string `json:"server_ip"`
}

func case_01() {

	jsonStr := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	var s ServerList
	json.Unmarshal([]byte(jsonStr), &s)

	fmt.Println(s)

	var f interface{}

	json.Unmarshal([]byte(jsonStr), &f)

	fmt.Println(f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
