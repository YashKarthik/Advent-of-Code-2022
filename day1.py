with open('./input1.txt') as f:
    food_list = f.read().split('\n')
    food_cal = 0
    cals_list = []
    max_cals_sum = 0

    for food in food_list:
        try:
            food_cal += int(food)
        except:
            cals_list.append(food_cal)
            food_cal = 0

    for i in range(3):
        max_cals_sum += max(cals_list)
        cals_list.remove(max(cals_list))


    print(max_cals_sum)
