# Constructing The Paradox That Wasn't

## Which World Do You Live In?

There exists a world W. In that world there exist many families. All families have exactly one mother and exactly one father. Each family has exactly two children. Each child is either a boy or a girl - the gender being determined at random at conception. 50% of all children are boys, 50% are girls. There are 4-types of family, classified by the gender of the children sorted in birth order: BB, BG, GB, GG. There are equally many BB, BG, GB and GG families. 

Consider the following statements which may or may not be true. You can assume the first is true.

* S: Within W there exists a family Y, consisting of mother M, a daughter G, and one other child C of unknown gender.
* Cg: C is a girl.
* Gg: G has a sister.
* Mgg: M has 2 daughters.
* Ygg: Y is a GG family.
* Wgg: If, and only if, all of Cg, Gg, Mgg, Ygg are true.

Observe that:

	Cg <=> Gg <=> Mgg <=> Ygg <=> Wgg

	where: a <=> b means a => b and b => a.

Q1: Determine P(Ygg|S).

	P(Ygg|S) = P(S|Ygg).P(Ygg)
	           ---------------
	                P(S)

    P(S) = P(S|Ybb).P(Ybb)+P(S|Ybg).P(Ybg)+P(S|Ygb).P(Ygb)+P(S|Ygg).P(Ygg)

    P(S|Ygg) = 1
    P(S|Ygb) = 1
    P(S|Ybg) = 1
    P(S|Ybb) = 0

    P(Ybb) = P(Ybg) = P(Ygg) = P(Ybb) = 1/4


    P(Ygg|S) =       1/4
               ---------------
               1/4 + 1/4 + 1/4

             = 1/3

Q2: Determine P(Mgg|S).

	P(Mgg|S) = P(S|Mgg).P(Mgg)
	           ---------------
	                P(S)

    P(S) = P(S|Mbb).P(Mbb)+P(S|Mbg).P(Mbg)+P(S|Mgb).P(Mgb)+P(S|Mgg).P(Mgg)

    P(S|Mgg) = 1
    P(S|Mgb) = 1
    P(S|Mbg) = 1
    P(S|Mbb) = 0

    P(Mbb) = P(Mbg) = P(Mgg) = P(Mbb) = 1/4

    P(Mgg|S) = 1/4
               -----
               1/4+1/4+1/4

             = 1/3

Q3: Determine P(Gg|S).

	P(Gg|S) = P(S|Gg).P(Gg)
	          -------------
	              P(S)

    P(S) = P(S|Gg).P(Gg)+P(S|Gb).P(Gb)

    P(S|Gg) = 1
    P(S|Gb) = 1

    P(Gg) = 1/2
    P(Gb) = 1 - P(Gg) = 1/2

    P(Gg|S) =    1/2
              ---------
              1/2 + 1/2

            = 1/2

Q4: Determine P(Cg|S).

	P(Cg|S) = P(S|Cg).P(Cg)
	          -------------
	              P(S)

    P(S) = P(S|Cg).P(Cg)+P(S|Cb).P(Cb)

    P(S|Cg) = 1
    P(S|Cb) = 1

    P(Cg) = 1/2
    P(Cb) = 1 - P(Cg) = 1/2

    P(Cg|S) =    1/2
              ---------
              1/2 + 1/2

            = 1/2

Q5. Determine P(Wgg|S).

	P(Wgg|S) = P(Ygg|S) = 1/3

but, also:

	P(Wgg|S) = P(Cgg|S) = 1/2

Why?

---

# No, There Is No Paradox

However, there is no paradox - just errors. The next sections document attempts to fix the errors
that lead to the inconsistency.

## First Attempt

The inconsistencies between Q1,Q2 and Q3,Q4 can be resolved if we assume the priors P(Gg) and P(Cg) are determined by P(Ygg|S) = 1/3 on the basis that we expect that Gg is true iff Ygg is true. For example:

	P(Gg|S) = P(S|Gg).P(Gg)
	          -------------
	              P(S)

    P(S) = P(S|Gg).P(Gg)+P(S|Gb).P(Gb)

    P(S|Gg) = 1
    P(S|Gb) = 1

    P(Gg) = 1/3
    P(Gb) = 1 - P(Gg) = 2/3

    P(Gg|S) =    1/3
              ---------
              1/3 + 1/3

            = 1/3

A similar argument can be used to revise P(Cg|S) to 1/3 also. 

So, we resolved the apparent paradox in the previous solution to Q5 and there was some justification for doing so since the equivalence of Gg and Ygg seemed sound. 

**Note, however, that the reasoning of the above attempt was faulty!**

## Second Attempt

The first attempt to fix the inconsistency produced the correct answer, however, the method was flawed. 
There is no justification for modifying the priors in this case, and certainly not by deriving it from the answer we obtained from the other methods - effectively assuming the answer.

The actual error is that the formulae for the likiehoods were incorrect and failed to take account of the fact that both P(S|Gg) and P(S|Gb) should be different given that there is only one way S|Gg can occur and 2 ways S|Gb can occur.

To correct this error, note that Gg <=> Ygg and Ygg is the only Y case where both Gg and S can be simultaneously true. Conversely, Gb <=> Ygb or Ybg and S can be true if either Ygb or Ybg is true. So, we can use these facts to derive P(S|Gg) and P(S|Gb) and thus P(Gg|S).

	P(Gg|S) = P(S|Gg).P(Gg)
	          -------------
	              P(S)

	P(S) = P(S|Gg).P(Gg)+P(S|Gb).P(Gb)

	P(S|Gg) =          P(Ygg)
		      ------------------------
			  P(Ygg) + P(Ygb) + P(Ybg)

			=       1/4
			  ---------------
			  1/4 + 1/4 + 1/4

			= 1/3

	P(S|Gb) =     P(Ygb) + P(Ybg)
		      ------------------------
			  P(Ygg) + P(Ygb) + P(Ybg)

			=       2/4
			  ---------------
			  1/4 + 1/4 + 1/4

			= 2/3


	P(Gb) = P(Gg) = 1/2

	P(Gg|S) =       1/2 * 1/3
	          ---------------------
	          1/2 * 1/3 + 1/2 * 2/3

	        = 1/3

A similar argument can be made for P(Cg|S), yielding the consistent result of 1/3 for Q4 also.