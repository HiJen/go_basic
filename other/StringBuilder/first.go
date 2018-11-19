package main

import (
	"bytes"
	"fmt"
	"strings"
)

//--- +号拼接------------------------------------------------------------
func StringPlus() string {
	var s string
	s += "昵称" + "ajunlab" + "\n"
	s += "博客" + "http://www.flysnow.org/" + "\n"
	s += "微信公众号" + ":" + "ajun"
	return s
}

//---fmt 拼接------------------------------------------------------------

func StringFmt() string {
	return fmt.Sprint("昵称", ":", "\n",
		"博客", ":",
		"http://www.flysnow.org/", "\n",
		"微信公众号", ":", "flysnow_org")
}

//---Join 拼接-----------------------------------------------------------

func StringJoin() string {
	s := []string{
		"昵称", ":", "\n",
		"博客", ":",
		"http://www.flysnow.org/", "\n",
		"微信公众号", ":", "flysnow_org",
	}
	return strings.Join(s, "")
}

//---buffer 拼接-----------------------------------------------------------

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("废墟去腥")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.flysnow.org/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("ajun")
	return b.String()
}

//---builder 拼接-----------------------------------------------------------

func StringBuilder() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("废墟去腥")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.flysnow.org/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("ajun")
	return b.String()
}
