## Gator
RSS feed aggregator

# Requirements
- postgresql
- go

# Installation
Maybe use "go install"?

# Configuration
~/.gatorconfig.json
```json
{
    "db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
}
```

# Usage
- login \<username\> : Login as previously registered user
- register \<username\> : Register new user
- reset : DELETE EVERYTHING!!!
- users : Get list of users
- agg : Aggregation loop (infinite loop for collecting posts)
- addfeed \<name\> \<url\> : Add feed with name to be collected from url
- feeds : List all feeds
- follow \<url\> : Subscribe user to feed
- following : List subscribed feeds for current user
- unfollow \<url\> : Unsubscribe from feed
- browse \[limit\] : List posts (max. limit; default 2)
