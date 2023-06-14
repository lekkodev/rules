package parser

var functionVisitors = map[string]func(*ASTBuilderV3, *CallExpContext) interface{}{
	// Add new function name -> visitor functions here
	// e.g. "bucket": (*ASTBuilderV3).visitBucket
}
