# This repo contains all go related code written for [blog](https://www.vivekakupatni.com/)

Uses [taskfile](https://taskfile.dev/) to run examples and test cases


## Running the test_cases

```bash
(ins)-> task test
task: [test] go test -v ./...
?       github.com/apatniv/website_golang       [no test files]
=== RUN   TestComputeSCC
=== RUN   TestComputeSCC/linear_chain
=== RUN   TestComputeSCC/single_isolated_nodes
=== RUN   TestComputeSCC/complete_3_node_cycle
=== RUN   TestComputeSCC/4_components_example
--- PASS: TestComputeSCC (0.00s)
    --- PASS: TestComputeSCC/linear_chain (0.00s)
    --- PASS: TestComputeSCC/single_isolated_nodes (0.00s)
    --- PASS: TestComputeSCC/complete_3_node_cycle (0.00s)
    --- PASS: TestComputeSCC/4_components_example (0.00s)
PASS
ok      github.com/apatniv/website_golang/graphs        (cach
```

## Running the examples


### Strongly connected components Example
```bash
(ins)-> task example
task: [example] go run main.go
2024/02/17 09:11:25 INFO Running the SCC example
2024/02/17 09:11:25 DEBUG Dfs status node=c discover=4 finish=5
2024/02/17 09:11:25 DEBUG Dfs status node=d discover=3 finish=6
2024/02/17 09:11:25 DEBUG Dfs status node=b discover=2 finish=7
2024/02/17 09:11:25 DEBUG Dfs status node=a discover=1 finish=8
2024/02/17 09:11:25 INFO SCC Results components="[[b a] [c d]]"
```