package ikimono

type Ikimono struct {
	Name    string
	Atack   int
	Defence int
	Speed   int
}

func NewIkimono(name string, atack, defence, speed int) *Ikimono {
	newIkimono := &Ikimono{Name: name, Atack: atack, Defence: defence, Speed: speed}
	return newIkimono
}

func NewIkimonoOnlyName(name string) *Ikimono {
	newIkimono := &Ikimono{Name: name}
	return newIkimono
}

func (i *Ikimono) ShowIkimonoName() string {
	return i.Name
}

func CalcDamage(ikimonoA, ikimonoB *Ikimono) int {
	return ikimonoA.Atack - ikimonoB.Defence
}
