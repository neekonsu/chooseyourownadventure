package chooseyourownadventure

import (
	"io"
	"reflect"
	"testing"
)

func TestMakeStory(t *testing.T) {
	type args struct {
		r *io.Reader
	}
	tests := []struct {
		name string
		args args
		want Story
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeStory(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}
