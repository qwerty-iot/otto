package otto

import (
	"encoding/binary"
	"math"
)

func builtinDataView(call FunctionCall) Value {
	return objectValue(builtinNewDataViewNative(call.runtime, call.ArgumentList))
}

func builtinNewDataView(self *object, argumentList []Value) Value {
	return objectValue(builtinNewDataViewNative(self.runtime, argumentList))
}

func builtinNewDataViewNative(runtime *runtime, argumentList []Value) *object {
	if len(argumentList) == 1 {
		firstArgument := argumentList[0]
		if firstArgument.IsObject() && firstArgument.Class() == "ArrayBuffer" {
			return runtime.newDataView(firstArgument)
		}
	}
	// TODO-JWW not sure this is valid
	return runtime.newDataView(toValue(nil))
}

func builtinDataView_getInt8(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length, true)

	strBuffer := thisObject.get("buffer").object().get("buffer").string()

	return int8Value(int8(strBuffer[index]))
}

func builtinDataView_getUint8(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length, true)

	strBuffer := thisObject.get("buffer").object().get("buffer").string()

	return uint8Value(strBuffer[index])
}

func builtinDataView_getInt16(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-2, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint16
	if littleEndian {
		val = binary.LittleEndian.Uint16(bb[index:])
	} else {
		val = binary.BigEndian.Uint16(bb[index:])

	}
	return int16Value(int16(val))
}

func builtinDataView_getUint16(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-2, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint16
	if littleEndian {
		val = binary.LittleEndian.Uint16(bb[index:])
	} else {
		val = binary.BigEndian.Uint16(bb[index:])

	}
	return uint16Value(val)
}

func builtinDataView_getInt32(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-4, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint32
	if littleEndian {
		val = binary.LittleEndian.Uint32(bb[index:])
	} else {
		val = binary.BigEndian.Uint32(bb[index:])

	}
	return int32Value(int32(val))
}

func builtinDataView_getUint32(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-4, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint32
	if littleEndian {
		val = binary.LittleEndian.Uint32(bb[index:])
	} else {
		val = binary.BigEndian.Uint32(bb[index:])

	}
	return uint32Value(val)
}

func builtinDataView_getBigInt64(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-8, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint64
	if littleEndian {
		val = binary.LittleEndian.Uint64(bb[index:])
	} else {
		val = binary.BigEndian.Uint64(bb[index:])

	}
	return int64Value(int64(val))
}

func builtinDataView_getBigUint64(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-8, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint64
	if littleEndian {
		val = binary.LittleEndian.Uint64(bb[index:])
	} else {
		val = binary.BigEndian.Uint64(bb[index:])

	}
	return uint64Value(val)
}

func builtinDataView_getFloat32(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-4, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint32

	if littleEndian {
		val = binary.LittleEndian.Uint32(bb[index:])
	} else {
		val = binary.BigEndian.Uint32(bb[index:])

	}
	return float32Value(math.Float32frombits(val))
}

func builtinDataView_getFloat64(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-8, true)

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 1).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	var val uint64
	if littleEndian {
		val = binary.LittleEndian.Uint64(bb[index:])
	} else {
		val = binary.BigEndian.Uint64(bb[index:])

	}
	return float64Value(math.Float64frombits(val))
}

func builtinDataView_setInt8(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 127 || val < -128 {
		panic(call.runtime.panicRangeError("int8 out of range"))
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	bb[index] = byte(val)
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setUint8(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 255 || val < 0 {
		panic(call.runtime.panicRangeError("uint8 out of range"))
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	bb[index] = byte(val)
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setInt16(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-2, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 32767 || val < -32768 {
		panic(call.runtime.panicRangeError("int16 out of range"))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint16(bb[index:], uint16(val))
	} else {
		binary.BigEndian.PutUint16(bb[index:], uint16(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setUint16(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-2, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 65535 || val < 0 {
		panic(call.runtime.panicRangeError("uint16 out of range"))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint16(bb[index:], uint16(val))
	} else {
		binary.BigEndian.PutUint16(bb[index:], uint16(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setInt32(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-4, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 2147483647 || val < -2147483648 {
		panic(call.runtime.panicRangeError("int32 out of range"))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint32(bb[index:], uint32(val))
	} else {
		binary.BigEndian.PutUint32(bb[index:], uint32(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setUint32(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-4, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 4294967295 || val < 0 {
		panic(call.runtime.panicRangeError("uint32 out of range"))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint32(bb[index:], uint32(val))
	} else {
		binary.BigEndian.PutUint32(bb[index:], uint32(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setBigInt64(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-8, true)
	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if val > 9223372036854775807 || val < -9223372036854775808 {
		panic(call.runtime.panicRangeError("int64 out of range"))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint64(bb[index:], uint64(val))
	} else {
		binary.BigEndian.PutUint64(bb[index:], uint64(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setBigUint64(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-8, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToInteger()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}
	if uint64(val) > 18446744073709551615 || val < 0 {
		panic(call.runtime.panicRangeError("uint64 out of range"))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint64(bb[index:], uint64(val))
	} else {
		binary.BigEndian.PutUint64(bb[index:], uint64(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setFloat32(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-4, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToFloat()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint32(bb[index:], uint32(math.Float32bits(float32(val))))
	} else {
		binary.BigEndian.PutUint32(bb[index:], uint32(math.Float32bits(float32(val))))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}

func builtinDataView_setFloat64(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	index := valueToRangeIndex(valueOfArrayIndex(call.ArgumentList, 0), length-8, true)

	val, err := valueOfArrayIndex(call.ArgumentList, 1).ToFloat()
	if err != nil {
		panic(call.runtime.panicTypeError(err.Error()))
	}

	littleEndian, err := valueOfArrayIndex(call.ArgumentList, 2).ToBoolean()
	if err != nil {
		panic(err)
	}

	strBuffer := thisObject.get("buffer").object().get("buffer").string()
	bb := []byte(strBuffer)
	if littleEndian {
		binary.LittleEndian.PutUint64(bb[index:], math.Float64bits(val))
	} else {
		binary.BigEndian.PutUint64(bb[index:], math.Float64bits(val))
	}
	thisObject.get("buffer").object().put("buffer", stringValue(string(bb)), true)
	return Value{}
}
