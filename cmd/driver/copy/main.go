package main


import(
	"sync"
	//"fmt"
	"time"
	"flag"
	"syscall"
	"fmt"
	//"runtime"
	//"runtime/debug"
	"os"
	"strconv"
	"log"
	"bufio"
)


func main(){
	var outPut string
	flag.StringVar(&outPut, "o","./default_output.txt","output file")
	iat := flag.Int("a",10,"IAT degree")
	var policy string
	flag.StringVar(&policy, "p","d","scheduling policys: d: Default r: RL")
	var inPut string
	flag.StringVar(&inPut, "i","test2","input trace")
	//err := syscall.NewExec("./fib.py", 1, []string{"./fib.py","20","1"}, os.Environ())
	//fmt.Println(err)
	flag.Parse()
	flag.Usage()
	if policy == "d"{
		ExecTraceDefault(inPut, outPut, *iat)
	}else if policy == "r"{
		ExecTraceRL(inPut, outPut, *iat)
	}
}

func ExecTraceRL(inPut string, outPut string, iat int){
	trace, num := GetTrace(inPut, iat)
	for i:= 0; i < len(trace); i++{
		a := trace[i]
		err := syscall.NewExec("./fib.py", a.Pred, []string{"./fib.py",strconv.Itoa(a.Para),strconv.Itoa(a.Pred), outPut}, os.Environ())
		if err != nil{
			log.Fatal("Error: New Execve")
		}
		if i < len(trace) - 1{
			time.Sleep(time.Duration(trace[i+1].Start - trace[i].Start)*time.Millisecond)
		}
	}
	for{
                time.Sleep(time.Duration(1*time.Second))
                if CountFile(outPut) >= num{
                        break
                }
        }
}

func ExecTraceDefault(inPut string, outPut string, iat int){
	trace, num := GetTrace(inPut, iat)
	wg := sync.WaitGroup{}
	wg.Add(1)
        for i:= 0; i < len(trace); i++{
                a := trace[i]
		fmt.Println("loop",i,len(trace))
		go syscall.Exec("/bin/ls", []string{"-al"}, os.Environ())
		go syscall.Exec("/bin/ls", []string{"-al"}, os.Environ())
		Exec(0, a.Para, outPut)
                //err := syscall.NewExec("./fib.py", 0, []string{"./fib.py",strconv.Itoa(a.Para),strconv.Itoa(a.Pred), outPut}, os.Environ())
                if i < len(trace) - 1{
			fmt.Println(trace[i+1].Start - trace[i].Start)
                        time.Sleep(time.Duration(trace[i+1].Start - trace[i].Start)*time.Millisecond)
		}
        }
	fmt.Println("loop")
	go func(){
		for{
			time.Sleep(time.Duration(1*time.Second))
			fmt.Println(CountFile(outPut), num)
			if CountFile(outPut) >= 1000{
				fmt.Println(CountFile(outPut), num)
				wg.Done()
				break
			}
		}
	}()
	wg.Wait()
	fmt.Println("this is the end")
}

func Exec(pred int, para int, out string){
	fmt.Println("Here")
	go syscall.NewExec("./fib.py", pred, []string{"./fib.py",strconv.Itoa(para),strconv.Itoa(pred), out}, os.Environ())
}

func CountFile(fileName string) int {
	f, err := os.Open(fileName)
	if err!= nil{
		fmt.Println(0,"count")
		return 0
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	fmt.Println(len(result), "count")
	return len(result)
}
