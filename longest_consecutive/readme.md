# Overview
An integer array is given. Find the length of the longest consecutive sequence in it. Let’s see an example to understand it

## Example 1

Input: [4,7,3,8,2,1]
Output: 4
Reason: The longest consecutive sequence is [1,2,3,4]

## Example 2

Input: [4,7,3,8,2,1,9,24,10,11]
Output: 5
Reason: The longest consecutive sequence is [7,8,9,10,11]

# Solution
The naive idea is to sort the array and return the longest consecutive sequence. But we can do this in O(n) time. Idea is to use a hash here. Below is the idea

- Store each of the numbers in the hash.
- Then starting with each number, find the length of the longest consecutive sequence beginning at that number. So if a number let’s say is n. We try to find n+1, n+2… in the hash and get the length of the longest consecutive sequence  starting from that number
- If the length find in step 2 is greater than the previous longest consecutive sequence found, then update the longest consecutive sequence length