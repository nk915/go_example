package main

import (
	"encoding/json"
	"fmt"
)

// 이벤트 인터페이스 정의
type Event interface {
	Apply(AggregateRoot)
}

// AggregateRoot 인터페이스 정의
type AggregateRoot interface {
	ApplyEvent(Event)
}

// 사용자가 생성된 이벤트 정의
type UserCreated struct {
	ID   string
	Name string
}

func (uc *UserCreated) Apply(ar AggregateRoot) {
	u := ar.(*User)
	u.ID = uc.ID
	u.Name = uc.Name
}

// 사용자 이름이 변경된 이벤트 정의
type UserNameUpdated struct {
	Name string
}

func (unu *UserNameUpdated) Apply(ar AggregateRoot) {
	u := ar.(*User)
	u.Name = unu.Name
}

// 사용자 Aggregate Root 정의
type User struct {
	ID    string
	Name  string
	Events []Event // 이벤트 저장
}

func NewUser(id string, name string) *User {
	u := &User{}
	u.apply(&UserCreated{id, name})
	return u
}

func (u *User) UpdateName(name string) {
	u.apply(&UserNameUpdated{name})
}

func (u *User) apply(e Event) {
	e.Apply(u)
	u.Events = append(u.Events, e)
}

func (u *User) ApplyEvent(e Event) {
	u.apply(e)
}

func (u *User) Serialize() ([]byte, error) {
	return json.Marshal(u.Events)
}

func main() {
	// 사용자 생성
	user := NewUser("1", "John")

	// 사용자 이름 변경
	user.UpdateName("Jane")

	// 이벤트 serialize
	data, err := user.Serialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
