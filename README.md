# Mini project 01 in Distributed systems
## Group members
Group name: We know less and less the more we know

| Name                     | ITU email   |
|--------------------------|-------------|
| Albert Rise Nielsen      | albn@itu.dk |
| Amalie Bøgild Sørensen   | abso@itu.dk |
| Andreas Nicolaj Tietgen  | anti@itu.dk |
| Asger Brødslev Mathiasen | asgm@itu.dk |


## Implementation details
This mini project is a solution to the Philosophers problem written in Go, using the sync library.

Instantiation of philosophers and forks is done through a public method, which for forks take an id, and then sets it's defaults. The same is done for the philosophers, which also take 2 pointers to it's left and right forks.

The fork structs have an internal sync.Mutex, which is wrapped with a Lock and Unlock method. The forks Lock method sets the forks status to "not free" and locks the Mutex. While the Unlock method increases the forks number of times used, sets it's status to "free" then unlocks the Mutex.

Each individual philosopher, when ready to eat, locking the left fork, when locked it locks the right one, using the forks public Lock and Unlock methods, eats, adds to its internal eating counter and then unlocks the forks. This is done in a thread that is started on initiation.

Another thread is started for each philosopher, on instantiation, which allows communication with the philosopher no matter it's current state. The same approach is used on the forks.

The communication to forks and philosophers is done using channels, which recieve and int, and output and int to a different channel. The supported in messages and their outputs:
```
0 - Returns the fork or philosophers id
1 - Returns the amount of times eaten for philosophers, and times used for forks
2 - The current state of the fork or philosopher. It outputs a 0 or 1, the 0 denotes "not eating" for philosophers and "is not free" for forks. 1 denotes "eating" for philosophers and "is free" for forks
```

Each message is wrapped in individual functions on the struct as well as an aggregate function called GetStatus, for easier access.

## How to use
Simply run 
```
go run .
```
in the cloned repo.
