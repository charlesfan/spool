package spool

type Dispatcher interface {
	Next([]Printer, map[string]string) Printer
}

// RoundRobin
type dispatcher struct {
	idx int
}

func (d *dispatcher) Next(printers []Printer, meta map[string]string) Printer {
	if len(printers) == 0 {
		return nil
	}
	p := printers[d.idx%len(printers)]
	d.idx++
	return p
}
