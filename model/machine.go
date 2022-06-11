package model

type Machine struct {
	amount int
}

func NewMachine(amount int) *Machine {
	return &Machine{amount}
}

func (machine *Machine) GetAmount() int {
	return machine.amount
}

func (machine *Machine) SetAmount(amount int) bool {
	machine.amount = amount
	return true
}
