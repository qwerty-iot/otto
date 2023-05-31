package otto

func (runtime *runtime) newArrayBufferObject(length uint32) *object {
	self := runtime.newObject()
	self.class = "ArrayBuffer"
	self.defineProperty("byteLength", uint32Value(length), 0100, false)
	self.defineProperty("buffer", stringValue(string(make([]byte, length))), 0100, false)
	return self
}

func (runtime *runtime) newArrayBuffer(length uint32) *object {
	self := runtime.newArrayBufferObject(length)
	self.prototype = runtime.global.ArrayBufferPrototype
	return self
}

func (runtime *runtime) newArrayBufferOf(bytes string) *object {
	self := runtime.newArrayBuffer(uint32(len(bytes)))
	self.defineProperty("buffer", stringValue(bytes), 0111, false)
	return self
}

/*
func arrayBufferDefineOwnProperty(self *_object, name string, descriptor _property, throw bool) bool {
	lengthProperty := self.getOwnProperty("byteLength")
	lengthValue, valid := lengthProperty.value.(Value)
	if !valid {
		panic("ArrayBuffer.byteLength != Value{}")
	}
	length := lengthValue.value.(uint32)
	if name == "length" {
		if descriptor.value == nil {
			return objectDefineOwnProperty(self, name, descriptor, throw)
		}
		newLengthValue, isValue := descriptor.value.(Value)
		if !isValue {
			panic(self.runtime.panicTypeError())
		}
		newLength := arrayUint32(self.runtime, newLengthValue)
		descriptor.value = toValue_uint32(newLength)
		if newLength > length {
			return objectDefineOwnProperty(self, name, descriptor, throw)
		}
		if !lengthProperty.writable() {
			goto Reject
		}
		newWritable := true
		if descriptor.mode&0700 == 0 {
			// If writable is off
			newWritable = false
			descriptor.mode |= 0100
		}
		if !objectDefineOwnProperty(self, name, descriptor, throw) {
			return false
		}
		for newLength < length {
			length--
			if !self.delete(strconv.FormatInt(int64(length), 10), false) {
				descriptor.value = toValue_uint32(length + 1)
				if !newWritable {
					descriptor.mode &= 0077
				}
				objectDefineOwnProperty(self, name, descriptor, false)
				goto Reject
			}
		}
		if !newWritable {
			descriptor.mode &= 0077
			objectDefineOwnProperty(self, name, descriptor, false)
		}
	} else if index := stringToArrayIndex(name); index >= 0 {
		if index >= int64(length) && !lengthProperty.writable() {
			goto Reject
		}
		if !objectDefineOwnProperty(self, strconv.FormatInt(index, 10), descriptor, false) {
			goto Reject
		}
		if index >= int64(length) {
			lengthProperty.value = toValue_uint32(uint32(index + 1))
			objectDefineOwnProperty(self, "byteLength", *lengthProperty, false)
			return true
		}
	}
	return objectDefineOwnProperty(self, name, descriptor, throw)
Reject:
	if throw {
		panic(self.runtime.panicTypeError())
	}
	return false
}
*/
