# chanmerger

Merger is a tool to automate the creation of method that merges multiple channels into one.

## Installation

Use the below command to install channel merger:

```shell
go install github.com/patrickio/chanmerger
```

## Quick Start

For example, given the type chan int, run this command

```shell
chanmerger -type=int
```

in the same diretory will create a file named chan_int_merge.go, in same package, containing a definition of 

```go
func mergeIntChan(in ...chan int) chan int
```

That method will receive all values from input channels, and send them to a new ouput channel, which is called fanin.

### Import types support

It's safe to create channel merging method of import types, when import pathes should be exists in the same package.

For example, given the channel type `foo.Bar`, run this command to generate method:

```shell
chanmerger -type=foo.Bar
```

Also, we can add go generate command in source code file where need the channel merge method:

```go
//go:generate chanmerger -type=foo.Bar

package hello

import "github.com/patrickio/foo"
```