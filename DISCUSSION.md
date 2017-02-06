# ANSWERS

The answers to the problems are: 

* 1A: 33% or 1/3 
* 1B: ~50% or 1/2-e for some e < 1/6, depending on the probability parameter
* 2A: 50% or 1/2
* 2B: 50% or 1/2

# BAYESIAN ANALYSIS

To more completely understand the answer to 1B, it is helpful to do a Bayesian 
analysis of each question.

First, some terms:

    p = P(S)

means the probability of statement S being true is p.

    p = P(S|F)

means, given statement F is true, the probability that statement S is true is p.

Recapitulatng Bayes Theorem:

    P(H|E) = P(E|H).P(H)
             -----------
                P(E)
where:

    H = the statement of the hypothesis
    E = the statement of the evidence
    P(H) = the probability of the hypothesis in absence of the evidence evidence
    P(E) = probability of the evidence, independent of the hypothesis
    P(E|H) = probabity of the evidence, given the hypothesis is true
    P(H|E) = revised probability of the hypothesis, given the evidence.

Now, some definitions:

* **BB** is the statement that the family is a boy-boy family.
* **BG** is the statement that the family is a boy-girl family.
* **GB** is the statement that the family is a girl-boy family.
* **GG** is the statement that the family is a girl-girl family.
* **Gf** is the statement that a girl is named Florida.
* **D**  is the statement that the family has at least one daughter.
* **Df** is the statement that the family has a daughter named Florida.
* **Gs** is the statement that the sibling is a girl.

The Bayesian answer to 1A is:

    P(GG|D) = P(D|GG)*P(GG)
              -------------
                  P(D)

            = 1 * 1/4
              --------
                3/4

            = 1/3

The Bayesian answer to 1B is:

                P(Df|GG)*P(GG)
    P(GG|Df) =  --------------
                    P(Df)

    Let q = P(BG) = P(GB) = P(GG)

    Let p = P(Gf)

    P(Df|GG) = 2p * (1-p) + p^2 = p * (2-p)

    P(Df|BG) = P(Df|GB) = p

    P(Df)    = P(Df|BG) * P(BG) + P(Df|GB) * P(GB) + P(Df|GG) * P(GG)

             = (2p + p * (2-p)) * q 

             = p * (4-p) * q


               p * (2-p) * q
    P(GG|Df) = -------------
               p * (4-p) * q

               2 - p   2 - P(Gf)
             = ----- = ---------
               4 - p   4 - P(Gf)

    If P(Gf) is small, then P(GG|Df) is ~= 1/2.

    If P(Gf) is 1 then P(GG|Df) is 1/3 which comports the answer of with 1A.

The answer to 2A is:

    P(Gs) = 1/2

The answer to 2B is:

    P(Gs) = 1/2

# DISCUSSION
The differences between Problem 1 and Problem 2 are explained by the fact that the first problem counts 2-child families whereas the second problem counts daughters of 2-child families and there are 4/3 as many daughters as their are families (because 2-girl families skew the count).

There is a curious difference between the answers to 1B and 2B. The answer for 1B depends on the probability of a girl being named Florida, whereas 2B does not. It is interesting to ponder why the probability matters for case 1B but not for case 2B when, in some sense, the questions are the same. 2B could be turned into a question about a family simply by adding the words "... Florida who is the daughter of 2-child family. What is the chance the 
family she is a daughter of is a girl-girl family."

In fact, you can do this with the simulator by adding a `--count-families` flag which causes the simulator to count the family of each girl in the selected population and to do so at most once. This transforms the stream of
girls into a stream of families.

Observe first how the 2B answer is stable even if the probability is varied. At this point we are still counting 
girls:

    $ florida --girls --florida --probability 0.001
    [girls] n=1000000, c=1026, gg=519, gg/c=50.6%

    $ florida --girls --florida --probability 0.5
    [girls] n=1000000, c=499831, gg=249530, gg/c=49.9%    

    $ florida --girls --florida --probability 1
    [girls] n=1000000, c=999610, gg=497864, gg/c=49.8%

Now, let's add the --count-families flag and observe what happens when we count each family of selected girls just once and we vary the probability of girls being named "Florida":

    $ florida --girls --florida --count-families --probability 1
    [families] n=1000000, c=750278, gg=249059, gg/c=33.2%    

    $ florida --girls --florida --count-families --probability 0.5
    [families] n=1000000, c=437237, gg=187253, gg/c=42.8%

    $ florida --girls --florida --count-families --probability 0.001
    [families] n=1000000, c=1025, gg=517, gg/c=50.4%

In other words, even in 2B, the variation of 1B is in the data, provided you filter the data it in the right way.