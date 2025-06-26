package main

import "fmt"

type Student struct {
    Name  string
    Grade int
}

func main() {
    s1 := Student{"John", 85}
    s2 := &s1
    s3 := s2
    
    s2.Grade = 90
    s3.Name = "Jane"
    
    fmt.Printf("s1: %+v\n", s1)
    fmt.Printf("s2: %+v\n", *s2)
    fmt.Printf("s3: %+v\n", *s3)
}

/* Analysis :
   
   s1 is instance from Student object while s2 is a pointer to s1 and s3 equivalent to s2.
   When s2.Grade = 90, it is modifying the value of s1.Grad
   When s3.Name = "Jane", it is modifying the value of s1.Name
   So, s1, s2, and s3 are equivalent and both are pointing to the same memory location.


   when the program try to print out these variable, s1 will print the object content, s2 will print 
   the s1 address, but with asterisk operator, it will be converted into the actual object.
   Same as to s2 because s3 contain s2 where the value of s2 refers to s1 location 

*/