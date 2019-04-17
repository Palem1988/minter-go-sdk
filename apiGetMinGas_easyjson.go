// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package mintersdk

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

func easyjsonF835620fDecodeGithubComValidatorCenterMinterGoSdkMin(in *jlexer.Lexer, out *min_gas) {
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
		case "jsonrpc":
			out.JSONRPC = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "result":
			out.Result = int64(in.Int64Str())
		case "error":
			(out.Error).UnmarshalEasyJSON(in)
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
func easyjsonF835620fEncodeGithubComValidatorCenterMinterGoSdkMin(out *jwriter.Writer, in min_gas) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.JSONRPC))
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64Str(int64(in.Result))
	}
	{
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Error).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v min_gas) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF835620fEncodeGithubComValidatorCenterMinterGoSdkMin(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v min_gas) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF835620fEncodeGithubComValidatorCenterMinterGoSdkMin(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *min_gas) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF835620fDecodeGithubComValidatorCenterMinterGoSdkMin(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *min_gas) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF835620fDecodeGithubComValidatorCenterMinterGoSdkMin(l, v)
}