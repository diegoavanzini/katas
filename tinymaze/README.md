# tiny-maze

Alice found herself very tiny and wandering around Wonderland.  Even
the grass around her seemed like a maze.

This is a tiny maze solver.

A maze is represented by a matrix

```clojure
[[:S 0 1]
 [1  0 1]
 [1  0 :E]]
```

- _S_ : start of the maze
- _E_ : end of the maze
- _1_ : This is a wall that you cannot pass through
- _0_ : A free space that you can move through.

The goal is the get to the end of the maze.  A solved maze will have a
_:x_ in the start, the path, and the end of the maze, like this.

```clojure
[[:x :x 1]
 [1  :x 1]
 [1  :x :x]]
```


