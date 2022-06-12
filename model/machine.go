package model

type Machine struct {
	inventory int
}

func NewMachine(inventory int) *Machine {
	return &Machine{inventory}
}

func (machine *Machine) GetInventory() int {
	return machine.inventory
}

func (machine *Machine) SetInventory(inventory int) int {
	machine.inventory = inventory
	return machine.inventory
}

func (machine *Machine) Add(amount int) int {
	machine.inventory += amount
	return machine.inventory
}
