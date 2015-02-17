package main

import (
    "fmt"
    "github.com/golang/protobuf/proto"
    "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) Enum() *Person_PhoneType {
	p := new(Person_PhoneType)
	*p = x
	return p
}
func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (x *Person_PhoneType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Person_PhoneType_value, data, "Person_PhoneType")
	if err != nil {
		return err
	}
	*x = Person_PhoneType(value)
	return nil
}

type Person struct {
	Name             *string               `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id               *int32                `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Email            *string               `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone            []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phone" json:"phone,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *Person) GetPhone() []*Person_PhoneNumber {
	if m != nil {
		return m.Phone
	}
	return nil
}

type Person_PhoneNumber struct {
	Number           *string           `protobuf:"bytes,1,req,name=number" json:"number,omitempty"`
	Type             *Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=main.Person_PhoneType,def=1" json:"type,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}

const Default_Person_PhoneNumber_Type Person_PhoneType = Person_HOME

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Person_PhoneNumber_Type
}

type AddressBook struct {
	Person           []*Person `protobuf:"bytes,1,rep,name=person" json:"person,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}

func (m *AddressBook) GetPerson() []*Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}
func main() {
    person := &Person{
        Name: proto.String("Pankaj"),
        Id: proto.Int(1),
    }
    addressbook := &AddressBook{}
    addressbook.Person = []*Person {person}
    fmt.Printf("Hi");
    
}
