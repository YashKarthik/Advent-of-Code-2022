f = [l.strip('\n') for l in open('./input8.txt').readlines()]

visible = len(f[0])*2 + (len(f) - 2)*2
scenic_scores = []

for row_index in range(len(f)):
    if row_index == 0 or row_index == len(f) - 1:
        continue

    for col_index in range(len(f[row_index])):
        if col_index == 0 or col_index == len(f[row_index]) - 1:
            continue

        trees_left = [int(x) for x in f[row_index][:col_index]]
        trees_right = [int(x) for x in f[row_index][col_index + 1:]]
        trees_above = [int(row[col_index]) for row in f[:row_index]]
        trees_below = [int(row[col_index]) for row in f[row_index + 1:]]
        tree_height = int(f[row_index][col_index])

        if tree_height > max(trees_left) or int(f[row_index][col_index]) > max(trees_below) or int(f[row_index][col_index]) > max(trees_right) or int(f[row_index][col_index]) > max(trees_above):
            visible += 1

print(visible)

def main():
	f = open("./input8.txt","r")
	g = [line[0:len(line)-1] for line in f]
	[n,m] = [len(g),len(g[0])]

	vis= [[1 for i in range(m)] for i in range(n)]

	def valid(row, col):
		return row >= 0 and col >= 0 and row < n and col < m

	def iterate(row, col, d):
		stack = [[10,0]]
		i = 0
		while valid(row,col):
			while stack[-1][0] < int(g[row][col]):
				stack.pop()

			vis [row][col] *=(i-stack[-1][1])
			stack.append([int(g[row][col]), i])
			[row, col, i] = [row+d[0], col+d[1], i+1]
			
	for i in range(n):

		iterate(i,0,[0,1])
		iterate(i,m-1,[0,-1])

	for i in range(m):
		iterate(0,i,[1,0])
		iterate(n-1,i,[-1,0])

	print(max([max(vis[i]) for i in range(n)]))

main()
