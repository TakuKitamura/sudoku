
from itertools import product

def solve_sudoku(size, grid):
    R, C = size
    N = R * C
    X = ([("rc", rc) for rc in product(range(N), range(N))] +
         [("rn", rn) for rn in product(range(N), range(1, N + 1))] +
         [("cn", cn) for cn in product(range(N), range(1, N + 1))] +
         [("bn", bn) for bn in product(range(N), range(1, N + 1))])
    Y = dict()
    for r, c, n in product(range(N), range(N), range(1, N + 1)):
        b = (r // R) * R + (c // C) # Box number
        Y[(r, c, n)] = [
            ("rc", (r, c)),
            ("rn", (r, n)),
            ("cn", (c, n)),
            ("bn", (b, n))]

    X, Y = exact_cover(X, Y)
    for i, row in enumerate(grid):
        for j, n in enumerate(row):
            if n:
                select(X, Y, (i, j, n))

    b = solve(X, Y, [])
    for (r, c, n) in b:
        grid[r][c] = n
    return grid

def exact_cover(X, Y):
    X = {j: set() for j in X}
    for i, row in Y.items():
        for j in row:
            X[j].add(i)
    return X, Y

def solve(X, Y, solution):
    while True:
        if X == {}:
            return solution
        else:
            c = min(X, key=lambda c: len(X[c]))
            r = list(X[c])[0]
            solution.append(r)
            select(X, Y, r)

def select(X, Y, r):
    cols = []
    for j in Y[r]:
        for i in X[j]:
            for k in Y[i]:
                if k != j:
                    X[k].remove(i)
        cols.append(X.pop(j))
    return cols

def deselect(X, Y, r, cols):
    for j in reversed(Y[r]):
        X[j] = cols.pop()
        for i in X[j]:
            for k in Y[i]:
                if k != j:
                    X[k].add(i)

if __name__ == "__main__":
    grid = [
    [0,6,1,0,0,7,0,0,3],
    [0,9,2,0,0,3,0,0,0],
    [0,0,0,0,0,0,0,0,0],
    [0,0,8,5,3,0,0,0,0],
    [0,0,0,0,0,0,5,0,4],
    [5,0,0,0,0,8,0,0,0],
    [0,4,0,0,0,0,0,0,1],
    [0,0,0,1,6,0,8,0,0],
    [6,0,0,0,0,0,0,0,0]
    ]
    a = solve_sudoku((3, 3), grid)
    print(a)
