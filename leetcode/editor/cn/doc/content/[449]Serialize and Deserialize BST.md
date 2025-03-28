<p>Serialization is converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.</p>

<p>Design an algorithm to serialize and deserialize a <b>binary search tree</b>. There is no restriction on how your serialization/deserialization algorithm should work. You need to ensure that a binary search tree can be serialized to a string, and this string can be deserialized to the original tree structure.</p>

<p><b>The encoded string should be as compact as possible.</b></p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<pre><strong>Input:</strong> root = [2,1,3]
<strong>Output:</strong> [2,1,3]
</pre>
<p><strong class="example">Example 2:</strong></p> 
<pre><strong>Input:</strong> root = []
<strong>Output:</strong> []
</pre> 
<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in the tree is in the range <code>[0, 10<sup>4</sup>]</code>.</li> 
 <li><code>0 &lt;= Node.val &lt;= 10<sup>4</sup></code></li> 
 <li>The input tree is <strong>guaranteed</strong> to be a binary search tree.</li> 
</ul>

<div><div>Related Topics</div><div><li>树</li><li>深度优先搜索</li><li>广度优先搜索</li><li>设计</li><li>二叉搜索树</li><li>字符串</li><li>二叉树</li></div></div><br><div><li>👍 555</li><li>👎 0</li></div>