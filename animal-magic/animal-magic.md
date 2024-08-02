# Animal Magic

## Introduction

rescource:
[Randomeness](../randomness/randomness.md)

## Instructions

Elaine is working on a new children's game that features animals and magic wands. It is time to code functions for rolling a die, generating random wand energy and shuffling a slice.

### Task 1: Roll a die.

mplement a `RollADie` function.

This will be the traditional twenty-sided die with numbers 1 to 20.
```
d := RollADie() // d will be assigned a r
```

### Task 2: Generate wand energy.
Implement a GenerateWandEnergy function. The wand energy should be a random floating point number equal or greater than 0.0 and less than 12.0.
```
f := GenerateWandEnergy()  // f will be assigned a random float64, 0.0 <= f < 12.0

```

### Task 3: Shuffle a slice.

The game features eight different animals:

    ant
    beaver
    cat
    dog
    elephant
    fox
    giraffe
    hedgehog

Write a function ShuffleAnimals that returns a slice with the eight animals in random order.