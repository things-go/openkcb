package openkcb

import (
	"bytes"
	"context"
	"encoding/json"
)

func Invoke[R Validator, T any](c *Client, ctx context.Context, transCode string, req *R, bizNo string) (rsp *T, err error) {
	result, err := invoke[R, T](c, ctx, transCode, req, bizNo)
	if err != nil {
		return nil, Error{Code: "-1", Msg: err.Error()}
	}
	if result.Code == "000000" {
		return result.Data, nil
	}

	return nil, Error{
		Code: result.Code,
		Msg:  result.Msg,
	}
}

func invoke[R Validator, T any](c *Client, ctx context.Context, transCode string, req *R, bizNo string) (rsp *Response[T], err error) {
	err = (*req).Validate()
	if err != nil {
		return nil, err
	}

	rsp = &Response[T]{}

	var rr ResponseRaw
	if err = c.Invoke(ctx, transCode, req, &rr, bizNo); err != nil {
		return nil, err
	}

	rsp.Flag = rr.Flag
	rsp.Code = rr.Code
	rsp.Msg = rr.Msg

	if rr.Data != "" {
		//var data T
		//if err = json.Unmarshal([]byte(rr.Data), &data); err != nil {
		//	return
		//}

		dec := json.NewDecoder(bytes.NewReader([]byte(rr.Data)))
		dec.UseNumber() // 避免精度丢失
		var data T
		if err = dec.Decode(&data); err != nil {
			return
		}

		rsp.Data = &data
	}

	return
}
