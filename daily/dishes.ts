function groupingDishes(dishes: string[][]): string[][] {
  const ingMap = new Map<string, string[]>();
  for (const dish of dishes) {
    const name = dish[0];
    for (const ingredient of dish.slice(1)) {
      if (!ingMap.has(ingredient)) {
        ingMap.set(ingredient, [name]);
      } else {
        ingMap.get(ingredient).push(name);
      }
    }
  }
  const result: string[][] = [];
  for (const [name, ingredients] of ingMap) {
    if (ingredients.length < 2) {
      continue;
    }
    result.push([name, ...ingredients.sort()]);
  }
  return result.sort((a, b) => (a[0] < b[0] ? -1 : a[0] === b[0] ? 0 : 1));
}
/*
You are given a list dishes, where each element consists of a list of strings beginning with the name of the dish, followed by all the ingredients used in preparing it. You want to group the dishes by ingredients, so that for each ingredient you'll be able to find all the dishes that contain it (if there are at least 2 such dishes).

Return an array where each element is a list beginning with the ingredient name, followed by the names of all the dishes that contain this ingredient. The dishes inside each list should be sorted lexicographically, and the result array should be sorted lexicographically by the names of the ingredients.
*/
