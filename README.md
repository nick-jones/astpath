# astpath

Quick hack that provides XPath querying over Go ASTs. Inspired by [hchasestevens/astpath](https://github.com/hchasestevens/astpath)
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
$ astpath --template='{{.XML}}' '/File' test/test.go | xmllint --format - | xmllint --format -
```

```xml    
<?xml version="1.0" encoding="UTF-8"?>
<File pos-start="1" pos-end="239">
  <Ident name="test" pos-start="9" pos-end="13"/>
  <GenDecl pos-start="15" pos-end="43">
    <ImportSpec pos-start="25" pos-end="30">
      <BasicLit kind="STRING" value="log" pos-start="25" pos-end="30"/>
    </ImportSpec>
    <ImportSpec pos-start="32" pos-end="41">
      <BasicLit kind="STRING" value="strings" pos-start="32" pos-end="41"/>
    </ImportSpec>
  </GenDecl>
  <FuncDecl pos-start="45" pos-end="239">
    <Ident name="repeatConditional" pos-start="50" pos-end="67">
      <Object name="repeatConditional" kind="func"/>
    </Ident>
    <FuncType pos-start="45" pos-end="119">
      <FieldList pos-start="67" pos-end="112">
        <Field pos-start="68" pos-end="78">
          <Ident name="str" pos-start="68" pos-end="71">
            <Object name="str" kind="var"/>
          </Ident>
          <Ident name="string" pos-start="72" pos-end="78"/>
        </Field>
        <Field pos-start="80" pos-end="89">
          <Ident name="count" pos-start="80" pos-end="85">
            <Object name="count" kind="var"/>
          </Ident>
          <Ident name="int" pos-start="86" pos-end="89"/>
        </Field>
        <Field pos-start="91" pos-end="111">
          <Ident name="fn" pos-start="91" pos-end="93">
            <Object name="fn" kind="var"/>
          </Ident>
          <FuncType pos-start="94" pos-end="111">
            <FieldList pos-start="98" pos-end="106">
              <Field pos-start="99" pos-end="105">
                <Ident name="string" pos-start="99" pos-end="105"/>
              </Field>
            </FieldList>
            <FieldList pos-start="107" pos-end="111">
              <Field pos-start="107" pos-end="111">
                <Ident name="bool" pos-start="107" pos-end="111"/>
              </Field>
            </FieldList>
          </FuncType>
        </Field>
      </FieldList>
      <FieldList pos-start="113" pos-end="119">
        <Field pos-start="113" pos-end="119">
          <Ident name="string" pos-start="113" pos-end="119"/>
        </Field>
      </FieldList>
    </FuncType>
    <BlockStmt pos-start="120" pos-end="239">
      <IfStmt pos-start="123" pos-end="196">
        <CallExpr pos-start="126" pos-end="133">
          <Ident name="fn" pos-start="126" pos-end="128">
            <Object name="fn" kind="var"/>
          </Ident>
          <Ident name="str" pos-start="129" pos-end="132">
            <Object name="str" kind="var"/>
          </Ident>
        </CallExpr>
        <BlockStmt pos-start="134" pos-end="196">
          <ExprStmt pos-start="138" pos-end="157">
            <CallExpr pos-start="138" pos-end="157">
              <SelectorExpr pos-start="138" pos-end="149">
                <Ident name="log" pos-start="138" pos-end="141"/>
                <Ident name="Println" pos-start="142" pos-end="149"/>
              </SelectorExpr>
              <BasicLit kind="STRING" value="hit!" pos-start="150" pos-end="156"/>
            </CallExpr>
          </ExprStmt>
          <ReturnStmt pos-start="160" pos-end="193">
            <CallExpr pos-start="167" pos-end="193">
              <SelectorExpr pos-start="167" pos-end="181">
                <Ident name="strings" pos-start="167" pos-end="174"/>
                <Ident name="Repeat" pos-start="175" pos-end="181"/>
              </SelectorExpr>
              <Ident name="str" pos-start="182" pos-end="185">
                <Object name="str" kind="var"/>
              </Ident>
              <Ident name="count" pos-start="187" pos-end="192">
                <Object name="count" kind="var"/>
              </Ident>
            </CallExpr>
          </ReturnStmt>
        </BlockStmt>
      </IfStmt>
      <ExprStmt pos-start="198" pos-end="225">
        <CallExpr pos-start="198" pos-end="225">
          <SelectorExpr pos-start="198" pos-end="208">
            <Ident name="log" pos-start="198" pos-end="201"/>
            <Ident name="Printf" pos-start="202" pos-end="208"/>
          </SelectorExpr>
          <BasicLit kind="STRING" value="str = %s" pos-start="209" pos-end="219"/>
          <Ident name="str" pos-start="221" pos-end="224">
            <Object name="str" kind="var"/>
          </Ident>
        </CallExpr>
      </ExprStmt>
      <ReturnStmt pos-start="227" pos-end="237">
        <Ident name="str" pos-start="234" pos-end="237">
          <Object name="str" kind="var"/>
        </Ident>
      </ReturnStmt>
    </BlockStmt>
  </FuncDecl>
</File>
```

Extract import paths:

```
$ astpath '//ImportSpec' test.go
test.go:4:2 > "log"
test.go:5:2 > "strings"
```

... or

```
$ astpath --template='{{.XMLInner}}' '//ImportSpec/BasicLit/@value' test.go
log
strings
```

Locate `log.*` function calls:

```
$ astpath '//CallExpr/SelectorExpr[./Ident[1]/@name="log"]' test.go
test.go:10:3 > log.Println
test.go:13:2 > log.Printf
```

Locate `log.Printf` function calls:

```
$ astpath '//CallExpr/SelectorExpr[./Ident[1]/@name="log"][./Ident[2]/@name="Printf"]' test.go
test.go:13:2 > log.Printf
```



 