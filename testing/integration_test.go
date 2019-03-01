//// +build integration
//
//package testing
//
//import "testing"
//
//func TestCLIArgs(t *testing.T) {
//	tests := []struct {
//		cmd  string
//		args []string
//	}{
//		{"put", []string{"name", "Harley"}, "accepts put args"},
//		{"put", []string{"book", "read"}, "accepts put args"},
//		{"put", []string{"book", ""}, "error with empty string"},
//	}
//	//need to look at this example more tomorrow
//	for _, tt:=range tests{
//	t.Run(tt.cmd, func(t *testing.T){
//		result, err := os.
//	}
//}