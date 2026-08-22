# Merge sort

Merge sort follows divide and conquer method, where we first divide array to each halves untill it reaches to a single element(base case)
after splitting to two halves left and right, now just merge them in a sorted order merging to sorted array could be easier
since each single half is already sorted.
j
## EG
[2, 4, 5, 7, 1, 3, 6]
Consider arranging this in ascending order

- Separate the left and right first iteration, the result [2, 5, 4] [7, 1, 3, 6]
- Separate the left now into two, the result [2] [5, 4]
- Separate the above right now into two, the result [5] [4]
- Merge the above in sorted ascending order [4, 5]
- Merge the 2nd with above [2] [4, 5], results into [2, 4, 5]
- Seprate from step 2 [7, 1, 3, 6] into two, the result [7, 1] [3, 6]
- Seprate the above [7, 1] into two, the result [7] [1]
- Merge the above in sorted ascending order [1, 7]
- Seprate [3, 6] into two, the result [3] [6]
- Merge the above in sorted ascending order [3, 6]
- Merge the above in sorted ascending order [1, 3, 6, 7]
- Merge the overall [2, 4, 5] and [1, 3, 6, 7] into one, the result [1, 2, 3, 4, 5, 6, 7]

## Time complexity
