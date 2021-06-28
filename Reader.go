package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type node struct {
	name    string
	status  string
	roles   string
	age     string
	version string
}

func main() {
	cmd := exec.Command("kubectl", "get", "node")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	} else {

		strtmp := string(output)

		lines := strings.Split(strtmp, "\n")

		for _, str := range lines {
			fmt.Println(str)
			GetNode_(str)
		}
	}
}

func GetNodes() {
	cmd := exec.Command("kubectl", "get", "node")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	rows := strings.Split(string(output), "\n")

	for i := 0; i < len(rows); i++ {
		if i != 0 {
			node := InitNode(rows[i])
			//add
		}
	}

	return
}

func InitNode(para string) node {
	var idx int = 0
	var name, status, roles, age, version string

	strtemp := strings.Split(para, " ")

	for i := 0; i < len(strtemp); i++ {
		if strtemp[i] != "" {
			switch idx {
			case 0:
				name = strtemp[i]
			case 1:
				status = strtemp[i]
			case 2:
				roles = strtemp[i]
			case 3:
				age = strtemp[i]
			case 4:
				version = strtemp[i]
			}
			idx++
		}
	}

	temp := node{}
	temp.name = name
	temp.status = status
	temp.roles = roles
	temp.age = age
	temp.version = version

	return temp
}

func GetNode_(para string) node {

	var idx int = 0
	var name, status, roles, age, version string

	strtemp := strings.Split(para, " ")

	for i := 0; i < len(strtemp); i++ {
		if strtemp[i] != "" {
			switch idx {
			case 0:
				name = strtemp[i]
			case 1:
				status = strtemp[i]
			case 2:
				roles = strtemp[i]
			case 3:
				age = strtemp[i]
			case 4:
				version = strtemp[i]
			}
			idx++
		}
	}

	temp := node{}
	temp.name = name
	temp.status = status
	temp.roles = roles
	temp.age = age
	temp.version = version

	return temp
}
