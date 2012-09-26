# Dice
## Simple dice rolling for Go.

`dice` offers two basic methods for rolling dice:
```go
dice.Roll()
```
and
```go
dice.RollP()
```    
The first function receives a single `string` parameter which describes the dice roll.  For example, send `3d6` to roll three 6-sided dice, `1d20+1` to roll one 20-sided die and add 1 to the result, `1.5d6-1` to roll one and a half 6-sided rolls and subtract one from the result.  Half-dice are taken to be a single die of half the normal faces.  That is, `1.5d6-1` rolls one 6-sided die, one 3-sided die, and subtracts 1 from the sum of both dice.

The second function receives all the parameters of the dice roll separately (think of the `P` as "parameters", "pieces", or "parts").  The parameters are: `number` (the number of dice to roll), `faces` (how many sides, or faces, each die has), `adder` (any static quantity to add to the roll's total), and `half` (a boolean value indicating whether a half-die should be rolled in addition to the dice specified in the `number` parameter).  The equivalent to `dice.Roll("1.5d6-1")` is `dice.RollP(1,6,-1,true)`.

Both functions return a `DiceRoll` object:
```go
type DiceRoll struct {
    NumberOfDice int
    DieFaces     int
    Adder        int
    Half         bool
    Rolls        []int
    RawTotal     int
    Total        int
}
```    
`DiceRoll` objects have a method `Description()` which returns a description like the one used to roll with `dice.Roll()`.

Additionally, a `throwdice` runnable is included, which can take any number of dice roll descriptions as command-line arguments and prints the result for each. It will also read from stdin if no command-line arguments are given, stopping on EOF or "exit" or "quit" (case insensitive).

(Yes, it's all very simple and not extremely useful... it's mostly an exercise on Go on my part.)
