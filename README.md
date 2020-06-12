# astpath

Quick hack that provides XPath querying over Go ASTs. Inspired by [astpath](https://github.com/hchasestevens/astpath)
for Python.

## Installing

```
GO111MODULE=on go get github.com/nick-jones/astpath
```

## Usage

```
astpath <xpath-expr> <file-path>
```

The file path can be a directory, in which case all files in that directly are checked recursively.

## Examples

Taking a basic and crude example:

```go
package test

import (
	"log"
	"strings"
)

func repeatConditional(str string, count int, fn func(string) bool) string {
	if fn(str) {
		log.Println("hit!")
		return strings.Repeat(str, count)
	}
	log.Printf("str = %s", str)
	return str
}
```

To view the raw XML output for a single file:

```
$ astpath --print-mode="xml-inner" '/' test.go | xmllint --format -
```

```xml    
<?xml version="1.0" encoding="UTF-8"?>
<File pos-start="1" pos-end="239">
  <Ident pos-start="9" pos-end="13">
    <Name>test</Name>
  </Ident>
  <GenDecl pos-start="15" pos-end="43">
    <ImportSpec pos-start="25" pos-end="30">
      <BasicLit kind="STRING" value="log" pos-start="25" pos-end="30"/>
    </ImportSpec>
    <ImportSpec pos-start="32" pos-end="41">
      <BasicLit kind="STRING" value="strings" pos-start="32" pos-end="41"/>
    </ImportSpec>
  </GenDecl>
  <FuncDecl pos-start="45" pos-end="239">
    <Ident pos-start="50" pos-end="67">
      <Name>repeatConditional</Name>
      <Object name="repeatConditional" kind="func"/>
    </Ident>
    <FuncType pos-start="45" pos-end="119">
      <FieldList pos-start="67" pos-end="112">
        <Field pos-start="68" pos-end="78">
          <Ident pos-start="68" pos-end="71">
            <Name>str</Name>
            <Object name="str" kind="var"/>
          </Ident>
          <Ident pos-start="72" pos-end="78">
            <Name>string</Name>
          </Ident>
        </Field>
        <Field pos-start="80" pos-end="89">
          <Ident pos-start="80" pos-end="85">
            <Name>count</Name>
            <Object name="count" kind="var"/>
          </Ident>
          <Ident pos-start="86" pos-end="89">
            <Name>int</Name>
          </Ident>
        </Field>
        <Field pos-start="91" pos-end="111">
          <Ident pos-start="91" pos-end="93">
            <Name>fn</Name>
            <Object name="fn" kind="var"/>
          </Ident>
          <FuncType pos-start="94" pos-end="111">
            <FieldList pos-start="98" pos-end="106">
              <Field pos-start="99" pos-end="105">
                <Ident pos-start="99" pos-end="105">
                  <Name>string</Name>
                </Ident>
              </Field>
            </FieldList>
            <FieldList pos-start="107" pos-end="111">
              <Field pos-start="107" pos-end="111">
                <Ident pos-start="107" pos-end="111">
                  <Name>bool</Name>
                </Ident>
              </Field>
            </FieldList>
          </FuncType>
        </Field>
      </FieldList>
      <FieldList pos-start="113" pos-end="119">
        <Field pos-start="113" pos-end="119">
          <Ident pos-start="113" pos-end="119">
            <Name>string</Name>
          </Ident>
        </Field>
      </FieldList>
    </FuncType>
    <BlockStmt pos-start="120" pos-end="239">
      <IfStmt pos-start="123" pos-end="196">
        <CallExpr pos-start="126" pos-end="133">
          <Ident pos-start="126" pos-end="128">
            <Name>fn</Name>
            <Object name="fn" kind="var"/>
          </Ident>
          <Ident pos-start="129" pos-end="132">
            <Name>str</Name>
            <Object name="str" kind="var"/>
          </Ident>
        </CallExpr>
        <BlockStmt pos-start="134" pos-end="196">
          <ExprStmt pos-start="138" pos-end="157">
            <CallExpr pos-start="138" pos-end="157">
              <SelectorExpr pos-start="138" pos-end="149">
                <Ident pos-start="138" pos-end="141">
                  <Name>log</Name>
                </Ident>
                <Ident pos-start="142" pos-end="149">
                  <Name>Println</Name>
                </Ident>
              </SelectorExpr>
              <BasicLit kind="STRING" value="hit!" pos-start="150" pos-end="156"/>
            </CallExpr>
          </ExprStmt>
          <ReturnStmt pos-start="160" pos-end="193">
            <CallExpr pos-start="167" pos-end="193">
              <SelectorExpr pos-start="167" pos-end="181">
                <Ident pos-start="167" pos-end="174">
                  <Name>strings</Name>
                </Ident>
                <Ident pos-start="175" pos-end="181">
                  <Name>Repeat</Name>
                </Ident>
              </SelectorExpr>
              <Ident pos-start="182" pos-end="185">
                <Name>str</Name>
                <Object name="str" kind="var"/>
              </Ident>
              <Ident pos-start="187" pos-end="192">
                <Name>count</Name>
                <Object name="count" kind="var"/>
              </Ident>
            </CallExpr>
          </ReturnStmt>
        </BlockStmt>
      </IfStmt>
      <ExprStmt pos-start="198" pos-end="225">
        <CallExpr pos-start="198" pos-end="225">
          <SelectorExpr pos-start="198" pos-end="208">
            <Ident pos-start="198" pos-end="201">
              <Name>log</Name>
            </Ident>
            <Ident pos-start="202" pos-end="208">
              <Name>Printf</Name>
            </Ident>
          </SelectorExpr>
          <BasicLit kind="STRING" value="str = %s" pos-start="209" pos-end="219"/>
          <Ident pos-start="221" pos-end="224">
            <Name>str</Name>
            <Object name="str" kind="var"/>
          </Ident>
        </CallExpr>
      </ExprStmt>
      <ReturnStmt pos-start="227" pos-end="237">
        <Ident pos-start="234" pos-end="237">
          <Name>str</Name>
          <Object name="str" kind="var"/>
        </Ident>
      </ReturnStmt>
    </BlockStmt>
  </FuncDecl>
</File>
```

Extract import paths:

```
$ astpath '//ImportSpec/BasicLit' test.go 
test/test.go:4:2 > "log"
test/test.go:5:2 > "strings"
```

Locate `log.*` function calls:

```
$ astpath '//CallExpr/SelectorExpr[./Ident[1]/Name="log"]' test.go     
test/test.go:10:3 > log.Println
test/test.go:13:2 > log.Printf
```

Locate `log.Printf` function calls:

```
$ astpath '//CallExpr/SelectorExpr[./Ident[1]/Name="log"][./Ident[2]/Name="Printf"]' test.go
test/test.go:13:2 > log.Printf
```



 