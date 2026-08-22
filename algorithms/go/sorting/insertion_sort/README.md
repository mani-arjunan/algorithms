# Insertion sort

Insertion sort is more of a sorting a deck of cards from hand, assume u have `n` number of cards and then u took all in your hands and 
starts from the second card, you pick up and compare it to the card on its left, if it is out of order, shift that left card one spot right and
keep moving leftward, compare and shift, repeat for each next card, moving left to right through the hand.

## EG
[5, 2, 4]
Consider arranging this in ascending order

- Pick up 2 first, compare to 5, now 5 > 2 so do a swap [2, 5, 4]
- Now pick 4, compare to 5, now 5 > 4 so do a swap [2, 4, 5]

## Time complexity
