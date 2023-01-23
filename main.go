package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hoophq/pluginhooks"
)

type sqlManager struct {
	logger hclog.Logger
}

func (s *sqlManager) OnSessionOpen(p *pluginhooks.SesssionParams, resp *pluginhooks.SessionParamsResponse) error {
	s.logger.Info("opening session",
		"session", p.SessionID,
		"plugin-envvars", p.PluginEnvVars,
		"connection-name", p.ConnectionName,
		"connection-config", p.ConnectionConfig,
		"connection-type", p.ConnectionType,
		"connection-envs", p.ConnectionEnvVars,
		"connection-cmd", p.ConnectionCommand,
		"client-args", p.ClientArgs,
		"verb", p.ClientVerb,
		"user-id", p.UserID,
	)
	return nil
}

func (s *sqlManager) OnReceive(req *pluginhooks.Request, resp *pluginhooks.Response) error {
	s.logger.Info("on-receive",
		"session", req.SessionID,
		"packet-type", req.PacketType)
	return nil
}

func (s *sqlManager) OnSend(req *pluginhooks.Request, resp *pluginhooks.Response) error {
	s.logger.Info("on-send",
		"session", req.SessionID,
		"packet-type", req.PacketType)
	return nil
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:  hclog.Info,
		Output: os.Stderr,
	})
	logger.Info("starting demo plugin")
	pluginhooks.Serve(&sqlManager{logger: logger})
}
