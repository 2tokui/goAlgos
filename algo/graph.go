package algo

// graphs represents relationships between nodes/vertex with edges
// definition: A graph G=(V,E) is a set of vertices V
// and edges E where each edge (u,v) is a connection between verticies. u, v ∈ V

// Terminology: Neighbors - vertices u and v are neighbors if an edge (u,v) connects them
//              Degree    - number of edges connected to v
//              Path      - sequence of vertices connected by edges

// To organize a graph in the computer we can use an Adjacency Matrix
//   0 1 2 3
// 0 0 1 1 1
// 1 0 0 0 1
// 2 1 0 0 1 I don't even know how this looks like, I just threw random numbers
// 3 1 1 0 0
// We can also use an Adjacency List
// 0 -> 1, 2, 3
// 1 -> 3
// 2 -> 1, 3
// 3 -> 0, 1

