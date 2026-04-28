package ctx

type Scope struct {
	Parent *Scope
}

func NewCtxWithParent(parent *Scope) *Scope {
	return &Scope{
		Parent: parent,
	}
}

func NewCtx() *Scope {
	return NewCtxWithParent(nil)
}
