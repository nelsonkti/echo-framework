// Package config
// @Author fuzengyao
// @Date 2022-11-10 10:53:48
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nelsonkti/echo-framework/util/xencoding"
	"github.com/nelsonkti/echo-framework/util/xfile"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func unmarshal(info *xfile.FileInfo, target *map[string]interface{}) error {
	if codec := xencoding.GetCodec(info.Format); codec != nil {
		return codec.Unmarshal(info.Data, &target)
	}

	return errors.New(fmt.Sprintf("unsupported key: %s format: %s", info.Name, info.Format))
}

func marshalJSON(v interface{}) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	}
	return json.Marshal(v)
}

func unmarshalJSON(data []byte, v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, m)
	}

	return json.Unmarshal(data, v)
}
