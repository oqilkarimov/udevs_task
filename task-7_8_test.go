package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var nextID = incID()

// Tasks unit testing
var tests = []struct {
	t    Task
	want int
}{
	{Task{TaskID: nextID(), Title: "Task 1", CreatedDate: time.Now()}, 1},
	{Task{TaskID: nextID(), Title: "Task 2", CreatedDate: time.Now()}, 2},
	{Task{TaskID: nextID(), Title: "Task 3", CreatedDate: time.Now()}, 3},
	{Task{TaskID: nextID(), Title: "Task 4", CreatedDate: time.Now()}, 4},
	{Task{TaskID: nextID(), Title: "Task 5", CreatedDate: time.Now()}, 5},
	{Task{TaskID: nextID(), Title: "Task 6", CreatedDate: time.Now()}, 6},
}

func TestTaskCreate(t *testing.T) {
	id := nextID()
	task := Task{TaskID: id, Title: "Task 1"}
	if task.Title != "Task 1" || task.TaskID != id {
		t.Errorf("Task(%v, %v, %v)", task.TaskID, task.Title, task.CreatedDate)
	}
}

func TestTaskUpdate(t *testing.T) {
	id1 := nextID()
	id2 := nextID()
	t1 := Task{
		TaskID:      id1,
		Title:       "Task 1",
		CreatedDate: time.Now()}
	t2 := Task{
		TaskID:      id2,
		Title:       "Task 2",
		CreatedDate: time.Now()}
	t1.UPdate(t2)
	if t1.TaskID == t2.TaskID || t1.Title != t2.Title || t1.CreatedDate != t2.CreatedDate {
		t.Errorf("Task(%v %v %v)", t1.TaskID, t1.Title, t1.CreatedDate)
	}
}

func TestTaskListAddTask(t *testing.T) {
	tl := TaskList{}
	// testname := "Adding all sample tasks to tasklist"
	for _, tt := range tests {
		testname := fmt.Sprintf("TaskLIst.len: %v,TaskID: %v", tt.want, tt.t.TaskID)
		t.Run(testname, func(t *testing.T) {
			tl.AddTask(tt.t)
			if len(tl.Tasks) != tt.want {
				t.Errorf("got TaskList.len: %d, want: %d", len(tl.Tasks), tt.want)
			}
		})
	}
}

func TestDeleteTaskFromTaskIst(t *testing.T) {
	tl := TaskList{}
	for _, tt := range tests {
		tl.AddTask(tt.t)
	}

	tasksLen := len(tl.Tasks)
	tID := int32(rand.Intn(6))
	if err := tl.DeleteTask(tID); err != nil {
		t.Error(err)
	} else if len(tl.Tasks) != tasksLen-1 {
		t.Errorf("got TasksList.len: %d, want: %d", len(tl.Tasks), tasksLen-1)
	}
}

func TestUpdateTaskFromTaskList(t *testing.T) {
	tl := TaskList{}
	for _, tt := range tests {
		tl.AddTask(tt.t)
	}

	tasksLen := len(tl.Tasks)
	tID := int32(rand.Intn(6))
	tnewID := int32(rand.Intn(100))
	tnew := Task{
		TaskID:      tnewID,
		Title:       fmt.Sprintf("Task %v", tnewID),
		CreatedDate: time.Now(),
	}

	if err := tl.UpdateTask(tID, tnew); err != nil {
		t.Error(err)
	} else if len(tl.Tasks) != tasksLen {
		t.Errorf("got TasksList.len: %d, want: %d", len(tl.Tasks), tasksLen)
	} else if task, err := tl.GetTask(tID); err != nil {
		t.Error(err)
	} else if task.TaskID != tID || task.Title != tnew.Title {
		t.Errorf("Task.ID: %d,Task.Title: %s, Want: Task.ID: %d Task,Title: %s",
			task.TaskID,
			task.Title,
			tnew.TaskID,
			tnew.Title)
	}
}

// Contacts unit testing
var testsContacts = []struct {
	c Contact
}{
	{Contact{ContactID: nextID(), FullName: "Donald Trump", Mobile: 19007777777, Email: "trump@example.com"}},
	{Contact{ContactID: nextID(), FullName: "Emmanuel Macron", Mobile: 339002222222, Email: "makron@example.com"}},
	{Contact{ContactID: nextID(), FullName: "Angela Merkel", Mobile: 499003333333, Email: "merkel@example.com"}},
	{Contact{ContactID: nextID(), FullName: "Kim Chin-kyung", Mobile: 8509007777777, Email: "ironman@example.com"}},
}

func TestContactCreate(t *testing.T) {
	id := nextID()
	c := Contact{ContactID: id, FullName: "W Putin", Mobile: 799977700707, Email: "wp@example.com"}
	if c.FullName != "W Putin" || c.ContactID != id {
		t.Errorf("Contact(%v, %v, %v)", c.ContactID, c.FullName, c.Mobile)
	}
}

func TestContactUpdate(t *testing.T) {
	id1 := nextID()
	id2 := nextID()
	c1 := Contact{
		ContactID: id1,
		FullName:  "XYZ",
		Mobile:    222222222,
		Email:     "xyz@mail.com"}
	c2 := Contact{
		ContactID: id2,
		FullName:  "ABC",
		Mobile:    1331313133,
		Email:     "abc@mail.com"}
	c1.Update(c2)
	if c1.ContactID != id1 || c1.FullName != c2.FullName || c1.Mobile != c2.Mobile {
		t.Errorf("Contact(%v %v %v)", c1.ContactID, c1.FullName, c1.Mobile)
	}
}

func TestContactListAddContact(t *testing.T) {
	cl := ContactList{}
	i := 0
	// testname := "Adding all sample contacts to contactlist"
	for _, tt := range testsContacts {
		i++
		testname := fmt.Sprintf("ContactList.len: %v,ContactID: %v", i, tt.c.ContactID)
		t.Run(testname, func(t *testing.T) {
			cl.AddContact(tt.c)
			if len(cl.Contacts) != i {
				t.Errorf("got ContactLIst.len: %d, want: %d", len(cl.Contacts), i)
			}
		})
	}
}

func TestDeleteContactFromContactList(t *testing.T) {
	cl := ContactList{}
	for _, tt := range testsContacts {
		cl.AddContact(tt.c)
	}

	contactLen := len(cl.Contacts)
	cID := cl.Contacts[rand.Intn(contactLen)].ContactID
	if err := cl.DeleteContact(cID); err != nil {
		t.Error(err)
	} else if len(cl.Contacts) != contactLen-1 {
		t.Errorf("got ContactList.len: %d, want: %d", len(cl.Contacts), contactLen-1)
	}
}

func TestUpdateContactFromContactList(t *testing.T) {
	cl := ContactList{}
	for _, tt := range testsContacts {
		cl.AddContact(tt.c)
	}

	contactsLen := len(cl.Contacts)
	cID := cl.Contacts[rand.Intn(contactsLen)].ContactID
	cnewID := int32(rand.Intn(100))
	cnew := Contact{
		ContactID: cnewID,
		FullName:  "George W.Bush",
		Mobile:    199899444444,
		Email:     "gwb@circus.com",
	}

	if err := cl.UpdateContact(cID, cnew); err != nil {
		t.Error(err)
	} else if len(cl.Contacts) != contactsLen {
		t.Errorf("got ContactList.len: %d, want: %d", len(cl.Contacts), contactsLen)
	} else if c, err := cl.GetContact(cID); err != nil {
		t.Error(err)
	} else if c.ContactID != cID || c.FullName != cnew.FullName || c.Mobile != cnew.Mobile {
		t.Errorf("Contact.ID: %d,Contact.FullName: %s, Want: Contact.ID: %d Contact.FullName: %s",
			c.ContactID,
			c.FullName,
			c.ContactID,
			cnew.FullName)
	}
}
