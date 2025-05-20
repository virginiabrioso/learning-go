# Container/List Package

Here are the primary types and functions included in the `container/list` package:

## Types

- List
  - The main type representing a doubly linked list.

- Element
  - Represents a node in the list, containing the value and pointers to the next and previous elements.

## Functions

### Creating a List

- `listName := list.New()`: 
  - Creates and returns a new empty list.

### List Operations

- `PushBack(x)`: 
  - Adds a new element with the value `x` at the back of the list and returns the new element.

- `PushFront(x)`:
  - Adds a new element with the value `x` at the front of the list and returns the new element.

- `InsertAfter(x, element)`:
  - Inserts a new element with the value `x` after the specified mark element.

- `InsertBefore(x, element)`:
  - Inserts a new element with the value `x` before the specified mark element.

- `Remove(element)`:
  - Removes the specified element from the list and returns its value.

- `MoveToFront(element)`:
  - Moves the specified element to the front of the list.

- `MoveToBack(element)`:
  - Moves the specified element to the back of the list.

### Accessing Elements

- `listName.Front()`:
  - Returns the first element of the list or `nil` if the list is empty.
  - Can use `listName.Front().Value` to access element value

- `listName.Back()`:
  - Returns the last element of the list or `nil` if the list is empty.
  - Can use `listName.Back().Value` to access element value

- `listName.Len()`:
  - Returns the number of elements in the list.

- `listName.Init()`:
  - Initializes or clears the list, allowing it to be reused.

### Iteration

#### Move functions:
  - `listName.Next()`: 
    - Returns the next element in the list.

  - `listName.Prev()`: 
    - Returns the previous element in the list.
