package types

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Button) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zkbz = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zkbz uint32
	totalEncodedFields0zkbz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zkbz := totalEncodedFields0zkbz
	missingFieldsLeft0zkbz := maxFields0zkbz - totalEncodedFields0zkbz

	var nextMiss0zkbz int32 = -1
	var found0zkbz [maxFields0zkbz]bool
	var curField0zkbz string

doneWithStruct0zkbz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zkbz > 0 || missingFieldsLeft0zkbz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zkbz, missingFieldsLeft0zkbz, msgp.ShowFound(found0zkbz[:]), decodeMsgFieldOrder0zkbz)
		if encodedFieldsLeft0zkbz > 0 {
			encodedFieldsLeft0zkbz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zkbz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zkbz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zkbz = 0
			}
			for nextMiss0zkbz < maxFields0zkbz && (found0zkbz[nextMiss0zkbz] || decodeMsgFieldSkip0zkbz[nextMiss0zkbz]) {
				nextMiss0zkbz++
			}
			if nextMiss0zkbz == maxFields0zkbz {
				// filled all the empty fields!
				break doneWithStruct0zkbz
			}
			missingFieldsLeft0zkbz--
			curField0zkbz = decodeMsgFieldOrder0zkbz[nextMiss0zkbz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zkbz)
		switch curField0zkbz {
		// -- templateDecodeMsg ends here --

		case "Text":
			found0zkbz[0] = true
			z.Text, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Command":
			found0zkbz[1] = true
			z.Command, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zkbz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Button
var decodeMsgFieldOrder0zkbz = []string{"Text", "Command"}

var decodeMsgFieldSkip0zkbz = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z Button) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z Button) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Text"
	err = en.Append(0x82, 0xa4, 0x54, 0x65, 0x78, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Text)
	if err != nil {
		panic(err)
	}
	// write "Command"
	err = en.Append(0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Command)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Button) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Text"
	o = append(o, 0x82, 0xa4, 0x54, 0x65, 0x78, 0x74)
	o = msgp.AppendString(o, z.Text)
	// string "Command"
	o = append(o, 0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64)
	o = msgp.AppendString(o, z.Command)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Button) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Button) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1znaw = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1znaw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1znaw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1znaw := totalEncodedFields1znaw
	missingFieldsLeft1znaw := maxFields1znaw - totalEncodedFields1znaw

	var nextMiss1znaw int32 = -1
	var found1znaw [maxFields1znaw]bool
	var curField1znaw string

doneWithStruct1znaw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1znaw > 0 || missingFieldsLeft1znaw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1znaw, missingFieldsLeft1znaw, msgp.ShowFound(found1znaw[:]), unmarshalMsgFieldOrder1znaw)
		if encodedFieldsLeft1znaw > 0 {
			encodedFieldsLeft1znaw--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1znaw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1znaw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1znaw = 0
			}
			for nextMiss1znaw < maxFields1znaw && (found1znaw[nextMiss1znaw] || unmarshalMsgFieldSkip1znaw[nextMiss1znaw]) {
				nextMiss1znaw++
			}
			if nextMiss1znaw == maxFields1znaw {
				// filled all the empty fields!
				break doneWithStruct1znaw
			}
			missingFieldsLeft1znaw--
			curField1znaw = unmarshalMsgFieldOrder1znaw[nextMiss1znaw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1znaw)
		switch curField1znaw {
		// -- templateUnmarshalMsg ends here --

		case "Text":
			found1znaw[0] = true
			z.Text, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Command":
			found1znaw[1] = true
			z.Command, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1znaw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Button
var unmarshalMsgFieldOrder1znaw = []string{"Text", "Command"}

var unmarshalMsgFieldSkip1znaw = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Button) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Text) + 8 + msgp.StringPrefixSize + len(z.Command)
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *InboundChatMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields2zwtl = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zwtl uint32
	totalEncodedFields2zwtl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zwtl := totalEncodedFields2zwtl
	missingFieldsLeft2zwtl := maxFields2zwtl - totalEncodedFields2zwtl

	var nextMiss2zwtl int32 = -1
	var found2zwtl [maxFields2zwtl]bool
	var curField2zwtl string

doneWithStruct2zwtl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zwtl > 0 || missingFieldsLeft2zwtl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zwtl, missingFieldsLeft2zwtl, msgp.ShowFound(found2zwtl[:]), decodeMsgFieldOrder2zwtl)
		if encodedFieldsLeft2zwtl > 0 {
			encodedFieldsLeft2zwtl--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zwtl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zwtl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zwtl = 0
			}
			for nextMiss2zwtl < maxFields2zwtl && (found2zwtl[nextMiss2zwtl] || decodeMsgFieldSkip2zwtl[nextMiss2zwtl]) {
				nextMiss2zwtl++
			}
			if nextMiss2zwtl == maxFields2zwtl {
				// filled all the empty fields!
				break doneWithStruct2zwtl
			}
			missingFieldsLeft2zwtl--
			curField2zwtl = decodeMsgFieldOrder2zwtl[nextMiss2zwtl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zwtl)
		switch curField2zwtl {
		// -- templateDecodeMsg ends here --

		case "0":
			found2zwtl[0] = true
			z.SenderID, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "1":
			found2zwtl[1] = true
			z.SenderName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "2":
			found2zwtl[2] = true
			z.Message, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "3":
			found2zwtl[3] = true
			var zzkx uint32
			zzkx, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Command) >= int(zzkx) {
				z.Command = (z.Command)[:zzkx]
			} else {
				z.Command = make([]string, zzkx)
			}
			for zsik := range z.Command {
				z.Command[zsik], err = dc.ReadString()
				if err != nil {
					panic(err)
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss2zwtl != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InboundChatMessage
var decodeMsgFieldOrder2zwtl = []string{"0", "1", "2", "3"}

var decodeMsgFieldSkip2zwtl = []bool{false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *InboundChatMessage) fieldsNotEmpty(isempty []bool) uint32 {
	return 4
}

// EncodeMsg implements msgp.Encodable
func (z *InboundChatMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "0"
	err = en.Append(0x84, 0xa1, 0x30)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderID)
	if err != nil {
		panic(err)
	}
	// write "1"
	err = en.Append(0xa1, 0x31)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		panic(err)
	}
	// write "2"
	err = en.Append(0xa1, 0x32)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Message)
	if err != nil {
		panic(err)
	}
	// write "3"
	err = en.Append(0xa1, 0x33)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Command)))
	if err != nil {
		panic(err)
	}
	for zsik := range z.Command {
		err = en.WriteString(z.Command[zsik])
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *InboundChatMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "0"
	o = append(o, 0x84, 0xa1, 0x30)
	o = msgp.AppendString(o, z.SenderID)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.SenderName)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendString(o, z.Message)
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Command)))
	for zsik := range z.Command {
		o = msgp.AppendString(o, z.Command[zsik])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InboundChatMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *InboundChatMessage) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zlrh = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zlrh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zlrh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zlrh := totalEncodedFields3zlrh
	missingFieldsLeft3zlrh := maxFields3zlrh - totalEncodedFields3zlrh

	var nextMiss3zlrh int32 = -1
	var found3zlrh [maxFields3zlrh]bool
	var curField3zlrh string

doneWithStruct3zlrh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zlrh > 0 || missingFieldsLeft3zlrh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zlrh, missingFieldsLeft3zlrh, msgp.ShowFound(found3zlrh[:]), unmarshalMsgFieldOrder3zlrh)
		if encodedFieldsLeft3zlrh > 0 {
			encodedFieldsLeft3zlrh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zlrh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zlrh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zlrh = 0
			}
			for nextMiss3zlrh < maxFields3zlrh && (found3zlrh[nextMiss3zlrh] || unmarshalMsgFieldSkip3zlrh[nextMiss3zlrh]) {
				nextMiss3zlrh++
			}
			if nextMiss3zlrh == maxFields3zlrh {
				// filled all the empty fields!
				break doneWithStruct3zlrh
			}
			missingFieldsLeft3zlrh--
			curField3zlrh = unmarshalMsgFieldOrder3zlrh[nextMiss3zlrh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zlrh)
		switch curField3zlrh {
		// -- templateUnmarshalMsg ends here --

		case "0":
			found3zlrh[0] = true
			z.SenderID, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "1":
			found3zlrh[1] = true
			z.SenderName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "2":
			found3zlrh[2] = true
			z.Message, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "3":
			found3zlrh[3] = true
			if nbs.AlwaysNil {
				(z.Command) = (z.Command)[:0]
			} else {

				var ziqq uint32
				ziqq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Command) >= int(ziqq) {
					z.Command = (z.Command)[:ziqq]
				} else {
					z.Command = make([]string, ziqq)
				}
				for zsik := range z.Command {
					z.Command[zsik], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss3zlrh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InboundChatMessage
var unmarshalMsgFieldOrder3zlrh = []string{"0", "1", "2", "3"}

var unmarshalMsgFieldSkip3zlrh = []bool{false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InboundChatMessage) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.SenderID) + 2 + msgp.StringPrefixSize + len(z.SenderName) + 2 + msgp.StringPrefixSize + len(z.Message) + 2 + msgp.ArrayHeaderSize
	for zsik := range z.Command {
		s += msgp.StringPrefixSize + len(z.Command[zsik])
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Key) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zmkc string
		zmkc, err = dc.ReadString()
		(*z) = Key(zmkc)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z Key) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Key) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Key) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Key) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zybm string
		zybm, bts, err = nbs.ReadStringBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Key(zybm)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Key) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Location) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields4zhak = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zhak uint32
	totalEncodedFields4zhak, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zhak := totalEncodedFields4zhak
	missingFieldsLeft4zhak := maxFields4zhak - totalEncodedFields4zhak

	var nextMiss4zhak int32 = -1
	var found4zhak [maxFields4zhak]bool
	var curField4zhak string

doneWithStruct4zhak:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zhak > 0 || missingFieldsLeft4zhak > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zhak, missingFieldsLeft4zhak, msgp.ShowFound(found4zhak[:]), decodeMsgFieldOrder4zhak)
		if encodedFieldsLeft4zhak > 0 {
			encodedFieldsLeft4zhak--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zhak = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zhak < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zhak = 0
			}
			for nextMiss4zhak < maxFields4zhak && (found4zhak[nextMiss4zhak] || decodeMsgFieldSkip4zhak[nextMiss4zhak]) {
				nextMiss4zhak++
			}
			if nextMiss4zhak == maxFields4zhak {
				// filled all the empty fields!
				break doneWithStruct4zhak
			}
			missingFieldsLeft4zhak--
			curField4zhak = decodeMsgFieldOrder4zhak[nextMiss4zhak]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zhak)
		switch curField4zhak {
		// -- templateDecodeMsg ends here --

		case "Lat":
			found4zhak[0] = true
			z.Lat, err = dc.ReadFloat64()
			if err != nil {
				panic(err)
			}
		case "Lng":
			found4zhak[1] = true
			z.Lng, err = dc.ReadFloat64()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss4zhak != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Location
var decodeMsgFieldOrder4zhak = []string{"Lat", "Lng"}

var decodeMsgFieldSkip4zhak = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z Location) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z Location) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Lat"
	err = en.Append(0x82, 0xa3, 0x4c, 0x61, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Lat)
	if err != nil {
		panic(err)
	}
	// write "Lng"
	err = en.Append(0xa3, 0x4c, 0x6e, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Lng)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Location) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Lat"
	o = append(o, 0x82, 0xa3, 0x4c, 0x61, 0x74)
	o = msgp.AppendFloat64(o, z.Lat)
	// string "Lng"
	o = append(o, 0xa3, 0x4c, 0x6e, 0x67)
	o = msgp.AppendFloat64(o, z.Lng)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Location) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Location) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields5zhsp = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zhsp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zhsp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zhsp := totalEncodedFields5zhsp
	missingFieldsLeft5zhsp := maxFields5zhsp - totalEncodedFields5zhsp

	var nextMiss5zhsp int32 = -1
	var found5zhsp [maxFields5zhsp]bool
	var curField5zhsp string

doneWithStruct5zhsp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zhsp > 0 || missingFieldsLeft5zhsp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zhsp, missingFieldsLeft5zhsp, msgp.ShowFound(found5zhsp[:]), unmarshalMsgFieldOrder5zhsp)
		if encodedFieldsLeft5zhsp > 0 {
			encodedFieldsLeft5zhsp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zhsp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zhsp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zhsp = 0
			}
			for nextMiss5zhsp < maxFields5zhsp && (found5zhsp[nextMiss5zhsp] || unmarshalMsgFieldSkip5zhsp[nextMiss5zhsp]) {
				nextMiss5zhsp++
			}
			if nextMiss5zhsp == maxFields5zhsp {
				// filled all the empty fields!
				break doneWithStruct5zhsp
			}
			missingFieldsLeft5zhsp--
			curField5zhsp = unmarshalMsgFieldOrder5zhsp[nextMiss5zhsp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zhsp)
		switch curField5zhsp {
		// -- templateUnmarshalMsg ends here --

		case "Lat":
			found5zhsp[0] = true
			z.Lat, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Lng":
			found5zhsp[1] = true
			z.Lng, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss5zhsp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Location
var unmarshalMsgFieldOrder5zhsp = []string{"Lat", "Lng"}

var unmarshalMsgFieldSkip5zhsp = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Location) Msgsize() (s int) {
	s = 1 + 4 + msgp.Float64Size + 4 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Meetup) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields6znuh = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6znuh uint32
	totalEncodedFields6znuh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6znuh := totalEncodedFields6znuh
	missingFieldsLeft6znuh := maxFields6znuh - totalEncodedFields6znuh

	var nextMiss6znuh int32 = -1
	var found6znuh [maxFields6znuh]bool
	var curField6znuh string

doneWithStruct6znuh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6znuh > 0 || missingFieldsLeft6znuh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6znuh, missingFieldsLeft6znuh, msgp.ShowFound(found6znuh[:]), decodeMsgFieldOrder6znuh)
		if encodedFieldsLeft6znuh > 0 {
			encodedFieldsLeft6znuh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6znuh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6znuh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6znuh = 0
			}
			for nextMiss6znuh < maxFields6znuh && (found6znuh[nextMiss6znuh] || decodeMsgFieldSkip6znuh[nextMiss6znuh]) {
				nextMiss6znuh++
			}
			if nextMiss6znuh == maxFields6znuh {
				// filled all the empty fields!
				break doneWithStruct6znuh
			}
			missingFieldsLeft6znuh--
			curField6znuh = decodeMsgFieldOrder6znuh[nextMiss6znuh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6znuh)
		switch curField6znuh {
		// -- templateDecodeMsg ends here --

		case "ETag":
			found6znuh[0] = true
			z.ETag, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Title":
			found6znuh[1] = true
			z.Title, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Where":
			found6znuh[2] = true
			z.Where, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "When":
			found6znuh[3] = true
			z.When, err = dc.ReadTime()
			if err != nil {
				panic(err)
			}
		case "Map":
			found6znuh[4] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				z.Map = nil
			} else {
				if z.Map == nil {
					z.Map = new(Location)
				}
				const maxFields7zhrf = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields7zhrf uint32
				totalEncodedFields7zhrf, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft7zhrf := totalEncodedFields7zhrf
				missingFieldsLeft7zhrf := maxFields7zhrf - totalEncodedFields7zhrf

				var nextMiss7zhrf int32 = -1
				var found7zhrf [maxFields7zhrf]bool
				var curField7zhrf string

			doneWithStruct7zhrf:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft7zhrf > 0 || missingFieldsLeft7zhrf > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zhrf, missingFieldsLeft7zhrf, msgp.ShowFound(found7zhrf[:]), decodeMsgFieldOrder7zhrf)
					if encodedFieldsLeft7zhrf > 0 {
						encodedFieldsLeft7zhrf--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField7zhrf = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss7zhrf < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss7zhrf = 0
						}
						for nextMiss7zhrf < maxFields7zhrf && (found7zhrf[nextMiss7zhrf] || decodeMsgFieldSkip7zhrf[nextMiss7zhrf]) {
							nextMiss7zhrf++
						}
						if nextMiss7zhrf == maxFields7zhrf {
							// filled all the empty fields!
							break doneWithStruct7zhrf
						}
						missingFieldsLeft7zhrf--
						curField7zhrf = decodeMsgFieldOrder7zhrf[nextMiss7zhrf]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField7zhrf)
					switch curField7zhrf {
					// -- templateDecodeMsg ends here --

					case "Lat":
						found7zhrf[0] = true
						z.Map.Lat, err = dc.ReadFloat64()
						if err != nil {
							panic(err)
						}
					case "Lng":
						found7zhrf[1] = true
						z.Map.Lng, err = dc.ReadFloat64()
						if err != nil {
							panic(err)
						}
					default:
						err = dc.Skip()
						if err != nil {
							panic(err)
						}
					}
				}
				if nextMiss7zhrf != -1 {
					dc.PopAlwaysNil()
				}

			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss6znuh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Meetup
var decodeMsgFieldOrder6znuh = []string{"ETag", "Title", "Where", "When", "Map"}

var decodeMsgFieldSkip6znuh = []bool{false, false, false, false, false}

// fields of Location
var decodeMsgFieldOrder7zhrf = []string{"Lat", "Lng"}

var decodeMsgFieldSkip7zhrf = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Meetup) fieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// EncodeMsg implements msgp.Encodable
func (z *Meetup) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "ETag"
	err = en.Append(0x85, 0xa4, 0x45, 0x54, 0x61, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ETag)
	if err != nil {
		panic(err)
	}
	// write "Title"
	err = en.Append(0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Title)
	if err != nil {
		panic(err)
	}
	// write "Where"
	err = en.Append(0xa5, 0x57, 0x68, 0x65, 0x72, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Where)
	if err != nil {
		panic(err)
	}
	// write "When"
	err = en.Append(0xa4, 0x57, 0x68, 0x65, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.When)
	if err != nil {
		panic(err)
	}
	// write "Map"
	err = en.Append(0xa3, 0x4d, 0x61, 0x70)
	if err != nil {
		return err
	}
	if z.Map == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "Lat"
		err = en.Append(0x82, 0xa3, 0x4c, 0x61, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.Map.Lat)
		if err != nil {
			panic(err)
		}
		// write "Lng"
		err = en.Append(0xa3, 0x4c, 0x6e, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.Map.Lng)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Meetup) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "ETag"
	o = append(o, 0x85, 0xa4, 0x45, 0x54, 0x61, 0x67)
	o = msgp.AppendString(o, z.ETag)
	// string "Title"
	o = append(o, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "Where"
	o = append(o, 0xa5, 0x57, 0x68, 0x65, 0x72, 0x65)
	o = msgp.AppendString(o, z.Where)
	// string "When"
	o = append(o, 0xa4, 0x57, 0x68, 0x65, 0x6e)
	o = msgp.AppendTime(o, z.When)
	// string "Map"
	o = append(o, 0xa3, 0x4d, 0x61, 0x70)
	if z.Map == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "Lat"
		o = append(o, 0x82, 0xa3, 0x4c, 0x61, 0x74)
		o = msgp.AppendFloat64(o, z.Map.Lat)
		// string "Lng"
		o = append(o, 0xa3, 0x4c, 0x6e, 0x67)
		o = msgp.AppendFloat64(o, z.Map.Lng)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Meetup) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Meetup) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields8zqxd = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields8zqxd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields8zqxd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft8zqxd := totalEncodedFields8zqxd
	missingFieldsLeft8zqxd := maxFields8zqxd - totalEncodedFields8zqxd

	var nextMiss8zqxd int32 = -1
	var found8zqxd [maxFields8zqxd]bool
	var curField8zqxd string

doneWithStruct8zqxd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zqxd > 0 || missingFieldsLeft8zqxd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zqxd, missingFieldsLeft8zqxd, msgp.ShowFound(found8zqxd[:]), unmarshalMsgFieldOrder8zqxd)
		if encodedFieldsLeft8zqxd > 0 {
			encodedFieldsLeft8zqxd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField8zqxd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zqxd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss8zqxd = 0
			}
			for nextMiss8zqxd < maxFields8zqxd && (found8zqxd[nextMiss8zqxd] || unmarshalMsgFieldSkip8zqxd[nextMiss8zqxd]) {
				nextMiss8zqxd++
			}
			if nextMiss8zqxd == maxFields8zqxd {
				// filled all the empty fields!
				break doneWithStruct8zqxd
			}
			missingFieldsLeft8zqxd--
			curField8zqxd = unmarshalMsgFieldOrder8zqxd[nextMiss8zqxd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zqxd)
		switch curField8zqxd {
		// -- templateUnmarshalMsg ends here --

		case "ETag":
			found8zqxd[0] = true
			z.ETag, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Title":
			found8zqxd[1] = true
			z.Title, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Where":
			found8zqxd[2] = true
			z.Where, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "When":
			found8zqxd[3] = true
			z.When, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Map":
			found8zqxd[4] = true
			// default gPtr logic.
			if nbs.PeekNil(bts) && z.Map == nil {
				// consume the nil
				bts, err = nbs.ReadNilBytes(bts)
				if err != nil {
					return
				}
			} else {
				// read as-if the wire has bytes, letting nbs take care of nils.

				if z.Map == nil {
					z.Map = new(Location)
				}
				const maxFields9zeiv = 2

				// -- templateUnmarshalMsg starts here--
				var totalEncodedFields9zeiv uint32
				if !nbs.AlwaysNil {
					totalEncodedFields9zeiv, bts, err = nbs.ReadMapHeaderBytes(bts)
					if err != nil {
						panic(err)
						return
					}
				}
				encodedFieldsLeft9zeiv := totalEncodedFields9zeiv
				missingFieldsLeft9zeiv := maxFields9zeiv - totalEncodedFields9zeiv

				var nextMiss9zeiv int32 = -1
				var found9zeiv [maxFields9zeiv]bool
				var curField9zeiv string

			doneWithStruct9zeiv:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft9zeiv > 0 || missingFieldsLeft9zeiv > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zeiv, missingFieldsLeft9zeiv, msgp.ShowFound(found9zeiv[:]), unmarshalMsgFieldOrder9zeiv)
					if encodedFieldsLeft9zeiv > 0 {
						encodedFieldsLeft9zeiv--
						field, bts, err = nbs.ReadMapKeyZC(bts)
						if err != nil {
							panic(err)
							return
						}
						curField9zeiv = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss9zeiv < 0 {
							// set bts to contain just mnil (0xc0)
							bts = nbs.PushAlwaysNil(bts)
							nextMiss9zeiv = 0
						}
						for nextMiss9zeiv < maxFields9zeiv && (found9zeiv[nextMiss9zeiv] || unmarshalMsgFieldSkip9zeiv[nextMiss9zeiv]) {
							nextMiss9zeiv++
						}
						if nextMiss9zeiv == maxFields9zeiv {
							// filled all the empty fields!
							break doneWithStruct9zeiv
						}
						missingFieldsLeft9zeiv--
						curField9zeiv = unmarshalMsgFieldOrder9zeiv[nextMiss9zeiv]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField9zeiv)
					switch curField9zeiv {
					// -- templateUnmarshalMsg ends here --

					case "Lat":
						found9zeiv[0] = true
						z.Map.Lat, bts, err = nbs.ReadFloat64Bytes(bts)

						if err != nil {
							panic(err)
						}
					case "Lng":
						found9zeiv[1] = true
						z.Map.Lng, bts, err = nbs.ReadFloat64Bytes(bts)

						if err != nil {
							panic(err)
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							panic(err)
						}
					}
				}
				if nextMiss9zeiv != -1 {
					bts = nbs.PopAlwaysNil()
				}

			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss8zqxd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Meetup
var unmarshalMsgFieldOrder8zqxd = []string{"ETag", "Title", "Where", "When", "Map"}

var unmarshalMsgFieldSkip8zqxd = []bool{false, false, false, false, false}

// fields of Location
var unmarshalMsgFieldOrder9zeiv = []string{"Lat", "Lng"}

var unmarshalMsgFieldSkip9zeiv = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Meetup) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.ETag) + 6 + msgp.StringPrefixSize + len(z.Title) + 6 + msgp.StringPrefixSize + len(z.Where) + 5 + msgp.TimeSize + 4
	if z.Map == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 4 + msgp.Float64Size + 4 + msgp.Float64Size
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MeetupAuthCompletion) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields10zwcp = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zwcp uint32
	totalEncodedFields10zwcp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zwcp := totalEncodedFields10zwcp
	missingFieldsLeft10zwcp := maxFields10zwcp - totalEncodedFields10zwcp

	var nextMiss10zwcp int32 = -1
	var found10zwcp [maxFields10zwcp]bool
	var curField10zwcp string

doneWithStruct10zwcp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zwcp > 0 || missingFieldsLeft10zwcp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zwcp, missingFieldsLeft10zwcp, msgp.ShowFound(found10zwcp[:]), decodeMsgFieldOrder10zwcp)
		if encodedFieldsLeft10zwcp > 0 {
			encodedFieldsLeft10zwcp--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zwcp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zwcp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zwcp = 0
			}
			for nextMiss10zwcp < maxFields10zwcp && (found10zwcp[nextMiss10zwcp] || decodeMsgFieldSkip10zwcp[nextMiss10zwcp]) {
				nextMiss10zwcp++
			}
			if nextMiss10zwcp == maxFields10zwcp {
				// filled all the empty fields!
				break doneWithStruct10zwcp
			}
			missingFieldsLeft10zwcp--
			curField10zwcp = decodeMsgFieldOrder10zwcp[nextMiss10zwcp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zwcp)
		switch curField10zwcp {
		// -- templateDecodeMsg ends here --

		case "State":
			found10zwcp[0] = true
			{
				var zmwu []byte
				zmwu, err = dc.ReadBytes([]byte(z.State))
				z.State = MeetupAuthState(zmwu)
			}
			if err != nil {
				panic(err)
			}
		case "Tokens":
			found10zwcp[1] = true
			{
				var zfig []byte
				zfig, err = dc.ReadBytes([]byte(z.Tokens))
				z.Tokens = MeetupAuthTokens(zfig)
			}
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss10zwcp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MeetupAuthCompletion
var decodeMsgFieldOrder10zwcp = []string{"State", "Tokens"}

var decodeMsgFieldSkip10zwcp = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *MeetupAuthCompletion) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *MeetupAuthCompletion) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "State"
	err = en.Append(0x82, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes([]byte(z.State))
	if err != nil {
		panic(err)
	}
	// write "Tokens"
	err = en.Append(0xa6, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteBytes([]byte(z.Tokens))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MeetupAuthCompletion) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "State"
	o = append(o, 0x82, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendBytes(o, []byte(z.State))
	// string "Tokens"
	o = append(o, 0xa6, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73)
	o = msgp.AppendBytes(o, []byte(z.Tokens))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MeetupAuthCompletion) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MeetupAuthCompletion) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields11zmvv = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields11zmvv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields11zmvv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft11zmvv := totalEncodedFields11zmvv
	missingFieldsLeft11zmvv := maxFields11zmvv - totalEncodedFields11zmvv

	var nextMiss11zmvv int32 = -1
	var found11zmvv [maxFields11zmvv]bool
	var curField11zmvv string

doneWithStruct11zmvv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft11zmvv > 0 || missingFieldsLeft11zmvv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zmvv, missingFieldsLeft11zmvv, msgp.ShowFound(found11zmvv[:]), unmarshalMsgFieldOrder11zmvv)
		if encodedFieldsLeft11zmvv > 0 {
			encodedFieldsLeft11zmvv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField11zmvv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss11zmvv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss11zmvv = 0
			}
			for nextMiss11zmvv < maxFields11zmvv && (found11zmvv[nextMiss11zmvv] || unmarshalMsgFieldSkip11zmvv[nextMiss11zmvv]) {
				nextMiss11zmvv++
			}
			if nextMiss11zmvv == maxFields11zmvv {
				// filled all the empty fields!
				break doneWithStruct11zmvv
			}
			missingFieldsLeft11zmvv--
			curField11zmvv = unmarshalMsgFieldOrder11zmvv[nextMiss11zmvv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField11zmvv)
		switch curField11zmvv {
		// -- templateUnmarshalMsg ends here --

		case "State":
			found11zmvv[0] = true
			{
				var zesw []byte
				if nbs.AlwaysNil || msgp.IsNil(bts) {
					if !nbs.AlwaysNil {
						bts = bts[1:]
					}
					zesw = zesw[:0]
				} else {
					zesw, bts, err = nbs.ReadBytesBytes(bts, []byte(z.State))

					if err != nil {
						panic(err)
					}
				}
				if err != nil {
					panic(err)
				}
				z.State = MeetupAuthState(zesw)
			}
		case "Tokens":
			found11zmvv[1] = true
			{
				var znnh []byte
				if nbs.AlwaysNil || msgp.IsNil(bts) {
					if !nbs.AlwaysNil {
						bts = bts[1:]
					}
					znnh = znnh[:0]
				} else {
					znnh, bts, err = nbs.ReadBytesBytes(bts, []byte(z.Tokens))

					if err != nil {
						panic(err)
					}
				}
				if err != nil {
					panic(err)
				}
				z.Tokens = MeetupAuthTokens(znnh)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss11zmvv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MeetupAuthCompletion
var unmarshalMsgFieldOrder11zmvv = []string{"State", "Tokens"}

var unmarshalMsgFieldSkip11zmvv = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MeetupAuthCompletion) Msgsize() (s int) {
	s = 1 + 6 + msgp.BytesPrefixSize + len([]byte(z.State)) + 7 + msgp.BytesPrefixSize + len([]byte(z.Tokens))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MeetupAuthState) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zywi []byte
		zywi, err = dc.ReadBytes([]byte((*z)))
		(*z) = MeetupAuthState(zywi)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z MeetupAuthState) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes([]byte(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MeetupAuthState) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MeetupAuthState) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MeetupAuthState) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zstd []byte
		if nbs.AlwaysNil || msgp.IsNil(bts) {
			if !nbs.AlwaysNil {
				bts = bts[1:]
			}
			zstd = zstd[:0]
		} else {
			zstd, bts, err = nbs.ReadBytesBytes(bts, []byte((*z)))

			if err != nil {
				panic(err)
			}
		}
		if err != nil {
			panic(err)
		}
		(*z) = MeetupAuthState(zstd)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MeetupAuthState) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MeetupAuthTokens) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zcvl []byte
		zcvl, err = dc.ReadBytes([]byte((*z)))
		(*z) = MeetupAuthTokens(zcvl)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z MeetupAuthTokens) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes([]byte(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MeetupAuthTokens) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MeetupAuthTokens) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MeetupAuthTokens) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var ztlg []byte
		if nbs.AlwaysNil || msgp.IsNil(bts) {
			if !nbs.AlwaysNil {
				bts = bts[1:]
			}
			ztlg = ztlg[:0]
		} else {
			ztlg, bts, err = nbs.ReadBytesBytes(bts, []byte((*z)))

			if err != nil {
				panic(err)
			}
		}
		if err != nil {
			panic(err)
		}
		(*z) = MeetupAuthTokens(ztlg)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MeetupAuthTokens) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MeetupGroup) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields12zrjh = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zrjh uint32
	totalEncodedFields12zrjh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zrjh := totalEncodedFields12zrjh
	missingFieldsLeft12zrjh := maxFields12zrjh - totalEncodedFields12zrjh

	var nextMiss12zrjh int32 = -1
	var found12zrjh [maxFields12zrjh]bool
	var curField12zrjh string

doneWithStruct12zrjh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zrjh > 0 || missingFieldsLeft12zrjh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zrjh, missingFieldsLeft12zrjh, msgp.ShowFound(found12zrjh[:]), decodeMsgFieldOrder12zrjh)
		if encodedFieldsLeft12zrjh > 0 {
			encodedFieldsLeft12zrjh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zrjh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zrjh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zrjh = 0
			}
			for nextMiss12zrjh < maxFields12zrjh && (found12zrjh[nextMiss12zrjh] || decodeMsgFieldSkip12zrjh[nextMiss12zrjh]) {
				nextMiss12zrjh++
			}
			if nextMiss12zrjh == maxFields12zrjh {
				// filled all the empty fields!
				break doneWithStruct12zrjh
			}
			missingFieldsLeft12zrjh--
			curField12zrjh = decodeMsgFieldOrder12zrjh[nextMiss12zrjh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zrjh)
		switch curField12zrjh {
		// -- templateDecodeMsg ends here --

		case "Name":
			found12zrjh[0] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Urlname":
			found12zrjh[1] = true
			z.Urlname, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ID":
			found12zrjh[2] = true
			z.ID, err = dc.ReadInt()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zrjh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MeetupGroup
var decodeMsgFieldOrder12zrjh = []string{"Name", "Urlname", "ID"}

var decodeMsgFieldSkip12zrjh = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z MeetupGroup) fieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// EncodeMsg implements msgp.Encodable
func (z MeetupGroup) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Name"
	err = en.Append(0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		panic(err)
	}
	// write "Urlname"
	err = en.Append(0xa7, 0x55, 0x72, 0x6c, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Urlname)
	if err != nil {
		panic(err)
	}
	// write "ID"
	err = en.Append(0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ID)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MeetupGroup) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Name"
	o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Urlname"
	o = append(o, 0xa7, 0x55, 0x72, 0x6c, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Urlname)
	// string "ID"
	o = append(o, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MeetupGroup) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MeetupGroup) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields13zyps = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields13zyps uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zyps, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft13zyps := totalEncodedFields13zyps
	missingFieldsLeft13zyps := maxFields13zyps - totalEncodedFields13zyps

	var nextMiss13zyps int32 = -1
	var found13zyps [maxFields13zyps]bool
	var curField13zyps string

doneWithStruct13zyps:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zyps > 0 || missingFieldsLeft13zyps > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zyps, missingFieldsLeft13zyps, msgp.ShowFound(found13zyps[:]), unmarshalMsgFieldOrder13zyps)
		if encodedFieldsLeft13zyps > 0 {
			encodedFieldsLeft13zyps--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField13zyps = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss13zyps < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zyps = 0
			}
			for nextMiss13zyps < maxFields13zyps && (found13zyps[nextMiss13zyps] || unmarshalMsgFieldSkip13zyps[nextMiss13zyps]) {
				nextMiss13zyps++
			}
			if nextMiss13zyps == maxFields13zyps {
				// filled all the empty fields!
				break doneWithStruct13zyps
			}
			missingFieldsLeft13zyps--
			curField13zyps = unmarshalMsgFieldOrder13zyps[nextMiss13zyps]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zyps)
		switch curField13zyps {
		// -- templateUnmarshalMsg ends here --

		case "Name":
			found13zyps[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Urlname":
			found13zyps[1] = true
			z.Urlname, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ID":
			found13zyps[2] = true
			z.ID, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss13zyps != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MeetupGroup
var unmarshalMsgFieldOrder13zyps = []string{"Name", "Urlname", "ID"}

var unmarshalMsgFieldSkip13zyps = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MeetupGroup) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Urlname) + 3 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MeetupID) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zwlt string
		zwlt, err = dc.ReadString()
		(*z) = MeetupID(zwlt)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z MeetupID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MeetupID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MeetupID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MeetupID) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var znhr string
		znhr, bts, err = nbs.ReadStringBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = MeetupID(znhr)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MeetupID) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MeetupUserLink) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields14zegh = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields14zegh uint32
	totalEncodedFields14zegh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zegh := totalEncodedFields14zegh
	missingFieldsLeft14zegh := maxFields14zegh - totalEncodedFields14zegh

	var nextMiss14zegh int32 = -1
	var found14zegh [maxFields14zegh]bool
	var curField14zegh string

doneWithStruct14zegh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zegh > 0 || missingFieldsLeft14zegh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zegh, missingFieldsLeft14zegh, msgp.ShowFound(found14zegh[:]), decodeMsgFieldOrder14zegh)
		if encodedFieldsLeft14zegh > 0 {
			encodedFieldsLeft14zegh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField14zegh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zegh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zegh = 0
			}
			for nextMiss14zegh < maxFields14zegh && (found14zegh[nextMiss14zegh] || decodeMsgFieldSkip14zegh[nextMiss14zegh]) {
				nextMiss14zegh++
			}
			if nextMiss14zegh == maxFields14zegh {
				// filled all the empty fields!
				break doneWithStruct14zegh
			}
			missingFieldsLeft14zegh--
			curField14zegh = decodeMsgFieldOrder14zegh[nextMiss14zegh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zegh)
		switch curField14zegh {
		// -- templateDecodeMsg ends here --

		case "MeetupID":
			found14zegh[0] = true
			{
				var zafj string
				zafj, err = dc.ReadString()
				z.MeetupID = MeetupID(zafj)
			}
			if err != nil {
				panic(err)
			}
		case "UserID":
			found14zegh[1] = true
			{
				var zdos string
				zdos, err = dc.ReadString()
				z.UserID = UserID(zdos)
			}
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zegh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MeetupUserLink
var decodeMsgFieldOrder14zegh = []string{"MeetupID", "UserID"}

var decodeMsgFieldSkip14zegh = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z MeetupUserLink) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z MeetupUserLink) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "MeetupID"
	err = en.Append(0x82, 0xa8, 0x4d, 0x65, 0x65, 0x74, 0x75, 0x70, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(string(z.MeetupID))
	if err != nil {
		panic(err)
	}
	// write "UserID"
	err = en.Append(0xa6, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(string(z.UserID))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MeetupUserLink) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "MeetupID"
	o = append(o, 0x82, 0xa8, 0x4d, 0x65, 0x65, 0x74, 0x75, 0x70, 0x49, 0x44)
	o = msgp.AppendString(o, string(z.MeetupID))
	// string "UserID"
	o = append(o, 0xa6, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44)
	o = msgp.AppendString(o, string(z.UserID))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MeetupUserLink) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MeetupUserLink) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields15zyld = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields15zyld uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zyld, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft15zyld := totalEncodedFields15zyld
	missingFieldsLeft15zyld := maxFields15zyld - totalEncodedFields15zyld

	var nextMiss15zyld int32 = -1
	var found15zyld [maxFields15zyld]bool
	var curField15zyld string

doneWithStruct15zyld:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zyld > 0 || missingFieldsLeft15zyld > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zyld, missingFieldsLeft15zyld, msgp.ShowFound(found15zyld[:]), unmarshalMsgFieldOrder15zyld)
		if encodedFieldsLeft15zyld > 0 {
			encodedFieldsLeft15zyld--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField15zyld = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss15zyld < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zyld = 0
			}
			for nextMiss15zyld < maxFields15zyld && (found15zyld[nextMiss15zyld] || unmarshalMsgFieldSkip15zyld[nextMiss15zyld]) {
				nextMiss15zyld++
			}
			if nextMiss15zyld == maxFields15zyld {
				// filled all the empty fields!
				break doneWithStruct15zyld
			}
			missingFieldsLeft15zyld--
			curField15zyld = unmarshalMsgFieldOrder15zyld[nextMiss15zyld]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zyld)
		switch curField15zyld {
		// -- templateUnmarshalMsg ends here --

		case "MeetupID":
			found15zyld[0] = true
			{
				var zypi string
				zypi, bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
				z.MeetupID = MeetupID(zypi)
			}
		case "UserID":
			found15zyld[1] = true
			{
				var zmbc string
				zmbc, bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
				z.UserID = UserID(zmbc)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss15zyld != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MeetupUserLink
var unmarshalMsgFieldOrder15zyld = []string{"MeetupID", "UserID"}

var unmarshalMsgFieldSkip15zyld = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MeetupUserLink) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(string(z.MeetupID)) + 7 + msgp.StringPrefixSize + len(string(z.UserID))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *OutboundChatMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields16zajo = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zajo uint32
	totalEncodedFields16zajo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zajo := totalEncodedFields16zajo
	missingFieldsLeft16zajo := maxFields16zajo - totalEncodedFields16zajo

	var nextMiss16zajo int32 = -1
	var found16zajo [maxFields16zajo]bool
	var curField16zajo string

doneWithStruct16zajo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zajo > 0 || missingFieldsLeft16zajo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zajo, missingFieldsLeft16zajo, msgp.ShowFound(found16zajo[:]), decodeMsgFieldOrder16zajo)
		if encodedFieldsLeft16zajo > 0 {
			encodedFieldsLeft16zajo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zajo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zajo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zajo = 0
			}
			for nextMiss16zajo < maxFields16zajo && (found16zajo[nextMiss16zajo] || decodeMsgFieldSkip16zajo[nextMiss16zajo]) {
				nextMiss16zajo++
			}
			if nextMiss16zajo == maxFields16zajo {
				// filled all the empty fields!
				break doneWithStruct16zajo
			}
			missingFieldsLeft16zajo--
			curField16zajo = decodeMsgFieldOrder16zajo[nextMiss16zajo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zajo)
		switch curField16zajo {
		// -- templateDecodeMsg ends here --

		case "RecipientID":
			found16zajo[0] = true
			z.RecipientID, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Message":
			found16zajo[1] = true
			z.Message, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Map":
			found16zajo[2] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				z.Map = nil
			} else {
				if z.Map == nil {
					z.Map = new(Location)
				}
				const maxFields17zeqa = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields17zeqa uint32
				totalEncodedFields17zeqa, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft17zeqa := totalEncodedFields17zeqa
				missingFieldsLeft17zeqa := maxFields17zeqa - totalEncodedFields17zeqa

				var nextMiss17zeqa int32 = -1
				var found17zeqa [maxFields17zeqa]bool
				var curField17zeqa string

			doneWithStruct17zeqa:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft17zeqa > 0 || missingFieldsLeft17zeqa > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zeqa, missingFieldsLeft17zeqa, msgp.ShowFound(found17zeqa[:]), decodeMsgFieldOrder17zeqa)
					if encodedFieldsLeft17zeqa > 0 {
						encodedFieldsLeft17zeqa--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField17zeqa = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss17zeqa < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss17zeqa = 0
						}
						for nextMiss17zeqa < maxFields17zeqa && (found17zeqa[nextMiss17zeqa] || decodeMsgFieldSkip17zeqa[nextMiss17zeqa]) {
							nextMiss17zeqa++
						}
						if nextMiss17zeqa == maxFields17zeqa {
							// filled all the empty fields!
							break doneWithStruct17zeqa
						}
						missingFieldsLeft17zeqa--
						curField17zeqa = decodeMsgFieldOrder17zeqa[nextMiss17zeqa]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField17zeqa)
					switch curField17zeqa {
					// -- templateDecodeMsg ends here --

					case "Lat":
						found17zeqa[0] = true
						z.Map.Lat, err = dc.ReadFloat64()
						if err != nil {
							panic(err)
						}
					case "Lng":
						found17zeqa[1] = true
						z.Map.Lng, err = dc.ReadFloat64()
						if err != nil {
							panic(err)
						}
					default:
						err = dc.Skip()
						if err != nil {
							panic(err)
						}
					}
				}
				if nextMiss17zeqa != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Buttons":
			found16zajo[3] = true
			var zavc uint32
			zavc, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Buttons) >= int(zavc) {
				z.Buttons = (z.Buttons)[:zavc]
			} else {
				z.Buttons = make([]Button, zavc)
			}
			for zwzg := range z.Buttons {
				const maxFields18zcuh = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields18zcuh uint32
				totalEncodedFields18zcuh, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft18zcuh := totalEncodedFields18zcuh
				missingFieldsLeft18zcuh := maxFields18zcuh - totalEncodedFields18zcuh

				var nextMiss18zcuh int32 = -1
				var found18zcuh [maxFields18zcuh]bool
				var curField18zcuh string

			doneWithStruct18zcuh:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft18zcuh > 0 || missingFieldsLeft18zcuh > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zcuh, missingFieldsLeft18zcuh, msgp.ShowFound(found18zcuh[:]), decodeMsgFieldOrder18zcuh)
					if encodedFieldsLeft18zcuh > 0 {
						encodedFieldsLeft18zcuh--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField18zcuh = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss18zcuh < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss18zcuh = 0
						}
						for nextMiss18zcuh < maxFields18zcuh && (found18zcuh[nextMiss18zcuh] || decodeMsgFieldSkip18zcuh[nextMiss18zcuh]) {
							nextMiss18zcuh++
						}
						if nextMiss18zcuh == maxFields18zcuh {
							// filled all the empty fields!
							break doneWithStruct18zcuh
						}
						missingFieldsLeft18zcuh--
						curField18zcuh = decodeMsgFieldOrder18zcuh[nextMiss18zcuh]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField18zcuh)
					switch curField18zcuh {
					// -- templateDecodeMsg ends here --

					case "Text":
						found18zcuh[0] = true
						z.Buttons[zwzg].Text, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Command":
						found18zcuh[1] = true
						z.Buttons[zwzg].Command, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					default:
						err = dc.Skip()
						if err != nil {
							panic(err)
						}
					}
				}
				if nextMiss18zcuh != -1 {
					dc.PopAlwaysNil()
				}

			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss16zajo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of OutboundChatMessage
var decodeMsgFieldOrder16zajo = []string{"RecipientID", "Message", "Map", "Buttons"}

var decodeMsgFieldSkip16zajo = []bool{false, false, false, false}

// fields of Location
var decodeMsgFieldOrder17zeqa = []string{"Lat", "Lng"}

var decodeMsgFieldSkip17zeqa = []bool{false, false}

// fields of Button
var decodeMsgFieldOrder18zcuh = []string{"Text", "Command"}

var decodeMsgFieldSkip18zcuh = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *OutboundChatMessage) fieldsNotEmpty(isempty []bool) uint32 {
	return 4
}

// EncodeMsg implements msgp.Encodable
func (z *OutboundChatMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "RecipientID"
	err = en.Append(0x84, 0xab, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RecipientID)
	if err != nil {
		panic(err)
	}
	// write "Message"
	err = en.Append(0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Message)
	if err != nil {
		panic(err)
	}
	// write "Map"
	err = en.Append(0xa3, 0x4d, 0x61, 0x70)
	if err != nil {
		return err
	}
	if z.Map == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "Lat"
		err = en.Append(0x82, 0xa3, 0x4c, 0x61, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.Map.Lat)
		if err != nil {
			panic(err)
		}
		// write "Lng"
		err = en.Append(0xa3, 0x4c, 0x6e, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.Map.Lng)
		if err != nil {
			panic(err)
		}
	}
	// write "Buttons"
	err = en.Append(0xa7, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Buttons)))
	if err != nil {
		panic(err)
	}
	for zwzg := range z.Buttons {
		// map header, size 2
		// write "Text"
		err = en.Append(0x82, 0xa4, 0x54, 0x65, 0x78, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Buttons[zwzg].Text)
		if err != nil {
			panic(err)
		}
		// write "Command"
		err = en.Append(0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Buttons[zwzg].Command)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OutboundChatMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "RecipientID"
	o = append(o, 0x84, 0xab, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44)
	o = msgp.AppendString(o, z.RecipientID)
	// string "Message"
	o = append(o, 0xa7, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	// string "Map"
	o = append(o, 0xa3, 0x4d, 0x61, 0x70)
	if z.Map == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "Lat"
		o = append(o, 0x82, 0xa3, 0x4c, 0x61, 0x74)
		o = msgp.AppendFloat64(o, z.Map.Lat)
		// string "Lng"
		o = append(o, 0xa3, 0x4c, 0x6e, 0x67)
		o = msgp.AppendFloat64(o, z.Map.Lng)
	}
	// string "Buttons"
	o = append(o, 0xa7, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Buttons)))
	for zwzg := range z.Buttons {
		// map header, size 2
		// string "Text"
		o = append(o, 0x82, 0xa4, 0x54, 0x65, 0x78, 0x74)
		o = msgp.AppendString(o, z.Buttons[zwzg].Text)
		// string "Command"
		o = append(o, 0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64)
		o = msgp.AppendString(o, z.Buttons[zwzg].Command)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OutboundChatMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *OutboundChatMessage) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields19zvtd = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19zvtd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19zvtd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19zvtd := totalEncodedFields19zvtd
	missingFieldsLeft19zvtd := maxFields19zvtd - totalEncodedFields19zvtd

	var nextMiss19zvtd int32 = -1
	var found19zvtd [maxFields19zvtd]bool
	var curField19zvtd string

doneWithStruct19zvtd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zvtd > 0 || missingFieldsLeft19zvtd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zvtd, missingFieldsLeft19zvtd, msgp.ShowFound(found19zvtd[:]), unmarshalMsgFieldOrder19zvtd)
		if encodedFieldsLeft19zvtd > 0 {
			encodedFieldsLeft19zvtd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19zvtd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zvtd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19zvtd = 0
			}
			for nextMiss19zvtd < maxFields19zvtd && (found19zvtd[nextMiss19zvtd] || unmarshalMsgFieldSkip19zvtd[nextMiss19zvtd]) {
				nextMiss19zvtd++
			}
			if nextMiss19zvtd == maxFields19zvtd {
				// filled all the empty fields!
				break doneWithStruct19zvtd
			}
			missingFieldsLeft19zvtd--
			curField19zvtd = unmarshalMsgFieldOrder19zvtd[nextMiss19zvtd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zvtd)
		switch curField19zvtd {
		// -- templateUnmarshalMsg ends here --

		case "RecipientID":
			found19zvtd[0] = true
			z.RecipientID, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Message":
			found19zvtd[1] = true
			z.Message, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Map":
			found19zvtd[2] = true
			// default gPtr logic.
			if nbs.PeekNil(bts) && z.Map == nil {
				// consume the nil
				bts, err = nbs.ReadNilBytes(bts)
				if err != nil {
					return
				}
			} else {
				// read as-if the wire has bytes, letting nbs take care of nils.

				if z.Map == nil {
					z.Map = new(Location)
				}
				const maxFields20zaki = 2

				// -- templateUnmarshalMsg starts here--
				var totalEncodedFields20zaki uint32
				if !nbs.AlwaysNil {
					totalEncodedFields20zaki, bts, err = nbs.ReadMapHeaderBytes(bts)
					if err != nil {
						panic(err)
						return
					}
				}
				encodedFieldsLeft20zaki := totalEncodedFields20zaki
				missingFieldsLeft20zaki := maxFields20zaki - totalEncodedFields20zaki

				var nextMiss20zaki int32 = -1
				var found20zaki [maxFields20zaki]bool
				var curField20zaki string

			doneWithStruct20zaki:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft20zaki > 0 || missingFieldsLeft20zaki > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft20zaki, missingFieldsLeft20zaki, msgp.ShowFound(found20zaki[:]), unmarshalMsgFieldOrder20zaki)
					if encodedFieldsLeft20zaki > 0 {
						encodedFieldsLeft20zaki--
						field, bts, err = nbs.ReadMapKeyZC(bts)
						if err != nil {
							panic(err)
							return
						}
						curField20zaki = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss20zaki < 0 {
							// set bts to contain just mnil (0xc0)
							bts = nbs.PushAlwaysNil(bts)
							nextMiss20zaki = 0
						}
						for nextMiss20zaki < maxFields20zaki && (found20zaki[nextMiss20zaki] || unmarshalMsgFieldSkip20zaki[nextMiss20zaki]) {
							nextMiss20zaki++
						}
						if nextMiss20zaki == maxFields20zaki {
							// filled all the empty fields!
							break doneWithStruct20zaki
						}
						missingFieldsLeft20zaki--
						curField20zaki = unmarshalMsgFieldOrder20zaki[nextMiss20zaki]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField20zaki)
					switch curField20zaki {
					// -- templateUnmarshalMsg ends here --

					case "Lat":
						found20zaki[0] = true
						z.Map.Lat, bts, err = nbs.ReadFloat64Bytes(bts)

						if err != nil {
							panic(err)
						}
					case "Lng":
						found20zaki[1] = true
						z.Map.Lng, bts, err = nbs.ReadFloat64Bytes(bts)

						if err != nil {
							panic(err)
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							panic(err)
						}
					}
				}
				if nextMiss20zaki != -1 {
					bts = nbs.PopAlwaysNil()
				}

			}
		case "Buttons":
			found19zvtd[3] = true
			if nbs.AlwaysNil {
				(z.Buttons) = (z.Buttons)[:0]
			} else {

				var zmhx uint32
				zmhx, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Buttons) >= int(zmhx) {
					z.Buttons = (z.Buttons)[:zmhx]
				} else {
					z.Buttons = make([]Button, zmhx)
				}
				for zwzg := range z.Buttons {
					const maxFields21zuki = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields21zuki uint32
					if !nbs.AlwaysNil {
						totalEncodedFields21zuki, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft21zuki := totalEncodedFields21zuki
					missingFieldsLeft21zuki := maxFields21zuki - totalEncodedFields21zuki

					var nextMiss21zuki int32 = -1
					var found21zuki [maxFields21zuki]bool
					var curField21zuki string

				doneWithStruct21zuki:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft21zuki > 0 || missingFieldsLeft21zuki > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft21zuki, missingFieldsLeft21zuki, msgp.ShowFound(found21zuki[:]), unmarshalMsgFieldOrder21zuki)
						if encodedFieldsLeft21zuki > 0 {
							encodedFieldsLeft21zuki--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField21zuki = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss21zuki < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss21zuki = 0
							}
							for nextMiss21zuki < maxFields21zuki && (found21zuki[nextMiss21zuki] || unmarshalMsgFieldSkip21zuki[nextMiss21zuki]) {
								nextMiss21zuki++
							}
							if nextMiss21zuki == maxFields21zuki {
								// filled all the empty fields!
								break doneWithStruct21zuki
							}
							missingFieldsLeft21zuki--
							curField21zuki = unmarshalMsgFieldOrder21zuki[nextMiss21zuki]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField21zuki)
						switch curField21zuki {
						// -- templateUnmarshalMsg ends here --

						case "Text":
							found21zuki[0] = true
							z.Buttons[zwzg].Text, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Command":
							found21zuki[1] = true
							z.Buttons[zwzg].Command, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								panic(err)
							}
						}
					}
					if nextMiss21zuki != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19zvtd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of OutboundChatMessage
var unmarshalMsgFieldOrder19zvtd = []string{"RecipientID", "Message", "Map", "Buttons"}

var unmarshalMsgFieldSkip19zvtd = []bool{false, false, false, false}

// fields of Location
var unmarshalMsgFieldOrder20zaki = []string{"Lat", "Lng"}

var unmarshalMsgFieldSkip20zaki = []bool{false, false}

// fields of Button
var unmarshalMsgFieldOrder21zuki = []string{"Text", "Command"}

var unmarshalMsgFieldSkip21zuki = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *OutboundChatMessage) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.RecipientID) + 8 + msgp.StringPrefixSize + len(z.Message) + 4
	if z.Map == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 4 + msgp.Float64Size + 4 + msgp.Float64Size
	}
	s += 8 + msgp.ArrayHeaderSize
	for zwzg := range z.Buttons {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Buttons[zwzg].Text) + 8 + msgp.StringPrefixSize + len(z.Buttons[zwzg].Command)
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *URL) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zeeg string
		zeeg, err = dc.ReadString()
		(*z) = URL(zeeg)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z URL) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z URL) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *URL) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *URL) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zwtq string
		zwtq, bts, err = nbs.ReadStringBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = URL(zwtq)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z URL) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *UserID) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zbdg string
		zbdg, err = dc.ReadString()
		(*z) = UserID(zbdg)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z UserID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z UserID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *UserID) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zsyl string
		zsyl, bts, err = nbs.ReadStringBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = UserID(zsyl)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UserID) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Value) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var znaj string
		znaj, err = dc.ReadString()
		(*z) = Value(znaj)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z Value) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Value) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Value) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Value) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var ztqm string
		ztqm, bts, err = nbs.ReadStringBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Value(ztqm)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Value) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}
