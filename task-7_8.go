package main

import (
	"errors"
	"fmt"
	"time"
)

type Contact struct {
	ContactID int32
	FullName  string
	Mobile    int64
	Email     string
}

func NewContact(cID int32, fullname, email string, mobile int64) Contact {
	c := Contact{}
	c.ContactID = cID
	c.FullName = fullname
	c.Mobile = mobile
	c.Email = email
	return c
}

func (c *Contact) Update(c1 Contact) {
	c.FullName = c1.FullName
	c.Mobile = c1.Mobile
	c.Email = c1.Email

}

type ContactList struct {
	Contacts []Contact
}

func (cl *ContactList) AddContact(c Contact) {
	cl.Contacts = append(cl.Contacts, c)
}

func (cl *ContactList) DeleteContact(cID int32) error {
	for idx, contact := range cl.Contacts {
		if contact.ContactID == cID {
			cl.Contacts = append(cl.Contacts[:idx], cl.Contacts[idx+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}

func (cl *ContactList) UpdateContact(cID int32, cnew Contact) error {
	for idx, contact := range cl.Contacts {
		if contact.ContactID == cID {
			cl.Contacts[idx].FullName = cnew.FullName
			cl.Contacts[idx].Mobile = cnew.Mobile
			cl.Contacts[idx].Email = cnew.Email
			return nil
		}
	}
	return errors.New("Not updated")
}

func (cl *ContactList) GetContact(cID int32) (Contact, error) {
	for _, contact := range cl.Contacts {
		if contact.ContactID == cID {
			return contact, nil
		}
	}
	return Contact{}, errors.New("Not found")
}

func incID() func() int32 {
	var id int32
	return func() int32 {
		id++
		return id
	}
}

type Task struct {
	TaskID      int32
	Title       string
	CreatedDate time.Time
}

func (t *Task) UPdate(t1 Task) {
	t.Title = t1.Title
	t.CreatedDate = t1.CreatedDate
}

func NewTask(tID int32, title string) Task {
	return Task{TaskID: tID,
		Title:       title,
		CreatedDate: time.Now()}
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) AddTask(t Task) {
	tl.Tasks = append(tl.Tasks, t)
}

func (tl *TaskList) DeleteTask(tID int32) error {
	for idx, task := range tl.Tasks {
		if task.TaskID == tID {
			tl.Tasks = append(tl.Tasks[:idx], tl.Tasks[idx+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}

func (tl *TaskList) UpdateTask(tID int32, tnew Task) error {
	for idx, task := range tl.Tasks {
		if task.TaskID == tID {
			tl.Tasks[idx].Title = tnew.Title
			return nil
		}
	}
	return errors.New("Not updated")
}

func (tl *TaskList) GetTask(tID int32) (Task, error) {
	for _, task := range tl.Tasks {
		if task.TaskID == tID {
			return task, nil
		}
	}
	return Task{}, errors.New("Not found")
}

func main() {
	testContact()
	testTask()
}

func testTask() {
	nextID := incID()

	// Create Tasks
	t1 := NewTask(nextID(), "Spin up new VM")
	t2 := NewTask(nextID(), "Gracefully shutdown")
	fmt.Println(t1)
	fmt.Println(t2)
	// Add tasks to task list
	tl := TaskList{}
	tl.AddTask(t1)
	tl.AddTask(t2)
	fmt.Println(tl)

	// Update task in tasks list
	t3 := NewTask(nextID(), "Destroy Vm")
	tl.UpdateTask(1, t3)
	fmt.Println(tl)

	// Edit task
	t3.UPdate(t1)
	fmt.Println(t3)

	// Get task out of tasks list
	fmt.Println(tl.GetTask(1))

	// Drop task out of tasks list
	tl.DeleteTask(2)
	fmt.Println(tl)

}

func testContact() {

	nextID := incID()

	c1 := Contact{ContactID: nextID(),
		FullName: "Foo",
		Mobile:   998907777777,
		Email:    "foo@example.com"}

	c2 := Contact{ContactID: nextID(),
		FullName: "Bar",
		Mobile:   998909999999,
		Email:    "bar@example.com"}

	c3 := Contact{ContactID: nextID(),
		FullName: "Triple",
		Mobile:   9989011111111,
		Email:    "triple@example.com"}

	// Create contact instance
	c4 := NewContact(nextID(), "Donald Trump", "donald@hotmail.com", 1773333333)
	fmt.Println(c4)

	// Edit contact
	fmt.Println(c3)
	c3.Update(c4)
	fmt.Println(c3)

	cl := ContactList{}

	// Adding contacts to contact list
	fmt.Println(cl)
	cl.AddContact(c1)
	cl.AddContact(c2)
	cl.AddContact(c3)
	fmt.Println(cl)

	// Update contact instance in contacts list
	c5 := NewContact(nextID(), "Emmanuel Macron", "macron@hotmail.com", 488888888)

	if err := cl.UpdateContact(3, c5); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cl)
	}

	// Drop two contacts out of contact list
	cl.DeleteContact(1)
	cl.DeleteContact(2)
	fmt.Println(cl)

}
