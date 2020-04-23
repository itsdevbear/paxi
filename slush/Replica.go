package slush

import(
	"Slush/paxi"
	"Slush/paxi/log"
)
// Replica for Slush
type Replica struct {
	paxi.Node
	*Slush
}

// NewReplica generates new Slush replica
func NewReplica(id paxi.ID) *Replica {
	r := new(Replica)
	r.Node = paxi.NewNode(id)
	r.Slush = NewSlush(r)
	r.Register(paxi.Request{}, r.handleRequest)
	r.Register(Msg1{}, r.HandleMsg1)
	return r
}

func (r *Replica) handleRequest(m paxi.Request) {
	log.Debugf("Replica %s received %v\n", r.ID(), m)

	//r.Slush.SetQuerying(true)
	r.Slush.HandleRequest(m)
}
