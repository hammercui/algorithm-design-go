/**
 * Description:调表
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2022/08/17
 * Time: 14:41
 * Mail: hammercui@163.com
 * 每个节点多出随机个数的前进指针，
 * 插入与搜索更快，更节省空间
 */
package list

import "math/rand"

const DEFAULT_SKIP_LIST_P = 0.25
// SkipNode is an Element of a skipList.
type SkipNode struct {
	Value interface{}
	forward []*SkipNode
	currentLevel int // current node max level
	Score float64 //order set key
}

type SkipList struct {
	header *SkipNode // header is a dummy element
	length int // current skipList length，header not included
	maxLevel int // current skipList level，header not included
	p float32 //ratio for level
}

//p range(0,1)
func CreateSkipList(maxLevel int,p float32) *SkipList  {
	s := new(SkipList)
	s.header = new(SkipNode)
	s.maxLevel = maxLevel
	s.p = p

	//first column is header,set interface{},
	s.header.currentLevel = maxLevel
	s.header.forward = make([]*SkipNode,maxLevel)
	s.header.Value = new(interface{})
	s.length = 0
	return s
}

//generate node level by random
func (s *SkipList) randomLevel() int  {
	var level = 1 //default level
	for rand.Float32() < s.p && level < s.maxLevel {
		level++
	}
	return level
}

func (s *SkipList) newNode(value interface{},score float64,level int) *SkipNode  {
	return &SkipNode{
		Value:        value,
		forward:      make([]*SkipNode,s.maxLevel),
		currentLevel: level,
		Score:        score,
	}
}
func (s *SkipList) Insert(value interface{},score float64) (bool,error)  {
	level := s.randomLevel()
	newNode := s.newNode(value,score,level)
	p := s.header

	for l := s.maxLevel; l > 0; l-- {
		index := l - 1
		for p.forward[index] != nil && p.forward[index].Score < score    {
			p = p.forward[index]
		}
		//find the last Node which match user defined IsLess() condition in l level
		//insert new Node after the node
		// 在第l层找到应该插入的最佳位置，然后在该位置插入新的节点
		// level层以下都有这个节点
		if l <= level {
			//score相等，更新Value返回
			if p.forward[index] != nil && p.forward[index].Score == score {
				p.forward[index].Value = value
				return true,nil
			}
			// 相当于链表的插入操作，node旧的next交接给新node
			newNode.forward[index] = p.forward[index]
			// 当前node的next节点替换为新node
			p.forward[index] = newNode
		}
	}
	s.length++
	return true,nil
}


func (s *SkipList) Search(score float64) *SkipNode {
	p := s.header
	for l := s.maxLevel	; l >0 ; l-- {
		for p.forward[l-1] != nil && p.forward[l-1].Score < score{
			p = p.forward[l-1]
		}
	}
	//leve1位基础层，包含所有值
	p = p.forward[0]
	if p == nil || p.Score != score {
		return nil
	}
	return p
}

func (s *SkipList) Delete(score float64) (bool)  {
	//1先找到节点以及节点在每一层的前缀
	//缓存节点的前置节点
	var preNodes []*SkipNode = make([]*SkipNode,s.maxLevel)
	p := s.header
	for l := s.maxLevel; l >0 ; l-- {
		for p.forward[l-1] != nil && p.forward[l-1].Score < score {
			p = p.forward[l-1]
		}
		//每一层前置
		preNodes[l-1] = p
	}
	p = p.forward[0]
	//删除的节点不存在
	if p == nil || p.Score < score ||  score < p.Score{
		return false
	}

	//更新每一层前置的forward
	for l := p.currentLevel; l >0 ; l-- {
		index := l-1
		preNodes[index].forward[index] = p.forward[index]
	}
	s.length --
	return true
}






