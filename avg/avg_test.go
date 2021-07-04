package main

import "testing"

func TestAvg(t *testing.T){
	for _ , val := range []struct {
		nos [] int
		result int
	}{
		{[]int{1,2,3,4},2},
		{[]int{1,3,5},3},
		{[]int{1,2,3,10},4},
		{[]int{1},1},
		{[]int{},0},
	}{
		if Avg(val.nos...) == val.result {
			t.Logf("Success!! Value returned by Avg function %v is equal to %v",Avg(val.nos...) , val.result)
		}else{
			t.Errorf("fail!! Value returned by Avg function %v is not equal to %v",Avg(val.nos...) , val.result)
		}
	}
}