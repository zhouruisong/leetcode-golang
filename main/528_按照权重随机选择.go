package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

/*
给定一个正整数数组 w ，其中 w[i] 代表下标 i 的权重（下标从 0 开始），请写一个函数 pickIndex ，它可以随机地获取下标 i，选取下标 i 的概率与 w[i] 成正比。
例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。
也就是说，选取下标 i 的概率为 w[i] / sum(w) 。
*/

type Solution struct {
	nums []int
	sum  int
}

func Constructor5(w []int) Solution {
	sum := 0
	weightSums := make([]int, len(w))

	for i := 0; i < len(w); i++ {
		sum += w[i]
		weightSums[i] = sum
	}
	return Solution{weightSums, sum}
}

//找到小于等于随机值的最右侧值,取索引
func (this *Solution) PickIndex() int {
	target := rand.Intn(this.sum)
	left, right := 0, len(this.nums)
	for left < right {
		mid := left + (right-left)/2
		if target >= this.nums[mid] {
			left = mid + 1
		} else if target < this.nums[mid] {
			right = mid
		}
	}

	return left
}

//平滑的加权轮询,效果比较好
type LoadBalance interface {
	//选择一个后端Server
	//参数remove是需要排除选择的后端Server
	Select(remove []string) *Server
	//更新可用Server列表
	UpdateServers(servers []*Server)
}

type Server struct {
	Host   string //主机地址
	Name   string //主机名
	Weight int    //权重
	Online bool   //主机是否在线
}

type Weighted struct {
	Server          *Server
	Weight          int
	CurrentWeight   int //服务器目前的权重。一开始为0，之后会动态调整
	EffectiveWeight int
}

func (this *Weighted) String() string {
	return fmt.Sprintf("[%s][%d]", this.Server.Host, this.Weight)
}

type LoadBalanceWeightedRoundRobin struct {
	servers  []*Server
	weighted []*Weighted
}

func NewLoadBalanceWeightedRoundRobin(servers []*Server) *LoadBalanceWeightedRoundRobin {
	l := &LoadBalanceWeightedRoundRobin{}
	l.UpdateServers(servers)
	return l
}

func (this *LoadBalanceWeightedRoundRobin) UpdateServers(servers []*Server) {
	if len(this.servers) == len(servers) {
		for _, s := range servers {
			isEqual := false
			for _, old := range this.servers {
				if s.Host == old.Host && s.Weight == old.Weight && s.Online == old.Online {
					isEqual = true
					break
				}
			}
			if isEqual == false {
				goto build
			}
		}
		return
	}

build:
	weighted := make([]*Weighted, 0)
	for _, v := range servers {
		if v.Online == true {
			w := &Weighted{
				Server:          v,
				Weight:          v.Weight,
				CurrentWeight:   0,
				EffectiveWeight: v.Weight,
			}
			weighted = append(weighted, w)
		}
	}
	this.weighted = weighted
	this.servers = servers
}

func (this *LoadBalanceWeightedRoundRobin) Select(remove []string) *Server {
	if len(this.weighted) == 0 {
		return nil
	}
	w := this.nextWeighted(this.weighted, remove)
	if w == nil {
		return nil
	}
	return w.Server
}

func (this *LoadBalanceWeightedRoundRobin) nextWeighted(servers []*Weighted, remove []string) (best *Weighted) {
	total := 0
	for i := 0; i < len(servers); i++ {
		w := servers[i]
		if w == nil {
			continue
		}
		isFind := false
		for _, v := range remove {
			if v == w.Server.Host {
				isFind = true
			}
		}
		if isFind == true {
			continue
		}

		w.CurrentWeight += w.EffectiveWeight
		total += w.EffectiveWeight
		if w.EffectiveWeight < w.Weight {
			w.EffectiveWeight++
		}

		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}
	}
	if best == nil {
		return nil
	}
	best.CurrentWeight -= total
	return best
}

func (this *LoadBalanceWeightedRoundRobin) String() string {
	return "WeightedRoundRobin"
}

func main() {

	count := make([]int, 4)
	servers := make([]*Server, 0)
	servers = append(servers, &Server{Host: "0", Weight: 10, Online: true})
	servers = append(servers, &Server{Host: "1", Weight: 20, Online: true})
	servers = append(servers, &Server{Host: "2", Weight: 30, Online: true})
	servers = append(servers, &Server{Host: "3", Weight: 40, Online: true})
	lb := NewLoadBalanceWeightedRoundRobin(servers)

	for i := 0; i < 100000; i++ {
		s := lb.Select(nil)
		id, _ := strconv.Atoi(s.Host)
		count[id]++
	}
	fmt.Println(count)

	//效果没有前面的好
	count2 := make([]int, 4)
	w := Constructor5([]int{10, 20, 30, 40})

	for i := 0; i < 100000; i++ {
		id := w.PickIndex()
		count2[id]++
	}
	fmt.Println(count2)
}
