package command

func registerHandlers(commands Commands) {
	commands.Register("login", handlerLogin)
	commands.Register("register", handlerRegister)
	commands.Register("reset", handlerReset)
	commands.Register("users", handlerUsers)
	commands.Register("agg", handlerAggregate)
	commands.Register("addfeed", handlerAddFeed)
	commands.Register("feeds", handlerFeeds)
	commands.Register("follow", handlerFollow)
	commands.Register("following", handlerFollowing)
}
