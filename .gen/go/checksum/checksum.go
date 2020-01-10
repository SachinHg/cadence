// The MIT License (MIT)
// 
// Copyright (c) 2019 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw v1.20.2. DO NOT EDIT.
// @generated

package checksum

import (
	fmt "fmt"
	shared "github.com/uber/cadence/.gen/go/shared"
	multierr "go.uber.org/multierr"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type MutableStateChecksumPayload struct {
	CancelRequested              *bool                    `json:"cancelRequested,omitempty"`
	State                        *int16                   `json:"state,omitempty"`
	CloseStatus                  *int16                   `json:"closeStatus,omitempty"`
	LastWriteVersion             *int64                   `json:"lastWriteVersion,omitempty"`
	LastWriteEventID             *int64                   `json:"lastWriteEventID,omitempty"`
	LastFirstEventID             *int64                   `json:"lastFirstEventID,omitempty"`
	NextEventID                  *int64                   `json:"nextEventID,omitempty"`
	LastProcessedEventID         *int64                   `json:"lastProcessedEventID,omitempty"`
	SignalCount                  *int64                   `json:"signalCount,omitempty"`
	DecisionAttempt              *int32                   `json:"decisionAttempt,omitempty"`
	DecisionVersion              *int64                   `json:"decisionVersion,omitempty"`
	DecisionScheduledID          *int64                   `json:"decisionScheduledID,omitempty"`
	DecisionStartedID            *int64                   `json:"decisionStartedID,omitempty"`
	PendingTimerStartedIDs       []int64                  `json:"pendingTimerStartedIDs,omitempty"`
	PendingActivityScheduledIDs  []int64                  `json:"pendingActivityScheduledIDs,omitempty"`
	PendingSignalInitiatedIDs    []int64                  `json:"pendingSignalInitiatedIDs,omitempty"`
	PendingReqCancelInitiatedIDs []int64                  `json:"pendingReqCancelInitiatedIDs,omitempty"`
	PendingChildInitiatedIDs     []int64                  `json:"pendingChildInitiatedIDs,omitempty"`
	StickyTaskListName           *string                  `json:"stickyTaskListName,omitempty"`
	VersionHistories             *shared.VersionHistories `json:"VersionHistories,omitempty"`
}

type _List_I64_ValueList []int64

func (v _List_I64_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := wire.NewValueI64(x), error(nil)
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_I64_ValueList) Size() int {
	return len(v)
}

func (_List_I64_ValueList) ValueType() wire.Type {
	return wire.TI64
}

func (_List_I64_ValueList) Close() {}

// ToWire translates a MutableStateChecksumPayload struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *MutableStateChecksumPayload) ToWire() (wire.Value, error) {
	var (
		fields [20]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.CancelRequested != nil {
		w, err = wire.NewValueBool(*(v.CancelRequested)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.State != nil {
		w, err = wire.NewValueI16(*(v.State)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 15, Value: w}
		i++
	}
	if v.CloseStatus != nil {
		w, err = wire.NewValueI16(*(v.CloseStatus)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 16, Value: w}
		i++
	}
	if v.LastWriteVersion != nil {
		w, err = wire.NewValueI64(*(v.LastWriteVersion)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 21, Value: w}
		i++
	}
	if v.LastWriteEventID != nil {
		w, err = wire.NewValueI64(*(v.LastWriteEventID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 22, Value: w}
		i++
	}
	if v.LastFirstEventID != nil {
		w, err = wire.NewValueI64(*(v.LastFirstEventID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 23, Value: w}
		i++
	}
	if v.NextEventID != nil {
		w, err = wire.NewValueI64(*(v.NextEventID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 24, Value: w}
		i++
	}
	if v.LastProcessedEventID != nil {
		w, err = wire.NewValueI64(*(v.LastProcessedEventID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 25, Value: w}
		i++
	}
	if v.SignalCount != nil {
		w, err = wire.NewValueI64(*(v.SignalCount)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 26, Value: w}
		i++
	}
	if v.DecisionAttempt != nil {
		w, err = wire.NewValueI32(*(v.DecisionAttempt)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 35, Value: w}
		i++
	}
	if v.DecisionVersion != nil {
		w, err = wire.NewValueI64(*(v.DecisionVersion)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 36, Value: w}
		i++
	}
	if v.DecisionScheduledID != nil {
		w, err = wire.NewValueI64(*(v.DecisionScheduledID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 37, Value: w}
		i++
	}
	if v.DecisionStartedID != nil {
		w, err = wire.NewValueI64(*(v.DecisionStartedID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 38, Value: w}
		i++
	}
	if v.PendingTimerStartedIDs != nil {
		w, err = wire.NewValueList(_List_I64_ValueList(v.PendingTimerStartedIDs)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 45, Value: w}
		i++
	}
	if v.PendingActivityScheduledIDs != nil {
		w, err = wire.NewValueList(_List_I64_ValueList(v.PendingActivityScheduledIDs)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 46, Value: w}
		i++
	}
	if v.PendingSignalInitiatedIDs != nil {
		w, err = wire.NewValueList(_List_I64_ValueList(v.PendingSignalInitiatedIDs)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 47, Value: w}
		i++
	}
	if v.PendingReqCancelInitiatedIDs != nil {
		w, err = wire.NewValueList(_List_I64_ValueList(v.PendingReqCancelInitiatedIDs)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 48, Value: w}
		i++
	}
	if v.PendingChildInitiatedIDs != nil {
		w, err = wire.NewValueList(_List_I64_ValueList(v.PendingChildInitiatedIDs)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 49, Value: w}
		i++
	}
	if v.StickyTaskListName != nil {
		w, err = wire.NewValueString(*(v.StickyTaskListName)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 55, Value: w}
		i++
	}
	if v.VersionHistories != nil {
		w, err = v.VersionHistories.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 56, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_I64_Read(l wire.ValueList) ([]int64, error) {
	if l.ValueType() != wire.TI64 {
		return nil, nil
	}

	o := make([]int64, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := x.GetI64(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func _VersionHistories_Read(w wire.Value) (*shared.VersionHistories, error) {
	var v shared.VersionHistories
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a MutableStateChecksumPayload struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a MutableStateChecksumPayload struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v MutableStateChecksumPayload
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *MutableStateChecksumPayload) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.CancelRequested = &x
				if err != nil {
					return err
				}

			}
		case 15:
			if field.Value.Type() == wire.TI16 {
				var x int16
				x, err = field.Value.GetI16(), error(nil)
				v.State = &x
				if err != nil {
					return err
				}

			}
		case 16:
			if field.Value.Type() == wire.TI16 {
				var x int16
				x, err = field.Value.GetI16(), error(nil)
				v.CloseStatus = &x
				if err != nil {
					return err
				}

			}
		case 21:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.LastWriteVersion = &x
				if err != nil {
					return err
				}

			}
		case 22:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.LastWriteEventID = &x
				if err != nil {
					return err
				}

			}
		case 23:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.LastFirstEventID = &x
				if err != nil {
					return err
				}

			}
		case 24:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.NextEventID = &x
				if err != nil {
					return err
				}

			}
		case 25:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.LastProcessedEventID = &x
				if err != nil {
					return err
				}

			}
		case 26:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.SignalCount = &x
				if err != nil {
					return err
				}

			}
		case 35:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.DecisionAttempt = &x
				if err != nil {
					return err
				}

			}
		case 36:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.DecisionVersion = &x
				if err != nil {
					return err
				}

			}
		case 37:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.DecisionScheduledID = &x
				if err != nil {
					return err
				}

			}
		case 38:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.DecisionStartedID = &x
				if err != nil {
					return err
				}

			}
		case 45:
			if field.Value.Type() == wire.TList {
				v.PendingTimerStartedIDs, err = _List_I64_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 46:
			if field.Value.Type() == wire.TList {
				v.PendingActivityScheduledIDs, err = _List_I64_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 47:
			if field.Value.Type() == wire.TList {
				v.PendingSignalInitiatedIDs, err = _List_I64_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 48:
			if field.Value.Type() == wire.TList {
				v.PendingReqCancelInitiatedIDs, err = _List_I64_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 49:
			if field.Value.Type() == wire.TList {
				v.PendingChildInitiatedIDs, err = _List_I64_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 55:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.StickyTaskListName = &x
				if err != nil {
					return err
				}

			}
		case 56:
			if field.Value.Type() == wire.TStruct {
				v.VersionHistories, err = _VersionHistories_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a MutableStateChecksumPayload
// struct.
func (v *MutableStateChecksumPayload) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [20]string
	i := 0
	if v.CancelRequested != nil {
		fields[i] = fmt.Sprintf("CancelRequested: %v", *(v.CancelRequested))
		i++
	}
	if v.State != nil {
		fields[i] = fmt.Sprintf("State: %v", *(v.State))
		i++
	}
	if v.CloseStatus != nil {
		fields[i] = fmt.Sprintf("CloseStatus: %v", *(v.CloseStatus))
		i++
	}
	if v.LastWriteVersion != nil {
		fields[i] = fmt.Sprintf("LastWriteVersion: %v", *(v.LastWriteVersion))
		i++
	}
	if v.LastWriteEventID != nil {
		fields[i] = fmt.Sprintf("LastWriteEventID: %v", *(v.LastWriteEventID))
		i++
	}
	if v.LastFirstEventID != nil {
		fields[i] = fmt.Sprintf("LastFirstEventID: %v", *(v.LastFirstEventID))
		i++
	}
	if v.NextEventID != nil {
		fields[i] = fmt.Sprintf("NextEventID: %v", *(v.NextEventID))
		i++
	}
	if v.LastProcessedEventID != nil {
		fields[i] = fmt.Sprintf("LastProcessedEventID: %v", *(v.LastProcessedEventID))
		i++
	}
	if v.SignalCount != nil {
		fields[i] = fmt.Sprintf("SignalCount: %v", *(v.SignalCount))
		i++
	}
	if v.DecisionAttempt != nil {
		fields[i] = fmt.Sprintf("DecisionAttempt: %v", *(v.DecisionAttempt))
		i++
	}
	if v.DecisionVersion != nil {
		fields[i] = fmt.Sprintf("DecisionVersion: %v", *(v.DecisionVersion))
		i++
	}
	if v.DecisionScheduledID != nil {
		fields[i] = fmt.Sprintf("DecisionScheduledID: %v", *(v.DecisionScheduledID))
		i++
	}
	if v.DecisionStartedID != nil {
		fields[i] = fmt.Sprintf("DecisionStartedID: %v", *(v.DecisionStartedID))
		i++
	}
	if v.PendingTimerStartedIDs != nil {
		fields[i] = fmt.Sprintf("PendingTimerStartedIDs: %v", v.PendingTimerStartedIDs)
		i++
	}
	if v.PendingActivityScheduledIDs != nil {
		fields[i] = fmt.Sprintf("PendingActivityScheduledIDs: %v", v.PendingActivityScheduledIDs)
		i++
	}
	if v.PendingSignalInitiatedIDs != nil {
		fields[i] = fmt.Sprintf("PendingSignalInitiatedIDs: %v", v.PendingSignalInitiatedIDs)
		i++
	}
	if v.PendingReqCancelInitiatedIDs != nil {
		fields[i] = fmt.Sprintf("PendingReqCancelInitiatedIDs: %v", v.PendingReqCancelInitiatedIDs)
		i++
	}
	if v.PendingChildInitiatedIDs != nil {
		fields[i] = fmt.Sprintf("PendingChildInitiatedIDs: %v", v.PendingChildInitiatedIDs)
		i++
	}
	if v.StickyTaskListName != nil {
		fields[i] = fmt.Sprintf("StickyTaskListName: %v", *(v.StickyTaskListName))
		i++
	}
	if v.VersionHistories != nil {
		fields[i] = fmt.Sprintf("VersionHistories: %v", v.VersionHistories)
		i++
	}

	return fmt.Sprintf("MutableStateChecksumPayload{%v}", strings.Join(fields[:i], ", "))
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I16_EqualsPtr(lhs, rhs *int16) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I32_EqualsPtr(lhs, rhs *int32) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _List_I64_Equals(lhs, rhs []int64) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, lv := range lhs {
		rv := rhs[i]
		if !(lv == rv) {
			return false
		}
	}

	return true
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this MutableStateChecksumPayload match the
// provided MutableStateChecksumPayload.
//
// This function performs a deep comparison.
func (v *MutableStateChecksumPayload) Equals(rhs *MutableStateChecksumPayload) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_Bool_EqualsPtr(v.CancelRequested, rhs.CancelRequested) {
		return false
	}
	if !_I16_EqualsPtr(v.State, rhs.State) {
		return false
	}
	if !_I16_EqualsPtr(v.CloseStatus, rhs.CloseStatus) {
		return false
	}
	if !_I64_EqualsPtr(v.LastWriteVersion, rhs.LastWriteVersion) {
		return false
	}
	if !_I64_EqualsPtr(v.LastWriteEventID, rhs.LastWriteEventID) {
		return false
	}
	if !_I64_EqualsPtr(v.LastFirstEventID, rhs.LastFirstEventID) {
		return false
	}
	if !_I64_EqualsPtr(v.NextEventID, rhs.NextEventID) {
		return false
	}
	if !_I64_EqualsPtr(v.LastProcessedEventID, rhs.LastProcessedEventID) {
		return false
	}
	if !_I64_EqualsPtr(v.SignalCount, rhs.SignalCount) {
		return false
	}
	if !_I32_EqualsPtr(v.DecisionAttempt, rhs.DecisionAttempt) {
		return false
	}
	if !_I64_EqualsPtr(v.DecisionVersion, rhs.DecisionVersion) {
		return false
	}
	if !_I64_EqualsPtr(v.DecisionScheduledID, rhs.DecisionScheduledID) {
		return false
	}
	if !_I64_EqualsPtr(v.DecisionStartedID, rhs.DecisionStartedID) {
		return false
	}
	if !((v.PendingTimerStartedIDs == nil && rhs.PendingTimerStartedIDs == nil) || (v.PendingTimerStartedIDs != nil && rhs.PendingTimerStartedIDs != nil && _List_I64_Equals(v.PendingTimerStartedIDs, rhs.PendingTimerStartedIDs))) {
		return false
	}
	if !((v.PendingActivityScheduledIDs == nil && rhs.PendingActivityScheduledIDs == nil) || (v.PendingActivityScheduledIDs != nil && rhs.PendingActivityScheduledIDs != nil && _List_I64_Equals(v.PendingActivityScheduledIDs, rhs.PendingActivityScheduledIDs))) {
		return false
	}
	if !((v.PendingSignalInitiatedIDs == nil && rhs.PendingSignalInitiatedIDs == nil) || (v.PendingSignalInitiatedIDs != nil && rhs.PendingSignalInitiatedIDs != nil && _List_I64_Equals(v.PendingSignalInitiatedIDs, rhs.PendingSignalInitiatedIDs))) {
		return false
	}
	if !((v.PendingReqCancelInitiatedIDs == nil && rhs.PendingReqCancelInitiatedIDs == nil) || (v.PendingReqCancelInitiatedIDs != nil && rhs.PendingReqCancelInitiatedIDs != nil && _List_I64_Equals(v.PendingReqCancelInitiatedIDs, rhs.PendingReqCancelInitiatedIDs))) {
		return false
	}
	if !((v.PendingChildInitiatedIDs == nil && rhs.PendingChildInitiatedIDs == nil) || (v.PendingChildInitiatedIDs != nil && rhs.PendingChildInitiatedIDs != nil && _List_I64_Equals(v.PendingChildInitiatedIDs, rhs.PendingChildInitiatedIDs))) {
		return false
	}
	if !_String_EqualsPtr(v.StickyTaskListName, rhs.StickyTaskListName) {
		return false
	}
	if !((v.VersionHistories == nil && rhs.VersionHistories == nil) || (v.VersionHistories != nil && rhs.VersionHistories != nil && v.VersionHistories.Equals(rhs.VersionHistories))) {
		return false
	}

	return true
}

type _List_I64_Zapper []int64

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _List_I64_Zapper.
func (l _List_I64_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range l {
		enc.AppendInt64(v)
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of MutableStateChecksumPayload.
func (v *MutableStateChecksumPayload) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.CancelRequested != nil {
		enc.AddBool("cancelRequested", *v.CancelRequested)
	}
	if v.State != nil {
		enc.AddInt16("state", *v.State)
	}
	if v.CloseStatus != nil {
		enc.AddInt16("closeStatus", *v.CloseStatus)
	}
	if v.LastWriteVersion != nil {
		enc.AddInt64("lastWriteVersion", *v.LastWriteVersion)
	}
	if v.LastWriteEventID != nil {
		enc.AddInt64("lastWriteEventID", *v.LastWriteEventID)
	}
	if v.LastFirstEventID != nil {
		enc.AddInt64("lastFirstEventID", *v.LastFirstEventID)
	}
	if v.NextEventID != nil {
		enc.AddInt64("nextEventID", *v.NextEventID)
	}
	if v.LastProcessedEventID != nil {
		enc.AddInt64("lastProcessedEventID", *v.LastProcessedEventID)
	}
	if v.SignalCount != nil {
		enc.AddInt64("signalCount", *v.SignalCount)
	}
	if v.DecisionAttempt != nil {
		enc.AddInt32("decisionAttempt", *v.DecisionAttempt)
	}
	if v.DecisionVersion != nil {
		enc.AddInt64("decisionVersion", *v.DecisionVersion)
	}
	if v.DecisionScheduledID != nil {
		enc.AddInt64("decisionScheduledID", *v.DecisionScheduledID)
	}
	if v.DecisionStartedID != nil {
		enc.AddInt64("decisionStartedID", *v.DecisionStartedID)
	}
	if v.PendingTimerStartedIDs != nil {
		err = multierr.Append(err, enc.AddArray("pendingTimerStartedIDs", (_List_I64_Zapper)(v.PendingTimerStartedIDs)))
	}
	if v.PendingActivityScheduledIDs != nil {
		err = multierr.Append(err, enc.AddArray("pendingActivityScheduledIDs", (_List_I64_Zapper)(v.PendingActivityScheduledIDs)))
	}
	if v.PendingSignalInitiatedIDs != nil {
		err = multierr.Append(err, enc.AddArray("pendingSignalInitiatedIDs", (_List_I64_Zapper)(v.PendingSignalInitiatedIDs)))
	}
	if v.PendingReqCancelInitiatedIDs != nil {
		err = multierr.Append(err, enc.AddArray("pendingReqCancelInitiatedIDs", (_List_I64_Zapper)(v.PendingReqCancelInitiatedIDs)))
	}
	if v.PendingChildInitiatedIDs != nil {
		err = multierr.Append(err, enc.AddArray("pendingChildInitiatedIDs", (_List_I64_Zapper)(v.PendingChildInitiatedIDs)))
	}
	if v.StickyTaskListName != nil {
		enc.AddString("stickyTaskListName", *v.StickyTaskListName)
	}
	if v.VersionHistories != nil {
		err = multierr.Append(err, enc.AddObject("VersionHistories", v.VersionHistories))
	}
	return err
}

// GetCancelRequested returns the value of CancelRequested if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetCancelRequested() (o bool) {
	if v != nil && v.CancelRequested != nil {
		return *v.CancelRequested
	}

	return
}

// IsSetCancelRequested returns true if CancelRequested is not nil.
func (v *MutableStateChecksumPayload) IsSetCancelRequested() bool {
	return v != nil && v.CancelRequested != nil
}

// GetState returns the value of State if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetState() (o int16) {
	if v != nil && v.State != nil {
		return *v.State
	}

	return
}

// IsSetState returns true if State is not nil.
func (v *MutableStateChecksumPayload) IsSetState() bool {
	return v != nil && v.State != nil
}

// GetCloseStatus returns the value of CloseStatus if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetCloseStatus() (o int16) {
	if v != nil && v.CloseStatus != nil {
		return *v.CloseStatus
	}

	return
}

// IsSetCloseStatus returns true if CloseStatus is not nil.
func (v *MutableStateChecksumPayload) IsSetCloseStatus() bool {
	return v != nil && v.CloseStatus != nil
}

// GetLastWriteVersion returns the value of LastWriteVersion if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetLastWriteVersion() (o int64) {
	if v != nil && v.LastWriteVersion != nil {
		return *v.LastWriteVersion
	}

	return
}

// IsSetLastWriteVersion returns true if LastWriteVersion is not nil.
func (v *MutableStateChecksumPayload) IsSetLastWriteVersion() bool {
	return v != nil && v.LastWriteVersion != nil
}

// GetLastWriteEventID returns the value of LastWriteEventID if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetLastWriteEventID() (o int64) {
	if v != nil && v.LastWriteEventID != nil {
		return *v.LastWriteEventID
	}

	return
}

// IsSetLastWriteEventID returns true if LastWriteEventID is not nil.
func (v *MutableStateChecksumPayload) IsSetLastWriteEventID() bool {
	return v != nil && v.LastWriteEventID != nil
}

// GetLastFirstEventID returns the value of LastFirstEventID if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetLastFirstEventID() (o int64) {
	if v != nil && v.LastFirstEventID != nil {
		return *v.LastFirstEventID
	}

	return
}

// IsSetLastFirstEventID returns true if LastFirstEventID is not nil.
func (v *MutableStateChecksumPayload) IsSetLastFirstEventID() bool {
	return v != nil && v.LastFirstEventID != nil
}

// GetNextEventID returns the value of NextEventID if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetNextEventID() (o int64) {
	if v != nil && v.NextEventID != nil {
		return *v.NextEventID
	}

	return
}

// IsSetNextEventID returns true if NextEventID is not nil.
func (v *MutableStateChecksumPayload) IsSetNextEventID() bool {
	return v != nil && v.NextEventID != nil
}

// GetLastProcessedEventID returns the value of LastProcessedEventID if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetLastProcessedEventID() (o int64) {
	if v != nil && v.LastProcessedEventID != nil {
		return *v.LastProcessedEventID
	}

	return
}

// IsSetLastProcessedEventID returns true if LastProcessedEventID is not nil.
func (v *MutableStateChecksumPayload) IsSetLastProcessedEventID() bool {
	return v != nil && v.LastProcessedEventID != nil
}

// GetSignalCount returns the value of SignalCount if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetSignalCount() (o int64) {
	if v != nil && v.SignalCount != nil {
		return *v.SignalCount
	}

	return
}

// IsSetSignalCount returns true if SignalCount is not nil.
func (v *MutableStateChecksumPayload) IsSetSignalCount() bool {
	return v != nil && v.SignalCount != nil
}

// GetDecisionAttempt returns the value of DecisionAttempt if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetDecisionAttempt() (o int32) {
	if v != nil && v.DecisionAttempt != nil {
		return *v.DecisionAttempt
	}

	return
}

// IsSetDecisionAttempt returns true if DecisionAttempt is not nil.
func (v *MutableStateChecksumPayload) IsSetDecisionAttempt() bool {
	return v != nil && v.DecisionAttempt != nil
}

// GetDecisionVersion returns the value of DecisionVersion if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetDecisionVersion() (o int64) {
	if v != nil && v.DecisionVersion != nil {
		return *v.DecisionVersion
	}

	return
}

// IsSetDecisionVersion returns true if DecisionVersion is not nil.
func (v *MutableStateChecksumPayload) IsSetDecisionVersion() bool {
	return v != nil && v.DecisionVersion != nil
}

// GetDecisionScheduledID returns the value of DecisionScheduledID if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetDecisionScheduledID() (o int64) {
	if v != nil && v.DecisionScheduledID != nil {
		return *v.DecisionScheduledID
	}

	return
}

// IsSetDecisionScheduledID returns true if DecisionScheduledID is not nil.
func (v *MutableStateChecksumPayload) IsSetDecisionScheduledID() bool {
	return v != nil && v.DecisionScheduledID != nil
}

// GetDecisionStartedID returns the value of DecisionStartedID if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetDecisionStartedID() (o int64) {
	if v != nil && v.DecisionStartedID != nil {
		return *v.DecisionStartedID
	}

	return
}

// IsSetDecisionStartedID returns true if DecisionStartedID is not nil.
func (v *MutableStateChecksumPayload) IsSetDecisionStartedID() bool {
	return v != nil && v.DecisionStartedID != nil
}

// GetPendingTimerStartedIDs returns the value of PendingTimerStartedIDs if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetPendingTimerStartedIDs() (o []int64) {
	if v != nil && v.PendingTimerStartedIDs != nil {
		return v.PendingTimerStartedIDs
	}

	return
}

// IsSetPendingTimerStartedIDs returns true if PendingTimerStartedIDs is not nil.
func (v *MutableStateChecksumPayload) IsSetPendingTimerStartedIDs() bool {
	return v != nil && v.PendingTimerStartedIDs != nil
}

// GetPendingActivityScheduledIDs returns the value of PendingActivityScheduledIDs if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetPendingActivityScheduledIDs() (o []int64) {
	if v != nil && v.PendingActivityScheduledIDs != nil {
		return v.PendingActivityScheduledIDs
	}

	return
}

// IsSetPendingActivityScheduledIDs returns true if PendingActivityScheduledIDs is not nil.
func (v *MutableStateChecksumPayload) IsSetPendingActivityScheduledIDs() bool {
	return v != nil && v.PendingActivityScheduledIDs != nil
}

// GetPendingSignalInitiatedIDs returns the value of PendingSignalInitiatedIDs if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetPendingSignalInitiatedIDs() (o []int64) {
	if v != nil && v.PendingSignalInitiatedIDs != nil {
		return v.PendingSignalInitiatedIDs
	}

	return
}

// IsSetPendingSignalInitiatedIDs returns true if PendingSignalInitiatedIDs is not nil.
func (v *MutableStateChecksumPayload) IsSetPendingSignalInitiatedIDs() bool {
	return v != nil && v.PendingSignalInitiatedIDs != nil
}

// GetPendingReqCancelInitiatedIDs returns the value of PendingReqCancelInitiatedIDs if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetPendingReqCancelInitiatedIDs() (o []int64) {
	if v != nil && v.PendingReqCancelInitiatedIDs != nil {
		return v.PendingReqCancelInitiatedIDs
	}

	return
}

// IsSetPendingReqCancelInitiatedIDs returns true if PendingReqCancelInitiatedIDs is not nil.
func (v *MutableStateChecksumPayload) IsSetPendingReqCancelInitiatedIDs() bool {
	return v != nil && v.PendingReqCancelInitiatedIDs != nil
}

// GetPendingChildInitiatedIDs returns the value of PendingChildInitiatedIDs if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetPendingChildInitiatedIDs() (o []int64) {
	if v != nil && v.PendingChildInitiatedIDs != nil {
		return v.PendingChildInitiatedIDs
	}

	return
}

// IsSetPendingChildInitiatedIDs returns true if PendingChildInitiatedIDs is not nil.
func (v *MutableStateChecksumPayload) IsSetPendingChildInitiatedIDs() bool {
	return v != nil && v.PendingChildInitiatedIDs != nil
}

// GetStickyTaskListName returns the value of StickyTaskListName if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetStickyTaskListName() (o string) {
	if v != nil && v.StickyTaskListName != nil {
		return *v.StickyTaskListName
	}

	return
}

// IsSetStickyTaskListName returns true if StickyTaskListName is not nil.
func (v *MutableStateChecksumPayload) IsSetStickyTaskListName() bool {
	return v != nil && v.StickyTaskListName != nil
}

// GetVersionHistories returns the value of VersionHistories if it is set or its
// zero value if it is unset.
func (v *MutableStateChecksumPayload) GetVersionHistories() (o *shared.VersionHistories) {
	if v != nil && v.VersionHistories != nil {
		return v.VersionHistories
	}

	return
}

// IsSetVersionHistories returns true if VersionHistories is not nil.
func (v *MutableStateChecksumPayload) IsSetVersionHistories() bool {
	return v != nil && v.VersionHistories != nil
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "checksum",
	Package:  "github.com/uber/cadence/.gen/go/checksum",
	FilePath: "checksum.thrift",
	SHA1:     "c3ee77b53c2e06c35a3296cfeeeadf140711ed95",
	Includes: []*thriftreflect.ThriftModule{
		shared.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "// Copyright (c) 2019 Uber Technologies, Inc.\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\ninclude \"shared.thrift\"\n\nnamespace java com.uber.cadence\n\nstruct MutableStateChecksumPayload {\n    10: optional bool cancelRequested\n    15: optional i16 state\n    16: optional i16 closeStatus\n\n    21: optional i64 (js.type = \"Long\") lastWriteVersion\n    22: optional i64 (js.type = \"Long\") lastWriteEventID\n    23: optional i64 (js.type = \"Long\") lastFirstEventID\n    24: optional i64 (js.type = \"Long\") nextEventID\n    25: optional i64 (js.type = \"Long\") lastProcessedEventID\n    26: optional i64 (js.type = \"Long\") signalCount\n\n    35: optional i32 decisionAttempt\n    36: optional i64 (js.type = \"Long\") decisionVersion\n    37: optional i64 (js.type = \"Long\") decisionScheduledID\n    38: optional i64 (js.type = \"Long\") decisionStartedID\n\n    45: optional list<i64> pendingTimerStartedIDs\n    46: optional list<i64> pendingActivityScheduledIDs\n    47: optional list<i64> pendingSignalInitiatedIDs\n    48: optional list<i64> pendingReqCancelInitiatedIDs\n    49: optional list<i64> pendingChildInitiatedIDs\n\n    55: optional string stickyTaskListName\n    56: optional shared.VersionHistories VersionHistories\n}\n"
