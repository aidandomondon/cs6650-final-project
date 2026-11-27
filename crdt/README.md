This module contains the core CRDT implementation.

## Components

Conceptually, this design consists of a board, which is represented as a 2D array of pixels, and which supports one operation: adding a stroke. Adding a stroke involves specifying a list of coordinates in the 2D array of pixels to fill/color in.

These entities are represented by structs
- `Board`: [board.go](./board.go)
- `CoordinatePair`: [utils.go](./utils.go)

Specifying this as an interface, we have:

## Interface

```
+----------------------------------+            +----------------+
| Board                            |      +-----| CoordinatePair |
+----------------------------------+      |     +----------------+
| AddStroke(coor []CoordinatePair) |<-----+     | X int          |
|                                  |            | Y int          |
+----------------------------------+            +----------------+
```