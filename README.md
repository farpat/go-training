This repository contains a collection of Go programming exercises designed to help you practice and improve your Go programming skills.

[![Test](https://github.com/odavid/go-exercises/workflows/Test/badge.svg)](https://github.com/odavid/go-exercises/actions)


# Project Structure
Each exercise is contained in its own directory under `exercises/` and includes:
- `main.go`: Implementation of the exercise
- `test.go`: Unit tests for the exercise

# Creating a New Exercise
To create a new exercise, use the following make command:
```bash
make exercise NAME=exercisename
```
This will create a new directory `exercises/exercisename` with basic template files.

# Running Tests
To run tests for an exercise:
```bash
make test NAME=exercisename
```
