# NAME
florida - a simulation of the "a girl named florida" problem

# DESCRIPTION

This go-lang program simulates 2 variations of the "A girl named Florida" problem.

# Problem #1
A) You are told of a 2-child family that has at least one girl. What is the chance the family consists of 2 girls?

B) You are then told that one of the siblings is a girl whose name is Florida. What is the chance Florida's sibling is also a girl?

The default simulation reproduces this case and we find that the answer to A) is ~= 33% and the answer to B) is ~= 50%

	$ ./florida 
	n=1000000, c=749968, gg=249737 (0.332997), florida=1003 (0.001337), florida && gg=525 (0.523430)

You can vary the number of families generated with --families and the probability of a girl being named Florida with --probability. 

# Problem #2
A) You are told of a girl who has exactly one sibling. What is the chance the girl's sibling is also a girl?

B) You are then told that the girl's name is Florida. What is the chance that Florida's sibling is also a girl?

The answer to part A is given by a simulation that selects girls (`--girls`) - it allows each girl in a 2 girl family to be counted separately. The answer here is = 0.499523 ~=50%.  Note that the `florida && gg=1040 (0.673139)`
is not the answer to part B, rather it is the answer to a different question such as: 'You are then told either the girl or the girl's sibling's name is Florida. What is the chance that Florida's sibling is also a girl?'.

	$ ./florida --girls
	n=1000000, c=999790, gg=499418 (0.499523), florida=1545 (0.001545), florida && gg=1040 (0.673139)

The answer to part B is given by augmenting the `--girls` filter with a filter that only selects girls named Florida (`--florida`), yielding a result of 0.526580 ~= 50%.

	$ ./florida --girls --florida
	n=1000000, c=997, gg=525 (0.526580), florida=997 (1.000000), florida && gg=525 (0.526580)	