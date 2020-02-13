package redis

import (
	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_sdk_Ping(t *testing.T) {
	assert := assert.New(t)

	s, err := miniredis.Run()
	assert.NoError(err)
	defer s.Close()

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "ping redis server",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		c := NewRedisClient().
			SetAddress(s.Addr()).
			SetDatabase(0).
			SetPassword("").
			SetTimeout(5, 5).
			SetPoolSize(3).
			Call()

		t.Run(tt.name, func(t *testing.T) {
			if err := c.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sdk_Set(t *testing.T) {
	assert := assert.New(t)

	s, err := miniredis.Run()
	assert.NoError(err)
	defer s.Close()

	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "set value to redis success",
			args: args{
				key:        "name",
				value:      "danny ferian",
				expiration: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedisClient().
				SetAddress(s.Addr()).
				SetDatabase(0).
				SetPassword("").
				SetTimeout(5, 5).
				SetPoolSize(3).
				Call()

			if err := c.Set(tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sdk_Get(t *testing.T) {
	assert := assert.New(t)

	s, err := miniredis.Run()
	assert.NoError(err)
	defer s.Close()

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get data success",
			args: args{
				key: "name",
			},
			want:    "danny",
			wantErr: false,
		},
		{
			name: "get data failed",
			args: args{
				key: "user",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedisClient().
				SetAddress(s.Addr()).
				SetDatabase(0).
				SetPassword("").
				SetTimeout(5, 5).
				SetPoolSize(3).
				Call()

			if tt.want != nil {
				s.Set(tt.args.key, tt.want.(string))
			}

			got, err := c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
