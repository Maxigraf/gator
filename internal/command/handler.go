package command

func registerHandlers(commands Commands) {
	commands.Register("login", middlewareArgsRequired(handlerLogin))
	commands.Register("register", middlewareArgsRequired(handlerRegister))
	commands.Register("reset", handlerReset)
	commands.Register("users", handlerUsers)
	commands.Register("agg", handlerAggregate)
	commands.Register("addfeed", middlewareArgsCountRequired(middlewareLoggedIn(handlerAddFeed), 2))
	commands.Register("feeds", handlerFeeds)
	commands.Register("follow", middlewareArgsRequired(middlewareLoggedIn(handlerFollow)))
	commands.Register("following", middlewareLoggedIn(handlerFollowing))
	commands.Register("unfollow", middlewareArgsRequired(middlewareLoggedIn(handlerUnfollow)))
}
