---
title: Simulating Accretion Disks with CFD
publication_date: 2025-10-24
tags: [math, physics]
---


Accretion disks around compact objects, such as black holes or neutron stars, are some of the most fascinating astrophysical systems. Modeling these disks requires solving the full set of compressible hydrodynamics equations under strong gravitational fields.

## Governing Equations

The fundamental equations we solve in computational fluid dynamics (CFD) for an ideal gas are the **Navier-Stokes equations** coupled with gravity:

**Continuity equation:**

$$
\frac{\partial \rho}{\partial t} + \nabla \cdot (\rho \mathbf{v}) = 0
$$

That was it, cool right?

![Beautiful Landscape](https://images.unsplash.com/photo-1506744038136-46273834b3fb?w=800)


**Momentum equation:**

$$
\frac{\partial (\rho \mathbf{v})}{\partial t} + \nabla \cdot (\rho \mathbf{v} \mathbf{v} + P \mathbf{I}) = -\rho \nabla \Phi
$$

**Energy equation:**

$$
\frac{\partial E}{\partial t} + \nabla \cdot [(E+P)\mathbf{v}] = -\rho \mathbf{v} \cdot \nabla \Phi
$$

where $ \rho $ is the density, $ \mathbf{v} $ is the velocity, $ P $ is the pressure, $ E $ is the total energy density, and $ \Phi $ is the gravitational potential.

## The Role of Viscosity

Viscosity in accretion disks is often parameterized using the **Shakura-Sunyaev alpha model**:

$$
\nu = \alpha c_s H
$$

Here, $ \nu $ is the kinematic viscosity, $ c_s $ the sound speed, $ H $ the disk scale height, and $ \alpha $ is a dimensionless parameter between 0 and 1 that captures turbulent transport.

## Numerical Implementation

To solve these equations, we typically discretize the domain using a **finite volume method**, which ensures conservation of mass, momentum, and energy. In cylindrical coordinates, the radial inflow is governed by:

$$
v_r \sim - \frac{3 \nu}{2 r}
$$

which highlights how viscous transport drives matter inward toward the central object.

### Boundary Conditions

At the inner boundary, we often apply a **free inflow** condition, allowing matter to fall into the compact object. At the outer boundary, a **reflective or outflow** boundary is used depending on the simulation setup.

---

Simulating these systems provides insight into high-energy phenomena like **X-ray emissions**, **jet launching**, and **quasi-periodic oscillations**. Future work involves coupling CFD simulations with **radiative transfer** to better predict observable signatures.

