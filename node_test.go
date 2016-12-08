package clusterTest

import (
	"testing"
	"github.com/zhangfuwen/sshutils/host"
	"time"
	"os/exec"
	"sync"
	"strings"
)
var wg sync.WaitGroup

func TestWaitLogFor(t *testing.T) {
	var info []string
	if bs,err:= exec.Command("sh","-c","cat ./ssh.account").Output();err!=nil {
		t.Error(err)
		return
	}else{
		info = strings.Split(strings.TrimSpace(string(bs)),";")
	}
	var node = Node{
		host.NewHost(info[0],info[1],info[2]),
	}
	go func() {
		time.Sleep(3*time.Second)
		err:=exec.Command("sh","-c","rm ./test.log").Run()
		err=exec.Command("sh","-c", "echo 'haha' > /home/dean/test.log").Run()
		if err!=nil {
			t.Error(err)
		}
	}()
	if err:=node.WaitLogFor("/home/dean/test.log","haha", 10*time.Second); err!=nil {
		t.Error(err)
	}
}
