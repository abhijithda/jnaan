# Linked List containing a loop

In a single linked list of 'n' nodes, the last node `Nn` instead of pointing to 'NULL' is pointing to someother node `Nx` of the same list creating a loop.

`N1 -> N2 -> N3 - - -> Nx - - -> Nn -> Nx`

**Constraints:**

1. Data of the nodes cannot be altered,
1. Solution should be memory optimized as it needs to run on a system that has limited memory.
1. The system has unlimited processing power. Or assume you've unlimited time.

**Find the last node of the linked list** so that it's next pointer could be set to 'NULL'. I.e., return address of `Nn`.

`N1 -> N2 -> N3 - - -> Nx - - -> Nn -> NULL`
