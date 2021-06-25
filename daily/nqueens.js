function nQueens(n) {
    var board = new Array(n);
    for (var i = 0; i < n; i++) {
        board[i].fill(0);
    }
    console.log(board);
    if (n === 1)
        return [[1]];
    for (var i = 0; i < n; i++) { }
    var results = [];
    return results;
}
nQueens(4);
