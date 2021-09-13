package rln

/*
#include "./librln.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

type RLN struct {
	ptr *C.RLN_Bn256
}

func New(depth int, parameters []byte) (*RLN, error) {
	r := &RLN{}

	buf := toBuffer(parameters)

	size := int(unsafe.Sizeof(buf))
	in := (*C.Buffer)(C.malloc(C.size_t(size)))
	*in = buf

	if bool(C.new_circuit_from_params(C.ulong(depth), in, &r.ptr)) {
		return nil, errors.New("failed to initialize")
	}

	return r, nil
}

func (r *RLN) Hash(input []byte) []byte {
	size := int(unsafe.Sizeof(C.Buffer{}))
	in := (*C.Buffer)(C.malloc(C.size_t(size)))
	*in = toBuffer(input)

	out := (*C.Buffer)(C.malloc(C.size_t(size)))
	C.hash(r.ptr, in, &in.len, out)

	return C.GoBytes(unsafe.Pointer(out.ptr), C.int(out.len))
}

func (r *RLN) CircuitFromParams(depth int, parameters []byte) bool {
	ptr := r.ptr
	buf := toBuffer(parameters)
	return bool(C.new_circuit_from_params(C.ulong(depth), &buf, &ptr))
}

func (r *RLN) GenerateProof(input, output []byte) bool {
	inputBuf := toBuffer(input)
	outputBuf := toBuffer(output)

	return bool(C.generate_proof(r.ptr, &inputBuf, &outputBuf))
}

func (r *RLN) GenerateKey(buf []byte) bool {
	buffer := toBuffer(buf)
	return bool(C.key_gen(r.ptr, &buffer))
}

func (r *RLN) Verify(proof []byte, publicInputs []byte, result uint32) bool {
	proofBuf := toBuffer(proof)
	inputs := toBuffer(publicInputs)
	res := C.uint(result)
	return bool(C.verify(r.ptr, &proofBuf, &inputs, &res))
}

func toBuffer(data []byte) C.Buffer {
	dataPtr, dataLen := sliceToPtr(data)
	return C.Buffer{
		ptr: dataPtr,
		len: C.ulong(dataLen),
	}
}

func sliceToPtr(slice []byte) (*C.uchar, C.int) {
	if len(slice) == 0 {
		return nil, 0
	} else {
		return (*C.uchar)(unsafe.Pointer(&slice[0])), C.int(len(slice))
	}
}
