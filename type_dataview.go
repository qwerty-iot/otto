package otto

// support constructor start/len

func (runtime *_runtime) newDataViewObject(arrayBuffer Value) *_object {
	self := runtime.newObject()
	self.defineProperty("byteLength", arrayBuffer._object().get("byteLength"), 0100, false)
	self.defineProperty("byteOffset", toValue_int(0), 0100, false)
	self.defineProperty("buffer", arrayBuffer, 0100, false)
	return self
}

func (runtime *_runtime) newDataView(arrayBuffer Value) *_object {
	self := runtime.newDataViewObject(arrayBuffer)
	self.prototype = runtime.global.DataViewPrototype
	return self
}
