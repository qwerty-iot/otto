package otto

// support constructor start/len

func (runtime *runtime) newDataViewObject(arrayBuffer Value) *object {
	self := runtime.newObject()
	self.defineProperty("byteLength", arrayBuffer.object().get("byteLength"), 0100, false)
	self.defineProperty("byteOffset", intValue(0), 0100, false)
	self.defineProperty("buffer", arrayBuffer, 0100, false)
	return self
}

func (runtime *runtime) newDataView(arrayBuffer Value) *object {
	self := runtime.newDataViewObject(arrayBuffer)
	self.prototype = runtime.global.DataViewPrototype
	return self
}
