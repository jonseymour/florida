# Discussion - Mary, Jane, K problem

If you haven't done the problem yet, you might like to [try](MARY-JANE-K.md) it first.

---
## Answers

The discussion below assumes that the answers to Part A, Part B and Part C are.

 Part A: Q1/A) 1/2, Q2/C) 1/2, Q3/D) 1/2

 Part B: Q1/B) 1/3, Q2/A) 1/3, Q3/C) 1/3

 Part C: Q1/B) 1/3, Q2/A) 1/3, Q3/C) 1/3

The fact that A and B yield different answers is explained by the fact that the selection process is operating on daughters of 2-child families in part A and mothers of 2 child families in Part B. So, at this point there is no paradox or contradiction there - we get different results if we view the world through different goggles.

Part C is simply a repeat of Part B used to setup Part D but uses pseudonymous labels to simplify reasoning about this case. Not surprisingly the answers for Part B and Part C are the same.

## Discussion

Where it becomes interesting is Part D where we assert a correspondence between
the labels used in Part C and the terms of Part A.

According to Part A, K has a 1/2 chance of being a girl, according to Part C/D, K has a 1/3 chance of being a girl - apparently a contradiction.

### What does this apparent contradiction mean?

One way to deal with the contradiction is to deny it is really a contradiction as we did when we explained away the difference between Part A and Part B above - we were sampling populations, so the fact that we get
different answers is not surprising.

However, Part D is still troubling, since we have sampled two different populations that have yielded two different probabilistic models of the same concrete family. One asks us to believe the family will behave one way with probability 1/2, the other asks us to beieve they will behave that same way with probability of only 1/3. What now, are we meant believe about the probability that the family will behave that way? Do we believe the answer provided
by Part A, the answer provided by Part C or something else? If something else, then what?

It should be noted that if Part C is accurate about the "selection at random" that produces Mc, Jc and Kc, it would only produce an Mc, Jc and Kc that matches Mary, Jane and K very infrequently - with a probability proportional to the inverse of the population size. One interpretation of this observation is that the perverse implication of Part D occurs so infrequently we can ignore it. Maybe, but that feels like a copout.

## The Empirical Answer

One way to answer the question empirically is to run a simulation with the "florida" program which has been updated
with support for the kind of matching implied by Part D of the problem. In particular, it generates a large number of 2 person families and creates large slices of both families and daughters. It then samples each a large number of times and outputs whenever there is a match between a selection from the family pool and a selection from the daughter pool. 

The matching process can use two criteria - daughter matching and family matching. When daughter matching is used, then if a selected family has 2 girls, one of the two daughters is selected with a coin flip and that daughter is matched with one selected from the daughter pool, otherwise the family's only daughter is used for that purpose. When family matching is used, we match if the family selected from the family pool is the same as the family of the daughter selected from the daughter pool.

It turns out that the empirical answer to Part D is 1/3 when daughter matching is used and 1/2 when family matching is used. The relative size of these answers is not too surprising since it is harder to match on daughters than on families, but it is striking when one considers that Part A, which is concerned with daughters, yields answers of 1/2, but Part B which is concerned with mothers(/families), yields answers of 1/3. 