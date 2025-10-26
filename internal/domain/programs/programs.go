package programs

type Program struct {
	Executions []Execution
	Hour       Hour
}

type Programs struct {
	Daily       Daily
	Odd         Odd
	Even        Even
	Temperature Temperature
	Weekly      []Weekly
}
