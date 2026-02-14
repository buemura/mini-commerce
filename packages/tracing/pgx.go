package tracing

import (
	"context"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
)

// QueryOnlyTracer wraps otelpgx.Tracer but only exposes query-related
// interfaces. By not implementing pgx.ConnectTracer and pgxpool.AcquireTracer,
// pgx's type assertions will skip pool acquire and connect spans.
type QueryOnlyTracer struct {
	tracer *otelpgx.Tracer
}

func NewQueryOnlyTracer(opts ...otelpgx.Option) *QueryOnlyTracer {
	return &QueryOnlyTracer{tracer: otelpgx.NewTracer(opts...)}
}

// QueryTracer interface
func (t *QueryOnlyTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return t.tracer.TraceQueryStart(ctx, conn, data)
}

func (t *QueryOnlyTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	t.tracer.TraceQueryEnd(ctx, conn, data)
}

// BatchTracer interface
func (t *QueryOnlyTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	return t.tracer.TraceBatchStart(ctx, conn, data)
}

func (t *QueryOnlyTracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	t.tracer.TraceBatchQuery(ctx, conn, data)
}

func (t *QueryOnlyTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	t.tracer.TraceBatchEnd(ctx, conn, data)
}

// CopyFromTracer interface
func (t *QueryOnlyTracer) TraceCopyFromStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromStartData) context.Context {
	return t.tracer.TraceCopyFromStart(ctx, conn, data)
}

func (t *QueryOnlyTracer) TraceCopyFromEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromEndData) {
	t.tracer.TraceCopyFromEnd(ctx, conn, data)
}

// PrepareTracer interface
func (t *QueryOnlyTracer) TracePrepareStart(ctx context.Context, conn *pgx.Conn, data pgx.TracePrepareStartData) context.Context {
	return t.tracer.TracePrepareStart(ctx, conn, data)
}

func (t *QueryOnlyTracer) TracePrepareEnd(ctx context.Context, conn *pgx.Conn, data pgx.TracePrepareEndData) {
	t.tracer.TracePrepareEnd(ctx, conn, data)
}
