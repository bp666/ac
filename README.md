# Automation Control With Windows



## Examples

```go
package main

import "github.com/bp666/ac"

func main() {
    ac.Click(100, 100)

    ac.Sleep(2)
    
    ac.SendKey("hello")
    
    ac.SleepRandom(0, 3)
    
    ac.Copy("word")
    ac.HotKey("ctrl", "v")
}
```





## APIs

```go
// mouse
GetMousePos() Point
Click(x, y int32)
DClick(x, y int32)
RClick(x, y int32)
MClick(x, y int32)
Scroll(value int32)
HScroll(value int32)

// keyboard
Key(key string)
HotKey(keys ...string)

// clipboard
Copy(text string)

// sleep
Sleep(sec uint16)
SleepRand(left, right uint16)
```

