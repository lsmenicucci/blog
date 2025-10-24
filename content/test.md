---
title: My First Post
publication_date: 2025-10-23
tags: [golang, web, stuff, test, aaa, another]
---

Hey this is post!!!

Teste _italic_

Here is some code:

```fortran
program HighlightTest
    implicit none
    integer :: i, n
    real :: x, y
    character(len=20) :: name

    ! Variable initialization
    n = 5
    name = "FortranTest"

    print *, "Program starts:", name

    ! Loop example
    do i = 1, n
        x = real(i) * 1.5
        y = sqrt(x**2 + 2.0)
        print *, "i=", i, "x=", x, "y=", y
    end do

    ! Conditional example
    if (n > 3) then
        print *, "n is greater than 3"
    else
        print *, "n is 3 or less"
    end if

    ! Select case
    select case (n)
    case (1:3)
        print *, "n between 1 and 3"
    case (4:6)
        print *, "n between 4 and 6"
    case default
        print *, "n out of range"
    end select

    ! Subroutine call
    call greet(name)

contains

    subroutine greet(user)
        character(len=*), intent(in) :: user
        print *, "Hello,", user
    end subroutine greet

end program HighlightTest
```

## Golang code!

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum tincidunt elit in lorem posuere, id dictum augue faucibus. Cras cursus tortor nec elementum sodales. Nullam accumsan risus lorem, ac lacinia velit tristique quis. Etiam tempor bibendum eros a auctor. Proin pretium porttitor ipsum. Cras gravida volutpat tellus ut vehicula. Pellentesque non vestibulum elit.

Nam gravida nulla sit amet turpis malesuada feugiat. Nullam nisi purus, maximus in nulla sed, suscipit fringilla lacus. Duis sollicitudin placerat est. Etiam feugiat nisl ligula, at finibus quam molestie a. Aenean fringilla leo in justo porta ultricies. Pellentesque in blandit enim. Vivamus eget pellentesque dui. Vestibulum molestie vehicula sem, ut sollicitudin mi sollicitudin nec.

Duis efficitur mauris quis placerat vehicula. Ut vitae ante ac eros tincidunt interdum. Suspendisse potenti. Pellentesque leo purus, porta a lectus quis, egestas euismod eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Vivamus nunc felis, imperdiet quis molestie ac, egestas ac odio. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Cras efficitur, odio at dictum pretium, purus orci euismod enim, eu rutrum mauris ante a eros. Pellentesque nec luctus velit. Praesent gravida pellentesque eros pretium iaculis. Quisque vel congue nulla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Mauris pretium tellus id diam vestibulum, ut porta elit ullamcorper. Pellentesque in viverra urna, a laoreet eros.

```go 
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	var points []Point
	points = append(points, Point{3, 4}, Point{5, 12})

	for i, pt := range points {
		dist := pt.Distance()
		fmt.Printf("Point %d: (%.1f, %.1f), Distance: %.2f\n", i+1, pt.X, pt.Y, dist)
		if dist > 10 {
			fmt.Println("This point is far from the origin.")
		} else {
			fmt.Println("This point is close to the origin.")
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("Sum of numbers:", sum)

	result := factorial(5)
	fmt.Println("Factorial of 5:", result)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
```

## Python

Oh my god, who let the _snakes_ out?!

```python
#!/usr/bin/env python3

import math
from dataclasses import dataclass
from typing import List

@dataclass
class Point:
    x: float
    y: float

    def distance(self) -> float:
        return math.sqrt(self.x ** 2 + self.y ** 2)

def factorial(n: int) -> int:
    if n <= 1:
        return 1
    return n * factorial(n - 1)

def greet(name: str, times: int = 1) -> None:
    for _ in range(times):
        print(f"Hello, {name}!")

def main():
    points: List[Point] = [Point(3, 4), Point(5, 12)]

    for i, pt in enumerate(points, start=1):
        dist = pt.distance()
        print(f"Point {i}: ({pt.x}, {pt.y}), Distance: {dist:.2f}")
        if dist > 10:
            print("This point is far from the origin.")
        else:
            print("This point is close to the origin.")

    numbers = [1, 2, 3, 4, 5]
    total = sum(numbers)
    print(f"Sum of numbers: {total}")

    print(f"Factorial of 5: {factorial(5)}")
    greet("Alice", 2)

if __name__ == "__main__":
    main()
```
