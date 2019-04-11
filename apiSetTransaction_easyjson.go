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

func easyjsonB468fa26DecodeGithubComValidatorCenterMinterGoSdkSend(in *jlexer.Lexer, out *send_transaction) {
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
			(out.Result).UnmarshalEasyJSON(in)
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
func easyjsonB468fa26EncodeGithubComValidatorCenterMinterGoSdkSend(out *jwriter.Writer, in send_transaction) {
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
		(in.Result).MarshalEasyJSON(out)
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
func (v send_transaction) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB468fa26EncodeGithubComValidatorCenterMinterGoSdkSend(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v send_transaction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB468fa26EncodeGithubComValidatorCenterMinterGoSdkSend(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *send_transaction) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB468fa26DecodeGithubComValidatorCenterMinterGoSdkSend(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *send_transaction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB468fa26DecodeGithubComValidatorCenterMinterGoSdkSend(l, v)
}
func easyjsonB468fa26DecodeGithubComValidatorCenterMinterGoSdk(in *jlexer.Lexer, out *TransSendResponse) {
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
		case "code":
			out.Code = int(in.Int())
		case "log":
			out.Log = string(in.String())
		case "data":
			out.Data = string(in.String())
		case "hash":
			out.Hash = string(in.String())
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
func easyjsonB468fa26EncodeGithubComValidatorCenterMinterGoSdk(out *jwriter.Writer, in TransSendResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"log\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Log))
	}
	{
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Data))
	}
	{
		const prefix string = ",\"hash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Hash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TransSendResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB468fa26EncodeGithubComValidatorCenterMinterGoSdk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TransSendResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB468fa26EncodeGithubComValidatorCenterMinterGoSdk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TransSendResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB468fa26DecodeGithubComValidatorCenterMinterGoSdk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TransSendResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB468fa26DecodeGithubComValidatorCenterMinterGoSdk(l, v)
}
