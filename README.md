# fromcharcode #
A minimalist command line utility to parse fromCharCode() to stdout.  

Times out with error if no input is received within 3 seconds.

# Usage #
## With arguments ##
`fromcharcode 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100` 
```
Hello World
```
## With pipes ##
`echo 102, 111, 111, 32, 98, 97, 114|fromcharcode`
```
foo bar
```