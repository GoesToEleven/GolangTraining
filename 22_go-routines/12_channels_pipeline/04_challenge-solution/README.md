A student sent me the below note so I included the student's excellent solution to this project:

Todd, I reviewed your solution to 22_go-routines/04_challenge-solution in your Golang Programming class.

I believe your solution might not run 100 factorial computations concurrently and in parallel as the following statement will calculate the factorials sequentially, since it is receiving them - one by one -  from the pipeline:
out <- fact(n)

I used the sync package to create a solution that I believe will accomplish your goal. Let me know if I am missing something:

https://github.com/arsanjea/udemyTraining/blob/master/Exercises/exercise38.go

/////

# Definitions

## concurrency
a design pattern
go uses goroutines to create concurrent design patterns

## parallelism
running code from a program on more than one cpu
parallelism implies you have used concurrent design patterns

## sequentially
one thing happening after another, in sequence

# here are my thoughts

The original solution uses concurrent design patterns. It uses two different goroutines. The [control flow](https://en.wikipedia.org/wiki/Control_flow) of the program is no longer a straight top-to-bottom sequence. Different goroutines have been launched.

The program may or may not run in parallel. If the program is run on a machine with two or more cpus, the program has the potential to run in parallel. Each of our three goroutines (the two goroutines we launched, and main) could be running on different cpu cores.

Jean-Marc Arsan is correct: the program IS still running sequentially. Even though calculations are occuring in different goroutines, and potentially on different CPU cores, the sequence in which they occur is still sequential. 

goroutines allow synchronization of code.

In this original example, the code is synchronized.

Thank you, Jean-Marc, for your comment on this code!

I appreciate these discussions and hope the notes and your new code sample are helpful to everyone!

