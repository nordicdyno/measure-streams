# callstat assignment

Hi, thanks for the call just now! As discussed, here is the assignment. Let me know if you have any questions.

Thanks!
Charlotta

## Problem Set 1:

Conference calls are happening in real-time over the Internet and we get the network delay between each participants at regular intervals, typically once every 5 seconds. We need to calculate the median over a set of measurements for providing reliable information about the network delay. Unfortunately the call can be from a few minutes long to 10 hours. We've seen calls that last days. So, we can receive a lot a data and to calculate the median in real-time is an interesting problem.

For this problem set, assuming there are two participants in a call. The analytics pipeline receives the network delay measurements (these are typically, integers), and stores it in a sliding window. You need to implement the getMedian method, which calculates the median of the elements in the sliding window. The sliding window is variable for each test set.

### Test 1:

An example is given below, using a sliding window with length of 3.

The delay measurement arrive one-by-one (iteration) in the following order:
100, 102, 101, 110, 120, 115,

The sliding window should look like this at each iteration:

    100
    100, 102,
    100, 102, 101,
    102, 101, 110,
    101, 110, 120,
    110, 120, 115,

Output: after each iteration (use \r\n delimiter)

    -1
    101
    101
    102
    110
    115

Help:

* If only one element available in the sliding window the answer is -1.
* If n is odd then Median (M) = value of ((n + 1)/2)th item from a sorted array of length n.
* If n is even then Median (M) = value of [((n)/2)th item term + ((n)/2 + 1)th item term ] /2

Hint:

If you have a version, which works, try to improve it by focusing the time complexity.

Attached within are additional test vectors:

* Sliding window size for Test 2 is 100
* Sliding window size for Test 3 is 1000
* Sliding window size for Test 4 is 10000

The input file contains a value on each line.
The output file your program generates should have on each line the median of the sliding window.

### TEST 2:

What is the time-complexity of your solution? If you had more time, how would you improve it?
