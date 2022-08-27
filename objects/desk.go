package objects

type Desk struct {
	id       int
	name     string
	active   bool
	callback func(bool)
}

func (this *Desk) GetName() string {
	return this.name
}

func (this *Desk) SetName(name string) {
	this.name = name
}

func (this *Desk) GetId() int {
	return this.id
}

func (this *Desk) SetId(id int) {
	this.id = id
}

func (this *Desk) GetActive() bool {
	return this.active
}

func (this *Desk) SetActive(active bool) {
	this.callback(active)
	this.active = active
}

func (this *Desk) GetCallback() func(bool) {
	return this.callback
}

func (this *Desk) SetCallback(callback func(bool)) {
	this.callback = callback
}

func (this *Desk) Open(OnFail func(), OnSuccess func()) {
	if this.active {
		OnFail()
		return
	} else {
		OnSuccess()
		this.active = true
		return
	}
}
