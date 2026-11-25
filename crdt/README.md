This module contains the core CRDT implementation.

## Components

The main entity is a **board**. This is represented by a `Board` struct. This lives in [board.go](./board.go).

A board consists of a 2D slice of **pixels**. These are represented by a `Pixel` struct. This lives in [pixel.go](./pixel.go).