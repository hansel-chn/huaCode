<p>An array is <strong>squareful</strong> if the sum of every pair of adjacent elements is a <strong>perfect square</strong>.</p>

<p>Given an integer array nums, return <em>the number of permutations of </em><code>nums</code><em> that are <strong>squareful</strong></em>.</p>

<p>Two permutations <code>perm1</code> and <code>perm2</code> are different if there is some index <code>i</code> such that <code>perm1[i] != perm2[i]</code>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,17,8]
<strong>Output:</strong> 2
<strong>Explanation:</strong> [1,8,17] and [17,8,1] are the valid permutations.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,2,2]
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>1 &lt;= nums.length &lt;= 12</code></li> 
 <li><code>0 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li> 
</ul>

<div><div>Related Topics</div><div><li>位运算</li><li>数组</li><li>哈希表</li><li>数学</li><li>动态规划</li><li>回溯</li><li>状态压缩</li></div></div><br><div><li>👍 124</li><li>👎 0</li></div>