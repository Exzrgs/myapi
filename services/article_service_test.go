package services_test

import "testing"

func BenchmarkGetArticleService(b *testing.B) {
	id := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := s.GetArticleService(id)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
