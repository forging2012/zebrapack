// Code generated by ZEBRAPACK (github.com/glycerine/zebrapack). DO NOT EDIT.

package testdata

import (
	"github.com/glycerine/zebrapack/msgp"
)

// MSGPfieldsNotEmpty supports omitempty tags
func (z *A) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Sibs == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *A) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Bday"
	o = append(o, 0xa4, 0x42, 0x64, 0x61, 0x79)
	o = msgp.AppendTime(o, z.Bday)
	if !empty[2] {
		// string "phone"
		o = append(o, 0xa5, 0x70, 0x68, 0x6f, 0x6e, 0x65)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// string "Sibs"
		o = append(o, 0xa4, 0x53, 0x69, 0x62, 0x73)
		o = msgp.AppendInt(o, z.Sibs)
	}

	// string "GPA"
	o = append(o, 0xa3, 0x47, 0x50, 0x41)
	o = msgp.AppendFloat64(o, z.GPA)
	// string "Friend"
	o = append(o, 0xa6, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64)
	o = msgp.AppendBool(o, z.Friend)
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *A) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *A) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0zgensym_58b5c2649cecee38_1 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zgensym_58b5c2649cecee38_1 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zgensym_58b5c2649cecee38_1, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zgensym_58b5c2649cecee38_1 := totalEncodedFields0zgensym_58b5c2649cecee38_1
	missingFieldsLeft0zgensym_58b5c2649cecee38_1 := maxFields0zgensym_58b5c2649cecee38_1 - totalEncodedFields0zgensym_58b5c2649cecee38_1

	var nextMiss0zgensym_58b5c2649cecee38_1 int32 = -1
	var found0zgensym_58b5c2649cecee38_1 [maxFields0zgensym_58b5c2649cecee38_1]bool
	var curField0zgensym_58b5c2649cecee38_1 string

doneWithStruct0zgensym_58b5c2649cecee38_1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgensym_58b5c2649cecee38_1 > 0 || missingFieldsLeft0zgensym_58b5c2649cecee38_1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgensym_58b5c2649cecee38_1, missingFieldsLeft0zgensym_58b5c2649cecee38_1, msgp.ShowFound(found0zgensym_58b5c2649cecee38_1[:]), unmarshalMsgFieldOrder0zgensym_58b5c2649cecee38_1)
		if encodedFieldsLeft0zgensym_58b5c2649cecee38_1 > 0 {
			encodedFieldsLeft0zgensym_58b5c2649cecee38_1--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zgensym_58b5c2649cecee38_1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgensym_58b5c2649cecee38_1 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zgensym_58b5c2649cecee38_1 = 0
			}
			for nextMiss0zgensym_58b5c2649cecee38_1 < maxFields0zgensym_58b5c2649cecee38_1 && (found0zgensym_58b5c2649cecee38_1[nextMiss0zgensym_58b5c2649cecee38_1] || unmarshalMsgFieldSkip0zgensym_58b5c2649cecee38_1[nextMiss0zgensym_58b5c2649cecee38_1]) {
				nextMiss0zgensym_58b5c2649cecee38_1++
			}
			if nextMiss0zgensym_58b5c2649cecee38_1 == maxFields0zgensym_58b5c2649cecee38_1 {
				// filled all the empty fields!
				break doneWithStruct0zgensym_58b5c2649cecee38_1
			}
			missingFieldsLeft0zgensym_58b5c2649cecee38_1--
			curField0zgensym_58b5c2649cecee38_1 = unmarshalMsgFieldOrder0zgensym_58b5c2649cecee38_1[nextMiss0zgensym_58b5c2649cecee38_1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgensym_58b5c2649cecee38_1)
		switch curField0zgensym_58b5c2649cecee38_1 {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zgensym_58b5c2649cecee38_1[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday":
			found0zgensym_58b5c2649cecee38_1[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone":
			found0zgensym_58b5c2649cecee38_1[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs":
			found0zgensym_58b5c2649cecee38_1[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA":
			found0zgensym_58b5c2649cecee38_1[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend":
			found0zgensym_58b5c2649cecee38_1[5] = true
			z.Friend, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zgensym_58b5c2649cecee38_1 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of A
var unmarshalMsgFieldOrder0zgensym_58b5c2649cecee38_1 = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zgensym_58b5c2649cecee38_1 = []bool{false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.TimeSize + 6 + msgp.StringPrefixSize + len(z.Phone) + 5 + msgp.IntSize + 4 + msgp.Float64Size + 7 + msgp.BoolSize
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Big) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Big) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 5
	// string "Slice"
	o = append(o, 0x85, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
	for zgensym_58b5c2649cecee38_2 := range z.Slice {
		o, err = z.Slice[zgensym_58b5c2649cecee38_2].MSGPMarshalMsg(o) // not is.iface, gen/marshal.go:261
		if err != nil {
			return
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zgensym_58b5c2649cecee38_3, zgensym_58b5c2649cecee38_4 := range z.Transform {
		o = msgp.AppendInt(o, zgensym_58b5c2649cecee38_3)
		if zgensym_58b5c2649cecee38_4 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zgensym_58b5c2649cecee38_4.MSGPMarshalMsg(o) // not is.iface, gen/marshal.go:261
			if err != nil {
				return
			}
		}
	}
	// string "Myptr"
	o = append(o, 0xa5, 0x4d, 0x79, 0x70, 0x74, 0x72)
	if z.Myptr == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Myptr.MSGPMarshalMsg(o) // not is.iface, gen/marshal.go:261
		if err != nil {
			return
		}
	}
	// string "Myarray"
	o = append(o, 0xa7, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, 3)
	for zgensym_58b5c2649cecee38_5 := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zgensym_58b5c2649cecee38_5])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zgensym_58b5c2649cecee38_6 := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zgensym_58b5c2649cecee38_6])
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Big) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Big) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields7zgensym_58b5c2649cecee38_8 = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zgensym_58b5c2649cecee38_8 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zgensym_58b5c2649cecee38_8, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft7zgensym_58b5c2649cecee38_8 := totalEncodedFields7zgensym_58b5c2649cecee38_8
	missingFieldsLeft7zgensym_58b5c2649cecee38_8 := maxFields7zgensym_58b5c2649cecee38_8 - totalEncodedFields7zgensym_58b5c2649cecee38_8

	var nextMiss7zgensym_58b5c2649cecee38_8 int32 = -1
	var found7zgensym_58b5c2649cecee38_8 [maxFields7zgensym_58b5c2649cecee38_8]bool
	var curField7zgensym_58b5c2649cecee38_8 string

doneWithStruct7zgensym_58b5c2649cecee38_8:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zgensym_58b5c2649cecee38_8 > 0 || missingFieldsLeft7zgensym_58b5c2649cecee38_8 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgensym_58b5c2649cecee38_8, missingFieldsLeft7zgensym_58b5c2649cecee38_8, msgp.ShowFound(found7zgensym_58b5c2649cecee38_8[:]), unmarshalMsgFieldOrder7zgensym_58b5c2649cecee38_8)
		if encodedFieldsLeft7zgensym_58b5c2649cecee38_8 > 0 {
			encodedFieldsLeft7zgensym_58b5c2649cecee38_8--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField7zgensym_58b5c2649cecee38_8 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zgensym_58b5c2649cecee38_8 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zgensym_58b5c2649cecee38_8 = 0
			}
			for nextMiss7zgensym_58b5c2649cecee38_8 < maxFields7zgensym_58b5c2649cecee38_8 && (found7zgensym_58b5c2649cecee38_8[nextMiss7zgensym_58b5c2649cecee38_8] || unmarshalMsgFieldSkip7zgensym_58b5c2649cecee38_8[nextMiss7zgensym_58b5c2649cecee38_8]) {
				nextMiss7zgensym_58b5c2649cecee38_8++
			}
			if nextMiss7zgensym_58b5c2649cecee38_8 == maxFields7zgensym_58b5c2649cecee38_8 {
				// filled all the empty fields!
				break doneWithStruct7zgensym_58b5c2649cecee38_8
			}
			missingFieldsLeft7zgensym_58b5c2649cecee38_8--
			curField7zgensym_58b5c2649cecee38_8 = unmarshalMsgFieldOrder7zgensym_58b5c2649cecee38_8[nextMiss7zgensym_58b5c2649cecee38_8]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zgensym_58b5c2649cecee38_8)
		switch curField7zgensym_58b5c2649cecee38_8 {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found7zgensym_58b5c2649cecee38_8[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zgensym_58b5c2649cecee38_9 uint32
				zgensym_58b5c2649cecee38_9, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zgensym_58b5c2649cecee38_9) {
					z.Slice = (z.Slice)[:zgensym_58b5c2649cecee38_9]
				} else {
					z.Slice = make([]S2, zgensym_58b5c2649cecee38_9)
				}
				for zgensym_58b5c2649cecee38_2 := range z.Slice {
					bts, err = z.Slice[zgensym_58b5c2649cecee38_2].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Transform":
			found7zgensym_58b5c2649cecee38_8[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zgensym_58b5c2649cecee38_10 uint32
				zgensym_58b5c2649cecee38_10, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zgensym_58b5c2649cecee38_10 > 0 {
					z.Transform = make(map[int]*S2, zgensym_58b5c2649cecee38_10)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zgensym_58b5c2649cecee38_10 > 0 {
					var zgensym_58b5c2649cecee38_3 int
					var zgensym_58b5c2649cecee38_4 *S2
					zgensym_58b5c2649cecee38_10--
					zgensym_58b5c2649cecee38_3, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zgensym_58b5c2649cecee38_4 != nil {
							zgensym_58b5c2649cecee38_4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zgensym_58b5c2649cecee38_4 {
								zgensym_58b5c2649cecee38_4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zgensym_58b5c2649cecee38_4 == nil {
								zgensym_58b5c2649cecee38_4 = new(S2)
							}
							bts, err = zgensym_58b5c2649cecee38_4.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
						}
					}
					z.Transform[zgensym_58b5c2649cecee38_3] = zgensym_58b5c2649cecee38_4
				}
			}
		case "Myptr":
			found7zgensym_58b5c2649cecee38_8[2] = true
			if nbs.AlwaysNil {
				if z.Myptr != nil {
					z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Myptr {
						z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Myptr == nil {
						z.Myptr = new(S2)
					}
					bts, err = z.Myptr.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Myarray":
			found7zgensym_58b5c2649cecee38_8[3] = true
			var zgensym_58b5c2649cecee38_11 uint32
			zgensym_58b5c2649cecee38_11, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zgensym_58b5c2649cecee38_11 != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zgensym_58b5c2649cecee38_11}
				return
			}
			for zgensym_58b5c2649cecee38_5 := range z.Myarray {
				z.Myarray[zgensym_58b5c2649cecee38_5], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice":
			found7zgensym_58b5c2649cecee38_8[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zgensym_58b5c2649cecee38_12 uint32
				zgensym_58b5c2649cecee38_12, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zgensym_58b5c2649cecee38_12) {
					z.MySlice = (z.MySlice)[:zgensym_58b5c2649cecee38_12]
				} else {
					z.MySlice = make([]string, zgensym_58b5c2649cecee38_12)
				}
				for zgensym_58b5c2649cecee38_6 := range z.MySlice {
					z.MySlice[zgensym_58b5c2649cecee38_6], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss7zgensym_58b5c2649cecee38_8 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Big
var unmarshalMsgFieldOrder7zgensym_58b5c2649cecee38_8 = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip7zgensym_58b5c2649cecee38_8 = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_2 := range z.Slice {
		s += z.Slice[zgensym_58b5c2649cecee38_2].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zgensym_58b5c2649cecee38_3, zgensym_58b5c2649cecee38_4 := range z.Transform {
			_ = zgensym_58b5c2649cecee38_4
			_ = zgensym_58b5c2649cecee38_3
			s += msgp.IntSize
			if zgensym_58b5c2649cecee38_4 == nil {
				s += msgp.NilSize
			} else {
				s += zgensym_58b5c2649cecee38_4.MSGPMsgsize()
			}
		}
	}
	s += 6
	if z.Myptr == nil {
		s += msgp.NilSize
	} else {
		s += z.Myptr.MSGPMsgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_5 := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zgensym_58b5c2649cecee38_5])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_6 := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zgensym_58b5c2649cecee38_6])
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *S2) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 7
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *S2) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 7
	// string "beta"
	o = append(o, 0x87, 0xa4, 0x62, 0x65, 0x74, 0x61)
	o = msgp.AppendString(o, z.B)
	// string "ralph"
	o = append(o, 0xa5, 0x72, 0x61, 0x6c, 0x70, 0x68)
	o = msgp.AppendMapHeader(o, uint32(len(z.R)))
	for zgensym_58b5c2649cecee38_13, zgensym_58b5c2649cecee38_14 := range z.R {
		o = msgp.AppendString(o, zgensym_58b5c2649cecee38_13)
		o = msgp.AppendUint8(o, zgensym_58b5c2649cecee38_14)
	}
	// string "P"
	o = append(o, 0xa1, 0x50)
	o = msgp.AppendUint16(o, z.P)
	// string "Q"
	o = append(o, 0xa1, 0x51)
	o = msgp.AppendUint32(o, z.Q)
	// string "T"
	o = append(o, 0xa1, 0x54)
	o = msgp.AppendInt64(o, z.T)
	// string "arr"
	o = append(o, 0xa3, 0x61, 0x72, 0x72)
	o = msgp.AppendArrayHeader(o, 6)
	for zgensym_58b5c2649cecee38_15 := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zgensym_58b5c2649cecee38_15])
	}
	// string "MyTree"
	o = append(o, 0xa6, 0x4d, 0x79, 0x54, 0x72, 0x65, 0x65)
	if z.MyTree == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.MyTree.MSGPMarshalMsg(o) // not is.iface, gen/marshal.go:261
		if err != nil {
			return
		}
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *S2) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *S2) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields16zgensym_58b5c2649cecee38_17 = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields16zgensym_58b5c2649cecee38_17 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields16zgensym_58b5c2649cecee38_17, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft16zgensym_58b5c2649cecee38_17 := totalEncodedFields16zgensym_58b5c2649cecee38_17
	missingFieldsLeft16zgensym_58b5c2649cecee38_17 := maxFields16zgensym_58b5c2649cecee38_17 - totalEncodedFields16zgensym_58b5c2649cecee38_17

	var nextMiss16zgensym_58b5c2649cecee38_17 int32 = -1
	var found16zgensym_58b5c2649cecee38_17 [maxFields16zgensym_58b5c2649cecee38_17]bool
	var curField16zgensym_58b5c2649cecee38_17 string

doneWithStruct16zgensym_58b5c2649cecee38_17:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zgensym_58b5c2649cecee38_17 > 0 || missingFieldsLeft16zgensym_58b5c2649cecee38_17 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zgensym_58b5c2649cecee38_17, missingFieldsLeft16zgensym_58b5c2649cecee38_17, msgp.ShowFound(found16zgensym_58b5c2649cecee38_17[:]), unmarshalMsgFieldOrder16zgensym_58b5c2649cecee38_17)
		if encodedFieldsLeft16zgensym_58b5c2649cecee38_17 > 0 {
			encodedFieldsLeft16zgensym_58b5c2649cecee38_17--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField16zgensym_58b5c2649cecee38_17 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zgensym_58b5c2649cecee38_17 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss16zgensym_58b5c2649cecee38_17 = 0
			}
			for nextMiss16zgensym_58b5c2649cecee38_17 < maxFields16zgensym_58b5c2649cecee38_17 && (found16zgensym_58b5c2649cecee38_17[nextMiss16zgensym_58b5c2649cecee38_17] || unmarshalMsgFieldSkip16zgensym_58b5c2649cecee38_17[nextMiss16zgensym_58b5c2649cecee38_17]) {
				nextMiss16zgensym_58b5c2649cecee38_17++
			}
			if nextMiss16zgensym_58b5c2649cecee38_17 == maxFields16zgensym_58b5c2649cecee38_17 {
				// filled all the empty fields!
				break doneWithStruct16zgensym_58b5c2649cecee38_17
			}
			missingFieldsLeft16zgensym_58b5c2649cecee38_17--
			curField16zgensym_58b5c2649cecee38_17 = unmarshalMsgFieldOrder16zgensym_58b5c2649cecee38_17[nextMiss16zgensym_58b5c2649cecee38_17]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zgensym_58b5c2649cecee38_17)
		switch curField16zgensym_58b5c2649cecee38_17 {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found16zgensym_58b5c2649cecee38_17[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph":
			found16zgensym_58b5c2649cecee38_17[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zgensym_58b5c2649cecee38_18 uint32
				zgensym_58b5c2649cecee38_18, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zgensym_58b5c2649cecee38_18 > 0 {
					z.R = make(map[string]uint8, zgensym_58b5c2649cecee38_18)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zgensym_58b5c2649cecee38_18 > 0 {
					var zgensym_58b5c2649cecee38_13 string
					var zgensym_58b5c2649cecee38_14 uint8
					zgensym_58b5c2649cecee38_18--
					zgensym_58b5c2649cecee38_13, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgensym_58b5c2649cecee38_14, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zgensym_58b5c2649cecee38_13] = zgensym_58b5c2649cecee38_14
				}
			}
		case "P":
			found16zgensym_58b5c2649cecee38_17[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q":
			found16zgensym_58b5c2649cecee38_17[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T":
			found16zgensym_58b5c2649cecee38_17[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr":
			found16zgensym_58b5c2649cecee38_17[6] = true
			var zgensym_58b5c2649cecee38_19 uint32
			zgensym_58b5c2649cecee38_19, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zgensym_58b5c2649cecee38_19 != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zgensym_58b5c2649cecee38_19}
				return
			}
			for zgensym_58b5c2649cecee38_15 := range z.Arr {
				z.Arr[zgensym_58b5c2649cecee38_15], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree":
			found16zgensym_58b5c2649cecee38_17[7] = true
			if nbs.AlwaysNil {
				if z.MyTree != nil {
					z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.MyTree {
						z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.MyTree == nil {
						z.MyTree = new(Tree)
					}
					bts, err = z.MyTree.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss16zgensym_58b5c2649cecee38_17 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of S2
var unmarshalMsgFieldOrder16zgensym_58b5c2649cecee38_17 = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip16zgensym_58b5c2649cecee38_17 = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zgensym_58b5c2649cecee38_13, zgensym_58b5c2649cecee38_14 := range z.R {
			_ = zgensym_58b5c2649cecee38_14
			_ = zgensym_58b5c2649cecee38_13
			s += msgp.StringPrefixSize + len(zgensym_58b5c2649cecee38_13) + msgp.Uint8Size
		}
	}
	s += 2 + msgp.Uint16Size + 2 + msgp.Uint32Size + 2 + msgp.Int64Size + 4 + msgp.ArrayHeaderSize + (6 * (msgp.Float64Size)) + 7
	if z.MyTree == nil {
		s += msgp.NilSize
	} else {
		s += z.MyTree.MSGPMsgsize()
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z Sys) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 1
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z Sys) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 1
	// string "F"
	o = append(o, 0x81, 0xa1, 0x46)
	o, err = msgp.AppendIntf(o, z.F)
	if err != nil {
		return
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Sys) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Sys) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields20zgensym_58b5c2649cecee38_21 = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields20zgensym_58b5c2649cecee38_21 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields20zgensym_58b5c2649cecee38_21, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft20zgensym_58b5c2649cecee38_21 := totalEncodedFields20zgensym_58b5c2649cecee38_21
	missingFieldsLeft20zgensym_58b5c2649cecee38_21 := maxFields20zgensym_58b5c2649cecee38_21 - totalEncodedFields20zgensym_58b5c2649cecee38_21

	var nextMiss20zgensym_58b5c2649cecee38_21 int32 = -1
	var found20zgensym_58b5c2649cecee38_21 [maxFields20zgensym_58b5c2649cecee38_21]bool
	var curField20zgensym_58b5c2649cecee38_21 string

doneWithStruct20zgensym_58b5c2649cecee38_21:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft20zgensym_58b5c2649cecee38_21 > 0 || missingFieldsLeft20zgensym_58b5c2649cecee38_21 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft20zgensym_58b5c2649cecee38_21, missingFieldsLeft20zgensym_58b5c2649cecee38_21, msgp.ShowFound(found20zgensym_58b5c2649cecee38_21[:]), unmarshalMsgFieldOrder20zgensym_58b5c2649cecee38_21)
		if encodedFieldsLeft20zgensym_58b5c2649cecee38_21 > 0 {
			encodedFieldsLeft20zgensym_58b5c2649cecee38_21--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField20zgensym_58b5c2649cecee38_21 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss20zgensym_58b5c2649cecee38_21 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss20zgensym_58b5c2649cecee38_21 = 0
			}
			for nextMiss20zgensym_58b5c2649cecee38_21 < maxFields20zgensym_58b5c2649cecee38_21 && (found20zgensym_58b5c2649cecee38_21[nextMiss20zgensym_58b5c2649cecee38_21] || unmarshalMsgFieldSkip20zgensym_58b5c2649cecee38_21[nextMiss20zgensym_58b5c2649cecee38_21]) {
				nextMiss20zgensym_58b5c2649cecee38_21++
			}
			if nextMiss20zgensym_58b5c2649cecee38_21 == maxFields20zgensym_58b5c2649cecee38_21 {
				// filled all the empty fields!
				break doneWithStruct20zgensym_58b5c2649cecee38_21
			}
			missingFieldsLeft20zgensym_58b5c2649cecee38_21--
			curField20zgensym_58b5c2649cecee38_21 = unmarshalMsgFieldOrder20zgensym_58b5c2649cecee38_21[nextMiss20zgensym_58b5c2649cecee38_21]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField20zgensym_58b5c2649cecee38_21)
		switch curField20zgensym_58b5c2649cecee38_21 {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found20zgensym_58b5c2649cecee38_21[0] = true
			z.F, bts, err = nbs.ReadIntfBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss20zgensym_58b5c2649cecee38_21 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Sys
var unmarshalMsgFieldOrder20zgensym_58b5c2649cecee38_21 = []string{"F"}

var unmarshalMsgFieldSkip20zgensym_58b5c2649cecee38_21 = []bool{false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) MSGPMsgsize() (s int) {
	s = 1 + 2 + msgp.GuessSize(z.F)
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Tree) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Tree) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 3
	// string "Chld"
	o = append(o, 0x83, 0xa4, 0x43, 0x68, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
	for zgensym_58b5c2649cecee38_22 := range z.Chld {
		o, err = z.Chld[zgensym_58b5c2649cecee38_22].MSGPMarshalMsg(o) // not is.iface, gen/marshal.go:261
		if err != nil {
			return
		}
	}
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	// string "Par"
	o = append(o, 0xa3, 0x50, 0x61, 0x72)
	if z.Par == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Par.MSGPMarshalMsg(o) // not is.iface, gen/marshal.go:261
		if err != nil {
			return
		}
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Tree) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Tree) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields23zgensym_58b5c2649cecee38_24 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields23zgensym_58b5c2649cecee38_24 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields23zgensym_58b5c2649cecee38_24, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft23zgensym_58b5c2649cecee38_24 := totalEncodedFields23zgensym_58b5c2649cecee38_24
	missingFieldsLeft23zgensym_58b5c2649cecee38_24 := maxFields23zgensym_58b5c2649cecee38_24 - totalEncodedFields23zgensym_58b5c2649cecee38_24

	var nextMiss23zgensym_58b5c2649cecee38_24 int32 = -1
	var found23zgensym_58b5c2649cecee38_24 [maxFields23zgensym_58b5c2649cecee38_24]bool
	var curField23zgensym_58b5c2649cecee38_24 string

doneWithStruct23zgensym_58b5c2649cecee38_24:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft23zgensym_58b5c2649cecee38_24 > 0 || missingFieldsLeft23zgensym_58b5c2649cecee38_24 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft23zgensym_58b5c2649cecee38_24, missingFieldsLeft23zgensym_58b5c2649cecee38_24, msgp.ShowFound(found23zgensym_58b5c2649cecee38_24[:]), unmarshalMsgFieldOrder23zgensym_58b5c2649cecee38_24)
		if encodedFieldsLeft23zgensym_58b5c2649cecee38_24 > 0 {
			encodedFieldsLeft23zgensym_58b5c2649cecee38_24--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField23zgensym_58b5c2649cecee38_24 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss23zgensym_58b5c2649cecee38_24 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss23zgensym_58b5c2649cecee38_24 = 0
			}
			for nextMiss23zgensym_58b5c2649cecee38_24 < maxFields23zgensym_58b5c2649cecee38_24 && (found23zgensym_58b5c2649cecee38_24[nextMiss23zgensym_58b5c2649cecee38_24] || unmarshalMsgFieldSkip23zgensym_58b5c2649cecee38_24[nextMiss23zgensym_58b5c2649cecee38_24]) {
				nextMiss23zgensym_58b5c2649cecee38_24++
			}
			if nextMiss23zgensym_58b5c2649cecee38_24 == maxFields23zgensym_58b5c2649cecee38_24 {
				// filled all the empty fields!
				break doneWithStruct23zgensym_58b5c2649cecee38_24
			}
			missingFieldsLeft23zgensym_58b5c2649cecee38_24--
			curField23zgensym_58b5c2649cecee38_24 = unmarshalMsgFieldOrder23zgensym_58b5c2649cecee38_24[nextMiss23zgensym_58b5c2649cecee38_24]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField23zgensym_58b5c2649cecee38_24)
		switch curField23zgensym_58b5c2649cecee38_24 {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found23zgensym_58b5c2649cecee38_24[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zgensym_58b5c2649cecee38_25 uint32
				zgensym_58b5c2649cecee38_25, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zgensym_58b5c2649cecee38_25) {
					z.Chld = (z.Chld)[:zgensym_58b5c2649cecee38_25]
				} else {
					z.Chld = make([]Tree, zgensym_58b5c2649cecee38_25)
				}
				for zgensym_58b5c2649cecee38_22 := range z.Chld {
					bts, err = z.Chld[zgensym_58b5c2649cecee38_22].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Str":
			found23zgensym_58b5c2649cecee38_24[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par":
			found23zgensym_58b5c2649cecee38_24[2] = true
			if nbs.AlwaysNil {
				if z.Par != nil {
					z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Par {
						z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Par == nil {
						z.Par = new(S2)
					}
					bts, err = z.Par.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss23zgensym_58b5c2649cecee38_24 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Tree
var unmarshalMsgFieldOrder23zgensym_58b5c2649cecee38_24 = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip23zgensym_58b5c2649cecee38_24 = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_22 := range z.Chld {
		s += z.Chld[zgensym_58b5c2649cecee38_22].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
