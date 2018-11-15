package worker

import (
	"context"
)

type Worker interface {
	Try(context.Context) error
	Run(context.Context) (*Result, error)
}
