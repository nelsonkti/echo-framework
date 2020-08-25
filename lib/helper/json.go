package helper

import (
    "encoding/json"
    "echo-framework/lib/logger"
)

func JsonMarshal(v interface{}) string {
    bytes, err := json.Marshal(v)
    if err != nil {
        logger.Sugar.Error("json序列化：", err)
    }
    return Bytes2str(bytes)
}

func StructToMap(in interface{}) map[string]interface{} {
    var inInterface map[string]interface{}
    inrec, _ := json.Marshal(in)
    if err := json.Unmarshal(inrec, &inInterface); err != nil {
        return nil
    }
    return inInterface
}
