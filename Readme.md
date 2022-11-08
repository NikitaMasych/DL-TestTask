First of, lets reduce concretion and build abstract mathematical model.
From the given data we are able to build directed weighted graph.
Depending on the variant, weight may be: (NOTE: consider when we need best price and time simultaneously)
    1. Cost of the ride
    2. Time duration of the ride (NOTE: we need to take into account time between rides as well)
Hence, reading each line in the file, we fetch:
    1. Start node
    2. Final node
    3. Weight of the edge
Considering weight as a cost of the ride, If there exists another ride between the same stations, we select cheapest one.
If weight is time, it is a separate case.

How to represent the solution:
 Train ride numbers sequence (NOTE: consider, when there exist several equal)

KEY IDEA:
First we find all Hamilton paths in the graph and then sort them as needed.