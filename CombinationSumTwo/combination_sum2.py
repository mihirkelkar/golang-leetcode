"""
Sort the list in descending order.
Then start with the element that is smaller than the target.
Add it to a stack. If the sum is equal to the target. Add it to a result set.
if the sum is greater than the target, pop the number. else keep adding.
"""


def combinationsum(arr, target):
    results = set()
    arr = sorted(arr, reverse=True)
    for index, value in enumerate(arr):
        if value > target:
            pass
        else:
            stack = list()
            stack.append(value)
            if sum(stack) == target:
                results.add(str(stack[0]))
            else:
                cur_index = index + 1
                while cur_index < len(arr):
                    if arr[cur_index] + sum(stack) == target:
                        temp = ",".join([str(ii) for ii in sorted(stack + [arr[cur_index]])])
                        results.add(temp)
                    elif (arr[cur_index] + sum(stack)) < target:
                        stack.append(arr[cur_index])
                    cur_index += 1
    print(results)

combinationsum([2,5,2,1,2], 5)
combinationsum([10,1,2,7,6,1,5], 8)
