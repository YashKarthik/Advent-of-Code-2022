f = [l.strip('\n') for l in open('./input3.txt').readlines()]

priorities = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'

sum_priorities = 0
for index in range(0, len(f), 3):
    elf_1 = set(f[index])
    elf_2 = set(f[index + 1])
    elf_3 = set(f[index + 2])

    badge = elf_1.intersection(elf_2).intersection(elf_3)
    sum_priorities += 1 + priorities.find(badge.pop())

print(sum_priorities)
