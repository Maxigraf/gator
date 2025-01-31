package command

func registerHandlers(commands Commands) {
	commands.Register("login", handlerLogin)
	commands.Register("register", handlerRegister)
	commands.Register("reset", handlerReset)
	commands.Register("users", handlerUsers)
	commands.Register("agg", handlerAggregate)
	commands.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.Register("feeds", handlerFeeds)
	commands.Register("follow", middlewareLoggedIn(handlerFollow))
	commands.Register("following", middlewareLoggedIn(handlerFollowing))
}
