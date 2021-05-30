---
title: "Stochastic Series Expansion: General Framework"
date: 2021-05-30T00:09:27Z
draft: false
---

## Series Expansion

We begin introducing the fundamentals behind the stochastic series expansion. On to the definition of the partition function of a quantum system, we immediately find ourselves in front of a hairy problem:
$$
Z = Tr \ e^{-\beta H} 
$$
The computation of that trace above requires the knowledge of the whole set of eigenstates of the operator H (and hence, of $e^{\beta H} $).  Remember that we are generally dealing with a system with a large scale of interacting entities, diagonalizing such systems is a hard problem even for the simpler models! 

A possible approach could be obtained by approximately the exponential by a product of a series of exponentials each one containing a term of $H$ which we can diagonalize separately. This is would not be possible for non commuting terms on the Hamiltonian but the Trotter formula ensures that the exact exponential is recovered by having an infinite number of "slices". This approach is the entry point of the Suzuki-Trotter decomposition which introduces a new time dimension to our problem by slicing the exponential operator into small spaced time intervals. Because of the uncontrolled systematic error introduced by these slices, we will not take this path here now. Who knows in a future post?

Instead we will make no approximations on to $H$, we begin the computation by expanding the exponential in Taylor series and inserting $n - 1$ identities between each $H$:
$$
\begin{align}
Z &= \sum_{\psi_0} \left< \psi_0 \right| \left( \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!} H^n \right) \left| \psi_0 \right> \\\\\\
&= \sum_{\psi_0} \left< \psi_0 \right| \left( \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!} H \ \mathbb{1} \ H \dots H \ \mathbb{1} \ H \right) \left| \psi_0 \right>
\end{align}
$$

Now each identity $\mathbb{1}$ can be replaced by an individual resolution $\mathbb{1} = \sum \left| \psi_k \right> \left< \psi_k \right|$ and all the summations can be factored out:
$$
\begin{align}
Z &= \sum_{\psi_0} \left< \psi_0 \right| \left[
	\sum_{n = 0}^\infty \frac{(-\beta)^n}{n!}
	H \left( \sum_{\psi_1} \left| \psi_1 \right> \left< \psi_1 \right| \right)
	H \left( \sum_{\psi_2} \left| \psi_2 \right> \left< \psi_2 \right| \right)
	\dots
	\left( \sum_{\psi_{n - 1}} \left| \psi_{n - 1} \right> \left< \psi_{n-1} \right| \right) H
\right] \left| \psi_0 \right> \\\\
 &= \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!} \left[
 	\sum_{\psi_0} \left< \psi_0 \right|
	H \left( \sum_{\psi_1} \left| \psi_1 \right> \left< \psi_1 \right| \right)
	H \left( \sum_{\psi_2} \left| \psi_2 \right> \left< \psi_2 \right| \right)
	\dots
	\left( \sum_{\psi_{n - 1}} \left| \psi_{n - 1} \right> \left< \psi_{n-1} \right| \right) H
	\left| \psi_0 \right>
\right] \\\\
&= \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!}
 	\sum_{\psi_0} \sum_{\psi_1} \dots \sum_{\psi_{n-1}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right|  H
	\left| \psi_0 \right> \\\\
&= \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!}
 	\sum_{\{\psi_k\}_0^{n - 1}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right|  H
	\left| \psi_0 \right> \\\\
\end{align}
$$
Note now that we have finally now written the operators in term of matrix elements. Our partition function's expression could, in theory, be to express the transition probabilities required by the metropolis algorithm. However, we should point that the matrix elements on the above terms $\left< \psi_k \right|
	H  \left| \psi_{k + 1} \right>$ are not positive in general. This could eventually lead us to negative probabilities and consequently wrong estimative of the thermodynamic means. This is essentially the origin of the sign problem present in the description of frustrated and fermionic systems.  The positiveness of the above term can be guarantied by the insertion of a constant term in the Hamiltonian for bosonic spin systems in bipartite lattices (we will see this in details latter) leading us to restrict our discussion only to that case. We should point, however, that this restriction doesn't mean that our system cannot be simulated, but a specialized treatment should be addressed and the trade-off is an exponential convergence on the system's size in order to overcome the cancellation of the summation terms. 

Finally we will move the states summation next to the expansion order, which gives us the following form:
$$
\begin{align}
Z &= \sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n - 1}}
	\frac{(-\beta)^n}{n!} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right|  H
	\left| \psi_0 \right> \\\\
\end{align}
$$
The "gotcha" here is that all this summations will be evaluated stochastically, even the expansion order. Normally in a metropolis simulation the sample/configuration space is the same as the state space while here we sample a random expansion order $n$ and also a set of $n$ states. We can further write each Hamiltonian as a sum of individual interactions. Suppose that we are dealing with two body interactions and each entity of out system is connected by a bond indexed by $b$, if we have multiple types of interactions we can index each type with a variable $t$. The Hamiltonian is written as a sum in interaction types and in bonds:
$$
H = \sum_{t, b} H_{[t, b]}
$$
Each term $H_{[t,b]}$ could represent, for instance, an exchange term between two adjacent spins $H_{1, b} = S_{i(b)}S_{j(b)}$ where  both $i(b)$ and $j(b)$ maps the bond to the corresponding sites and the $t = 1$ would represent an exchange interaction type. Substituting the above expression into the last expression for $Z$ we have:
$$
\begin{align}
Z &= \sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n - 1}}
	\frac{(-\beta)^n}{n!} \left< \psi_0 \right|
	\left( \sum_{t_0, b_0} H_{[t_0, b_0]} \right)  \left| \psi_1 \right> \left< \psi_1 \right| 
	\left( \sum_{t_1, b_1} H_{[t_1, b_1]} \right)  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right|  \left( \sum_{t_{n - 1}, b_{n - 1}} H_{[t_{n - 1}, b_{n - 1}]} \right)
	\left| \psi_0 \right> \\\\
\end{align}
$$
 Because each summation is independent, all the summations factor out and the sample term is now a sequence of operators
$$
\begin{align}
Z &= \sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n - 1}} \sum_{\{t_m, b_m\}^{n - 1}_0 }
	\frac{(-\beta)^n}{n!} \left< \psi_0 \right|
	H_{[t_0, b_0]} \left| \psi_1 \right> \left< \psi_1 \right| 
	H_{[t_1, b_1]}  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right| H_{[t_{n - 1}, b_{n - 1}]}
	\left| \psi_0 \right> \\\\
\end{align}
$$
Note that we now sample not only on the order $n$ and the states $\{\psi_k\}$ but also on the type, bond pair $[t_m, b_m]$. This allow us to distinguish between different kinds of interactions which will be very useful when constructing the update schemes on an actual simulation. ]

## Internal Energy

Now we will derive an (rather fancy) expression for the internal energy $\left< E \right>$. From the statistical mechanics theory, we can calculate the mean energy by differentiating $\ln Z$ in terms of $\beta$:
$$
\begin{align}
\left< E \right> &= - \partial_{\beta} \ln Z \\\\
&= \frac{1}{Z} \sum_{\psi_0} \left< \psi_0 \right| H e^{-\beta H} \left| \psi_0 \right> \\\\
&= \frac{1}{Z} \sum_{\psi_0} \left< \psi_0 \right| H \left( \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!} H^n \right) \left| \psi_0 \right> \\\\
\end{align}
$$
Note that the above trace summation is exactly the one in the beginning of our expansion of $Z$ but with an extra $H$. Carrying out the same procedure we arrive on to the following expression:
$$
\begin{align}
\left< E \right> &= \frac{1}{Z} \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!}
 	\sum_{\psi_0} \sum_{\psi_1} \dots \sum_{\psi_{n-1}} \sum_{\psi_{n}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right| H \left| \psi_n \right>
	\left< \psi_{n} \right| H \left| \psi_0 \right> \\\\
&= \frac{1}{Z} \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!}
 	\sum_{\{\psi_k\}_0^{n}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right| H \left| \psi_n \right>
	\left< \psi_{n} \right| H \left| \psi_0 \right> 
\end{align}
$$
But with an extra summation and an extra matrix element on the product. In total, we have $n + 1$ matrix elements and summations for each term of order $n$ in beta. We can re-index the summation in $n$ so that the states summation have the same number of factors that the original term present in $Z$. With that in mind we proceed with defining a new $n \rightarrow n - 1$ so that the series lower limit transforms like $0 \rightarrow 1$:
$$
\begin{align}
\left< E \right>= \frac{1}{Z} \sum_{n = 1}^\infty \frac{(-\beta)^{n - 1}}{(n - 1)!}
 	\sum_{\{\psi_k\}_0^{n - 1}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right| H 
	\left| \psi_0 \right> 
\end{align}
$$
Now the states summation looks just like the one present in the denominator $Z$. But what about the missing factors in $\frac{(-\beta)^{n - 1}}{(n - 1)!}$ and the lower limit of the series? Note that we can write the previous factor as:
$$
\frac{(-\beta)^{n - 1}}{(n - 1)!} = \frac{(-\beta)^n}{n!} \frac{n}{(-\beta)}
$$
And note that this factor is equal to $0$ when $n = 0$, so we can safely start our series from $n = 0$:
$$
\begin{align}
\left< E \right> &= \frac{1}{Z} \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!}
	\frac{n}{(-\beta)}
 	\sum_{\{\psi_k\}_0^{n - 1}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{n - 1} \right| H 
	\left| \psi_0 \right> \\\\
&= \frac{1}{Z} \sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n - 1}}
	\frac{(-\beta)^n}{n!}
 	\left< \psi_0 \right|
		H  \left| \psi_1 \right> \left< \psi_1 \right| 
		H  \left| \psi_2 \right>
		\dots
		\left< \psi_{n - 1} \right| H 
	\left| \psi_0 \right>  \frac{n}{(-\beta)} \\\\
&= - \frac{\left< n \right>}{\beta}
\end{align}
$$
Note that the last passage used the fact that the weight multiplying the $-n/\beta$ factor is exactly the term that appears on the $Z$ summation. This is in fact the probability weight we will use on our simulation, but for now let us delight ourselves with the fancy expression that we arrived
$$
\left< E \right> = -\frac{\left< n \right>}{\beta}
$$
The above expression tels us that $n/\beta$ can be interpreted as an estimator of $E$ if the above weights are positive.

## Heat Capacity

For deriving an expression for the heat capacity, lets us recall how we've "corrected" the series summation terms of the internal energy expression so the probability weight present in $Z$ appears and we can identify the expression as a thermal mean. Because of the differentiation, we got an extra $H$ term inside the trace, so that to match the number of matrix elements with the expansion term's order we carried out a re-indexing $n \rightarrow n - 1$. If instead we had differentiated $k$ times in respect to $-\beta$ than we would be off by $k$ terms for each term. To correct that, we would need to introduce a new index $n \rightarrow n - k$ but that would lead to the same conclusions about the starting index and the $(-\beta)^n/n!$ factor, to correct that, we can rewrite this as:
$$
\frac{(-\beta)^{n - k}}{(n - k)!} = \frac{(-\beta)^n}{n!} \frac{n (n - 1)\dots (n - k + 1)}{(-\beta)^k}
$$
The $(-\beta)^n/n!$ factor can be absorbed as a weight present on the denominator $Z$ and the remaining product ensures that the series terms is zeros whenever n is less than $k - 1$ so that we can include those terms in the summation in order to fix the starting series index. We end up with:
$$
\begin{align}
\partial^k_{(-\beta)} \ln Z &= \frac{1}{Z} \partial^k_{(-\beta)}Z \\\\
&= \frac{\left< n \ (n - 1) \dots (n - k + 1) \right>}{(-\beta)^k}
\end{align}
$$


