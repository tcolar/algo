function climbingStaircase(n, k) {
  // k is max steps at once, n is total steps
  ways = {};
  for (i = 1; i <= n; i++) {
    ways[i] = [];
    // indirect ways
    for (j = 1; j <= k; j++) {
      w = ways[i - j];
      if (!w) continue;
      for (way of w) {
        ways[i].push([j, ...way]);
      }
    }
    if (i <= k) {
      ways[i].push([i]); // direct way to get there
    }
    //console.log(i, ways[i])
  }
  return ways[n] || [[]];
}

/*

You need to climb a staircase that has n steps, and you decide to get some extra exercise by jumping up the steps. 
You can cover at most k steps in a single jump. Return all the possible sequences of jumps that you could take to climb the staircase, sorted.

Example

For n = 4 and k = 2, the output should be

climbingStaircase(n, k) =
[[1, 1, 1, 1],
 [1, 1, 2],
 [1, 2, 1],
 [2, 1, 1],
 [2, 2]]
There are 4 steps in the staircase, and you can jump up 2 or fewer steps at a time. There are 5 potential sequences in which you jump up the stairs 
either 2 or 1 at a time.
*/
