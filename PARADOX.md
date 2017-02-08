# Which World Do You Live In?

There exists a world W. In that world there exist many families. All families have exactly one mother and exactly one father. Each family has exactly two children. Each child is either a boy or a girl - the gender being determined at random at conception. 50% of all children are boys, 50% are girls. There are 4-types of family, classified by the gender of the children sorted in birth order: BB, BG, GB, GG. B+G is describes the type that includes both BG and GB when the birth order is not important. There are equally many BB, BG, GB and GG families. There are twice as many B+G families as either BB or separately GG families.

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


    P(Ygg|S) = 1/4
               -----
               1/4+1/4+1/4

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

    P(S|Gg) = 1/2
    P(S|Gb) = 1/2

    P(Gg) = 1/2
    P(Gb) = 1 - P(Gg) = 1/2

    P(Gg|S) = 1/4
              -------
              1/4+1/4

            = 2/4

            = 1/2

Q4: Determine P(Cg|S).

	P(Cg|S) = P(S|Cg).P(Cg)
	          -------------
	              P(S)

    P(S) = P(S|Cg).P(Cg)+P(S|Cb).P(Cb)

    P(S|Cg) = 1/2
    P(S|Cb) = 1/2

    P(Cg) = 1/2
    P(Cb) = 1 - P(Cg) = 1/2

    P(Cg|S) = 1/4
              -------
              1/4+1/4

            = 2/4 

            = 1/2

Q5. Determine P(Wgg|S).

	P(Wgg|S) = P(Ygg|S) = 1/3

but, also:

	P(Wgg|S) = P(Cgg|S) = 1/2

Why?

---

The inconsistencies between Q1,Q2 and Q3,Q4 can be resolved if we assume the priors P(Gg) and P(Cg) are determined by P(Ygg|S) = 1/3 on the basis that we expect that Cg is true iff Yg is true. For example:

	P(Gg|S) = P(S|Gg).P(Gg)
	          -------------
	              P(S)

    P(S) = P(S|Gg).P(Gg)+P(S|Gb).P(Gb)

    P(S|Gg) = 1/2
    P(S|Gb) = 1/2

    P(Gg) = 1/3
    P(Gb) = 1 - P(Gg) = 2/3

    P(Gg|S) = 1/2 * 1/3
              ---------------------
              1/2 * 1/3 + 2/3 * 1/2

            = 1/3

A similar argument can be used to revise P(Cg|S) to 1/3 also. So, we resolved the apparent paradox in the previous
solution to Q5 and there was some justification for doing so since the equivalence of Cg anf Yg seemed sound. Still, if not for the contradiction, the previously selected priors could also have been justified. The decision
to switch the priors from 1/2 to 1/3 is, in some sense (I think), arbitrary.