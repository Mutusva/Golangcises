package algorithms

import "fmt"

func bitshift() {
   s := "056"
   d := int(s[1] - 48)
   sum := 0
   fmt.Printf("%v\n", d)
   fmt.Printf("sum = %v\n", sum)
   
   for _,v := range s {
     d = int(v - 48)
     fmt.Printf("d * 2 = %v\n", d << 1)
   }
}
