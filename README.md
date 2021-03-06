## Solutions to [Codewars](https://www.codewars.com/) problems :computer:

![build](https://travis-ci.org/edsoncelio/codewars-trainning.svg?branch=master)

## Problems

### [**7kyu**] [Is this a triangle?](https://www.codewars.com/kata/56606694ec01347ce800001b/train/go)

Implement a method that accepts 3 integer values a, b, c. The method should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

- [x] [Golang](solutions/triangle.go)
- [ ] Python 

Solved using the [triangle inequality teorem](http://www.mathwarehouse.com/geometry/triangles/triangle-inequality-theorem-rule-explained.php)

---

### [**8kyu**] [Remove First and Last Character](https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go)

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, the original string. You don't have to worry with strings with less than two characters.

- [x] [Golang](solutions/first_and_last.go)
- [ ] Python 
---

### [**8kyu**] [Even or Odd](https://www.codewars.com/kata/even-or-odd/train/go)

Create a function (or write a script in Shell) that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

- [x] [Golang](solutions/even_odd.go)
- [ ] Python 
---

### [8kyu] [Sum of positive](https://www.codewars.com/kata/sum-of-positive/train/go)

You get an array of numbers, return the sum of all of the positives ones.

`Example [1,-4,7,12] => 1 + 7 + 12 = 20`

Note: if there is nothing to sum, the sum is default to `0`.

- [x] [Golang](solutions/sum_positive.go)
- [ ] Python 
---

### [7kyu] [Are the numers in order?](https://www.codewars.com/kata/are-the-numbers-in-order/train/go)

In this Kata, your function receives an array of integers as input. Your task is to determine whether the numbers are in ascending order. An array is said to be in ascending order if there are no two adjacent integers where the left integer exceeds the right integer in value.

For the purposes of this Kata, you may assume that all inputs are valid, i.e. non-empty arrays containing only integers.

Note that an array of 1 integer is automatically considered to be sorted in ascending order since all (zero) adjacent pairs of integers satisfy the condition that the left integer does not exceed the right integer in value. An empty list is considered a degenerate case and therefore will not be tested in this Kata - feel free to raise an Issue if you see such a list being tested.
- [ ] Golang 
- [ ] Python

--- 

### [8 kyu] [Return Negative](https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go)


In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

Example:

```
MakeNegative(1)    # return -1
MakeNegative(-5)   # return -5
MakeNegative(0)    # return 0
```

Notes:

* The number can be negative already, in which case no change is required.
* Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.

- [x] [Golang](solutions/negative.go)
- [ ] Python 

---

### [7kyu] [Binary Addition](https://www.codewars.com/kata/551f37452ff852b7bd000139/train/python)
Implement a function that adds two numbers together and returns their sum in binary. The conversion can be done before, or after the addition.

The binary number returned should be a string.

- [ ] Golang
- [x] [Python]()

---

### [7kyu] [Shortest String](https://www.codewars.com/kata/shortest-word/train/python)
Simple, given a string of words, return the length of the shortest word(s).

String will never be empty and you do not need to account for different data types.

- [ ] Golang
- [x] [Python]()

---

### [6kyu][Find the unique number](https://www.codewars.com/kata/find-the-unique-number-1/train/python)
There is an array with some numbers. All numbers are equal except for one. Try to find it!

```
findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
```

It’s guaranteed that array contains more than 3 numbers.

The tests contain some very huge arrays, so think about performance.

- [ ] Golang
- [x] [Python]()

### [7kyu][Reverse words](https://www.codewars.com/kata/5259b20d6021e9e14c0010d4/train/python)
Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

Examples

```
"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"
```

- [ ] Golang
- [x] [Python]()
