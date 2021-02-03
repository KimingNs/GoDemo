package main

import "testing"

func TestMonster_Restore(t *testing.T) {
	type fields struct {
		Name  string
		Age   int
		Skill string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			"first", fields{Name: "TestOne", Age: 100, Skill: "FixSuccess"}, true,
		},
		{
			"second", fields{Name: "TestTwo", Age: 200, Skill: "FixBug"}, true,
		},
		{
			"third", fields{Name: "TestThree", Age: 300, Skill: "TestSuccess"}, true,
		},
		{
			"fourth", fields{Name: "TestFour", Age: 400, Skill: "TestFix"}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Monster{
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Skill: tt.fields.Skill,
			}
			if got := this.Store(); got != tt.want {
				t.Errorf("Store() = %v, want %v", got, tt.want)
			}
		})
	}
}
