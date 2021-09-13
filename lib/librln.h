#include <cstdarg>
#include <cstdint>
#include <cstdlib>
#include <ostream>
#include <new>

template<typename E = void>
struct RLN;

/// Buffer struct is taken from
/// https://github.com/celo-org/celo-threshold-bls-rs/blob/master/crates/threshold-bls-ffi/src/ffi.rs
struct Buffer {
  const uint8_t *ptr;
  uintptr_t len;
};

extern "C" {

bool new_circuit_from_params(uintptr_t merkle_depth,
                             const Buffer *parameters_buffer,
                             RLN<Bn256> **ctx);

bool generate_proof(const RLN<Bn256> *ctx, const Buffer *input_buffer, Buffer *output_buffer);

bool verify(const RLN<Bn256> *ctx,
            const Buffer *proof_buffer,
            const Buffer *public_inputs_buffer,
            uint32_t *result_ptr);

bool hash(const RLN<Bn256> *ctx,
          const Buffer *inputs_buffer,
          const uintptr_t *input_len,
          Buffer *output_buffer);

bool key_gen(const RLN<Bn256> *ctx, Buffer *keypair_buffer);

} // extern "C"
