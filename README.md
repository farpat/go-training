# Go Training Exercises

This repository contains a collection of Go programming exercises designed to help you practice and improve your Go programming skills.

## Exercises

1. **CountBits**: Count the number of bits set to 1 in a binary number.
2. **Divisors**: Find all divisors of a given number.
3. **MinMax**: Find the minimum and maximum values in a slice of integers.
4. **Prime**: Generate all prime numbers up to a given number using the Sieve of Eratosthenes algorithm.
5. **Reverse**: Reverse the order of words in a string.

## Project Structure

Each exercise is contained in its own directory under `exercises/` and includes:
- `main.go`: Implementation of the exercise
- `test.go`: Unit tests for the exercise

## Creating a New Exercise

To create a new exercise, use the following make command:

```bash
make exercise NAME=example
```

This will create a new directory `exercises/example` with basic template files.

## Running Tests

To run tests for an exercise:

```bash
cd exercises/exercise_name
go test
```

## Contributing

Feel free to add new exercises or improve existing ones by creating pull requests.

## License

This project is licensed under the MIT License.
