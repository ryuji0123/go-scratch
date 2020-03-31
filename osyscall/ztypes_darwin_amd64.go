package osyscall

type RawSockAddr struct {
	Len uint8
	Family uint8
	Data [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockAddr
	Pad [92]int8
}

type _Socklen uint32

const (
	SizeofSockaddrAny = 0x6c
)