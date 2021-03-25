// CAUTION: Generated by yy - DO NOT EDIT.

// Copyright 2015 The parser Authors. All rights reserved.  Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// This is a derived work base on the original at
//
// http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html
//
// The original work is
//
// Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.
//
// Grammar for the input to yacc.

package parser // import "github.com/skycoin/cx/cmd/goyacc/parser/yacc"

import (
	"fmt"
)

func ExampleAction() {
	fmt.Println(exampleAST(2, "%% a: { }"))
	// Output:
	// &parser.Action{
	// · Values: []*parser.ActionValue{ // len 1
	// · · 0: &parser.ActionValue{
	// · · · Pos: 7,
	// · · · Src: "{ }",
	// · · },
	// · },
	// · Token: example.y:1:7: '{' "{", Comments: [],
	// · Token2: example.y:1:9: '}' "}", Comments: [],
	// }
}

func ExampleDefinition() {
	fmt.Println(exampleAST(3, "%start a %error-verbose"))
	// Output:
	// &parser.Definition{
	// · Token: example.y:1:1: START "%start", Comments: [],
	// · Token2: example.y:1:8: IDENTIFIER "a", Comments: [],
	// }
}

func ExampleDefinition_case1() {
	fmt.Println(exampleAST(4, "%union{int i} %%"))
	// Output:
	// &parser.Definition{
	// · Value: "{int i}",
	// · Case: 1,
	// · Token: example.y:1:1: UNION "%union", Comments: [],
	// }
}

func ExampleDefinition_case2() {
	fmt.Println(exampleAST(6, "%{ %} %error-verbose"))
	// Output:
	// &parser.Definition{
	// · Value: " ",
	// · Case: 2,
	// · Token: example.y:1:1: LCURL "%{", Comments: [],
	// · Token2: example.y:1:4: RCURL "%}", Comments: [],
	// }
}

func ExampleDefinition_case3() {
	fmt.Println(exampleAST(7, "%left a %error-verbose"))
	// Output:
	// &parser.Definition{
	// · Nlist: []*parser.Name{ // len 1
	// · · 0: &parser.Name{
	// · · · Identifier: "a",
	// · · · Number: -1,
	// · · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · · },
	// · },
	// · Case: 3,
	// · NameList: &parser.NameList{
	// · · Name: &parser.Name{ /* recursive/repetitive pointee not shown */ },
	// · },
	// · ReservedWord: &parser.ReservedWord{
	// · · Case: 1,
	// · · Token: example.y:1:1: LEFT "%left", Comments: [],
	// · },
	// }
}

func ExampleDefinition_case4() {
	fmt.Println(exampleAST(8, "%left %error-verbose"))
	// Output:
	// &parser.Definition{
	// · Case: 4,
	// · ReservedWord: &parser.ReservedWord{
	// · · Case: 1,
	// · · Token: example.y:1:1: LEFT "%left", Comments: [],
	// · },
	// }
}

func ExampleDefinition_case5() {
	fmt.Println(exampleAST(9, "%error-verbose %error-verbose"))
	// Output:
	// &parser.Definition{
	// · Case: 5,
	// · Token: example.y:1:1: ERROR_VERBOSE "%error-verbose", Comments: [],
	// }
}

func ExampleDefinitionList() {
	fmt.Println(exampleAST(10, "%error-verbose") == (*DefinitionList)(nil))
	// Output:
	// true
}

func ExampleDefinitionList_case1() {
	fmt.Println(exampleAST(11, "%error-verbose %error-verbose"))
	// Output:
	// &parser.DefinitionList{
	// · Definition: &parser.Definition{
	// · · Case: 5,
	// · · Token: example.y:1:1: ERROR_VERBOSE "%error-verbose", Comments: [],
	// · },
	// }
}

func ExampleLiteralStringOpt() {
	fmt.Println(exampleAST(12, "%left a ,") == (*LiteralStringOpt)(nil))
	// Output:
	// true
}

func ExampleLiteralStringOpt_case1() {
	fmt.Println(exampleAST(13, "%left a \"b\" ,"))
	// Output:
	// &parser.LiteralStringOpt{
	// · Token: example.y:1:9: STRING_LITERAL "\"b\"", Comments: [],
	// }
}

func ExampleName() {
	fmt.Println(exampleAST(14, "%left a ,"))
	// Output:
	// &parser.Name{
	// · Identifier: "a",
	// · Number: -1,
	// · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// }
}

func ExampleName_case1() {
	fmt.Println(exampleAST(15, "%left a 98 ,"))
	// Output:
	// &parser.Name{
	// · Identifier: "a",
	// · Number: 98,
	// · Case: 1,
	// · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · Token2: example.y:1:9: NUMBER "98", Comments: [],
	// }
}

func ExampleNameList() {
	fmt.Println(exampleAST(16, "%left a ,"))
	// Output:
	// &parser.NameList{
	// · Name: &parser.Name{
	// · · Identifier: "a",
	// · · Number: -1,
	// · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · },
	// }
}

func ExampleNameList_case1() {
	fmt.Println(exampleAST(17, "%left a b ,"))
	// Output:
	// &parser.NameList{
	// · Name: &parser.Name{
	// · · Identifier: "a",
	// · · Number: -1,
	// · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · },
	// · NameList: &parser.NameList{
	// · · Case: 1,
	// · · Name: &parser.Name{
	// · · · Identifier: "b",
	// · · · Number: -1,
	// · · · Token: example.y:1:9: IDENTIFIER "b", Comments: [],
	// · · },
	// · },
	// }
}

func ExampleNameList_case2() {
	fmt.Println(exampleAST(18, "%left a , b ,"))
	// Output:
	// &parser.NameList{
	// · Name: &parser.Name{
	// · · Identifier: "a",
	// · · Number: -1,
	// · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · },
	// · NameList: &parser.NameList{
	// · · Case: 2,
	// · · Name: &parser.Name{
	// · · · Identifier: "b",
	// · · · Number: -1,
	// · · · Token: example.y:1:11: IDENTIFIER "b", Comments: [],
	// · · },
	// · · Token: example.y:1:9: ',' ",", Comments: [],
	// · },
	// }
}

func ExamplePrecedence() {
	fmt.Println(exampleAST(19, "%% a: |") == (*Precedence)(nil))
	// Output:
	// true
}

func ExamplePrecedence_case1() {
	fmt.Println(exampleAST(20, "%% a: %prec b"))
	// Output:
	// &parser.Precedence{
	// · Identifier: "b",
	// · Case: 1,
	// · Token: example.y:1:7: PREC "%prec", Comments: [],
	// · Token2: example.y:1:13: IDENTIFIER "b", Comments: [],
	// }
}

func ExamplePrecedence_case2() {
	fmt.Println(exampleAST(21, "%% a: %prec b { }"))
	// Output:
	// &parser.Precedence{
	// · Identifier: "b",
	// · Action: &parser.Action{
	// · · Values: []*parser.ActionValue{ // len 1
	// · · · 0: &parser.ActionValue{
	// · · · · Pos: 15,
	// · · · · Src: "{ }",
	// · · · },
	// · · },
	// · · Token: example.y:1:15: '{' "{", Comments: [],
	// · · Token2: example.y:1:17: '}' "}", Comments: [],
	// · },
	// · Case: 2,
	// · Token: example.y:1:7: PREC "%prec", Comments: [],
	// · Token2: example.y:1:13: IDENTIFIER "b", Comments: [],
	// }
}

func ExamplePrecedence_case3() {
	fmt.Println(exampleAST(22, "%% a: ;"))
	// Output:
	// &parser.Precedence{
	// · Case: 3,
	// · Token: example.y:1:7: ';' ";", Comments: [],
	// }
}

func ExampleReservedWord() {
	fmt.Println(exampleAST(23, "%token <"))
	// Output:
	// &parser.ReservedWord{
	// · Token: example.y:1:1: TOKEN "%token", Comments: [],
	// }
}

func ExampleReservedWord_case1() {
	fmt.Println(exampleAST(24, "%left <"))
	// Output:
	// &parser.ReservedWord{
	// · Case: 1,
	// · Token: example.y:1:1: LEFT "%left", Comments: [],
	// }
}

func ExampleReservedWord_case2() {
	fmt.Println(exampleAST(25, "%right <"))
	// Output:
	// &parser.ReservedWord{
	// · Case: 2,
	// · Token: example.y:1:1: RIGHT "%right", Comments: [],
	// }
}

func ExampleReservedWord_case3() {
	fmt.Println(exampleAST(26, "%nonassoc <"))
	// Output:
	// &parser.ReservedWord{
	// · Case: 3,
	// · Token: example.y:1:1: NONASSOC "%nonassoc", Comments: [],
	// }
}

func ExampleReservedWord_case4() {
	fmt.Println(exampleAST(27, "%type <"))
	// Output:
	// &parser.ReservedWord{
	// · Case: 4,
	// · Token: example.y:1:1: TYPE "%type", Comments: [],
	// }
}

func ExampleReservedWord_case5() {
	fmt.Println(exampleAST(28, "%precedence <"))
	// Output:
	// &parser.ReservedWord{
	// · Case: 5,
	// · Token: example.y:1:1: PRECEDENCE "%precedence", Comments: [],
	// }
}

func ExampleRule() {
	fmt.Println(exampleAST(29, "%% a: b:"))
	// Output:
	// &parser.Rule{
	// · Name: example.y:1:7: C_IDENTIFIER "b", Comments: [],
	// · Token: example.y:1:7: C_IDENTIFIER "b", Comments: [],
	// }
}

func ExampleRule_case1() {
	fmt.Println(exampleAST(30, "%% a: |"))
	// Output:
	// &parser.Rule{
	// · Name: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · Case: 1,
	// · Token: example.y:1:7: '|' "|", Comments: [],
	// }
}

func ExampleRuleItemList() {
	fmt.Println(exampleAST(31, "%% a:") == (*RuleItemList)(nil))
	// Output:
	// true
}

func ExampleRuleItemList_case1() {
	fmt.Println(exampleAST(32, "%% a: b"))
	// Output:
	// &parser.RuleItemList{
	// · Case: 1,
	// · Token: example.y:1:7: IDENTIFIER "b", Comments: [],
	// }
}

func ExampleRuleItemList_case2() {
	fmt.Println(exampleAST(33, "%% a: { }"))
	// Output:
	// &parser.RuleItemList{
	// · Action: &parser.Action{
	// · · Values: []*parser.ActionValue{ // len 1
	// · · · 0: &parser.ActionValue{
	// · · · · Pos: 7,
	// · · · · Src: "{ }",
	// · · · },
	// · · },
	// · · Token: example.y:1:7: '{' "{", Comments: [],
	// · · Token2: example.y:1:9: '}' "}", Comments: [],
	// · },
	// · Case: 2,
	// }
}

func ExampleRuleItemList_case3() {
	fmt.Println(exampleAST(34, "%% a: \"b\""))
	// Output:
	// &parser.RuleItemList{
	// · Case: 3,
	// · Token: example.y:1:7: STRING_LITERAL "\"b\"", Comments: [],
	// }
}

func ExampleRuleList() {
	fmt.Println(exampleAST(35, "%% a:"))
	// Output:
	// &parser.RuleList{
	// · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// }
}

func ExampleRuleList_case1() {
	fmt.Println(exampleAST(36, "%% a: |"))
	// Output:
	// &parser.RuleList{
	// · RuleList: &parser.RuleList{
	// · · Case: 1,
	// · · Rule: &parser.Rule{
	// · · · Name: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · · · Case: 1,
	// · · · Token: example.y:1:7: '|' "|", Comments: [],
	// · · },
	// · },
	// · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// }
}

func ExampleSpecification() {
	fmt.Println(exampleAST(37, "%% a:"))
	// Output:
	// &parser.Specification{
	// · Rules: []*parser.Rule{ // len 1
	// · · 0: &parser.Rule{
	// · · · Name: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · · · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · · },
	// · },
	// · RuleList: &parser.RuleList{
	// · · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · },
	// · Token: example.y:1:1: MARK "%%", Comments: [],
	// }
}

func ExampleTag() {
	fmt.Println(exampleAST(38, "%left %error-verbose") == (*Tag)(nil))
	// Output:
	// true
}

func ExampleTag_case1() {
	fmt.Println(exampleAST(39, "%left < a > %error-verbose"))
	// Output:
	// &parser.Tag{
	// · Token: example.y:1:7: '<' "<", Comments: [],
	// · Token2: example.y:1:9: IDENTIFIER "a", Comments: [],
	// · Token3: example.y:1:11: '>' ">", Comments: [],
	// }
}

func ExampleTail() {
	fmt.Println(exampleAST(40, "%% a: %%"))
	// Output:
	// &parser.Tail{
	// · Token: example.y:1:7: MARK "%%", Comments: [],
	// }
}

func ExampleTail_case1() {
	fmt.Println(exampleAST(41, "%% a:") == (*Tail)(nil))
	// Output:
	// true
}