package step

//
//import (
//	"encoding/json"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestUnmarshalSteps(t *testing.T) {
//	tests := []struct {
//		data     []byte
//		expected Steps
//		wantErr  bool
//	}{
//		{
//			data: []byte(`
//[
//  {"type": "task", "yes": true, "dir": "./foo/bar", "target": "build"},
//  {"type": "pwsh", "yes": true, "path": "./build.ps1"}
//]
//`),
//			expected: []Step{
//				&taskfileStep{
//					baseStep: baseStep{
//						Type: "task",
//					},
//					Target: "build",
//					Dir:    "./foo/bar",
//				},
//				&pwshStep{
//					baseStep: baseStep{
//						Type: "pwsh",
//					},
//					Path: "./build.ps1",
//				},
//			},
//			wantErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run("", func(t *testing.T) {
//			var steps Steps
//			err := json.Unmarshal(test.data, &steps)
//			if test.wantErr {
//				assert.Error(t, err)
//				return
//			}
//
//			assert.NoError(t, err)
//			assert.Equal(t, test.expected, steps)
//		})
//	}
//}
