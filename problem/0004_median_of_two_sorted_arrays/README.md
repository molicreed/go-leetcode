# 4.Median of Two Sorted Arrays

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

```
Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

## Answer

### 思路一：
首先将两个数组合并，因为只需要获取到中位数，当合并到第(m+n)/2+1个就可以停止了。
因为前面的数据都是不需要的，可以把储存的数组去掉，使用两个变量记录最后两个值即可。
时间复杂度：O(m+n)