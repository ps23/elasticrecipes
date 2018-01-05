package elasticrecipes

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		elasticserver string
		attempts int
		result bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"test0", args{"http://localhost:9200", 3, true}},
		{"test1", args{"http://localhost:1234", 3, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckStatus(tt.args.elasticserver, tt.args.attempts)
			if result != tt.args.result {
          t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", tt.args.elasticserver, tt.args.attempts, result, tt.args.result)
      }
		})
	}
}
