package main

import (
	"context"
	sharedUtils "shared/utils"
)

func main( ) {

	ctx, cancel := context.WithCancel(context.Background( ))
	defer cancel( )

	sharedUtils.CreateDbConnection(ctx)

	sharedUtils.CreateGRPCServer( )
}