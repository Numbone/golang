package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timeStamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID %d | Email %s\n Activity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out = out + fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timeStamp)
	}
	return out
}

func main() {
	t := time.Now()
	users := generateUser(1000)
	wg := &sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go saveUserInfo(user, wg)
	}
	wg.Wait()

	fmt.Println("TIME ELAPSED:", time.Since(t).String())
}

func generateUser(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{action: actions[rand.Intn(len(actions)-1)], timeStamp: time.Now()}
	}
	return logs
}

func saveUserInfo(user User, wg *sync.WaitGroup) error {
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", user.id)
	filename := fmt.Sprintf("logs/user%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}
	wg.Done()
	return err
}
