package test

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

type list []string

type ctx struct {
	l list
}

func TestStruct(t *testing.T) {
	command := "go"
	args := []string{"tool", "dist", "list"}
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("执行命令时发生错误：%v\n", err)
	}
	//fmt.Println(string(output))
	split := strings.Split(string(output), "\n")
	fmt.Println(len(split))
	m := make(map[string][]string)
	for i := 0; i < len(split); i++ {
		fmt.Println(i)
		//fmt.Println(split[i])
		s := split[i]
		if len(s) > 0 {
			fmt.Println(s)
			s2 := strings.Split(s, "/")
			k := s2[1]
			m2, exists := m[k]
			if exists {
				m[k] = append(m2, s2[0])
			} else {
				m[k] = append(make([]string, 0), s2[0])
			}
		}
	}
	//fmt.Printf("%v\n", m)
	for k, v := range m {
		fmt.Printf("key=%s, value=%v\n", k, v)
	}
}
