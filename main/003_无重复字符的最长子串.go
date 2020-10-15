package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 */

/*
//对应官方答案思想里面的方案2,3.
// 滑动串口解法1
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int i = 0,j = 0, n = 0, ans= 0;
        set<char> cont;
        n = s.size();
        set<char>::iterator iter;

        while(i< n && j<n){
            if( (iter=cont.find(s[j])) != cont.end() ){
                cont.erase(s[i++]);
            }
            else {

                cont.insert(s[j++]);
                ans = max(ans, j-i);
            }

        }
        return ans;
    }
};

//滑动窗口解法2(优化了方法1)
    int lengthOfLongestSubstring3(string s) {
        int n = s.length();
        map<char, int> Map_tmp;
        map<char, int>::iterator iter;
        int i= 0, ans = 0;

        for(int j = 0; j<n; j++){
           if( (iter = Map_tmp.find(s[j])) != Map_tmp.end() ){
               i = max(i, iter->second);
           }
            ans = max(ans, j-i+1);
            Map_tmp[(s[j])] = j+1;
        }
        return ans;
    }


*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	i := 0
	j := 0
	ans := 0
	m := make(map[uint8]uint8)

	for j < n {
		// try to extend the range [i, j]
		_, ok := m[s[j]]
		if ok {
			delete(m, s[i])
			i += 1
			//startPos = int(math.Max(float64(i), float64(startPos))) //不重复的起始位置
		} else {
			ans = int(math.Max(float64(ans), float64(j-i+1))) //字符串长度
			m[s[j]] = s[j]
			j += 1
		}
	}

	////打印
	//type kv struct {
	//	Key   string
	//	Value int
	//}
	//
	//var ss []kv
	//for k, v := range m {
	//	ss = append(ss, kv{string(k), int(v)})
	//}
	//
	//sort.Slice(ss, func(i, j int) bool {
	//	//return ss[i].Value > ss[j].Value // 降序
	//	return ss[i].Value < ss[j].Value // 升序
	//})
	//
	//for _, kv := range ss {
	//	fmt.Printf("%v, %d\n", kv.Key, kv.Value)
	//}

	fmt.Printf("ans= %+v\n", ans)
	return ans
}

func lengthOfLongestSubstring3(s string) int {
	n := len(s)
	Map_tmp := make(map[uint8]int)
	i, ans := 0, 0

	for j := 0; j < n; j++ {
		if _, ok := Map_tmp[s[j]]; ok {
			i = Max(i, Map_tmp[s[j]])
		}
		ans = Max(ans, j-i+1)
		Map_tmp[s[j]] = j + 1
	}

	//type kv struct {
	//	Key   string
	//	Value int
	//}
	//
	//var ss []kv
	//for k, v := range Map_tmp {
	//	ss = append(ss, kv{string(k), v})
	//}
	//
	//sort.Slice(ss, func(i, j int) bool {
	//	//return ss[i].Value > ss[j].Value // 降序
	//	return ss[i].Value < ss[j].Value // 升序
	//})
	//
	//for _, kv := range ss {
	//	fmt.Printf("%v, %d\n", kv.Key, kv.Value)
	//}

	return ans
}

func zichuan(s string) int {
	mp := make(map[uint8]uint8)
	ln := len(s)

	ans := 0
	i := 0
	j := 0
	for i < ln && j < ln {
		_, ok := mp[s[j]]
		if ok {
			delete(mp, s[i])
			i++
		} else {
			ans = Max(ans, j-i+1)
			mp[s[j]] = s[j]
			j++
		}
	}

	//fmt.Printf("ans:%v,st:%v,et:%v\n", ans, st, et)
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	x := "pwwkew"
	//x := "bbbbb"
	//x := "abcabcbb"
	//lengthOfLongestSubstring(x)
	fmt.Println(zichuan(x))
	//fmt.Println(lengthOfLongestSubstring1(x))
	//fmt.Println(lengthOfLongestSubstring3(x))
}
