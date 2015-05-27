package lawparser

type stateFn func(*interface{}) stateFn
