package matcher

import (
	"bytes"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"unsafe"
)

type FragmentFold struct {
	embed *bool //can be embedded within word
	value []byte
	size  int
}

func (d *FragmentFold) isEmbed() bool {
	if d.embed == nil {
		return true
	}
	return *d.embed
}
func (d *FragmentFold) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	if matchEnd <= cursor.InputSize {
		if MatchFold(d.value, cursor.Input, 0, cursor.Pos) {
			if d.isEmbed() {
				return d.size
			}
			if matchEnd >= len(cursor.Input) {
				return d.size
			}
			if !IsWhiteSpace(cursor.Input[matchEnd]) {
				return 0
			}
			return d.size
		}
	}
	return 0
}

type Fragment struct {
	embed *bool //can be embedded within word
	value []byte
	size  int
}

func (d *Fragment) isEmbed() bool {
	if d.embed == nil {
		return true
	}
	return *d.embed
}

func (d *Fragment) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	if matchEnd <= cursor.InputSize {
		if bytes.Equal(cursor.Input[cursor.Pos:matchEnd], d.value) {
			if d.isEmbed() {
				return d.size
			}
			if matchEnd >= len(cursor.Input) {
				return d.size
			}
			if !IsWhiteSpace(cursor.Input[matchEnd]) {
				return 0
			}
			return d.size
		}
	}
	return 0
}

// NewFragments creates FragmentFold matcher
func NewFragment(value string, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	if AssignOption(options, &caseOpt) && !caseOpt.Sensitive {
		return &FragmentFold{
			value: []byte(value),
			size:  len(value),
			embed: caseOpt.Embed,
		}
	}
	return &Fragment{
		value: []byte(value),
		size:  len(value),
		embed: caseOpt.Embed,
	}
}

const (
	uint16LowerMask = uint16(0x2020)
	uint32LowerMask = uint32(0x20202020)
	uint64LowerMask = uint64(0x2020202020202020)
)

// MatchFold returns true if source  is matched with target (case insensitive)
func MatchFold(target, source []byte, targetOffset, sourceOffset int) bool {
	bytesToCheck := len(target) - targetOffset
	if sourceOffset+bytesToCheck > len(source) {
		return false
	}
	if (target[targetOffset] | 0x20) != (source[sourceOffset] | 0x20) {
		return false
	}
	last := len(target) - 1
	if (target[last] | 0x20) != (source[sourceOffset+last] | 0x20) {
		return false
	}
	switch bytesToCheck {
	case 1, 2:
		return true
	case 3:
		x := *(*uint16)(unsafe.Pointer(&target[targetOffset+1])) | uint16LowerMask
		y := *(*uint16)(unsafe.Pointer(&source[sourceOffset+1])) | uint16LowerMask
		return x == y
	case 4:
		x := *(*uint32)(unsafe.Pointer(&target[targetOffset])) | uint32LowerMask
		y := *(*uint32)(unsafe.Pointer(&source[sourceOffset])) | uint32LowerMask
		return x == y
	case 5, 6:
		x := *(*uint32)(unsafe.Pointer(&target[targetOffset+1])) | uint32LowerMask
		y := *(*uint32)(unsafe.Pointer(&source[sourceOffset+1])) | uint32LowerMask
		return x == y
	case 7:
		if (target[targetOffset+1] | 0x20) != (source[sourceOffset+1] | 0x20) {
			return false
		}
		x := *(*uint32)(unsafe.Pointer(&target[targetOffset+2])) | uint32LowerMask
		y := *(*uint32)(unsafe.Pointer(&source[sourceOffset+2])) | uint32LowerMask
		return x == y
	case 8:
		x := *(*uint64)(unsafe.Pointer(&target[targetOffset])) | uint64LowerMask
		y := *(*uint64)(unsafe.Pointer(&source[sourceOffset])) | uint64LowerMask
		if x != y {
			return false
		}
		return true
	}
	repeat := bytesToCheck / 8
	for i := 0; i < repeat; i++ {
		x := *(*uint64)(unsafe.Pointer(&target[targetOffset])) | uint64LowerMask
		y := *(*uint64)(unsafe.Pointer(&source[sourceOffset])) | uint64LowerMask
		if x != y {
			return false
		}
		targetOffset += 8
		sourceOffset += 8
	}

	rem := bytesToCheck % 8
	if rem == 0 {
		return true
	}
	for i := 0; i < rem; i++ {
		if (target[targetOffset] | 0x20) != (source[sourceOffset] | 0x20) {
			return false
		}
		targetOffset++
		sourceOffset++
	}
	return true
}
