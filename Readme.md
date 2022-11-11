# _Solution:_
First of, lets reduce concretion and compose abstract mathematical model.

From the given data we are able to build directed graph, where nodes are represented as train stations.

Considering assignment specification, we are able to conclude that the problem is in TSP (Travelling Salesman Problem) class. Hence key idea:
1. Find all Hamiltonian paths
2. Select best by money cost or time consumption

# Realization and peculiarities:
As a matter of fact, TSP-like is NP-hard so where does not exist easy solution.
1. Hamiltonian paths are found with recursive backtracking-based algorithm.
2. Solution uses concurrency for computations optimization.
3. For selection are used custom algorithms.

* ## **By price:**
This request allows to build "Divide and Conquer" - based algorithm.
For each stations pair in Hamiltonian paths are found and left only the most cheap options of the train rides.
Then, having best solutions between separate components of the path, we obtain best variant(s) for current stations sequence.
Of all those paths select cheapest, which respectivelly is the answer to the assignment. 
* ## **By time:**
Here is not possible to use the same approach as in the selection by price, because of the waiting time between rides - it does not allow to divide them separately and just use ride duration as a weight. Hence, the only possible approach is bruteforce-based: compose all possible combinations of the rides and select the best.

# Deployment and testing:
1. Clone this repository:
```bash
git clone https://github.com/NikitaMasych/DL-TestTask
``` 
2. Switch to the repository directory:
```bash
cd ./DL-TestTask
```
3. Launch program:
```bash
make run
```
* In order to launch tests:
```bash
make test
```
# Project structure:

```
.
├── code
│   ├── cmd
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── graph
│   │   ├── graph.go
│   │   ├── graph_test.go
│   │   ├── hamiltonian_path.go
│   │   ├── hamiltonian_test.go
│   │   └── mocks.go
│   ├── plans
│   │   ├── best_price.go
│   │   ├── best_time.go
│   │   ├── price_test.go
│   │   └── time_test.go
│   └── utils
│       ├── auxiliary.go
│       └── repository.go
├── data
│   └─── schedule.csv
├── makefile
│── Readme.md
└── specs.md
```