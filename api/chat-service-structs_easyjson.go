// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonAfe46659Decode20191TheBangApi(in *jlexer.Lexer, out *ChatMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "timestamp":
			out.Timestamp = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "edited":
			out.Edited = bool(in.Bool())
		case "deleted":
			out.Deleted = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAfe46659Encode20191TheBangApi(out *jwriter.Writer, in ChatMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Timestamp))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"edited\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Edited))
	}
	{
		const prefix string = ",\"deleted\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Deleted))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAfe46659Encode20191TheBangApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAfe46659Encode20191TheBangApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAfe46659Decode20191TheBangApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAfe46659Decode20191TheBangApi(l, v)
}
