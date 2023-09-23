package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/omalloc/kratos-agent/api/agent"
	"github.com/omalloc/kratos-console/internal/biz"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type TaskServer struct {
	log     *log.Helper
	usecase *biz.AppRuntimeUsecase
	cli     agent.AgentClient
	ch      chan struct{}
	tracer  *tracing.Tracer
}

func NewTaskServer(logger log.Logger, cli agent.AgentClient, usecase *biz.AppRuntimeUsecase) *TaskServer {
	r := &TaskServer{
		log:     log.NewHelper(logger),
		usecase: usecase,
		ch:      make(chan struct{}, 1),
		cli:     cli,
		tracer:  tracing.NewTracer(trace.SpanKindServer),
	}
	return r
}

func (s *TaskServer) Start(ctx context.Context) error {
	s.log.WithContext(ctx).Infof("[TASK] server starting")
	go s.run()

	return nil
}

func (s *TaskServer) Stop(ctx context.Context) error {
	s.log.WithContext(ctx).Infof("[TASK] server stopping")
	return nil
}

func (s *TaskServer) run() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			if err := s.flushed(); err != nil {
				continue
			}
		case <-s.ch:
			return
		}
	}
}

func (s *TaskServer) flushed() error {
	ctx, span := s.tracer.Start(context.Background(), "/cron/sync.agent.discovery/Services", propagation.MapCarrier{})
	defer func() {
		s.tracer.End(ctx, span, nil, nil)
	}()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	reply, err := s.cli.ListService(ctx, &agent.ListServiceRequest{})
	if err != nil {
		return err
	}
	if err = s.usecase.Updates(ctx, reply.Data); err != nil {
		s.log.Infof("[TASK] fetch discovery services failed: %v", err)
		return err
	}
	return nil
}
