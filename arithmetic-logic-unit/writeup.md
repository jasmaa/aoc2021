# Arithmetic Logic Unit Write-Up

Reverse engineering the ALU assembly yields the code for `monad1` and `monad2`
in `monad.py` as well as the hard-coded values below:

```
arr1 = [10, 12, 15, -9, 15, 10, 14, -5, 14, -7, -12, -10, -1, -11]
arr2 = [ 15, 8,  2,  6, 13,  4,  1,  9,  5, 13,   9,   6,  2,   2]
arr3 = [  1, 1,  1, 26,  1,  1,  1, 26,  1, 26,  26,  26, 26,  26]
```

MONAD runs through each digit, c, of the model number, keeping track of a stack.
For each idx of the model number:
- `arr1[idx]` is summed with stack head. If sum is not equal to c, then stack gets pushed to.
- `arr2[idx]` is summed with c. The sum is pushed to stack if a push occurs.
- `arr3[idx]` determines if stack should pop or not.

Valid model #s needed to have an stack empty at the end of execution.

Model #s are 14 digit long. Since `c := [1, 9]`, every positive `arr1[idx] > 9`, and every `arr1[idx] > 0`,
pushing to the stack on a positive `arr1[idx]` is unavoidable. A pop would need to happen on a negative `arr1[idx]`.

There are also only 7 indices at which to pop and 7 positive `arr1[idx]`. This means that we cannot allow a push
to occur on a negative `arr1[idx]`.

Fortunately, the indices where `arr1[idx]` is negative line up with the indices where pops occur which
makes the problem easier to reason about. We know each push and pop will cancel each other out, so we
can pair up each push index with the pop index that will cancel out the paired push.

```
arr1  = [10, 12, 15, -9, 15, 10, 14, -5, 14, -7, -12, -10, -1, -11]
arr2  = [15,  8,  2,  6, 13,  4,  1,  9,  5, 13,   9,   6,  2,   2]
arr3  = [ 1,  1,  1, 26,  1,  1,  1, 26,  1, 26,  26,  26, 26,  26]
pairs = [ 0,  1,  2,  2,  3,  4,  5,  5,  6,  6,  4,    3,  1,   0]
```


## Part 1: Find largest valid model #

For each pair, we want to maximize c when a push occurs while ensuring that the push's pop pair
is able to cancel out the pushed value.

```
Pair 0:
push(a + 15), b = a + 15 - 11 := [1, 9]
a = 5
b = 9

Pair 1:
push(a + 8), b = a + 8 - 1 := [1, 9]
a = 2
b = 9

Pair 2:
push(a + 2), b = a + 2 - 9 := [1, 9]
a = 9
b = 2

Pair 3:
push(a + 13), b = a + 13 - 10 := [1, 9]
a = 6
b = 9

Pair 4:
push(a + 4), b = a + 4 - 12 := [1, 9]
a = 9
b = 1

Pair 5:
push(a + 1), b = a + 1 - 5 := [1, 9]
a = 9
b = 5

Pair 6:
push(a + 5), b = a + 5 - 7 := [1, 9]
a = 9
b = 7

Which forms: 52926995971999
```

## Part 2: Find smallest valid model #

Same procedure as part 1 but now minimizing a.

```
Pair 0:
push(a + 15), b = a + 15 - 11 := [1, 9]
a = 1
b = 5

Pair 1:
push(a + 8), b = a + 8 - 1 := [1, 9]
a = 1
b = 8

Pair 2:
push(a + 2), b = a + 2 - 9 := [1, 9]
a = 8
b = 1

Pair 3:
push(a + 13), b = a + 13 - 10 := [1, 9]
a = 1
b = 4

Pair 4:
push(a + 4), b = a + 4 - 12 := [1, 9]
a = 9
b = 1

Pair 5:
push(a + 1), b = a + 1 - 5 := [1, 9]
a = 5
b = 1

Pair 6:
push(a + 5), b = a + 5 - 7 := [1, 9]
a = 3
b = 1

Which forms: 11811951311485
```