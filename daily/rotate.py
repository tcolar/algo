# Note: Try to solve this task in-place (with O(1) additional memory), since this is what you'll be asked to do during an interview.

# You are given an n x n 2D matrix that represents an image. Rotate the image by 90 degrees (clockwise).

# Example

# For

# a = [[1, 2, 3],
#      [4, 5, 6],
#      [7, 8, 9]]
# the output should be

# rotateImage(a) =
#     [[7, 4, 1],
#      [8, 5, 2],
#      [9, 6, 3]]

def rotateImage(a):
    cols = len(a)
    rows = len(a[0])
    for x in range(0, cols//2):
        for y in range(x, cols-x-1):
            temp = a[x][y]
            a[x][y] = a[cols - 1 - y][x]
            a[cols - 1 - y][x] = a[cols - 1 - x][cols - 1 - y]
            a[cols - 1 - x][cols - 1 - y] = a[y][cols - 1 - x]
            a[y][cols - 1 - x] = temp


def translate(size, x, y):
    return y, size-1-x


a = [[1, 2, 3],
     [4, 5, 6],
     [7, 8, 9]]
rotateImage(a)
print(a)

b = [[6, 8, 7, 6, 10],
     [8, 9, 6, 10, 9],
     [6, 7, 3, 2, 6],
     [8, 9, 8, 9, 3],
     [2, 9, 2, 7, 7]]

rotateImage(b)
print(b)
