package messagebus

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessageBus(t *testing.T) {
	mb := New()

	topic := "test_topic"
	message := "test_message"

	t.Run("Subscribe and Publish", func(t *testing.T) {
		ch := mb.Subscribe(topic)
		require.NotNil(t, ch)

		mb.Publish(topic, message)

		received := <-ch
		assert.Equal(t, message, received)
	})

	t.Run("Unsubscribe", func(t *testing.T) {
		ch := mb.Subscribe(topic)
		require.NotNil(t, ch)

		mb.Unsubscribe(topic, ch)

		mb.Publish(topic, message)

		select {
		case received := <-ch:
			t.Fatalf("Expected no message received, but received %v", received)
		default:
			// no message received, test passed
		}
	})
}
