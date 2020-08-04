package main

import "fmt"
import "time"


// this is a comment

func main() {
    

    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2020, 01, 01, 00, 00, 00, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Weekday())

    var kalender[4][7]int
    fmt.Println(kalender)

    for i := 0; i<4; i++ {
        for j := 0; j<7; j++ {
        kalender[i][j]=(1+j)+(7*i);
        fmt.Print(kalender[i][j], " ")
        }
        fmt.Println("")
        }



    for i := then.Month(); i <= 12; i++ {
        fmt.Println("====================", i, "====================")
        fmt.Println("|Sun   |Mon   |Tue   |Wed   |Thur   |Fri   |Sat  |")
    
        for i := 0; i<4; i++ {
            for j := 0; j<7; j++ {
            kalender[i][j]=(1+j)+(7*i);
            fmt.Print(kalender[i][j], "\t")
            }
            fmt.Println("")
            }
    }

    
}