package ai

import (
    "../grid"
)

type AI struct {
    Grid *grid.Grid
    ResultChan chan *grid.Grid
    MoveChan chan int
    BestScore int
    BestMove int
}

func (ai *AI) Run(grid *grid.Grid, ch chan *grid.Grid, mv chan int) {
    ai.Grid = grid
    ai.ResultChan = ch
    ai.MoveChan = mv

    depth := 0
    for {
        ai.BestScore, ai.BestMove = ai.alphaBeta(grid, depth, -10000, 10000, 0, 0)
        depth++
    }

    return 
}

func (ai *AI) alphaBeta(grid *grid.Grid, depth int, alpha int, beta int, positions int, cutoffs int) (bestScore int, bestMove int) {
    directions := []int{1,2,3,4}
    for direction := range directions {
        _, ch, mv := ai.Grid.Clone()
        select {
        case <- ch:
            positions++
            if depth == 0 {
                return 0, direction
            }
        }
        mv <- direction
        // go ai.alphaBeta(newGrid, depth-1, alpha, beta, position, cutoffs)
    }

    return 0, 0
}