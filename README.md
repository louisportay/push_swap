# Push_Swap

## About

Given an **A** stack of *integers*, sort it using only a limited set of instructions with another empty stack **B**
You must generate the minimum possible number of instructions

The instruction set is as follows:
* Swap A: *sa*
	exchanges the two first integers at the top of stack A
* Swap B: *sb*
	exchanges the two first integers at the top of stack A
* Push A: *pa*
	pop the first integer of B and push it on top of A
* Push B: *pb*
	pop the first integer of A and push it on top of B
* Rotate A: *ra*
	move the first integer of A on the bottom of it
* Rotate B: *rb*
	move the first integer of B on the bottom of it
* Reverse Rotate A: *rra*
	move the last integer of A to the top of it
* Reverse Rotate B: *rrb*
	move the last integer of A to the top of it

Addtional instructions
* Swap Both: *ss*
	Swap A + Swap B
* Rotate Both: *rr*
	Rotate A + Rotate B
* Reverse Rotate Both: *rrr*
	Reverse Rotate A + Reverse Rotate B

More in-depth explanations in the PDF file located at the repository root

This project is about efficiently implementing a custom sorting algorithm
I chose Quicksort, it quite suits my needs.

## Usage

*solver* takes a list of integers and generate the instructions to sort them
*checker* takes a list of integers and reads a list of instructionsthen computes them and output whether the resulting list is sorted or not

ARGS="5 3 4 2 1"
solver $ARGS
[...]
checker $ARGS
[...]
OK

ARGS=`ruby -e "puts (-1000..1000).to_a.shuffle.join(' ')"`; solver $ARGS | checker $ARGS

## Rationale

Starting with the whole stack **A**, I will adopt different strategies regarding few criterions:
-Is is already sorted ? Or at least partly ?
-How big it is ?

In my implementation, the **Sorted** stack is located at the bottom of **A**, a simple *ra* will implicitly add any element on top of **A** to the **Sorted** stack 

I split my **A** stack with a *Pivot*, I define it from the bottom of **A** so that there are elements superior and inferior to it
Every element inferior to *Pivot* are pushed to **B**, the others are rotated to **A** bottom
Once **A** is depilated, I restore the previously rotated superior elements.
Now **A** contains only elements that are bigger than Pivot and **B** elements that are inferior
I will repeat the process on **B** by keeping track of the heights of different *Substacks*
Once I got a small enough *Substack* (<= 4), I add the elements to *sorted* list
Every substacks will be recursively depilated resulting in empty **A** and **B** lists

## Optimization Paths

Generate a *descendingly* sorted list (highest value first) at the beginning of the program, stored just above the *ascending* one (*sorted* list)
When the *sorted* list is full just *reverse rotate* it or *rotate* the *descendingly* sorted list
try a long list with only the two first elements swapped to get a feel for its necessity
This is a special case optimization

Write an algorithm for short lists (without using a *sorted* list)


## Notes

I'm starting to learn Go so I tried to be as idiomatic as possible and to apply best practices when possible.  
Do not hesitate to open an issue to point out anything I could do a more Go-ish way.  

It would be more appropriate to use the word 'height' while using stacks, i've kept 'len' for brevity in the code

## References

https://en.wikipedia.org/wiki/Tower_of_Hanoi
