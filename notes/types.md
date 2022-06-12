# Integers

## Sized integer types

   1. {int,uint}: the most efficient size for signed or unsigned integers on a particular platform
   2. {int,uint}{8,16,32,64}: 4 distinct sizes of integers

## Other types

   1. `rune`: a synonym for `int32`, but conventionally indicates a 
Unicode value.
   2. `byte`: a synonym for `uint8`, with emphasis that the value is raw data
instead of a small numeric quantity.
   3. `uintptr`: an unsigned integer type for pointers.