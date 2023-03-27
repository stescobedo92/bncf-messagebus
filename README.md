# bncf-messagebus

bncf-messagebus is a Go library that provides an easy and efficient way to send and receive messages through channels. It uses the standards defined by Google to create a robust and user-friendly implementation."

## Installation

To use this library, you can download it using the go get command:

```golang
go get github.com/username/messagebus
```

### Use

To use the bncf-messagebus library in your project, you must first import it in your code

```golang
import "github.com/username/messagebus"
```

### Create a Message Bus

To create an instance of Message Bus, you can use the New() function:

```golang
mb := messagebus.New()
```

To subscribe to a topic, you can use the Subscribe method:

```golang
ch := mb.Subscribe("topic_name")
```

To post a message in a topic, use the Publish method:

```golang
mb.Publish("topic_name", message)
```
Where "message" is the message that you want to send. The message can be any type of data.

To unsubscribe from a topic, use the Unsubscribe method.

```golang
mb.Unsubscribe("topic_name", ch)
```
Where "ch" is the channel that you want to remove from the list of channels associated with the topic.

## Examples

Example 1: Subscribe and Publish

```golang
mb := messagebus.New()

ch := mb.Subscribe("topic_name")

mb.Publish("topic_name", "Hello, world!")

received := <-ch

fmt.Println(received)
// Output: Hello, world!
```

Example 2: Subscribe and Publish Multiple Messages

```golang
mb := messagebus.New()

ch := mb.Subscribe("topic_name")

mb.Publish("topic_name", "Hello, world!")
mb.Publish("topic_name", "Goodbye, world!")

received1 := <-ch
received2 := <-ch

fmt.Println(received1, received2)
// Output: Hello, world! Goodbye, world!
```

Example 3: Subscribe to Multiple Topics

```golang
mb := messagebus.New()

ch1 := mb.Subscribe("topic1")
ch2 := mb.Subscribe("topic2")

mb.Publish("topic1", "Hello, world!")
mb.Publish("topic2", "Goodbye, world!")

received1 := <-ch1
received2 := <-ch2

fmt.Println(received1, received2)
// Output: Hello, world! Goodbye, world!
```

Example 4: Unsubscribe from a Topic

```golang
mb := messagebus.New()

ch1 := mb.Subscribe("topic1")
ch2 := mb.Subscribe("topic2")

mb.Publish("topic1", "Hello, world!")
mb.Publish("topic2", "Goodbye, world!")

mb.Unsubscribe("topic1", ch1)

received1 := <-ch1 // no messages received
received2 := <-ch2

fmt.Println(received1, received2)
// Output:
```