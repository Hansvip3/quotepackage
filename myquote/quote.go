package myquote

import  (
"fmt"
"rsc.io/quote"
)
func PrintHello() {
    fmt.Println("Hello, Hello")
}
func PrintQuote() {
    fmt.Println(quote.Go())
    fmt.Println(quote.Glass())
    fmt.Println(quote.Opt())
    fmt.Println(quote.Hello())
}