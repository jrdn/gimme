package source

//func TestParseSource(t *testing.T) {
//	// make a test file
//	tempFile, err := os.CreateTemp("", "test*")
//	require.NoError(t, err)
//
//	tests := []struct {
//		input   string
//		want    Source
//		wantErr bool
//	}{
//		{
//			input: "org/repo@branch",
//			want: &gitSource{
//				pullURL: "https://github.com/org/repo.git",
//				version: "branch",
//			},
//			wantErr: false,
//		},
//		{
//			input: "https://github.com/org/repo.git",
//			want: &gitSource{
//				pullURL: "https://github.com/org/repo.git",
//			},
//			wantErr: false,
//		},
//		{
//			input: "https://github.com/org/repo.git@branch",
//			want: &gitSource{
//				pullURL: "https://github.com/org/repo.git",
//				version: "branch",
//			},
//			wantErr: false,
//		},
//		{
//			input:   "/foo/bar/baz",
//			want:    &pathSource{},
//			wantErr: true,
//		},
//		{
//			input: tempFile.Name(),
//			want: &pathSource{
//				path: tempFile.Name(),
//			},
//			wantErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run("", func(t *testing.T) {
//			result, err := ParseSource(test.input)
//
//			if test.wantErr {
//				assert.Error(t, err)
//				return
//			}
//
//			assert.NoError(t, err)
//			assert.IsType(t, test.want, result)
//			assert.Equal(t, test.want, result)
//		})
//	}
//}
