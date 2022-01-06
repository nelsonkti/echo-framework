/**
** @创建时间 : 2022/1/4 17:31
** @作者 : fzy
 */
package yaml

import (
	"echo-framework/util/xencoding"
	"gopkg.in/yaml.v3"
)

const Name = "yaml"

func init() {
	xencoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
