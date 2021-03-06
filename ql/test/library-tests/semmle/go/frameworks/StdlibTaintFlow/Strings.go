// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import (
	"io"
	"strings"
)

func TaintStepTest_StringsFields_B0I0O0(sourceCQL interface{}) interface{} {
	fromString656 := sourceCQL.(string)
	intoString414 := strings.Fields(fromString656)
	return intoString414
}

func TaintStepTest_StringsFieldsFunc_B0I0O0(sourceCQL interface{}) interface{} {
	fromString518 := sourceCQL.(string)
	intoString650 := strings.FieldsFunc(fromString518, nil)
	return intoString650
}

func TaintStepTest_StringsJoin_B0I0O0(sourceCQL interface{}) interface{} {
	fromString784 := sourceCQL.([]string)
	intoString957 := strings.Join(fromString784, "")
	return intoString957
}

func TaintStepTest_StringsJoin_B0I1O0(sourceCQL interface{}) interface{} {
	fromString520 := sourceCQL.(string)
	intoString443 := strings.Join(nil, fromString520)
	return intoString443
}

func TaintStepTest_StringsMap_B0I0O0(sourceCQL interface{}) interface{} {
	fromString127 := sourceCQL.(string)
	intoString483 := strings.Map(nil, fromString127)
	return intoString483
}

func TaintStepTest_StringsNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	fromString989 := sourceCQL.(string)
	intoReader982 := strings.NewReader(fromString989)
	return intoReader982
}

func TaintStepTest_StringsNewReplacer_B0I0O0(sourceCQL interface{}) interface{} {
	fromString417 := sourceCQL.(string)
	intoReplacer584 := strings.NewReplacer(fromString417)
	return intoReplacer584
}

func TaintStepTest_StringsRepeat_B0I0O0(sourceCQL interface{}) interface{} {
	fromString991 := sourceCQL.(string)
	intoString881 := strings.Repeat(fromString991, 0)
	return intoString881
}

func TaintStepTest_StringsReplace_B0I0O0(sourceCQL interface{}) interface{} {
	fromString186 := sourceCQL.(string)
	intoString284 := strings.Replace(fromString186, "", "", 0)
	return intoString284
}

func TaintStepTest_StringsReplace_B0I1O0(sourceCQL interface{}) interface{} {
	fromString908 := sourceCQL.(string)
	intoString137 := strings.Replace("", "", fromString908, 0)
	return intoString137
}

func TaintStepTest_StringsReplaceAll_B0I0O0(sourceCQL interface{}) interface{} {
	fromString494 := sourceCQL.(string)
	intoString873 := strings.ReplaceAll(fromString494, "", "")
	return intoString873
}

func TaintStepTest_StringsReplaceAll_B0I1O0(sourceCQL interface{}) interface{} {
	fromString599 := sourceCQL.(string)
	intoString409 := strings.ReplaceAll("", "", fromString599)
	return intoString409
}

func TaintStepTest_StringsSplit_B0I0O0(sourceCQL interface{}) interface{} {
	fromString246 := sourceCQL.(string)
	intoString898 := strings.Split(fromString246, "")
	return intoString898
}

func TaintStepTest_StringsSplitAfter_B0I0O0(sourceCQL interface{}) interface{} {
	fromString598 := sourceCQL.(string)
	intoString631 := strings.SplitAfter(fromString598, "")
	return intoString631
}

func TaintStepTest_StringsSplitAfterN_B0I0O0(sourceCQL interface{}) interface{} {
	fromString165 := sourceCQL.(string)
	intoString150 := strings.SplitAfterN(fromString165, "", 0)
	return intoString150
}

func TaintStepTest_StringsSplitN_B0I0O0(sourceCQL interface{}) interface{} {
	fromString340 := sourceCQL.(string)
	intoString471 := strings.SplitN(fromString340, "", 0)
	return intoString471
}

func TaintStepTest_StringsTitle_B0I0O0(sourceCQL interface{}) interface{} {
	fromString290 := sourceCQL.(string)
	intoString758 := strings.Title(fromString290)
	return intoString758
}

func TaintStepTest_StringsToLower_B0I0O0(sourceCQL interface{}) interface{} {
	fromString396 := sourceCQL.(string)
	intoString707 := strings.ToLower(fromString396)
	return intoString707
}

func TaintStepTest_StringsToLowerSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	fromString912 := sourceCQL.(string)
	intoString718 := strings.ToLowerSpecial(nil, fromString912)
	return intoString718
}

func TaintStepTest_StringsToTitle_B0I0O0(sourceCQL interface{}) interface{} {
	fromString972 := sourceCQL.(string)
	intoString633 := strings.ToTitle(fromString972)
	return intoString633
}

func TaintStepTest_StringsToTitleSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	fromString316 := sourceCQL.(string)
	intoString145 := strings.ToTitleSpecial(nil, fromString316)
	return intoString145
}

func TaintStepTest_StringsToUpper_B0I0O0(sourceCQL interface{}) interface{} {
	fromString817 := sourceCQL.(string)
	intoString474 := strings.ToUpper(fromString817)
	return intoString474
}

func TaintStepTest_StringsToUpperSpecial_B0I0O0(sourceCQL interface{}) interface{} {
	fromString832 := sourceCQL.(string)
	intoString378 := strings.ToUpperSpecial(nil, fromString832)
	return intoString378
}

func TaintStepTest_StringsToValidUTF8_B0I0O0(sourceCQL interface{}) interface{} {
	fromString541 := sourceCQL.(string)
	intoString139 := strings.ToValidUTF8(fromString541, "")
	return intoString139
}

func TaintStepTest_StringsToValidUTF8_B0I1O0(sourceCQL interface{}) interface{} {
	fromString814 := sourceCQL.(string)
	intoString768 := strings.ToValidUTF8("", fromString814)
	return intoString768
}

func TaintStepTest_StringsTrim_B0I0O0(sourceCQL interface{}) interface{} {
	fromString468 := sourceCQL.(string)
	intoString736 := strings.Trim(fromString468, "")
	return intoString736
}

func TaintStepTest_StringsTrimFunc_B0I0O0(sourceCQL interface{}) interface{} {
	fromString516 := sourceCQL.(string)
	intoString246 := strings.TrimFunc(fromString516, nil)
	return intoString246
}

func TaintStepTest_StringsTrimLeft_B0I0O0(sourceCQL interface{}) interface{} {
	fromString679 := sourceCQL.(string)
	intoString736 := strings.TrimLeft(fromString679, "")
	return intoString736
}

func TaintStepTest_StringsTrimLeftFunc_B0I0O0(sourceCQL interface{}) interface{} {
	fromString839 := sourceCQL.(string)
	intoString273 := strings.TrimLeftFunc(fromString839, nil)
	return intoString273
}

func TaintStepTest_StringsTrimPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	fromString982 := sourceCQL.(string)
	intoString458 := strings.TrimPrefix(fromString982, "")
	return intoString458
}

func TaintStepTest_StringsTrimRight_B0I0O0(sourceCQL interface{}) interface{} {
	fromString506 := sourceCQL.(string)
	intoString213 := strings.TrimRight(fromString506, "")
	return intoString213
}

func TaintStepTest_StringsTrimRightFunc_B0I0O0(sourceCQL interface{}) interface{} {
	fromString468 := sourceCQL.(string)
	intoString219 := strings.TrimRightFunc(fromString468, nil)
	return intoString219
}

func TaintStepTest_StringsTrimSpace_B0I0O0(sourceCQL interface{}) interface{} {
	fromString265 := sourceCQL.(string)
	intoString971 := strings.TrimSpace(fromString265)
	return intoString971
}

func TaintStepTest_StringsTrimSuffix_B0I0O0(sourceCQL interface{}) interface{} {
	fromString320 := sourceCQL.(string)
	intoString545 := strings.TrimSuffix(fromString320, "")
	return intoString545
}

func TaintStepTest_StringsBuilderString_B0I0O0(sourceCQL interface{}) interface{} {
	fromBuilder566 := sourceCQL.(strings.Builder)
	intoString497 := fromBuilder566.String()
	return intoString497
}

func TaintStepTest_StringsBuilderWrite_B0I0O0(sourceCQL interface{}) interface{} {
	fromByte274 := sourceCQL.([]byte)
	var intoBuilder783 strings.Builder
	intoBuilder783.Write(fromByte274)
	return intoBuilder783
}

func TaintStepTest_StringsBuilderWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	fromString905 := sourceCQL.(string)
	var intoBuilder389 strings.Builder
	intoBuilder389.WriteString(fromString905)
	return intoBuilder389
}

func TaintStepTest_StringsReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader198 := sourceCQL.(strings.Reader)
	var intoByte477 []byte
	fromReader198.Read(intoByte477)
	return intoByte477
}

func TaintStepTest_StringsReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader544 := sourceCQL.(strings.Reader)
	var intoByte382 []byte
	fromReader544.ReadAt(intoByte382, 0)
	return intoByte382
}

func TaintStepTest_StringsReaderReset_B0I0O0(sourceCQL interface{}) interface{} {
	fromString715 := sourceCQL.(string)
	var intoReader179 strings.Reader
	intoReader179.Reset(fromString715)
	return intoReader179
}

func TaintStepTest_StringsReaderWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader366 := sourceCQL.(strings.Reader)
	var intoWriter648 io.Writer
	fromReader366.WriteTo(intoWriter648)
	return intoWriter648
}

func TaintStepTest_StringsReplacerReplace_B0I0O0(sourceCQL interface{}) interface{} {
	fromString544 := sourceCQL.(string)
	var mediumObjCQL strings.Replacer
	intoString484 := mediumObjCQL.Replace(fromString544)
	return intoString484
}

func TaintStepTest_StringsReplacerWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	fromString824 := sourceCQL.(string)
	var intoWriter754 io.Writer
	var mediumObjCQL strings.Replacer
	mediumObjCQL.WriteString(intoWriter754, fromString824)
	return intoWriter754
}

func RunAllTaints_Strings() {
	{
		source := newSource(0)
		out := TaintStepTest_StringsFields_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_StringsFieldsFunc_B0I0O0(source)
		sink(1, out)
	}
	{
		source := newSource(2)
		out := TaintStepTest_StringsJoin_B0I0O0(source)
		sink(2, out)
	}
	{
		source := newSource(3)
		out := TaintStepTest_StringsJoin_B0I1O0(source)
		sink(3, out)
	}
	{
		source := newSource(4)
		out := TaintStepTest_StringsMap_B0I0O0(source)
		sink(4, out)
	}
	{
		source := newSource(5)
		out := TaintStepTest_StringsNewReader_B0I0O0(source)
		sink(5, out)
	}
	{
		source := newSource(6)
		out := TaintStepTest_StringsNewReplacer_B0I0O0(source)
		sink(6, out)
	}
	{
		source := newSource(7)
		out := TaintStepTest_StringsRepeat_B0I0O0(source)
		sink(7, out)
	}
	{
		source := newSource(8)
		out := TaintStepTest_StringsReplace_B0I0O0(source)
		sink(8, out)
	}
	{
		source := newSource(9)
		out := TaintStepTest_StringsReplace_B0I1O0(source)
		sink(9, out)
	}
	{
		source := newSource(10)
		out := TaintStepTest_StringsReplaceAll_B0I0O0(source)
		sink(10, out)
	}
	{
		source := newSource(11)
		out := TaintStepTest_StringsReplaceAll_B0I1O0(source)
		sink(11, out)
	}
	{
		source := newSource(12)
		out := TaintStepTest_StringsSplit_B0I0O0(source)
		sink(12, out)
	}
	{
		source := newSource(13)
		out := TaintStepTest_StringsSplitAfter_B0I0O0(source)
		sink(13, out)
	}
	{
		source := newSource(14)
		out := TaintStepTest_StringsSplitAfterN_B0I0O0(source)
		sink(14, out)
	}
	{
		source := newSource(15)
		out := TaintStepTest_StringsSplitN_B0I0O0(source)
		sink(15, out)
	}
	{
		source := newSource(16)
		out := TaintStepTest_StringsTitle_B0I0O0(source)
		sink(16, out)
	}
	{
		source := newSource(17)
		out := TaintStepTest_StringsToLower_B0I0O0(source)
		sink(17, out)
	}
	{
		source := newSource(18)
		out := TaintStepTest_StringsToLowerSpecial_B0I0O0(source)
		sink(18, out)
	}
	{
		source := newSource(19)
		out := TaintStepTest_StringsToTitle_B0I0O0(source)
		sink(19, out)
	}
	{
		source := newSource(20)
		out := TaintStepTest_StringsToTitleSpecial_B0I0O0(source)
		sink(20, out)
	}
	{
		source := newSource(21)
		out := TaintStepTest_StringsToUpper_B0I0O0(source)
		sink(21, out)
	}
	{
		source := newSource(22)
		out := TaintStepTest_StringsToUpperSpecial_B0I0O0(source)
		sink(22, out)
	}
	{
		source := newSource(23)
		out := TaintStepTest_StringsToValidUTF8_B0I0O0(source)
		sink(23, out)
	}
	{
		source := newSource(24)
		out := TaintStepTest_StringsToValidUTF8_B0I1O0(source)
		sink(24, out)
	}
	{
		source := newSource(25)
		out := TaintStepTest_StringsTrim_B0I0O0(source)
		sink(25, out)
	}
	{
		source := newSource(26)
		out := TaintStepTest_StringsTrimFunc_B0I0O0(source)
		sink(26, out)
	}
	{
		source := newSource(27)
		out := TaintStepTest_StringsTrimLeft_B0I0O0(source)
		sink(27, out)
	}
	{
		source := newSource(28)
		out := TaintStepTest_StringsTrimLeftFunc_B0I0O0(source)
		sink(28, out)
	}
	{
		source := newSource(29)
		out := TaintStepTest_StringsTrimPrefix_B0I0O0(source)
		sink(29, out)
	}
	{
		source := newSource(30)
		out := TaintStepTest_StringsTrimRight_B0I0O0(source)
		sink(30, out)
	}
	{
		source := newSource(31)
		out := TaintStepTest_StringsTrimRightFunc_B0I0O0(source)
		sink(31, out)
	}
	{
		source := newSource(32)
		out := TaintStepTest_StringsTrimSpace_B0I0O0(source)
		sink(32, out)
	}
	{
		source := newSource(33)
		out := TaintStepTest_StringsTrimSuffix_B0I0O0(source)
		sink(33, out)
	}
	{
		source := newSource(34)
		out := TaintStepTest_StringsBuilderString_B0I0O0(source)
		sink(34, out)
	}
	{
		source := newSource(35)
		out := TaintStepTest_StringsBuilderWrite_B0I0O0(source)
		sink(35, out)
	}
	{
		source := newSource(36)
		out := TaintStepTest_StringsBuilderWriteString_B0I0O0(source)
		sink(36, out)
	}
	{
		source := newSource(37)
		out := TaintStepTest_StringsReaderRead_B0I0O0(source)
		sink(37, out)
	}
	{
		source := newSource(38)
		out := TaintStepTest_StringsReaderReadAt_B0I0O0(source)
		sink(38, out)
	}
	{
		source := newSource(39)
		out := TaintStepTest_StringsReaderReset_B0I0O0(source)
		sink(39, out)
	}
	{
		source := newSource(40)
		out := TaintStepTest_StringsReaderWriteTo_B0I0O0(source)
		sink(40, out)
	}
	{
		source := newSource(41)
		out := TaintStepTest_StringsReplacerReplace_B0I0O0(source)
		sink(41, out)
	}
	{
		source := newSource(42)
		out := TaintStepTest_StringsReplacerWriteString_B0I0O0(source)
		sink(42, out)
	}
}
