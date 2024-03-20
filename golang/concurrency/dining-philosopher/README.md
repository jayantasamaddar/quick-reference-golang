# The Dining Philosopher Problem

- A classic computer science problem introduced by Djikstra in 1965.
- Five philosophers live in a house together, and they always dine together at the same table, always sitting in the same place.
- They always eat a special kind of spaghetti which requires two forks.
- There are two forks next to each plate, that means no two neighbours can be eating at the same time.

---

# Solving the Problem

- We will need a Philosopher struct denoting each philosopher. Each philosopher has a left spoon and a right spoon. When one philosopher uses the spoons, these particular spoons will get locked.

---
