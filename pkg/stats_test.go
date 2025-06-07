package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStats_Observe(t *testing.T) {
	t.Run("observe a few values and verify frequencies are tracked correctly", func(t *testing.T) {
		stats := NewStats()
		stats.Observe("google.com")
		stats.Observe("google.com")
		stats.Observe("google.com")

		stats.Observe("amazon.com")
		stats.Observe("amazon.com")

		assert.Equal(t, stats.valueFrequency["google.com"], 3)
		assert.Equal(t, stats.valueFrequency["amazon.com"], 2)
	})
}

func TestStats_Top(t *testing.T) {
	t.Run("observe a few values and verify frequencies are tracked correctly", func(t *testing.T) {
		stats := NewStats()
		stats.Observe("google.com")
		stats.Observe("google.com")
		stats.Observe("google.com")
		stats.Observe("google.com")

		stats.Observe("amazon.com")
		stats.Observe("amazon.com")
		stats.Observe("amazon.com")

		stats.Observe("facebook.com")

		stats.Observe("gmail.com")

		stats.Observe("netflix.com")
		stats.Observe("netflix.com")

		assert.Equal(t, stats.Top(3), []string{
			"google.com: 4",
			"amazon.com: 3",
			"netflix.com: 2",
		})
	})
}
