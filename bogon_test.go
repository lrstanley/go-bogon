package bogon

import (
	"net"
	"reflect"
	"testing"
)

func TestDefaultRanges(t *testing.T) {
	got := DefaultRanges()

	if got == nil || len(got) < 1 {
		t.Fatal("DefaultRanges() returned invalid results")
	}
}

func TestBogon_Ranges(t *testing.T) {
	tests := []struct {
		name string
		b    *Bogon
		want []*net.IPNet
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Ranges(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bogon.Ranges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func dummyBogon() *Bogon {
	b, err := New([]string{"127.0.0.1/32", "10.0.0.0/8", "11.0.0.0/24"})
	if err != nil {
		panic(err)
	}

	return b
}

const dummyString = "127.0.0.1/32 10.0.0.0/8 11.0.0.0/24"

func TestBogon_String(t *testing.T) {
	b := dummyBogon()

	if b.String() != dummyString {
		t.Errorf("Bogon.String() = %v, want %v", b.String(), dummyString)
	}
}

func TestBogon_Is(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name               string
		b                  *Bogon
		args               args
		wantIsIn           bool
		wantRepresentation string
	}{
		{"is single", dummyBogon(), args{ip: "127.0.0.1"}, true, "127.0.0.1/32"},
		{"is past single", dummyBogon(), args{ip: "127.0.0.2"}, false, ""},
		{"not in", dummyBogon(), args{ip: "1.2.3.4"}, false, ""},
		{"in range", dummyBogon(), args{ip: "10.0.0.100"}, true, "10.0.0.0/8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsIn, gotRepresentation := tt.b.Is(tt.args.ip)
			if gotIsIn != tt.wantIsIn {
				t.Errorf("Bogon.Is() gotIsIn = %v, want %v", gotIsIn, tt.wantIsIn)
			}
			if gotRepresentation != tt.wantRepresentation {
				t.Errorf("Bogon.Is() gotRepresentation = %v, want %v", gotRepresentation, tt.wantRepresentation)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		cidrList []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Bogon
		wantErr bool
	}{
		{"valid", args{cidrList: []string{"10.0.0.0/8"}}, &Bogon{ipRange: []*net.IPNet{MustCIDR("10.0.0.0/8")}}, false},
		{"invalid", args{cidrList: []string{"10.0.0.0/1000"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.cidrList)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name               string
		args               args
		wantIsIn           bool
		wantRepresentation string
	}{
		{"is bogon", args{ip: "10.1.2.3"}, true, "10.0.0.0/8"},
		{"is bogon", args{ip: "11.1.2.3"}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsIn, gotRepresentation := Is(tt.args.ip)
			if gotIsIn != tt.wantIsIn {
				t.Errorf("Is() gotIsIn = %v, want %v", gotIsIn, tt.wantIsIn)
			}
			if gotRepresentation != tt.wantRepresentation {
				t.Errorf("Is() gotRepresentation = %v, want %v", gotRepresentation, tt.wantRepresentation)
			}
		})
	}
}
