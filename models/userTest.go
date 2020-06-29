package models

type UserTest struct {
	Id   int
	Name string
}

func GetUser() map[string]UserTest {
	var a UserTest
	a.Id = 1
	a.Name = "testname"
	u := make(map[string]UserTest)
	u["test"] = a
	return u
}
