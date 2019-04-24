package simplifypath

import (
	"container/list"
	"strings"
)

func simplifyPath(path string) string {
	var newPath string

	queue := split(path)
	if queue.Len() == 0 {
		newPath += "/"
	} else {
		// Combine the strings in the queue
		for queue.Len() != 0 {
			token := queue.Front().Value.(string)
			newPath += "/" + token
			queue.Remove(queue.Front())
		}
	}

	return newPath
}

func split(path string) *list.List {
	// Split path by "/"
	tokens := strings.Split(path, "/")

	// Create a queue to store tokens
	queue := list.New()
	for _, token := range tokens {
		if len(token) > 0 {
			// Jump "."
			if strings.Compare(token, ".") == 0 {
				continue
			}

			// Process ".." and remove the previous token
			if strings.Compare(token, "..") == 0 {
				if queue.Len() != 0 {
					queue.Remove(queue.Back())
				}
			} else { // Push token into queue
				queue.PushBack(token)
			}
		}
	}
	return queue
}
