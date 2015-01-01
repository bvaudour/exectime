exectime
========

exectime is a simple toolkit intended to measure the time of execution of a command.

**Installation**

1. Clone the repository:
   `git clone https://github.com/bvaudour/exectime.git`

2. Set the GOPATH environment variable:
   `export GOPATH=$GOPATH: <your_folder>/exectime`

3. Compile the sources: 
   `go build -o <path_of_the_binary>/exectime <your_folder>/exectime/src/exectime/exectime.go`

4. Enjoy!

**Usage**

    exectime <command> [args_of_<command>]


**Use as library**

You can use exectime to benchmark go codes, too.

***Example:***

    package main
    
    import (
    	"fmt"
    	"exectime/timer"
    )
    
    func f1(arg string) {
    	fmt.Println("Hello", arg)
    }
    
    func f2() {
    	fmt.Println("Hello World")
    }
    
    func main() {
    	// Simple way
    	t := timer.New()
    	t.Start()
    	f1("toto")
    	t.End()
    	fmt.Println("Time of f1 execution:", t)
    	fmt.Println("Time in seconds:", t.Seconds())
    
    	// Timer with function execution
    	tf := timer.NewFunc(f2)
    	tf.Exec() // Equivalent to: t := timer.New(); t.Start(); f2(); t.End()
    	fmt.Println("Time of f2 execution:", tf)
    	fmt.Println("Time in µs:", ts.Microseconds())
    }
