package main


import(
        //"fmt"
        "flag"
        "syscall"
        //"fmt"
        //"runtime"
        //"runtime/debug"
        "os"
        "strconv"
)

func main(){
        //err := syscall.NewExec("./fib.py", 1, []string{"./fib.py","20","1"}, os.Environ())
        //fmt.Println(err)
	pred := flag.Int("p",0,"pred")
	var id string
	flag.StringVar(&id, "i","","id")
	var para string
        flag.StringVar(&para, "a","","para")
	var o string
	flag.StringVar(&o, "o","","output")
        flag.Parse()
	syscall.NewExec("./fib.py", *pred, []string{"./fib.py",strconv.Itoa(*pred),id,para, o}, os.Environ())
        //flag.Usage()
}
