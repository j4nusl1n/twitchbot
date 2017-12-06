package twitchbot

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		userName string
		token    string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		// TODO: Add test cases.
		{
			name: "UserInitBasic",
			args: args{
				userName: "hello",
				token:    "world",
			},
			want: &User{
				name:  "hello",
				token: "world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.userName, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	type args struct {
		u   User
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "TestConn",
			args: args{
				u: User{
					name:  os.Getenv("NICK"),
					token: os.Getenv("TOKEN"),
				},
				url: "irc.chat.twitch.tv:6667",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.u)
			Connect(tt.args.u, tt.args.url)
		})
	}
}
