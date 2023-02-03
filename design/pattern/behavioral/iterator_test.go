package behavioral

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {

	var users []*Node
	for i := 0; i <5;i++{
		users = append(users,&Node{name:fmt.Sprintf("user:%d",i)})
	}

	userCollection := &UserCollection{users: users}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("user name: %s\n", user.name)
	}
}
