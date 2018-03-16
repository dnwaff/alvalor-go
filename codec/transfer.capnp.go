// Code generated by capnpc-go. DO NOT EDIT.

package codec

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Transfer struct{ capnp.Struct }

// Transfer_TypeID is the unique identifier for the type Transfer.
const Transfer_TypeID = 0xfebd493a48891da9

func NewTransfer(s *capnp.Segment) (Transfer, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Transfer{st}, err
}

func NewRootTransfer(s *capnp.Segment) (Transfer, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Transfer{st}, err
}

func ReadRootTransfer(msg *capnp.Message) (Transfer, error) {
	root, err := msg.RootPtr()
	return Transfer{root.Struct()}, err
}

func (s Transfer) String() string {
	str, _ := text.Marshal(0xfebd493a48891da9, s.Struct)
	return str
}

func (s Transfer) From() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Transfer) HasFrom() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Transfer) SetFrom(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Transfer) To() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Transfer) HasTo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Transfer) SetTo(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Transfer) Amount() uint64 {
	return s.Struct.Uint64(0)
}

func (s Transfer) SetAmount(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Transfer_List is a list of Transfer.
type Transfer_List struct{ capnp.List }

// NewTransfer creates a new list of Transfer.
func NewTransfer_List(s *capnp.Segment, sz int32) (Transfer_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Transfer_List{l}, err
}

func (s Transfer_List) At(i int) Transfer { return Transfer{s.List.Struct(i)} }

func (s Transfer_List) Set(i int, v Transfer) error { return s.List.SetStruct(i, v.Struct) }

func (s Transfer_List) String() string {
	str, _ := text.MarshalList(0xfebd493a48891da9, s.List)
	return str
}

// Transfer_Promise is a wrapper for a Transfer promised by a client call.
type Transfer_Promise struct{ *capnp.Pipeline }

func (p Transfer_Promise) Struct() (Transfer, error) {
	s, err := p.Pipeline.Struct()
	return Transfer{s}, err
}

const schema_b660749b49284997 = "x\xda4\xc8\xafJ\xc7P\x18\xc6\xf1\xe7y\xcf\x992" +
	"p\x8c\xb1%A\xacb\x10\xad+.(\xec\x80\xe2^" +
	"X\xd7!.\xb9?\xcc\xe3E\x18\x0c\x061x\x05&" +
	"\x9b\xcd\xe2\x85yd\xe0\xaf}\xbe\xdf\xe3kVr\x12" +
	"]\xc5\x80\xdeG[\xe1c\xef\xa9.\xdd\xf7/4%" +
	"\xc3\x9b;p\xef\xfe\xe6\x0b\x91l\x03\xf9\xb3\xfd\xc9_" +
	"\xed\xaa\x17\xfb\x89\x80\xfd\xe0\x97n|\xe8\xef\x169\xba" +
	"\xed\xe6q.\xdb\xffFC\xea\x8e\xb1\x80%\x90\x9d\x1f" +
	"\x02Z\x19\xea\x850#\x0b\xae\xd3\xed\x02zf\xa8\x8d" +
	"\x90RP\x80\xec\xb2\x04\xb46\xd4V\x98\xf6\xcb40" +
	"\x810\x01\x8d\x9f6<\xed\x86\xe9q\xf4\x8c!\x8c\xc1" +
	"\xbf\x00\x00\x00\xff\xff~:&\xa7"

func init() {
	schemas.Register(schema_b660749b49284997,
		0xfebd493a48891da9)
}
