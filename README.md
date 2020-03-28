# Push_Swap

##About

This project is about efficiently implementing a custom sorting algorithm
I picked Quicksort, it quite suits my needs.

## Rationale


## Optimization Paths

Use a temporary Rotate stack to further avoid, if not completely, prepend slices operations

## TODO

complete helper message in 'checker'
rename SortStacks --> stacks.Sorter
extract SubStack type to dedicated package and put it in the stack package --> rename stacks.Splitter and create interface
rename rotateSorted -> pushASorted
rename pushSorted -> pushBSorted

write a function for lists smaller than 5

## Notes

I'm starting to learn Go so I tried to be as idiomatic as possible and to apply best practices when possible.
Do not hesitate to open an issue to point out anything I could do a more Go-ish way.

It would be more appropriate to use the word 'height' while using stacks, i've kept 'len' for brevity in the code

## References

https://en.wikipedia.org/wiki/Tower_of_Hanoi
