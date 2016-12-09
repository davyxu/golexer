package golexer

import (
	"testing"
)

func TestKVParser(t *testing.T) {

	ParseKV(`protobuf:"varint,1,opt,name=AutoID,json=autoID" json:"AutoID,omitempty"`, func(key, value string) bool {

		t.Logf("'%s' = '%s'", key, value)

		return true
	})

}
