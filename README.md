## codewars-trainning
Solutions to codewars problems

---

Template:
[nível] [pergunta] <br>
[instruções] <br>
[link] <br>
[solução]

---

[**7kyu**] [Is this a triangle?](https://www.codewars.com/kata/56606694ec01347ce800001b/train/go)

Implement a method that accepts 3 integer values a, b, c. The method should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

**Solução:**
```
package kata

func IsTriangle(a, b, c int) bool {
  
   for a > 0 && b> 0 || c > 0{ 
      if a + b > c && b + c > a && a + c > b{ 
        return true
      }else{ return false}
    }
    return false
}
```
Using [triangle inequality teorem](http://www.mathwarehouse.com/geometry/triangles/triangle-inequality-theorem-rule-explained.php)

---
