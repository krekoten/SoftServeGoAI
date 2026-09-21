# Lesson 01

Each folder is a separate Go module. Run the commands from inside the folder.

## hello_world

Prints a fixed greeting.

```
cd hello_world
go run .
```

## greeting_manual

Prints a greeting for the name passed as the first argument. Without a name it asks for one.

```
cd greeting_manual
go run . Marian
```

## greeting_ai

Same as `greeting_manual`, but without a name it prints a usage message to stderr and exits with status 1.

```
cd greeting_ai
go run . Marian
```
