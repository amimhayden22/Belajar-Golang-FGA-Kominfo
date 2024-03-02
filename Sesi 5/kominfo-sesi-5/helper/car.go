package helper

import "fmt"

type Car struct {
	Price          uint
	Brand          string
	machineVersion string
}

func (c *Car) SetMachineVersion(role string, machineVersion string) {
	if role != "admin" {
		panic("you are not authorized to access this resource")
	}

	c.machineVersion = machineVersion
}

func (c Car) MachineVersion(role string) {
	fmt.Printf("machine version => %s\n", c.machineVersion)
}