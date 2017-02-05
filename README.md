# NAME
florida - a simulation of the "a girl named florida" problem

# DESCRIPTION

This go-lang program simulates 2 variations of the "A girl named Florida" problem.

# Variation#1
The first variation is:

A) You are told that there is a 2-child family and at least one child is a girl. What is the chance the other child is also a girl?

B) You are then told that one of the siblings is a girl whose name is Florida. What is chance her sibling is also a girl?

The default simulation reproduces this case and we find that the answer to A) is ~33% and the answer to B) is ~50%

	$ ./florida 
	n=750022, gg=249761 (0.333005), florida=998 (0.001331), florida && gg=521 (0.522044)

# Variation#2
The second variation is:

A) You are told of a girl who has exactly one sibling, what is the chance that the sibling is also a girl?

B) You then told that one of the two siblings' names is Florida. What is the chance that Florida's sibling is also a girl.

C) You are then told that actually, it is the girl of question A that is called Florida. What is the chance that the sibling is also a girl?

The answer to part A is given by a simulation that selects girls - it allows each girl in a 2 girl family
to be counted separately. The answer is ~50%. 

The answer to part B is also given by this simulation - here it is the value 0.685442 ~= 2/3.

	$ ./florida --girls
	n=999262, gg=498756 (0.499124), florida=1540 (0.001541), florida && gg=1014 (0.658442)

The answer to part C is given by augmenting girls filter with a filter that only selects girls named Florida, yielding a result of 0.520000 ~= 50%.

	$ ./florida --girls --florida
	n=1000, gg=520 (0.520000), florida=1000 (1.000000), florida && gg=520 (0.520000)
