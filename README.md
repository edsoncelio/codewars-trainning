## Solutions to codewars problems

![badge](https://www.codewars.com/users/edsoncelio/badges/large)

---

## Template:

```
[ranking] <question> 
[instructions]
[solution]
```

## Questions

#### [**7kyu**] [Is this a triangle?](https://www.codewars.com/kata/56606694ec01347ce800001b/train/go)

Implement a method that accepts 3 integer values a, b, c. The method should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

**Solution:**

[triangle.go](solutions/triangle.go)

Using [triangle inequality teorem](http://www.mathwarehouse.com/geometry/triangles/triangle-inequality-theorem-rule-explained.php)

---

#### [**8kyu**] [Remove First and Last Character](https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go)

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, the original string. You don't have to worry with strings with less than two characters.

**Solution:**
```
package kata

func RemoveChar(word string) string {
  return word[1:(len(word)-1)]
}
```

---

#### [**8kyu**][Even or Odd](https://www.codewars.com/kata/even-or-odd/train/go)

Create a function (or write a script in Shell) that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
**Solution:**
```
package kata

func EvenOrOdd(number int) string {
  if number % 2 == 0 {
    return "Even"
  }else{ return "Odd"}
}
```

---

#### [8kyu] [Sum of positive](https://www.codewars.com/kata/sum-of-positive/train/go)

You get an array of numbers, return the sum of all of the positives ones.

`Example [1,-4,7,12] => 1 + 7 + 12 = 20`

Note: if there is nothing to sum, the sum is default to `0`.

**Solution:**

```
package kata

func PositiveSum(numbers []int) int {
   var sum int = 0
   for _, value := range numbers{
     if value > 0 { 
       sum =  sum + value
     }
   } 
   return sum
}
```

---

#### [7kyu] [Are the numers in order?](https://www.codewars.com/kata/are-the-numbers-in-order/train/go)

In this Kata, your function receives an array of integers as input. Your task is to determine whether the numbers are in ascending order. An array is said to be in ascending order if there are no two adjacent integers where the left integer exceeds the right integer in value.

For the purposes of this Kata, you may assume that all inputs are valid, i.e. non-empty arrays containing only integers.

Note that an array of 1 integer is automatically considered to be sorted in ascending order since all (zero) adjacent pairs of integers satisfy the condition that the left integer does not exceed the right integer in value. An empty list is considered a degenerate case and therefore will not be tested in this Kata - feel free to raise an Issue if you see such a list being tested.

**Solution:**

--- 



