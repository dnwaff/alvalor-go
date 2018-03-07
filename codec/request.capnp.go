// Code generated by capnpc-go. DO NOT EDIT.

package codec

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Request struct{ capnp.Struct }

// Request_TypeID is the unique identifier for the type Request.
const Request_TypeID = 0xe8d311873fc93b93

func NewRequest(s *capnp.Segment) (Request, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Request{st}, err
}

func NewRootRequest(s *capnp.Segment) (Request, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Request{st}, err
}

func ReadRootRequest(msg *capnp.Message) (Request, error) {
	root, err := msg.RootPtr()
	return Request{root.Struct()}, err
}

func (s Request) String() string {
	str, _ := text.Marshal(0xe8d311873fc93b93, s.Struct)
	return str
}

func (s Request) Ids() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.DataList{List: p.List()}, err
}

func (s Request) HasIds() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Request) SetIds(v capnp.DataList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewIds sets the ids field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Request) NewIds(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Request_List is a list of Request.
type Request_List struct{ capnp.List }

// NewRequest creates a new list of Request.
func NewRequest_List(s *capnp.Segment, sz int32) (Request_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Request_List{l}, err
}

func (s Request_List) At(i int) Request { return Request{s.List.Struct(i)} }

func (s Request_List) Set(i int, v Request) error { return s.List.SetStruct(i, v.Struct) }

func (s Request_List) String() string {
	str, _ := text.MarshalList(0xe8d311873fc93b93, s.List)
	return str
}

// Request_Promise is a wrapper for a Request promised by a client call.
type Request_Promise struct{ *capnp.Pipeline }

func (p Request_Promise) Struct() (Request, error) {
	s, err := p.Pipeline.Struct()
	return Request{s}, err
}

const schema_88b5d31dcdf19e4d = "x\xda2pet`2d\xf5\xe7d`\x08\xcca" +
	"e\xfb?\xd9\xfa\xa4}\xbb\xe0\xe5\x17\x0c\x82|\x8c\xff" +
	"}\xe7}<+{yk\x07\x03+#;\x03\x83p" +
	"/\xcb&\xe1\xa9, \xd6D\x16{\x86\xff\x0c\xf2\xff" +
	"\x8bR\x0bKS\x8bK\xf4\x18\x93\x13\x0b\xf2\x0a\xac\x82" +
	"R\xe5\xc1\xfc\x00F\xc6@\x16f\x16\x06\x06\x16F\x06" +
	"\x06A^%\x06\x86@\x0ef\xc6@\x15&F\xf6\xcc" +
	"\x94bF>\x06\xc6\x00fFF^\x06&\x10\x13\x10" +
	"\x00\x00\xff\xffI\xaf\x1ee"

func init() {
	schemas.Register(schema_88b5d31dcdf19e4d,
		0xe8d311873fc93b93)
}