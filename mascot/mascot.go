package mascot

type mascotObj struct {
	mascotType string
	mascotName string
}

func (m mascotObj) GetMascotName() string {
	return m.mascotName
}

func (m mascotObj) GetMascotType() string {
	return m.mascotType
}

func (m mascotObj) SetMascotName() mascotObj {
	m.mascotName = "Krampus"
	return m
}

func (m mascotObj) SetMascotType() mascotObj {
	m.mascotType = "Evil Christmas Creature"
	return m
}

func GetMascot() mascotObj {
	var m mascotObj = mascotObj{}
	m = m.SetMascotName()
	m = m.SetMascotType()
	return m
}
