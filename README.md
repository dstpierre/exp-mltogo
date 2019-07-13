# exp-mltogo
An experimentation of an Elm-ish syntax that transpile to Go code.

### Status

*It's not working at the moment*.  It's just a starting point.

### An example

```elm
package main

import log exposing {Println}

main = Println "first flang program"
```

Transpile to

```go
package main

import log
func main() {
	log.Println"first flang program"
}
```
