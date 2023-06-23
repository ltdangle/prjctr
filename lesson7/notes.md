1. For unbuffered channels there always has to be a goroutine that reads the channel values invoked prior to channel write.
In other words: It's okay to send to an unbuffered channel from the main goroutine, but you need to ensure that there is another goroutine ready to receive the data, or else it will cause a deadlock.
