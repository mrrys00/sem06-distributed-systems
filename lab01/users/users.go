package users

import (
	"lab01/logs"
)

func (u *User) exit(usersList *Users) {
	logs.LogTrace("Go to remove user")
	usersList.RemoveUser(u)
}

func (us *Users) AddUser(u *User) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	us.users = append(us.users, u)
	logs.LogDebug(nil, "New user added")
}

func (us *Users) RemoveUser(u *User) {
	us.userMutex.Lock()
	defer us.userMutex.Unlock()

	for i, user := range us.users {
		if user != u {
			continue
		}

		us.users = append(us.users[0:i], us.users[i:]...)
	}
}
