This repository contains a collection of Go programming exercises designed to help you practice and improve your Go programming skills.

[![Test](https://github.com/farpat/go-training/workflows/Test/badge.svg)](https://github.com/farpat/go-training/actions)

**Table of Contents**
- [Project Structure](#project-structure)
- [Creating a New Exercise](#creating-a-new-exercise)
- [Running Tests](#running-tests)

# Requirements
- Go +1.24. To install, run on Ubuntu (or similar for another OS): 
```bash
sudo apt update && sudo apt install golang-go
```
- Make. To install, run on Ubuntu (or similar for another OS): 
```bash
sudo apt update && sudo apt install make
```


# Project Structure
Each exercise is contained in its own directory under `exercises/` structured as follows:
```
 exercises
│   ├── exercise_name_1
│   │   ├── main.go // Implementation of the exercise 1
│   │   └── main_test.go // Unit tests for the exercise 1
│   └── exercise_name_2
│   │   ├── main.go // Implementation of the exercise 2
│   │   └── main_test.go // Unit tests for the exercise 2
│   └── etc.
```

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
