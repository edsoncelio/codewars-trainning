package main

func PositiveSum(numbers []int) int {
   var sum int = 0
   for _, value := range numbers{
     if value > 0 { 
       sum =  sum + value
     }else{
         sum = sum - value
     }
   } 
   return sum
}
