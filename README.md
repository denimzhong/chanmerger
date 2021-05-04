# chanmerger

Merger is a tool to automate the creation of method that merges multiple channels into one.

For example, given the type chan int, run this command

```shell
chanmerger -type='chan int'
```

in the same diretory will create a file named chan_int_merge.go, in same package, containing a definition of 

```go
func mergeIntChan(in ...chan int) chan int
```

That method will receive all values from input channels, and send them to a new ouput channel, which is called fanin.
