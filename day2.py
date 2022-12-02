with open('./input2.txt') as f:
    
    txt = f.readlines()
    total_points = 0

    vals = {
        'A': 1,
        'B': 2,
        'C': 3,
        'X': 1,
        'Y': 2,
        'Z': 3,
    }

    for line in txt:
        line = line.strip('\n')
        points = 0
        points += vals[line[-1]]

        # if the value is same => draw
        if vals[line[-1]] == vals[line[0]]:
            points += 3
        
        # if difference > 1 => not adjecent => smaller number wins (rock beats scissors)
        elif abs(vals[line[-1]] - vals[line[0]]) == 2:
            if vals[line[0]] < vals[line[-1]]:
                points += 0
            else: points += 6

        # adjacent => greater beats smaller
        elif vals[line[-1]] > vals[line[0]]:
            points += 6

        total_points += points

    print(total_points)


# part 2

def custom_add(inp):
    if inp == 3: return 1
    return inp + 1

def custom_sub(inp):
    if inp == 1: return 3
    return inp - 1

with open('./input2.txt') as f:
    
    txt = f.readlines()
    total_points = 0

    vals = {
        'A': 1,
        'B': 2,
        'C': 3,
        'X': 0,
        'Y': 3,
        'Z': 6,
    }

    for line in txt:
        line = line.strip('\n')
        points = 0

        # add points for the outcome
        points += vals[line[-1]]

        # if need to lose
        if line[-1] == 'X':
            points += custom_sub(vals[line[0]])
        # if draw
        elif line[-1] == 'Y':
            points += vals[line[0]]
        elif line[-1] == 'Z':
            points += custom_add(vals[line[0]])

        total_points += points

    print(total_points)
