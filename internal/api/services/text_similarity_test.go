package services

import (
	"reflect"
	"testing"

	"github.com/jason-adam/text-similarity/internal/api/models"
)

func TestSimilarityService_Calculate(t *testing.T) {
	type args struct {
		t *models.TextPair
	}
	tests := []struct {
		name    string
		s       *SimilarityService
		args    args
		want    *models.SimilarityScore
		wantErr bool
	}{
		{
			name: "hello world",
			s:    &SimilarityService{},
			args: args{
				t: &models.TextPair{
					FirstText:  "hello world",
					SecondText: "hi world",
				},
			},
			want: &models.SimilarityScore{
				Score: 0.6363636,
			},
			wantErr: false,
		},
		{
			name: "sample 1 vs sample 2",
			s:    &SimilarityService{},
			args: args{
				t: &models.TextPair{
					FirstText:  "The easiest way to earn points with Fetch Rewards is to just shop for the products you already love. If you have any participating brands on your receipt, you'll get points based on the cost of the products. You don't need to clip any coupons or scan individual barcodes. Just scan each grocery receipt after you shop and we'll find the savings for you.",
					SecondText: "The easiest way to earn points with Fetch Rewards is to just shop for the items you already buy. If you have any eligible brands on your receipt, you will get points based on the total cost of the products. You do not need to cut out any coupons or scan individual UPCs. Just scan your receipt after you check out and we will find the savings for you.",
				},
			},
			want: &models.SimilarityScore{
				Score: 0.8130312,
			},
			wantErr: false,
		},
		{
			name: "sample 1 vs sample 3",
			s:    &SimilarityService{},
			args: args{
				t: &models.TextPair{
					FirstText:  "The easiest way to earn points with Fetch Rewards is to just shop for the products you already love. If you have any participating brands on your receipt, you'll get points based on the cost of the products. You don't need to clip any coupons or scan individual barcodes. Just scan each grocery receipt after you shop and we'll find the savings for you.",
					SecondText: "We are always looking for opportunities for you to earn more points, which is why we also give you a selection of Special Offers. These Special Offers are opportunities to earn bonus points on top of the regular points you earn every time you purchase a participating brand. No need to pre-select these offers, we'll give you the points whether or not you knew about the offer. We just think it is easier that way.",
				},
			},
			want: &models.SimilarityScore{
				Score: 0.2801932,
			},
			wantErr: false,
		},
		{
			name: "sample 2 vs sample 3",
			s:    &SimilarityService{},
			args: args{
				t: &models.TextPair{
					FirstText:  "The easiest way to earn points with Fetch Rewards is to just shop for the items you already buy. If you have any eligible brands on your receipt, you will get points based on the total cost of the products. You do not need to cut out any coupons or scan individual UPCs. Just scan your receipt after you check out and we will find the savings for you.",
					SecondText: "We are always looking for opportunities for you to earn more points, which is why we also give you a selection of Special Offers. These Special Offers are opportunities to earn bonus points on top of the regular points you earn every time you purchase a participating brand. No need to pre-select these offers, we'll give you the points whether or not you knew about the offer. We just think it is easier that way.",
				},
			},
			want: &models.SimilarityScore{
				Score: 0.29710144,
			},
			wantErr: false,
		},
		{
			name: "empty string error",
			s:    &SimilarityService{},
			args: args{
				t: &models.TextPair{
					FirstText:  "",
					SecondText: "",
				},
			},
			want: &models.SimilarityScore{
				Score: 0.0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimilarityService{}
			got, err := s.Calculate(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimilarityService.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimilarityService.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 value min test",
			args: args{
				args: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.args...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 value max test",
			args: args{
				args: []int{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.args...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalize(t *testing.T) {
	type args struct {
		s int
		t *models.TextPair
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "normalize test",
			args: args{
				s: 5,
				t: &models.TextPair{
					FirstText:  "hello world",
					SecondText: "hi world",
				},
			},
			want: 0.5454545,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
