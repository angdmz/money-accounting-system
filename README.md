# money-accounting-system
Golang money accounting system API REST
Built on go 1.12.7

### Installation
- clone the project
- on shell:
```sh
go get -u github.com/flosch/pongo2
go get -u github.com/gorilla/mux
go get github.com/sean-public/fast-skiplist
go build main.go
./money-accounting-system &
```

### Decisions in design of solution
1) For solving the REST part, I went for gomux but for no particular reason, it may have been gin gonic too
2) For the representation of the current balance, I used a uint64 and on credit and debit operations I had to lock the balance and atomically make the operation 
3) For the representation of the tx history as it had to be on memory, I had to use a data structure with fast operations that worked nice in concurrent environment, 
    1) skip list was my solution because of its fast operation times in practice, and because it puts locks only in minimum nodes in the list hence allowing other threads to use the resource with more freedom
    2) knowing that I had to list all the elements in an endpoint, I used a double linked list with pointers to the txs, as a redundance, and implementation of skip list I used was not easy for iteration, so the double linked list was just for ease 
4) For generating the ids, I decided to simply use a random int generator, taking a bet on non repetition because of large range in uint64
5) For the UI, I looked for a HTML template engine, a django-like one because I'm more familiar with it, and bootstrap for some style
6) I decided to let global variables in the code but I should have developed an abstraction that represents the transaction service and encapsulates my data structures

### Invariant on global variables
1) History is a map implemented on skip list with the global ids as keys, and the transactions as values
2) Every element in IdList is pointer to a tx that is a value in History
3) Every value in History is referenced as a node in IdList, hence History length and IdList length are the same
3) The sum of all credit minus the debits is equal to Balance
4) Balance is non negative

