# Parsly -  parsing utility.


The goal of this project is to simplify implementing parsers with 
a tokenizer with a set of commonly use token matchers.

## Usage

To build simple parser:
  - define token lexical registry matchers
```go
 ```
  - create a tokenizer
  - add parsing logic with MatchAfterWhitespace, MatchAny or Match

```go

    var Whitespace = parsly.NewToken(0, "Whitespace", matcher.NewWhiteSpace())
    var Number = parsly.NewToken(1, "Number", matcher.NewNumber())
    var Term = parsly.NewToken(2, "Term", matcher.NewCharset("+-"))
    var Factor = parsly.NewToken(3, "Factor", matcher.NewCharset("*/"))
 
  
    func Parse(input []byte) (root *Expression, err error) {
        cursor := parsly.NewCursor("", input, 0)
        root = &Expression{}
        expression := root
        matched := cursor.MatchAfterOptional(Whitespace, Number)
        if matched.Code != Number.Code {
            return nil, cursor.NewError(Number)
        }
        value, _ := matched.Float(cursor)
        expression.LeftOp = NewValue(value)
    
        for ; ; {
            matched = cursor.MatchAfterOptional(Whitespace, Factor, Term)
            if matched.Code == parsly.EOF {
                break
            }
            operator := matched.Text(cursor)
            if expression.Operator != "" {
                expression.RightOp = &Operand{Expression: &Expression{LeftOp: expression.RightOp}}
                expression = expression.RightOp.Expression
            }
            expression.Operator = operator
    
            matched := cursor.MatchAfterOptional(Whitespace, Number)
            if matched.Code != Number.Code {
                return nil, cursor.NewError(Number)
            }
            value, _ := matched.Float(cursor)
            expression.RightOp = NewValue(value)
        }
        return root, nil
    }
    
```

See [example](example) basic arithmetic AST expression parser 

## Matchers
   This project implements the following [matchers](matcher):
   
   - [char](matcher/char.go) to match individual rune
   - [charset](matcher/charset.go) to match set of runes
   - [digit](matcher/digit.go) to match a digit
   - [letter](matcher/digit.go) to match a letter
   - [fragment](matcher/fragment.go) to match a fragment
   - [fragments](matcher/fragments.go) to match a set of fragments
   - [number](matcher/number.go) to match a number
   - [numberSign](matcher/numbersign.go) to match a number sign (+/-)
   - [quote](matcher/quote.go) to match a quoted fragment  
   - [whitespace](matcher/whitespace.go) to match a whitespaces