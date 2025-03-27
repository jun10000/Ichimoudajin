package component

type ActorCom struct {
	name string
}

func NewActorCom(name string) *ActorCom {
	return &ActorCom{
		name: name,
	}
}

func (c *ActorCom) GetName() string {
	return c.name
}
