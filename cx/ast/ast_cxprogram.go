package ast

import (
	"errors"
	"fmt"

	"github.com/skycoin/cx/cx/constants"
	"github.com/skycoin/cx/cx/types"
)

// CXProgram is used to represent a full CX program.
//
// It is the root data structure for the declarations of all functions,
// variables and data structures.
//
type CXProgram struct {
	// Metadata
	//Remove Path //moved to cx/globals
	//Path string // Path to the CX project in the filesystem

	Stack StackSegmentStruct
	Data  DataSegmentStruct
	Heap  HeapSegmentStruct

	// Contents
	Packages map[string]*CXPackage // Packages in a CX program

	// Runtime information
	ProgramInput []*CXArgument // OS input arguments
	//ProgramOutput []*CXArgument // outputs to the OS
	Memory []byte // Used when running the program

	CallStack   []CXCall      // Collection of function calls
	CallCounter types.Pointer // What function call is the currently being executed in the CallStack
	Terminated  bool          // Utility field for the runtime. Indicates if a CX program has already finished or not.
	Version     string        // CX version used to build this CX program.

	// Used by the REPL and cxgo
	CurrentPackage *CXPackage // Represents the currently active package in the REPL or when parsing a CX file.
	ProgramError   error
}

type StackSegmentStruct struct {
	//TODO: Add StackStartsAt
	Size    types.Pointer // This field stores the size of a CX program's stack
	Pointer types.Pointer // At what byte the current stack frame is
}

type DataSegmentStruct struct {
	Size     types.Pointer // This field stores the size of a CX program's data segment size
	StartsAt types.Pointer // Offset at which the data segment starts in a CX program's memory
}

type HeapSegmentStruct struct {
	Size     types.Pointer // This field stores the size of a CX program's heap
	StartsAt types.Pointer // Offset at which the heap starts in a CX program's memory (normally the stack size)
	Pointer  types.Pointer // At what offset a CX program can insert a new object to the heap
}

// MakeProgram ...
func MakeProgram() *CXProgram {
	minHeapSize := minHeapSize()
	newPrgrm := &CXProgram{
		Packages:  make(map[string]*CXPackage, 0),
		CallStack: make([]CXCall, constants.CALLSTACK_SIZE),
		Memory:    make([]byte, constants.STACK_SIZE+minHeapSize),
		Stack: StackSegmentStruct{
			Size: constants.STACK_SIZE,
		},
		Data: DataSegmentStruct{
			StartsAt: constants.STACK_SIZE,
		},
		Heap: HeapSegmentStruct{
			Size:    minHeapSize,
			Pointer: constants.NULL_HEAP_ADDRESS_OFFSET, // We can start adding objects to the heap after the NULL (nil) bytes.
		},
	}
	return newPrgrm
}

// ----------------------------------------------------------------
//                         `CXProgram` Package handling

// AddPackage ...
func (cxprogram *CXProgram) AddPackage(mod *CXPackage) {
	if cxprogram.Packages[mod.Name] != nil {
		return
	}

	cxprogram.Packages[mod.Name] = mod
	cxprogram.CurrentPackage = mod
}

// RemovePackage ...
func (cxprogram *CXProgram) RemovePackage(modName string) {
	// If doesnt exist, return
	if cxprogram.Packages[modName] == nil {
		return
	}

	// Check if it is the current pkg so when it
	// is deleted, it will be replaced with new pkg
	isCurrentPkg := cxprogram.Packages[modName] == cxprogram.CurrentPackage

	// Delete package
	delete(cxprogram.Packages, modName)

	// This means that we're removing the package set to be the CurrentPackage.
	// If it is removed from the program's map of packages, cxprogram.CurrentPackage
	// would be pointing to a package meant to be collected by the GC.
	// We fix this by pointing to random package in the program's map of packages.
	if isCurrentPkg {
		for _, pkg := range cxprogram.Packages {
			cxprogram.CurrentPackage = pkg
			break
		}
	}

}

// ----------------------------------------------------------------
//                             `CXProgram` Selectors

// SetCurrentCxProgram sets `PROGRAM` to the the receiver `prgrm`. This is a utility function used mainly
// by CX chains. `PROGRAM` is used in multiple parts of the CX runtime as a convenience; instead of having
// to pass around a parameter of type CXProgram, the CX program currently being run is accessible through
// `PROGRAM`.

//Very strange
//Beware whenever this function is called
func (cxprogram *CXProgram) SetCurrentCxProgram() (*CXProgram, error) {
	PROGRAM = cxprogram
	return PROGRAM, nil
}

// GetCurrentCxProgram returns the CX program assigned to global variable `PROGRAM`.
// This function is mainly used for CX chains.
/*
func GetCurrentCxProgram() (*CXProgram, error) {
	if PROGRAM == nil {
		return nil, fmt.Errorf("a CX program has not been loaded")
	}
	return PROGRAM, nil
}
*/

// ----------------------------------------------------------------
//                             `CXProgram` Getters

// Only two users, both in cx/execute.go
func (cxprogram *CXProgram) SelectPackage(name string) (*CXPackage, error) {
	if cxprogram.Packages[name] == nil {
		return nil, fmt.Errorf("Package '%s' does not exist", name)
	}

	cxprogram.CurrentPackage = cxprogram.Packages[name]
	return cxprogram.Packages[name], nil
}

// GetCurrentPackage ...
func (cxprogram *CXProgram) GetCurrentPackage() (*CXPackage, error) {
	if cxprogram.CurrentPackage == nil {
		return nil, errors.New("current package is nil")
	}
	return cxprogram.CurrentPackage, nil
}

// GetCurrentStruct ...
func (cxprogram *CXProgram) GetCurrentStruct() (*CXStruct, error) {
	if cxprogram.CurrentPackage == nil {
		return nil, errors.New("current package is nil")
	}

	if cxprogram.CurrentPackage.CurrentStruct == nil {
		return nil, errors.New("current struct is nil")

	}

	return cxprogram.CurrentPackage.CurrentStruct, nil
}

// GetCurrentFunction ...
func (cxprogram *CXProgram) GetCurrentFunction() (*CXFunction, error) {
	if cxprogram.CurrentPackage == nil {
		return nil, errors.New("current package is nil")
	}

	if cxprogram.CurrentPackage.CurrentFunction != nil {
		return nil, errors.New("current function is nil")
	}

	return cxprogram.CurrentPackage.CurrentFunction, nil

}

// GetCurrentExpression ...
func (cxprogram *CXProgram) GetCurrentExpression() (*CXExpression, error) {
	if cxprogram.CurrentPackage == nil {
		return nil, errors.New("current package is nil")
	}

	if cxprogram.CurrentPackage.CurrentFunction == nil {
		return nil, errors.New("current function is nil")
	}

	if cxprogram.CurrentPackage.CurrentFunction.CurrentExpression == nil {
		return nil, errors.New("current expression is nil")
	}

	return cxprogram.CurrentPackage.CurrentFunction.CurrentExpression, nil
}

// GetPackage ...
func (cxprogram *CXProgram) GetPackage(pkgName string) (*CXPackage, error) {
	if cxprogram.Packages[pkgName] == nil {
		return nil, fmt.Errorf("package '%s' not found", pkgName)
	}

	return cxprogram.Packages[pkgName], nil
}

// GetStruct ...
func (cxprogram *CXProgram) GetStruct(strctName string, pkgName string) (*CXStruct, error) {
	pkg, err := cxprogram.GetPackage(pkgName)
	if err != nil {
		return nil, err
	}

	strct, err := pkg.GetStruct(strctName)
	if err != nil {
		return nil, err
	}

	return strct, nil
}

// GetFunction ...
func (cxprogram *CXProgram) GetFunction(fnName string, pkgName string) (*CXFunction, error) {
	pkg, err := cxprogram.GetPackage(pkgName)
	if err != nil {
		return nil, err
	}

	fn, err := pkg.GetFunction(fnName)
	if err != nil {
		return nil, err
	}

	return fn, nil
}

// GetCurrentCall returns the current CXCall
//TODO: What does this do?
//TODO: Only used in OP_JMP
func (cxprogram *CXProgram) GetCurrentCall() *CXCall {
	return &cxprogram.CallStack[cxprogram.CallCounter]
}

// GetGlobal ...
/*
func (cxprogram *CXProgram) GetGlobal(name string) (*CXArgument, error) {
	mod, err := cxprogram.GetCurrentPackage()
	if err != nil {
		return nil, err
	}

	var foundArgument *CXArgument
	for _, def := range mod.Globals {
		if def.Name == name {
			foundArgument = def
			break
		}
	}

	for _, imp := range mod.Imports {
		for _, def := range imp.Globals {
			if def.Name == name {
				foundArgument = def
				break
			}
		}
	}

	if foundArgument == nil {
		return nil, fmt.Errorf("global '%s' not found", name)
	}
	return foundArgument, nil
}
*/

/*
// GetCurrentOpCode returns the current OpCode
func (cxprogram *CXProgram) GetCurrentOpCode() int {
	return cxprogram.GetCurrentExpression2().Operator.OpCode
}
*/

/*
//not used
func (cxprogram *CXProgram) GetFramePointer() int {
	return cxprogram.GetCurrentCall().FramePointer
}
*/

// ----------------------------------------------------------------
//                             `CXProgram` Debugging

// PrintAllObjects prints all objects in a program
//
func (cxprogram *CXProgram) PrintAllObjects() {
	fp := types.Pointer(0)

	for c := types.Pointer(0); c <= cxprogram.CallCounter; c++ {
		op := cxprogram.CallStack[c].Operator

		for _, ptr := range op.ListOfPointers {
			heapOffset := types.Read_ptr(cxprogram.Memory, fp+ptr.Offset)

			var byts []byte

			if ptr.CustomType != nil {
				// then it's a pointer to a struct
				// use CustomStruct to match the fields against the bytes
				// for _, fld := range ptr.Fields {

				// }

				byts = types.Get_obj_data(cxprogram.Memory, heapOffset, ptr.CustomType.Size)
			}

			// var currLengths []int
			// var currCustom *CXStruct

			// for c := len(ptr.DeclarationSpecifiers) - 1; c >= 0; c-- {
			// 	// we need to go backwards in here

			// 	switch ptr.DeclarationSpecifiers[c] {
			// 	case DECL_POINTER:
			// 		// we might not need to do anything
			// 	case DECL_ARRAY:
			// 		currLengths = ptr.Lengths
			// 	case DECL_SLICE:
			// 	case DECL_STRUCT:
			// 		currCustom = ptr.CustomType
			// 	case DECL_BASIC:
			// 	}
			// }

			// if len(ptr.Lengths) > 0 {
			// 	fmt.Println("ARRAY")
			// }

			// if ptr.CustomType != nil {
			// 	fmt.Println("STRUCT")
			// }

			fmt.Println("declarat", ptr.DeclarationSpecifiers)

			fmt.Println("obj", ptr.ArgDetails.Name, ptr.CustomType, cxprogram.Memory[heapOffset:heapOffset+op.Size], byts)
		}

		fp += op.Size
	}
}

// PrintProgram prints the abstract syntax tree of a CX program in a
// human-readable format.
func (cxprogram *CXProgram) PrintProgram() {
	fmt.Println(ToString(cxprogram))
}