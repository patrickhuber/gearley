package charts

type chart struct {
	sets []Set
}

type Chart interface {
	Add(index int, state State) bool
	Get(index int) Set
	Sets() []Set
}

func NewChart() Chart {
	return &chart{
		sets: make([]Set, 0),
	}
}

func (c *chart) Add(index int, state State) bool {
	set := c.Get(index)
	return set.Add(state)
}

func (c *chart) Get(index int) Set {
	if index < len(c.sets) {
		return c.sets[index]
	}

	set := NewSet()
	c.sets = append(c.sets, set)
	return set
}

func (c *chart) Sets() []Set {
	return c.sets
}
