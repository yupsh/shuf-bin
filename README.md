# yup-shuf

```
NAME:
   shuf - generate random permutations

USAGE:
   shuf [OPTIONS] [FILE]
      shuf -e [OPTIONS] [ARG...]
      shuf -i LO-HI [OPTIONS]

      Write a random permutation of the input lines to standard output.
      With no FILE, or when FILE is -, read standard input.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --head-count value, -n value   output at most COUNT lines (default: 0)
   --input-range value, -i value  treat each number LO through HI as an input line
   --random-source value          get random bytes from FILE
   --echo, -e                     treat each ARG as an input line (default: false)
   --zero-terminated, -z          line delimiter is NUL, not newline (default: false)
   --repeat, -r                   output lines can be repeated (default: false)
   --help, -h                     show help
```
