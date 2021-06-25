function climbingStairs(n: number): number {
  const ways = {
    1: 1,
    2: 2,
  };
  for (let i = 3; i <= n; i++) {
    ways[i] = (ways[i - 1] || 0) + (ways[i - 2] || 0);
  }
  // console.log(ways)
  return ways[n];
}
/*
You are climbing a staircase that has n steps. You can take the steps either 1 or 2 at a time. 
Calculate how many distinct ways you can climb to the top of the staircase.
*/
