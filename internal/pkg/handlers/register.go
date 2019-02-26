package handlers

// Register tracks all the current handlers
// add new handlers to the array to augment the
// processing of items in the
var Register = []Expander{
	&ResourceGroupResourceExpander{},
	&SubscriptionExpander{},
	&ActionExpander{},
	&SwaggerResourceExpander{},
	&DeploymentsExpander{},
}
