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
4. Project uses Docker, PostgreSQL:
From schedule csv data is composed table in deployed via docker image database. 
It is done in order to make a good use of necessary for solution queries. 

* ## **By price:**
This request allows to build "Divide and Conquer"-based algorithm.
For each departure/arrival stations pair in the Hamiltonian path is found the minimal money cost.
Then, having best solutions between separate components of the path, we obtain minimal money cost for current stations sequence.
Of all those paths select cheapest and compose possible train rides, which, respectivelly, is the answer to the assignment. 

* ## **By time:**
Here is not possible to use the same approach as in the selection by price, because of the waiting time between rides - it does not allow to divide them separately and just use ride duration as a weight. Breaking example:

* 00:00:00 - > 01:00:00, 02:00:00 -> 03:00:00
* 00:00:00 - > 00:30:00, 01:30:00 -> 03:00:00

Here we have four same by overall time solutions, but using the previous algorithm, we obtain only one:

* 00:00:00 - > 00:30:00, 02:00:00 -> 03:00:00

Hence, the only possible approach is bruteforce-based: compose all possible combinations of the rides for each path, calculate minimal time duration and compose train rides for the best path.

 _Considering database calls overhead and none bonuses of queries there are used raw records from the csv, stored as variable_ 

# Deployment and testing:
* Prerequisites: Make sure you have docker daemon running
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
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── postgres.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── graph
│   │   ├── graph.go
│   │   ├── graph_test.go
│   │   ├── hamiltonian_path.go
│   │   └── hamiltonian_test.go
│   ├── models
│   │   └── route.go
│   ├── plans
│   │   ├── best_price.go
│   │   ├── best_time.go
│   │   └── time_test.go
│   └── utils
│       ├── auxiliary.go
│       └── repository.go
├── data
│   ├── schedule.csv
│   └── trains.sql
├── docker-compose.yml
├── makefile
├── Readme.md
└── specs.md
```