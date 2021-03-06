// Code generated by capnpc-go. DO NOT EDIT.

package codec

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Batch struct{ capnp.Struct }

// Batch_TypeID is the unique identifier for the type Batch.
const Batch_TypeID = 0x864e51a95b5bacee

func NewBatch(s *capnp.Segment) (Batch, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Batch{st}, err
}

func NewRootBatch(s *capnp.Segment) (Batch, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Batch{st}, err
}

func ReadRootBatch(msg *capnp.Message) (Batch, error) {
	root, err := msg.RootPtr()
	return Batch{root.Struct()}, err
}

func (s Batch) String() string {
	str, _ := text.Marshal(0x864e51a95b5bacee, s.Struct)
	return str
}

func (s Batch) Transactions() (Transaction_List, error) {
	p, err := s.Struct.Ptr(0)
	return Transaction_List{List: p.List()}, err
}

func (s Batch) HasTransactions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Batch) SetTransactions(v Transaction_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTransactions sets the transactions field to a newly
// allocated Transaction_List, preferring placement in s's segment.
func (s Batch) NewTransactions(n int32) (Transaction_List, error) {
	l, err := NewTransaction_List(s.Struct.Segment(), n)
	if err != nil {
		return Transaction_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Batch_List is a list of Batch.
type Batch_List struct{ capnp.List }

// NewBatch creates a new list of Batch.
func NewBatch_List(s *capnp.Segment, sz int32) (Batch_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Batch_List{l}, err
}

func (s Batch_List) At(i int) Batch { return Batch{s.List.Struct(i)} }

func (s Batch_List) Set(i int, v Batch) error { return s.List.SetStruct(i, v.Struct) }

func (s Batch_List) String() string {
	str, _ := text.MarshalList(0x864e51a95b5bacee, s.List)
	return str
}

// Batch_Promise is a wrapper for a Batch promised by a client call.
type Batch_Promise struct{ *capnp.Pipeline }

func (p Batch_Promise) Struct() (Batch, error) {
	s, err := p.Pipeline.Struct()
	return Batch{s}, err
}

const schema_e60ceb912269b107 = "x\xda2pct`2d\xf5\xe7d`\x08\xcca" +
	"e\xfb\xffnMt\xf4\xca@\xbf6\x06A\x1e\xc6\xff" +
	"\xec\x1b3\x95&\xbe\xe6y\xc6\xc0\xca\xc8\xce\xc0 \xdc" +
	"\xcb2Ix*\x0b\x885\x91\xc5\x9e\xe1?\x83\xfc\xff" +
	"\xa4\xc4\x92\xe4\x0c\xbd\xe4D\xc6\x82\xbc\x02+\xa7\xc4\x92" +
	"d\xc6\x8c\x00F\xc6@\x16f\x16\x06\x06\x16F\x06\x06" +
	"A\xde,\x06\x86@\x1ef\xc6@\x0d&\xc6\xff%E" +
	"\x89y\xc5\x89\xc9%\x0c\xfc\x99\xf9y\xc5\x8c|\x0c\x8c" +
	"\x01\xcc\x8c\x8c\x02\xff\xe3\x995/\xb2\xdd\x9c\xff\x86\x81" +
	"\x81\x11$\x08\x08\x00\x00\xff\xff\x18\xc0$\xdf"

func init() {
	schemas.Register(schema_e60ceb912269b107,
		0x864e51a95b5bacee)
}
