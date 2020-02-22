package tests

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestShowFlagBility(tt *testing.T) {
	// 实际中应该用更好的变量名
	var (
		h bool

		v, V bool
		t, T bool
		q    bool

		s string
		p string
		c string
		g string
	)

	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	flag.BoolVar(&q, "q", false, "suppress non-error messages during configuration testing")
	// 另一种绑定方式
	// q *bool
	//q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] ] [-g directives]

Options:
	`)
		flag.PrintDefaults()
	}

	flag.Parse()

	flag.Usage()

}