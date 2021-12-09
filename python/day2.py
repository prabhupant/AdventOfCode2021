# Reading the file

with open("../inputs/day1.txt", encoding="utf-8") as f:
    arr = list(map(int, f.readlines()))

prev = arr[0] + arr[1] + arr[2]
count = 0

for i in range(1, len(arr)-2):
    curr = arr[i] + arr[i+1] + arr[i+2]

    if curr > prev:    
        count = count + 1

    prev = curr

print(count)

