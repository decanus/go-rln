package rln

/*
#cgo LDFLAGS: -L${SRCDIR}/../../lib/ -llibrln -ldl -lm
#include "./../../lib/librln.h"
*/
import "C"
import "unsafe"

type Buffer struct {
	Ptr *[]byte
	Len uint
}

type RLN struct {
	ptr *C.RLN_Bn256
}

//func (r *RLN) CircuitFromParams(depth int, parameters []byte) bool {
//	return C.new_circuit_from_params(depth, toBuffer(parameters), r.ptr)
//}
//
//func (r *RLN) GenerateProof(input, output []byte) bool {
//	//bool generate_proof(const struct RLN_Bn256 *ctx,
//	//	const struct Buffer *input_buffer,
//	//struct Buffer *output_buffer);
//
//	return C.generate_proof(r.ptr, toBuffer(input), toBuffer(output));
//}

func (r *RLN) GenerateKey() {
	C.key_gen(r.ptr, &Buffer{})
	//bool key_gen(const struct RLN_Bn256 *ctx, struct Buffer *keypair_buffer);

}

func toBuffer(data []byte) C.Buffer {
	dataPtr, dataLen := sliceToPtr(data)
	return C.Buffer{
		ptr: dataPtr,
		len: C.int(dataLen),
	}
}

func sliceToPtr(slice []byte) (*C.uchar, C.int) {
	if len(slice) == 0 {
		return nil, 0
	} else {
		return (*C.uchar)(unsafe.Pointer(&slice[0])), C.int(len(slice))
	}
}
