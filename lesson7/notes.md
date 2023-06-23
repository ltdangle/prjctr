1. For unbuffered channels there always has to be a goroutine that reads the channel values invoked prior to channel write.
2. Send operation in its own goroutine, so it doesn't block the main goroutine.
