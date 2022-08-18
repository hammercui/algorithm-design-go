package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipList(t *testing.T) {
	a := assert.New(t)
	//create
	numList := CreateSkipList(5,0.5)
	a.Equal(numList.maxLevel,5)

	//insert
	length := 10.0
	for i := 1.0; i < length + 1; i++ {
		_,err :=numList.Insert(i,i);
		a.Nil(err,"insert fail")
	}
	a.Equal(numList.length,int(length))

	//traverse
	skipListTraverse(numList)

	//search
	result := numList.Search(9)
	a.Equal(result.Score,9.0)
	fmt.Printf("search result:%+v",result)

	//delete
	delRes :=numList.Delete(7)
	a.Equal(delRes,true)
	fmt.Println(" \n after delete 7：")
	skipListTraverse(numList)
}

func skipListTraverse(s *SkipList)  {
	p := s.header
	for l := s.maxLevel; l > 0 ; l-- {
		fmt.Printf("level%d:",l)
		for  {
			if p != nil {
				if p.Score > 0 {
					fmt.Printf("%.f", p.Score)
				}
				if p.forward[l-1] != nil {
					fmt.Print("-->")
				}
				p = p.forward[l-1]
			}else{
				break
			}
		}
		p = s.header
		fmt.Print("\n")
	}
}

