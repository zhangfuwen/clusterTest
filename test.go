package clusterTest

import (
	"reflect"
	"runtime"
)

type SubTest struct {
	name string
	doTest func()bool
	success bool
}

func (s * SubTest) Start() {
	s.success = s.doTest()
}
type Test struct {
	name string
	subtests []SubTest
}
func NewTest(name string) *Test {
	var t = Test {
		name,
		make([]SubTest,0),
	}
	return &t
}

func (t *Test)AddSubTest(name string,f func()bool) *Test {
	var subTestName = name
	if subTestName=="" {
		subTestName = getFunctionName(f)
	}
	t.subtests = append(t.subtests, SubTest{
		subTestName,
		f,
		false,
	})
	return t
}

func getFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
func (t *Test) Start() {
	for _, subTest := range t.subtests {
		subTest.Start()
	}
}

func (t *Test) Conclude(printf func(format string, a ...interface{}) (n int, err error)) {
	var nSucceeded int = 0
	printf("Test %d conclusion:", t.name)
	for _, subTest := range t.subtests {
		printf("\t%s\t:%v",subTest.name, subTest.success)
		if subTest.success {
			nSucceeded++
		}
	}

	printf("Total:%d/%d succeeded for test %s",nSucceeded, len(t.subtests),t.name)
}