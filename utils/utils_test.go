package utils

import "testing"

func TestPaginate(t *testing.T) {
    type args struct {
        page        int
        pageSize    int
        recordCount int
    }
    tests := []struct {
        name        string
        args        args
        wantLow     int
        wantHigh    int
        wantPageNum int
    }{
        {
            name: "TestPaginate1",
            args: args{1,5,7},
            wantLow: 0,
            wantHigh: 5,
            wantPageNum: 1,
        },
        {
            name: "TestPaginate2",
            args: args{0,5,7},
            wantLow: 0,
            wantHigh: 5,
            wantPageNum: 1,
        },
        {
            name: "TestPaginate3",
            args: args{5,10,7},
            wantLow: 0,
            wantHigh: 7,
            wantPageNum: 1,
        },
        {
            name: "TestPaginate4",
            args: args{5,5,7},
            wantLow: 5,
            wantHigh: 7,
            wantPageNum: 2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotLow, gotHigh, gotPageNum, _ := Paginate(tt.args.page, tt.args.pageSize, tt.args.recordCount)
            if gotLow != tt.wantLow {
                t.Errorf("Paginate() gotLow = %v, want %v", gotLow, tt.wantLow)
            }
            if gotHigh != tt.wantHigh {
                t.Errorf("Paginate() gotHigh = %v, want %v", gotHigh, tt.wantHigh)
            }
            if gotPageNum != tt.wantPageNum {
                t.Errorf("Paginate() gotPageNum = %v, want %v", gotPageNum, tt.wantPageNum)
            }
        })
    }
}