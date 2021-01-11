# errors package

simple structured error handling, tracing, and composition

```go
import(
    "fmt"
    "github.com/JKhawaja/errors"
)

func outerFunc(arg1 int) error {
    err := innerfunc(arg1)
    if err != nil {
        // wrap an error
        return errors.New(err, nil)
    }
    
    return nil
}

func innerFunc(arg1 int) error {
    if arg1 < 0 {
        // define a simple error, with arbitrary variable values in scope
        return errors.New(fmt.Errorf("negative argument"), map[string]interface{}{
            "arg1": arg1,
        })
    }
    return nil
}

func main() {
    err := outerFunc(-1)
    if err != nil {
        // log the trace of the error
        fmt.Println(errors.NewTrace(err).Error())
    }
}
```
