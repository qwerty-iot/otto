package otto

import "github.com/davecgh/go-spew/spew"

func builtinArrayBuffer(call FunctionCall) Value {
	return objectValue(builtinNewArrayBufferNative(call.runtime, call.ArgumentList))
}

func builtinNewArrayBuffer(self *object, argumentList []Value) Value {
	return objectValue(builtinNewArrayBufferNative(self.runtime, argumentList))
}

func builtinNewArrayBufferNative(runtime *runtime, argumentList []Value) *object {
	if len(argumentList) == 1 {
		firstArgument := argumentList[0]
		if firstArgument.IsNumber() {
			return runtime.newArrayBuffer(arrayUint32(runtime, firstArgument))
		}
	}
	// TODO-JWW not sure this is valid
	return runtime.newArrayBuffer(0)
}

func builtinArrayBuffer_slice(call FunctionCall) Value {
	thisObject := call.thisObject()

	length := int64(toUint32(thisObject.get("byteLength")))
	start, end := rangeStartEnd(call.ArgumentList, length, false)

	if start >= end {
		// Always an empty array
		return objectValue(call.runtime.newArrayBuffer(0))
	}
	buffer := thisObject.get("buffer").string()

	spew.Dump([]byte(buffer), start, end)

	return objectValue(call.runtime.newArrayBufferOf(buffer[start:end]))
}
