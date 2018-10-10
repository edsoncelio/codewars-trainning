package kata

func IsTriangle(a, b, c int) bool {
  
   for a > 0 && b> 0 || c > 0{ 
      if a + b > c && b + c > a && a + c > b{ 
        return true
      }else{ return false}
    }
    return false
}
