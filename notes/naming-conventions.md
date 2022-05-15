# Naming Conventions

## Packages

Packages are always in lowercase.  For example,
   * bufio
   * fmt
   * os

## Length

Short names are preferred, though there's no limit on length.

## Camel case, not snake case.

### Acronyms are all UPPERCASE

Acronyms such as "ASCII" or "HTML" don't get converted to "Ascii" or "Html".

### Case of the first letter 

If a variable is declared inside a function, it's local to that function.

If a variable is outside a function, then it's visible to the whole file.  Furthermore,
   1. If the first letter of a variable is UPPERCASE, it's *exported* and visible to functions 
outside its own package.
   2. This is also true to the functions, for example, "Printf"



