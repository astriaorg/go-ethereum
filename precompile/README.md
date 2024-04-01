# Writing a Precompile Contract

1. Create Solidity interface in `contracts/interfaces`, e.g, IExampleContract.sol

2. Generate bindings with `./gen.sh`

3. Implement the precompile in Go. The struct should implement the `StatefulPrecompiledContract` interface and methods defined in the Solidity interface.

See NativeMinter as an example implementation
