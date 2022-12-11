package shitbox

import (
	"github.com/ilotterytea/shitbox/prisma/db"
)

var (
	DBClient = db.NewClient()
)
