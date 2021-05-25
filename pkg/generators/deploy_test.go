package generators

import (
	"os"
	"shiploader/apis/apps"
	"testing"
)


func TestGenerateDeployment(t *testing.T) {
	type args struct {
		app  apps.App
		dest string
		w    WriterInterface
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestHappyPath",
			wantErr: false,
			args: args{
				app:  apps.App{
					Name: "foo",
					Namespace: "default",
					Image: "foo:latest",
					Port: 8080,
				},
				dest: "/tmp",
				w:    &MockWriter{Err: nil},
			},
		},
		{
			name: "TestFailurePath",
			wantErr: true,
			args: args{
				app:  apps.App{},
				dest: "/tmp",
				w:    &MockWriter{Err: os.ErrNotExist},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateDeployment(tt.args.app, tt.args.dest, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("GenerateDeployment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
