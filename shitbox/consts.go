package shitbox

import (
	"context"

	"github.com/ilotterytea/shitbox/db"
)

var (
	DBClient = db.NewClient()
	DBCtx    = context.Background()
)
