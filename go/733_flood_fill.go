/*
You are given an image represented by an m x n grid of integers image, where image[i][j] represents the pixel value of the image. You are also given three integers sr, sc, and color. Your task is to perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill:

    1. Begin with the starting pixel and change its color to color.
    2. Perform the same process for each pixel that is directly adjacent (pixels that share a side with the original pixel, either horizontally or vertically) and shares the same color as the starting pixel.
    3. Keep repeating this process by checking neighboring pixels of the updated pixels and modifying their color **if it matches the original color of the starting pixel**.
    4. The process stops when there are no more adjacent pixels of the original color to update.

Return the modified image after performing the flood fill.

Example 1:

Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2

Output: [[2,2,2],[2,2,0],[2,0,1]]
*/

package main


func floodFill(image [][]int, sr, sc, color int) [][]int {
	// image[sr][sc] == (1, 1)
	// [ 1 1 1 ]
	// [ 1 1 0 ]
	// [ 1 0 0 ]

	// NOTE: we fill a color **iff it matches the original color of the starting pixel**

	/* Init */
	startingColor := image[sr][sc]
	prevRow := sr-1
	prevCol := sc-1
	nextRow := sr+1
	nextCol := sc+1
	numRows := len(image)
	numCols := len(image[0])

	// 1. fill in first color
	image[sr][sc] = color

	// back up one (vertical)
	if prevRow >= 0 && image[prevRow][sc] == startingColor {
		image[prevRow][sc] = color
		floodFill(image, prevRow, sc, color)
	}
	// back left one (horizontal)
	if prevCol >= 0 && image[sr][prevCol] == startingColor {
		image[sr][prevCol] = color
		floodFill(image, sr, prevCol, color)
	}

	// up one (vertical)
	if prevRow < numRows && image[nextRow][sc] == startingColor {
		image[nextRow][sc] = color
		floodFill(image, nextRow, sc, color)
	}
	// left one (horizontal)
	if prevCol < numCols && image[sr][nextCol] == startingColor {
		image[sr][nextCol] = color
		floodFill(image, sr, nextCol, color)
	}
	
	return image
}

