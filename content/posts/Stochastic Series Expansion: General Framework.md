---
title: "Stochastic Series Expansion: General Framework"
date: 2021-05-30T00:09:27Z
draft: false
---

We seek to establish the basic mathematical grounding for simulating thermodynamic quantities of many-body quantum systems. Specifically, we aim to implement a program that can simulate spin-1/2 systems. For accomplishing that task, I've chooses to work with the stochastic series expansion framework which will end up being implemented on top of a metropolis algorithm together with loop updates[^1]. 

The developed content was produced from my notes when studding the amazing "Computational Studies of Quantum Spin Systems" from Sandvik[^2] which were a fantastic guide to this topic. Thanks, Sandvik!

## Expanding the exponential

We begin introducing the fundamentals behind the stochastic series expansion. On to the definition of the partition function of a quantum system, we immediately find ourselves in front of a hairy problem:
$$
Z = Tr \ e^{-\beta H} 
$$
The computation of that trace above requires the knowledge of the whole set of eigenstates of the operator H (and hence, of $e^{\beta H} $).  Remember that we are generally dealing with a system with a large scale of interacting entities, diagonalizing such systems is a hard problem even for the simpler models! 

A possible approach could be obtained by approximately the exponential by a product of a series of exponentials each one containing a term of $H$ which we can diagonalize separately. In principle this is would not be possible for non commuting terms on the Hamiltonian but the Trotter formula ensures that the exact exponential is recovered by having an infinite number of "slices". This idea is the entry point of the Suzuki-Trotter decomposition which introduces a new time dimension to our problem by slicing the exponential operator into small spaced time intervals. Because of the uncontrolled systematic error introduced by these slices, we will not take this path here now. Who knows in a future post?

Instead we will make no approximations regarding $H$, we begin the computation by expanding the exponential in Taylor series and inserting $n - 1$ identities between each $H$:
$$
\begin{align}
Z &= \sum_{\psi_0} \left< \psi_0 \right| \left( \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!} H^n \right) \left| \psi_0 \right> \\\\\\
&= \sum_{\psi_0} \left< \psi_0 \right| \left( \sum_{n = 0}^\infty \frac{(-\beta)^n}{n!} 
	\underbrace{ H \ \mathbb{1} \ H \dots H \ \mathbb{1} \ H }_{\text{$n - 1$ times $\mathbb{1}$}} \right)  \left| \psi_0 \right>
\end{align}
$$

Now each identity $\mathbb{1}$ can be replaced by a resolution $\mathbb{1} = \sum \left| \psi_k \right> \left< \psi_k \right|$ and all the summations can be factored out:
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
	H  \left| \psi_{k + 1} \right>$ are not positive in general. This could eventually lead us to configurations with negative probabilities and consequently wrong estimative of the thermodynamic means. This is essentially the origin of the sign problem present in the description of frustrated and fermionic systems.  The positiveness of the above term can be guarantied by the insertion of a constant term in the Hamiltonian for spin-1/2 systems in bipartite lattices (we will see this in details latter) leading us to restrict our discussion only to that case. We should point, however, that this restriction doesn't mean that our system cannot be simulated, but a specialized treatment should be addressed and still the problem turns out to take exponential complexity regarding the number of particles[^3]. 

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
The "gotcha" here is that all this summations will be evaluated stochastically, even the expansion order. Normally in a metropolis simulation the sample/configuration space is the same as the state space while here we sample a random expansion order $n$ and also a set of $n$ states. 

## Sampling on Hamiltonian terms

We can further write each Hamiltonian as a sum of individual interactions. Suppose that we are dealing with two body interactions and each entity of out system is connected by a bond indexed by $b$, if we have multiple types of interactions we can index each type with a variable $t$. The Hamiltonian is written as a sum in interaction types and in bonds:
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
 Because each summation is independent, all the summations factor out and the sampled term is now a sequence of operators
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
Note that we now sample not only on the order $n$ and the states $\{\psi_k\}$ but also on the type, bond pair $[t_m, b_m]$. This allow us to distinguish between different kinds of interactions which will be very useful when constructing the update schemes on an actual simulation. 

## Measuring physical observables

If you've read the last section, than you realized that we can also sample under $H$ terms. For the sake of simplicity, will not include the extra summation on the internal energy and heat capacity estimators deviation. The resulting formulas still valid if the bond $H_{[t, b]}$ were sampled stochastically, and can easy be verified by a substitution of
$$
H = \sum_{\{t_m, b_m\}} H_{[t_m, b_m]}
$$
right before the identification of the ratio of summations with a thermal average $\left< \dots\right>$. Also, a new argument on the weights should be included as they now depend on the type and bond of the sampled operators:
$$
W_n(\{\psi_k\}) \rightarrow W_n(\{\psi_k\}, \{t_m, b_m\})
$$


### Internal Energy

The expression for the internal energy $\left< E \right>$ is a rather fancy one. From the statistical mechanics theory, we can calculate the mean energy by differentiating $\ln Z$ in terms of $\beta$:
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
But with an extra summation and an extra matrix element on the product. In total, we have $n + 1$ matrix elements and summations for each term of order $n$ in $\beta$. We can re-index the summation in $n$ so that the states summation have the same number of factors that the original term present in $Z$. With that in mind we proceed with defining a new $n \rightarrow m - 1$ so that the series lower limit transforms like $0 \rightarrow 1$:
$$
\begin{align}
\left< E \right>= \frac{1}{Z} \sum_{m = 1}^\infty \frac{(-\beta)^{m - 1}}{(m - 1)!}
 	\sum_{\{\psi_k\}_0^{m - 1}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{m - 1} \right| H 
	\left| \psi_0 \right> 
\end{align}
$$
Now the states summation looks just like the one present in the denominator $Z$. But what about the missing factors in $\frac{(-\beta)^{m - 1}}{(m - 1)!}$ and the lower limit of the series? Note that we can write the previous factor as:
$$
\frac{(-\beta)^{m - 1}}{(m - 1)!} = \frac{(-\beta)^m}{m!} \frac{m}{(-\beta)}
$$
And note that this factor is equal to $0$ when $m = 0$, so we can safely start our series from $m = 0$:
$$
\begin{align}
\left< E \right> &= \frac{1}{Z} \sum_{m = 0}^\infty \frac{(-\beta)^m}{m!}
	\frac{m}{(-\beta)}
 	\sum_{\{\psi_k\}_0^{m - 1}} \left< \psi_0 \right|
	H  \left| \psi_1 \right> \left< \psi_1 \right| 
	H  \left| \psi_2 \right>
	\dots
	\left< \psi_{m - 1} \right| H 
	\left| \psi_0 \right> \\\\
&= \frac{1}{Z} \sum_{m = 0}^\infty \sum_{\{\psi_k\}_0^{m - 1}}
	\frac{(-\beta)^m}{m!}
 	\left< \psi_0 \right|
		H  \left| \psi_1 \right> \left< \psi_1 \right| 
		H  \left| \psi_2 \right>
		\dots
		\left< \psi_{m - 1} \right| H 
	\left| \psi_0 \right>  \frac{m}{(-\beta)} \\\\
&= \frac{1}{Z} \sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n - 1}}
	\underbrace{
	\frac{(-\beta)^n}{n!}
 	\left< \psi_0 \right|
		H  \left| \psi_1 \right> \left< \psi_1 \right| 
		H  \left| \psi_2 \right>
		\dots
		\left< \psi_{n - 1} \right| H 
	\left| \psi_0 \right>}_{W_n(\{\psi_k\})}  \frac{n}{(-\beta)} \\\\
&= - \frac{\left< n \right>}{\beta}
\end{align}
$$
Where we renamed $m \rightarrow n$ from the second to the third line only for sake of notation (the symbol used has no importance, instead we should refer to this quantity always as the expansion order). Also note that the weight $W_n(\{\psi_k\})$ is the exactly the summation term of $Z$. This is in fact the probability weight we will use on our simulation, but for now let us delight ourselves with the fancy expression that we arrived
$$
\left< E \right> = -\frac{\left< n \right>}{\beta}
$$
The above expression tels us that $n/\beta$ can be interpreted as an estimator of $E$ if the above weights are positive (and can be viewed as unnormalized probabilities).

### Heat Capacity

For deriving an expression for the heat capacity, lets us begin writing its definition in terms of the partition function:
$$
\begin{align}
c_v &= \partial_T \left< E \right> \\\\
&= \partial_T (-\partial_\beta \ln Z) \\\\
&= -\partial_T(1/T) \partial_\beta \partial_\beta \ln Z \\\\
&= \beta^2 \partial^2_\beta\ln Z \\\\
&= \beta^2 \partial_\beta \left( \frac{1}{Z} \partial_\beta Z \right) \\\\
&= \beta^2 \left( -\frac{1}{Z^2} (\partial_\beta Z)(\partial_\beta Z) + \frac{1}{Z} \partial^2_\beta Z \right) \\\\
&= \beta^2 \left( -(\partial_\beta \ln Z)(\partial_\beta \ln Z) + \frac{1}{Z} \partial^2_\beta Z \right) \\\\
&= \beta^2 \left( -\left< E\right>^2 + \frac{1}{Z} \partial^2_\beta Z \right) \\\\
\end{align}
$$
Now we need to calculate the second term inside the parenthesis $(\partial^2_\beta Z)/Z$. Instead of expanding that expression for the second derivative in the denominator, we will derive a formula for an arbitrary derivative order $q$ . For this, recall how we've "corrected" the series summation terms of the internal energy expression. In that case we had a ratio of two summations. because of the differentiation, we got an extra $H$ term inside the summation of the numerator. The trick used is to do some algebra so we can identify the same configuration weight that appears on the denominator terms. We also needed to re-index the series because the summation index $n$ on the numerator had to match the respective weight (a sequence of matrix elements of $H$). For each term of order $n$ the weight involved a sequence of $n + 1$ matrix elements (ranging from $0$ to $n$). If instead we had differentiated $q$ times in respect to $-\beta$ than we would be off by $k$ terms for each term. To correct that, we would need to introduce a new index $n - q\rightarrow m$  which would fix the one to one correspondence between the term order and the number of matrix elements. That would lead to the same conclusions about the $(-\beta)^n/n!$ factor and the starting index (the differentiations wipes the first $q$ terms). Note how that´s fixed by identifying a weight $W_n$ for a term with order $n$:
$$
\frac{(-\beta)^{n - q}}{(n - q)!} = \frac{(-\beta)^n}{n!} \frac{n (n - 1)\dots (n - q + 1)}{(-\beta)^q} = W_n \frac{n (n - 1)\dots (n - q + 1)}{(-\beta)^q}
$$
You can follow the described sequence in the following steps (notice that we only modify the numerator):
$$
\begin{align}
\frac{1}{Z}\partial^q_\beta Z &= \frac{1}{Z} \partial^q_\beta \left( 
	\sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n}} \frac{(-\beta)^n}{n!}
 	 \prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right>
	\right) \\\\
&= \left( 
	\sum_{n = k}^\infty \sum_{\{\psi_k\}_0^{n}} \frac{(-\beta)^{n - q}}{(n - q)!}
 	\prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right>
	\right) \Bigg/ \left( 
	\sum_{n = 0}^\infty  \sum_{\{\psi_k\}_0^{n}}\frac{(-\beta)^n}{n!}
 	\prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right>
	\right) \\\\
&= \left( 
	\sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n}} 
	\frac{(-\beta)^n}{n!} \frac{n (n - 1)\dots (n - q + 1)}{(-\beta)^q}
 	\prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right>
	\right) \Bigg/ \left( 
	\sum_{n = 0}^\infty  \sum_{\{\psi_k\}_0^{n}}\frac{(-\beta)^n}{n!}
 	\prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right>
	\right) \\\\
&= \left( 
	\sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n}} 
	\frac{n (n - 1)\dots (n - q + 1)}{(-\beta)^q} 
	\underbrace{ \frac{(-\beta)^n}{n!}
 	\prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right> }_{W_n(\{ \psi_k\})}
	\right) \Bigg/ \left( 
	\sum_{n = 0}^\infty  \sum_{\{\psi_k\}_0^{n}}
	\underbrace{ \frac{(-\beta)^n}{n!}
 	\prod_{k} \left< \psi_k \right| H  \left| \psi_{k + 1} \right> }_{W_n(\{ \psi_k\})}
	\right) \\\\
&= \left( 
	\sum_{n = 0}^\infty \sum_{\{\psi_k\}_0^{n}} 
	\frac{n (n - 1)\dots (n - q + 1)}{(-\beta)^q}
 	W_n(\{ \psi_k\})
	\right) \Bigg/ \left( 
	\sum_{n = 0}^\infty  \sum_{\{\psi_k\}_0^{n}} W_n(\{ \psi_k\})
	\right) \\\\
&= \frac{\left< n \ (n - 1) \dots (n - k + 1) \right>}{(-\beta)^k}
\end{align}
$$
Note that in the third line the series starting index changed $n = k \rightarrow n = 0$, we were only able to do this because the quantity inside the thermal mean is equally 0 if $n < k$, so we are not missing nor adding terms.

Finally we can take $k = 2$ for the case of the specific heat:
$$
\begin{align}
c_v &= \beta^2 \left( -\left< E\right>^2 + \frac{1}{Z} \partial^2_\beta Z \right) \\\\
&= \beta^2 \left( -\left< E\right>^2 + \frac{\left<n(n - 1)\right>}{\beta^2}\right) \\\\
&= \beta^2 \left( -\left< E\right>^2 + \frac{\left<n^2 - n\right>}{\beta^2}\right) \\\\
&= \beta^2 \left( -\frac{\left< n\right>^2}{\beta^2} + \frac{\left<n^2 \right> - \left<n\right>}{\beta^2}\right) \\\\
&= \left< n^2\right> - \left<n \right>^2 - \left<n \right>
\end{align}
$$
Again, note that we made no assumptions about the nature of $H$ to arrive at this result.

If the weights are positive, we can interpret them as probabilities (with normalization given by $Z$), so we can re-write the above expression in a more "statistical language" as
$$
c_v = Var[n] - E[n]
$$

## Truncating the series

Every formula derived until here could be perfectly used to compute the macroscopic properties of quantum system by the means of a computer. However, one should point out the corresponding data structure used as the "current configuration" on the context of a metropolis simulation. The procedure followed on that algorithm consists on constantly updating an actual configuration according to random moves obeying some transitions probabilities. If that probabilities had correctly chosen, we are them guarantied that the sampled configuration follows the equilibrium distribution and statistical measurements can be made using the generated configuration. Because of the constantly updates on the current configuration, we may choose a data structure that can afford such often write operations. Lets check on a table what are the variables that describe a sampled configuration on the SSE framework:

| **Configuration/Sampled variable**                           | Data size |
| ------------------------------------------------------------ | --------- |
| $n$                                                          | $1$       |
| set of states $\{\psi_k\}$                                   | $n$       |
| string of operators $\prod_k \left< \psi_k \right\| H_k \left\| \psi_{k + 1}\right>$ | $n$       |
| type of each operator $H_{[t_k, b_k]} \rightarrow t_k$       | $n$       |
| bond of each operator $H_{[t_k, b_k]} \rightarrow b_k$       | $n$       |

Note that there are 3 quantities whose size may vary with the expansion order $n$. Thats because the operators string is constantly growing/shrinking as we update $n$. Obviously we could use a list as a data structure when implementing that because of the mutable size. But mutable size is not a preferably approach when considering performance costs (not only they have linear access time but the data is not stored contiguously in memory). When designing high performance simulations we should try to keep data as compactly as possible and minimize the access time across by storing it contiguously. Also, not having to rely on constant copies operations or rewrites is something to keep in mind.

Its expected that out Markov chain fluctuate around a mean $\left<n\right>$ so that truncating the series expansions on a maximum $L > n_{max}$ would enable us to use a vector data structure for the above three sampled quantities. The maximum size $L$ should be calculated by fixing the simulation parameters (such as $\beta$), evaluating the mean $\left< n \right>$ and resizing to the new size $L \rightarrow \alpha \left< n\right>$ where something like $\alpha = 1.3$ should do it. The above truncation does not introduce any systematic errors to our simulation, in fact, configurations for $n > L are rarely reachable.

We can also take a step further and require that every operator's sequence have a fixed size $L$, that can be done inserting $L - n$ identity operators between the present operators. Doing so, the sum in $n$ is not required in our $Z$ expansion anymore. Note however, that an old configuration with $n$ operators may be mapped to a series of sequences which have the same operators in the same order but with $L - n$ identities distributed. In other words our configuration have being copied $L \choose n$ times according to the ways of distributing the sequence along the $L$ slots (you can also think the other way around and distribute the identities $L \choose L - n$).  To account for that, we divide the weight by the $L \choose n$ factor:
$$
\begin{align}
Z &= \sum_{\{\psi_k\}_0^{L - 1}} \sum_{\{t_m, b_m\}^{L - 1}_0 }
	\frac{(-\beta)^n}{n!} \frac{n! (L - n)!}{L!}
	\left< \psi_0 \right| H_{[t_0, b_0]} \left| \psi_1 \right> 
	\left< \psi_1 \right| H_{[t_1, b_1]}  \left| \psi_2 \right>
	\dots
	\left< \psi_{L - 1} \right| H_{[t_{L - 1}, b_{L - 1}]}
	\left| \psi_0 \right> \\\\
 &= \sum_{\{\psi_k\}_0^{L - 1}} \sum_{\{t_m, b_m\}^{L - 1}_0 }
	\frac{(-\beta)^n (L - n)!}{L!}
	\left< \psi_0 \right| H_{[t_0, b_0]} \left| \psi_1 \right> 
	\left< \psi_1 \right| H_{[t_1, b_1]}  \left| \psi_2 \right>
	\dots
	\left< \psi_{L - 1} \right| H_{[t_{L - 1}, b_{L - 1}]}
	\left| \psi_0 \right> \\\\
\end{align}
$$
Where now a special value for $t_m$ can be attributed so that it represents the presence of a identity operator (for example $H_{[0, b]} = \mathbb{1}$) and $n$ represents the number of non identity operators in the string. 

The final configuration weight can be finally written in terms of the operator sequence and the intermediary states:
$$
W(\{\psi_k\}, \{t_m, b_m\}) = \frac{(-\beta)^n (L - n)!}{L!} 
	\prod_{k = 0}^L
	\left< \psi_k \right| H_{[t_0, b_0]} \left| \psi_{k + 1} \right> 
$$
Where we point-out that periodic boundary conditions on the "time" component should be satisfied due to the initial trace on the $Z$  definition. Hence we have that $\left| \psi_{L + 1} \right> = \left| \psi_{0} \right>$.

[^1]: [Stochastic series expansion method with operator-loop update](https://journals.aps.org/prb/abstract/10.1103/PhysRevB.59.R14157)
[^2]: [Computational Studies of Quantum Spin Systems](https://aip.scitation.org/doi/abs/10.1063/1.3518900)
[^3]: [Computational complexity and fundamental limitations to fermionic quantum Monte Carlo simulations](https://journals.aps.org/prl/abstract/10.1103/PhysRevLett.94.170201)

