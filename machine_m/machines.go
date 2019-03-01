package machine_m

import "fmt"

type Machine struct {
	Ip   string
	Host string
}

type VsphMachine struct {
	Machine    Machine
	Data_store string
}

func (m Machine) GetName() string {
	return "dddd"
}

func Changeip(m Machine) {
	m.Ip = "changed in ip"
}

func Changeipp(m *Machine) {
	m.Ip = "changed in ipp"
}

func Testy() {
	m := Machine{"ip", "asddd"}
	p := &m
	fmt.Println(m.Ip)

	m.Ip = "really"
	fmt.Println(m.Ip)

	p.Ip = "newip"

	fmt.Println(m.Ip)

	t := m
	t.Ip = "tip"

	fmt.Println(m.Ip)
	fmt.Println(t.Ip)

	fmt.Printf("addr of m is %v and address of t is %v \n", p, &t)

	Changeip(m)
	Changeipp(&t)

	fmt.Printf("addr of m is %v and address of t is %v \n", p, &t)
}
