## Leap Year

Write a function that returns true or false depending on whether its input integer is a leap year or not.

The following rules determine if an integer is a leap year:

- It is NOT a leap year if it is NOT divisible by 4.
- It is a leap year if it is divisible by 4.
- It is a leap year if it is divisible by 400.
- It is NOT a leap year if it is divisible by 100 but NOT by 400.

For example:

- 1997 is not a leap year (not divisible by 4)
- 1996 is a leap year (divisible by 4)
- 1600 is a leap year (divisible by 400)
- 1800 is NOT a leap year (divisible by 4, divisible by 100, NOT divisible by 400)

---

## Run the Application


To run the tests run `go test`

To run the application just run `go run ./leapyear -year [year]`

```
> go run ./leapyear.go -year 12
> true
```

```
> go run ./leapyear.go -year 11
> false
```

