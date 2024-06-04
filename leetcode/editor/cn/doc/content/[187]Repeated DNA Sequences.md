<p>The <strong>DNA sequence</strong> is composed of a series of nucleotides abbreviated as <code>'A'</code>, <code>'C'</code>, <code>'G'</code>, and <code>'T'</code>.</p>

<ul> 
 <li>For example, <code>"ACGAATTCCG"</code> is a <strong>DNA sequence</strong>.</li> 
</ul>

<p>When studying <strong>DNA</strong>, it is useful to identify repeated sequences within the DNA.</p>

<p>Given a string <code>s</code> that represents a <strong>DNA sequence</strong>, return all the <strong><code>10</code>-letter-long</strong> sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in <strong>any order</strong>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<pre><strong>Input:</strong> s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
<strong>Output:</strong> ["AAAAACCCCC","CCCCCAAAAA"]
</pre>
<p><strong class="example">Example 2:</strong></p> 
<pre><strong>Input:</strong> s = "AAAAAAAAAAAAA"
<strong>Output:</strong> ["AAAAAAAAAA"]
</pre> 
<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>s[i]</code> is either <code>'A'</code>, <code>'C'</code>, <code>'G'</code>, or <code>'T'</code>.</li> 
</ul>

<div><div>Related Topics</div><div><li>位运算</li><li>哈希表</li><li>字符串</li><li>滑动窗口</li><li>哈希函数</li><li>滚动哈希</li></div></div><br><div><li>👍 586</li><li>👎 0</li></div>