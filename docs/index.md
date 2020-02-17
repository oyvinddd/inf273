# Mandatory Assignments

## Assignment 1a

As a solution representation to this problem, I’ve chosen to use an adjacency list data structure. This involves an array of size N + 1, where N is the number of ports to visit plus one for the starting/home position.

Further, each index contains a linked list of ports that are visited explicitly from the given child (see graphical representation below for better understanding). If there is no linked list in a given index in the main array, this means that no port was visited explicitly by a child.

![Figure 1a](resources/figure_1a.png)

## Assignment 1b

One possible solution representation for this problem would be to use a matrix. Number of columns in the matrix would be N + 1, where N is the number of nodes visited by the truck (+1 since start is visited two times). The number of rows in the matrix is equal to 1 (truck row) + 2 * the number of drones used by the truck.

Two rows are dedicated to each drone, where the first one contains positions where the drone is outgoing (to serve the customer), and the second row contains positions where the drone is incoming (from the customer). See below matrix for how this looks for the instance in the lecture slide.

![Figure 1b](resources/figure_1b.png)

Note that if we specify that a given drone only can visit one customer before returning to the truck, we could simplify the matrix to only contain one row for each drone (the row containing the outgoing drones).

## Assignment 2

## Assignment 3