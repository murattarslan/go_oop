package objects

type Callback interface {
	OnFail()
	OnSuccess()
}
