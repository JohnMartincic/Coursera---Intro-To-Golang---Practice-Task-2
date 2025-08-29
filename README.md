# Coursera - Golang Project 2 - Enhanced Fare Calculator

This repository contains the second project from the Golang for Beginners series on Coursera. This project builds upon the original travel fare calculator, adding key features to improve the user experience and introduce new concepts from the Go standard library.

## 🎯 Project Goal

The goal of this project was to enhance the existing command-line fare calculator by making it more user-friendly and robust. The application still calculates travel fares but now includes improved input handling and provides users with more information upfront.

## ✨ Key Features & Enhancements

This version of the application introduces two main improvements:

* **Dynamic Cabin Class Display**: The program now dynamically lists the available cabin class choices for the user. This is accomplished by iterating through the predefined list of cabin classes stored in the application's data structure.

* **Case-Insensitive Input**: User input for origin, destination, and cabin class is now case-insensitive. This was achieved by utilizing the Go standard `strings` package to convert all user input to uppercase, ensuring it matches the format of the stored data.

## 💡 Key Concepts Learned & Applied

This project reinforced previous concepts and introduced new ones, particularly from Go's standard library:

* **Using the `strings` Package**: Gained experience with standard library packages by using `strings` to manipulate and normalize user input.
* **Looping for Display**: Applied `for` loops with the `range` keyword to iterate over data structures and display information to the user, making the program more interactive.
* **Building on Existing Code**: Practiced reading and extending an existing Go codebase to add new functionality without breaking previous features.

## 🔗 Project Lineage

This project is a direct continuation and enhancement of the first travel fare calculator. The foundational code from that project was used as the starting point for this one.

* **[Original Travel Fare Calculator Project](https://github.com/JohnMartincic/Coursera---Intro-To-Golang---Fare-menu)**

## 🎓 Course Context

This project is the second assignment for the Coursera course linked below, designed to introduce standard library packages and build upon foundational Go programming concepts.

* **Course Link:** [Golang for Beginners: Data Types, Functions, and Packages](https://www.coursera.org/projects/golang-beginners-data-types-functions-packages)
