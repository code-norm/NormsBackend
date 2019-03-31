package classes

type Med struct {
	Name    string
	Gender  string
	Race    string
	Age     int
	Weight  int
	History string
}

type User struct {
	Uname  string
	Uemail string
	Symp   [12]Symptom
}

type Usertemp struct {
	Username string
	Email    string
}

type Symptom struct {
	Name          string
	Checked       string
	Notifications string
}
