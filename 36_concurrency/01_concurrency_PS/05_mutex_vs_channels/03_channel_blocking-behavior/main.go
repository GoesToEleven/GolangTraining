package main
 
import (  
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	
	mutex := make(chan bool, 1)
	
	
	for i:=1; i < 10;i++ { 
		for j:=1; j<10;j++ {
			mutex <- true		// puts bool on channel
			go func() { 
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex 		// takes bool off channel
			}()
		}
	}
	fmt.Scanln()
}

/*
the channel acts like a mutex in this example
it blocks the code from executing
a bool is put on the channel
if program execution gets back to wanting to put another bool on the channel
the program will pause there until the bool that is already on the chennel is taken off

NOT GOOD PRODUCTION CODE
just interesting to see that channels can behave like mutexes
*/