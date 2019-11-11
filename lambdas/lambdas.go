package lambdas

type (
	// Pipe - the pipe struct.
	Pipe struct {
		stages []*Stage
	}

	// Stage - what pipes are made of
	Stage struct {
		typ    string
		flags  []string
		entity interface{}
	}

	// Design - what pipes makes
	Design struct{}
)

/*
	Can we expect json all the way?
	Are there anyway we can make this dynamic?
	- Like for request A use these stages, and these for request B...
	 - Perhaps if we have a json version of the request?
	 - But what are then the penalty for potentially recreating the pipe on each request?
	 - Which can be avoided if we're the slightest bit clever... tbh
	Can we support pipes that does not return anything?
	Can we support pipes that are not started/triggered with design.Push()?
	Can we have pluggable support for new fancy stages?

	pipe := NewPipe()
	pipe.Run(Post(http://...), ResponseEntity1)
	pipe.Run(RPC(some.topic), ResponseEntity2)
	pipe.Run(Put(http://...), ResponseEntity3)
	design := pipe.Build()

	ResponseEntity3, err := design.Push(RequestEntity1)
*/

// NewPipe - creates a new pipe
func NewPipe() *Pipe {
	return &Pipe{}
}

// Run - append a stage
func (p *Pipe) Run(stage *Stage) *Pipe {
	p.stages = append(p.stages, stage)

	return p
}

// Build - finish the pipeline
func (p *Pipe) Build() *Design {
	// iterate p.stages
	// switch on typ
	// call corresponding higher order function with settings
	// add returned function into Design
	// return design
	return &Design{}
}
