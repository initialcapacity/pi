# Pi Estimation

Estimates the value of π using a Monte Carlo simulation.

## Method

We generate random points within the unit square and count how many fall within a circle of radius 1 centered at the origin.
The ratio of the number of points within the circle to the total number of points approaches π/4.

![Illustration of a Monte Carlo simulation](./readme_images/monte-carlo.svg)
