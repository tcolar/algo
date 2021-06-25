/*
2 pts safety
3 pts fd goal
7 pts TD

write code that finds the number of combinations for a given score
*/

function findCombis(score) {
  points = Array(score + 1).fill(0);
  points[0] = 1;
  for (i = 2; i < points.length; i++) {
    points[i] += points[i - 2];
  }
  for (i = 3; i < points.length; i++) {
    points[i] += points[i - 3];
  }
  for (i = 7; i < points.length; i++) {
    points[i] += points[i - 7];
  }
  console.log(points);
  return points[score];
}

console.log(findCombis(12));
