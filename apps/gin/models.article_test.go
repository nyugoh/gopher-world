package main

import (
	"reflect"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	tests := []struct {
		name     string
		wantArts []Article
	}{
		{
			name: "t1",
			wantArts: []Article{
				{Id: "1", Title: "Article 1", Content: "Content 1",},
				{Id: "2", Title: "Article 2", Content: "Content 2",},
				{Id: "3", Title: "Article 3", Content: "Content 3",},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArts := GetAllArticles(); !reflect.DeepEqual(gotArts, tt.wantArts) {
				t.Errorf("GetAllArticles() = %v, want %v", gotArts, tt.wantArts)
			}
		})
	}
}
