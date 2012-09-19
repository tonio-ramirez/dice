/*
Throwdice rolls virtual dice and prints the result.

It reads dice-descriptions from command-line arguments and rolls
each of them, printing the results to stdout.

For example:
    throwdice 3d6+1
could output:
    14 (3, 6, 4)
Where 14 is the total, and 3, 6, and 4 are the individual dice rolls.

When invoked with no command-line arguments, it reads from
standart input instead, one description per line.  In this case,
it will stop reading when receiving EOF or the words "exit" or
"quit" (case insensitive).
*/
package documentation
