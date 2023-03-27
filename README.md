# bncf-messagebus

bncf-messagebus is a Go library that provides an easy and efficient way to send and receive messages through channels. It uses the standards defined by Google to create a robust and user-friendly implementation."

## Installation

To use this library, you can download it using the go get command:

```
go get github.com/username/messagebus

```

### Use

To use the bncf-messagebus library in your project, you must first import it in your code

```
import "github.com/username/messagebus"
```

### Create a Message Bus

To create an instance of Message Bus, you can use the New() function:

```
mb := messagebus.New()
```

To subscribe to a topic, you can use the Subscribe method:

```
ch := mb.Subscribe("topic_name")
```

