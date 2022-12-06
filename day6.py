f = open('./input6.txt').read()

for i in range(len(f)):
    marker_list = f[i:i+4]
    marker_set = set(marker_list)

    message_list = f[i:i+14]
    message_set = set(message_list)

    if len(marker_list) == len(marker_set): print('Marker end:', i+4, marker_list);
    if len(message_list) == len(message_set): print('Message marker end:', i+14, message_list); break
