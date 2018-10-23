package main

import (
	"reflect"
	"testing"
)

func TestPlugin_versionCommand(t *testing.T) {
	type fields struct {
		Config Config
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "show version",
			fields: fields{},
			want:   []string{"up", "version"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{
				Config: tt.fields.Config,
			}
			if got := p.versionCommand().Args; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plugin.versionCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
