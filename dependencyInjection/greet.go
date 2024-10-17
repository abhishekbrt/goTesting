package dependencyInjection

import (
	//"bytes"
	"fmt"
	"io"
)

func Greet(w io.Writer,name string)  {
	// w.Write([]byte("Hello "))
	// w.Write([]byte(name))
	// fmt.Println(w)
	fmt.Fprintf(w,"Hello %s",name)

}
