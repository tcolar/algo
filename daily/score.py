def num_cumbs(finalScore, playScores):
    numForScore = [[1]+[0] * finalScore for _ in playScores]
    print(numForScore)
    for i in range(len(playScores)):
        for j in range(1, finalScore+1):
            without = (numForScore[i-1][j] if i >= 1 else 0)
            withit = (numForScore[i][j-playScores[i]]
                      if j >= playScores[i] else 0)
            print(i, j, withit, without)
        numForScore[i][j] = (without+withit)
    return numForScore[-1][-1]


print(num_cumbs(12, [2, 3, 7]))
