// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package httpclient

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	_v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient(in *jlexer.Lexer, out *getTokenResponse) {
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
		case "expiresIn":
			out.ExpiresIn = int(in.Int())
		case "token":
			out.Token = string(in.String())
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
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient(out *jwriter.Writer, in getTokenResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"expiresIn\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ExpiresIn))
	}
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix)
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getTokenResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getTokenResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getTokenResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getTokenResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient1(in *jlexer.Lexer, out *getBranchesResponse) {
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
		case "branches":
			if in.IsNull() {
				in.Skip()
				out.Branches = nil
			} else {
				in.Delim('[')
				if out.Branches == nil {
					if !in.IsDelim(']') {
						out.Branches = make([]int, 0, 8)
					} else {
						out.Branches = []int{}
					}
				} else {
					out.Branches = (out.Branches)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int
					v1 = int(in.Int())
					out.Branches = append(out.Branches, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient1(out *jwriter.Writer, in getBranchesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"branches\":"
		out.RawString(prefix[1:])
		if in.Branches == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Branches {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getBranchesResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getBranchesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getBranchesResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getBranchesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient1(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient2(in *jlexer.Lexer, out *downloadDocumentResponse) {
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
		case "document":
			if in.IsNull() {
				in.Skip()
				out.Document = nil
			} else {
				out.Document = in.Bytes()
			}
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
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient2(out *jwriter.Writer, in downloadDocumentResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"document\":"
		out.RawString(prefix[1:])
		out.Base64Bytes(in.Document)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v downloadDocumentResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v downloadDocumentResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *downloadDocumentResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *downloadDocumentResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient2(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient3(in *jlexer.Lexer, out *createMultipartUploadResponse) {
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
		case "additionalErrors":
			if in.IsNull() {
				in.Skip()
				out.AdditionalErrors = nil
			} else {
				if out.AdditionalErrors == nil {
					out.AdditionalErrors = new(_v1.AdditionalErrors)
				}
				easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, out.AdditionalErrors)
			}
		case "data":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in, &out.Data)
		case "error":
			out.ErrorFlag = bool(in.Bool())
		case "errorText":
			out.ErrorText = string(in.String())
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
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient3(out *jwriter.Writer, in createMultipartUploadResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"additionalErrors\":"
		out.RawString(prefix[1:])
		if in.AdditionalErrors == nil {
			out.RawString("null")
		} else {
			easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, *in.AdditionalErrors)
		}
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out, in.Data)
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.Bool(bool(in.ErrorFlag))
	}
	{
		const prefix string = ",\"errorText\":"
		out.RawString(prefix)
		out.String(string(in.ErrorText))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createMultipartUploadResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createMultipartUploadResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createMultipartUploadResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createMultipartUploadResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleS3Httpclient3(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in *jlexer.Lexer, out *_v1.CreateMultipartUploadData) {
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
		case "upload_id":
			out.UploadID = string(in.String())
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
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out *jwriter.Writer, in _v1.CreateMultipartUploadData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"upload_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UploadID))
	}
	out.RawByte('}')
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in *jlexer.Lexer, out *_v1.AdditionalErrors) {
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
		case "name":
			out.Name = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
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
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out *jwriter.Writer, in _v1.AdditionalErrors) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	out.RawByte('}')
}
