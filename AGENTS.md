# Project Overview

This repo contains a series of homework assignments for the course Go with GenAI.

# GitHub

github.com/krekoten/SoftServeGoAI/

# Tech Stack

 * Go 1.27.1

# CLI Commands

 * Each lesson folder is a separate Go module. Run all commands from within the specific module folder, for example
   `lesson-01/greeting_ai`
 * Run app `go run .`
 * Run full test suite `go test`. Use `-v` flag to get detailed output.
 * Run focused test `go test -run <TestName>`

# Code Style

 * Use idiomatic code style
 * Run `go fmt` after all changes

# Testing

 * Tests should be in the same folder and follow standard scheme `<file>_test.go` where `<file>` is the name of the file with
   the code for which tests are written or executed
 * Always run tests if they exist after making changes and running `go fmt`
 * Use focused tests first and always do final verification with the full test suite

# Git

 * Do not commit, push or do any destructive or irreversible actions
