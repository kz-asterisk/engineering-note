package ikimono

import (
	"testing"
)

func Test_NewIkimono(t *testing.T) {
	type args struct {
		name    string
		atack   int
		defence int
		speed   int
	}
	tests := []struct {
		name   string
		args   args
		expect *Ikimono
	}{
		{
			name:   "normal",
			args:   args{name: "hayabusa", atack: 10, defence: 10, speed: 10},
			expect: &Ikimono{Name: "hayabusa", Atack: 10, Defence: 10, Speed: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIkimono(tt.args.name, tt.args.atack, tt.args.defence, tt.args.speed); got.Name != tt.expect.Name &&
				got.Atack != tt.expect.Atack && got.Defence != tt.expect.Defence && got.Speed != tt.expect.Speed {
				t.Errorf("ShowIkimonoName() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
func Test_ShowIkimonoName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name:   "normal",
			args:   args{name: "hayabusa"},
			expect: "hayabusa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIkimonoOnlyName(tt.args.name).ShowIkimonoName(); got != tt.expect {
				t.Errorf("ShowIkimonoName() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func Test_CalcDamage(t *testing.T) {
	type args struct {
		ikimonoA *Ikimono
		ikimonoB *Ikimono
	}
	tests := []struct {
		name   string
		args   args
		expect int
	}{
		{
			name: "normal",
			args: args{ikimonoA: &Ikimono{Name: "kaba", Atack: 10, Defence: 5, Speed: 3},
				ikimonoB: &Ikimono{Name: "Panda", Atack: 9, Defence: 5, Speed: 5},
			},
			expect: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ikimonoA := NewIkimono(tt.args.ikimonoA.Name, tt.args.ikimonoA.Atack, tt.args.ikimonoA.Defence, tt.args.ikimonoA.Speed)
			ikimonoB := NewIkimono(tt.args.ikimonoB.Name, tt.args.ikimonoB.Atack, tt.args.ikimonoB.Defence, tt.args.ikimonoB.Speed)
			if got := CalcDamage(ikimonoA, ikimonoB); got != tt.expect {
				t.Errorf("ShowIkimonoName() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
