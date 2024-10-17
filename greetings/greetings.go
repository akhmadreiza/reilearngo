package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	//message := fmt.Sprint(randomGreeting()) //to test error
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages_map := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages_map[name] = message
	}
	
	return messages_map, nil
}

func DayGreeting(name string) string {
	message := fmt.Sprintf("Good %v, %v. Welcome.", time_wording(), name)
	return message;
}

func time_wording() string {
	hr := time.Now().Hour()
	if hr < 12 {
		return "Morning"
	} else if hr < 18 {
		return "Afternoon"
	} else {
		return "Night"
	}
}

func randomGreeting() string {
	greets := []string {
		"Hello, %v. Welcome.",
		"Nice to see you, %v. Wish you had a good day.",
		"Hi, %v. You're amazing!",
	}
	return greets[rand.Intn(len(greets))]
}