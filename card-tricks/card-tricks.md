# Card Tricks

## Instructions

As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.


When practicing with her cards, Elyse likes to start with her favorite three cards of the deck: 2, 6 and 9. Write a function FavoriteCards that returns a slice with those cards in that order.

### Task 1: Create a Slice with certain cards.
```
cards := FavoriteCards()
fmt.Println(cards)
// Output: [2 6 9]
```
### Task 2: Retrieve a card from a stack

Return the card at position index from the given stack.
```
card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
```
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return -1:
```
card := GetItem([]int{1, 2, 4, 1}, 10) 
```

#### Note
```
By convention in Go, an error is returned instead of returning an "out-of-band" value. Here the "out-of-band" value is -1 when a positive integer is expected. When returning an error, it's considered idiomatic to return the zero value with the error. Returning an error with the proper return value will be covered in a future exercise.
```
### Task 3: Exchange a card in the stack

Exchange the card at position index with the new card provided and return the adjusted stack. Note that this will modify the input slice which is the expected behavior.
```
index := 2
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards)
// Output: [1 2 6 1]
```
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:
```
index := -1
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards)
// Output: [1 2 4 1 6]
```
### Task 4: Add cards to the top of the stack

Add the card(s) specified in the value parameter at the top of the stack.
```
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice, 5, 1)
fmt.Println(cards)
// Output: [5 1 3 2 6 4 8]
```
If no argument is given for the value parameter, then the result equals the original slice.
```
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice)
fmt.Println(cards)
// Output: [3 2 6 4 8]
```
### Task 5: Remove a card from the stack

Remove the card at position index from the stack and return the stack. Note that this may modify the input slice which is ok.
```
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
fmt.Println(cards)
// Output: [3 2 4 8]
```
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:
```
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
fmt.Println(cards)
// Output: [3 2 6 4 8]
```