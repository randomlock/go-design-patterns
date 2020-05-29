package creational

import "fmt"

type vehicle struct {
	wheel, color int
}


type VehicleBuilder interface {
	wheel(int) VehicleBuilder
	color(int) VehicleBuilder
	display()
}

type Car struct {
	vehicle
}

func (c *Car) wheel(i int) VehicleBuilder {
	c.vehicle.wheel = i
	return c
}

func (c *Car) color(i int) VehicleBuilder {
	c.vehicle.color = i
	return c
}

func (c *Car) display() {
	fmt.Printf("The car has %d color and %d wheel", c.vehicle.color, c.vehicle.wheel)
}

type Bus struct {
	vehicle
}

func (b *Bus) wheel(i int) VehicleBuilder {
	b.vehicle.wheel = i
	return b
}

func (b *Bus) color(i int) VehicleBuilder {
	b.vehicle.color = i
	return b
}

func (b *Bus) display() {
	fmt.Println(fmt.Sprintf("The car has %d color and %d wheel\n\n", b.vehicle.color, b.vehicle.wheel))
}



type Director struct {
	builder VehicleBuilder
}


func init() {
	var vehicle VehicleBuilder
	vehicle = &Car{}
	vehicle.wheel(1).color(2).display()

	vehicle = &Bus{}
	vehicle.wheel(2).color(2).display()
}
