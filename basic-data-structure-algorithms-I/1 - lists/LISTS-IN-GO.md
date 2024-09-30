# Container/List Package

Here are the primary types and functions included in the `container/list` package:

## Types

### List
- The main type representing a doubly linked list.

### Element
- Represents a node in the list, containing the value and pointers to the next and previous elements.

## Functions

### Creating a List
- **New()**: 
  - Creates and returns a new empty list.

### List Operations
- **PushBack(v interface{}) *Element**: 
  - Adds a new element with the value `v` at the back of the list and returns the new element.

- **PushFront(v interface{}) *Element**: 
  - Adds a new element with the value `v` at the front of the list and returns the new element.

- **InsertAfter(value interface{}, mark *Element) *Element**: 
  - Inserts a new element with the value `value` after the specified mark element.

- **InsertBefore(value interface{}, mark *Element) *Element**: 
  - Inserts a new element with the value `value` before the specified mark element.

- **Remove(e *Element) interface{}**: 
  - Removes the specified element from the list and returns its value.

- **MoveToFront(e *Element)**: 
  - Moves the specified element to the front of the list.

- **MoveToBack(e *Element)**: 
  - Moves the specified element to the back of the list.

### Accessing Elements
- **Front() *Element**: 
  - Returns the first element of the list or `nil` if the list is empty.

- **Back() *Element**: 
  - Returns the last element of the list or `nil` if the list is empty.

- **Len() int**: 
  - Returns the number of elements in the list.

- **Init() *List**: 
  - Initializes or clears the list, allowing it to be reused.

### Iteration
- **Move functions**:
  - **Next() *Element**: Returns the next element in the list.
  - **Prev() *Element**: Returns the previous element in the list.
