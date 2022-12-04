f = [l.strip('\n') for l in open('./input4.txt').readlines()]

test = [
    '2-4,6-8',
    '2-3,4-5',
    '5-7,7-9',
    '2-8,3-7',
    '6-6,4-6',
    '2-6,4-8',
]

is_subset = 0
do_intersect = 0

for line in f:
    as_1 = line[:line.find(',')]
    as_2 = line[1 + line.find(','):]
    
    r_1 = range(int(as_1[:as_1.find('-')]), int(as_1[1 + as_1.find('-'):]) + 1)
    r_2 = range(int(as_2[:as_2.find('-')]), int(as_2[1 + as_2.find('-'):]) + 1)

    subset = set(r_1).issubset(set(r_2)) or set(r_2).issubset(set(r_1))
    intersection = set(r_1).intersection(set(r_2))
    if subset: is_subset += 1
    if intersection.__len__() > 0: do_intersect += 1

print(is_subset, do_intersect)
