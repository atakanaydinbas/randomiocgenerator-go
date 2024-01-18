# RandomIOCGenerator
RandomIOCGenerator is a simple library for generating random IOC's. 

## Usage

To use RandomIOCGenerator in your Go projects, import it directly from GitHub:
- ``` go get github.com/atakanaydinbas/randomiocgenerator-go```

```go
package main

import (
    "fmt"
    generator "github.com/atakanaydinbas/randomiocgenerator-go"
)

func main() {
    //It creates random IPv4
    randomIP := generator.GenerateRandomIPv4()
    //It created random domain
    randomDomain := generator.GenerateRandomDomain()
}
```