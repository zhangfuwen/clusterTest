package github.com/zhangfuwen/clusterTest

import (
	"github.com/zhangfuwen/sshutils/host"
	"bufio"
	"regexp"
	"errors"
	"time"
)

type Node struct {
	host.Host
}

/**
	WaitLogFor waits for specified log line that matches 'regexString', for 'howLong'
	Close and return error if time is out
 */
func (node *Node) WaitLogFor(filepath string, regexString string, howLong time.Duration) error {
	regex,err := regexp.Compile(regexString)
	if err!=nil {
		return err
	}
	rc, err:=node.Tailf(filepath)
	if err!=nil {
		return err
	}
	lineReader := bufio.NewReader(rc)
	timeout := make (chan bool)
	go func() {
		time.Sleep(howLong) // sleep one second
		timeout <- true
	}()
	lineChan := make(chan string)
	go func() {
		for {
			line, err := lineReader.ReadString('\n')
			if err == nil {
				lineChan<-line
			}
		}
	}()
	select {
	case line:=  <- lineChan:
		if regex.FindString(line) != "" {
			rc.Close()
			return nil
		}
	case <- timeout:
		rc.Close()
		return errors.New("timeout")
	}
	return nil
}

