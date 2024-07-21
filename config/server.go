package config

import (
	"time"
)

type ServerConfig struct {
	// common
	MuxTransport                    bool                            `yaml:"mux_transport"`                       // For WithMuxTransport
	ReadWriteTimeout                *TimeInterval                   `yaml:"read_write_timeout"`                  // For WithReadWriteTimeout
	ExitWaitTime                    *TimeInterval                   `yaml:"exit_wait_time"`                      // For WithExitWaitTime
	MaxConnIdleTime                 *TimeInterval                   `yaml:"max_conn_idle_time"`                  // For WithMaxConnIdleTime
	StatsLevel                      int                             `yaml:"stats_level"`                         // For WithStatsLevel
	ServiceAddr                     *NetAddr                        `yaml:"service_addr"`                        // For WithServiceAddr
	GRPCWriteBufferSize             uint32                          `yaml:"grpc_write_buffer_size"`              // For WithGRPCWriteBufferSize
	GRPCReadBufferSize              uint32                          `yaml:"grpc_read_buffer_size"`               // For WithGRPCReadBufferSize
	GRPCInitialWindowSize           uint32                          `yaml:"grpc_initial_window_size"`            // For WithGRPCInitialWindowSize
	GRPCInitialConnWindowSize       uint32                          `yaml:"grpc_initial_conn_window_size"`       // For WithGRPCInitialConnWindowSize
	GRPCKeepaliveParams             *GRPCKeepaliveParams            `yaml:"grpc_keepalive_params"`               // For WithGRPCKeepaliveParams
	GRPCKeepaliveEnforcementPolicy  *GRPCKeepaliveEnforcementPolicy `yaml:"grpc_keepalive_enforcement_policy"`   // For WithGRPCKeepaliveEnforcementPolicy
	GRPCMaxConcurrentStreams        uint32                          `yaml:"grpc_max_concurrent_streams"`         // For WithGRPCMaxConcurrentStreams
	GRPCMaxHeaderListSize           uint32                          `yaml:"grpc_max_header_list_size"`           // For WithGRPCMaxHeaderListSize
	ContextBackup                   *ContextBackup                  `yaml:"context_backup"`                      // For WithContextBackup
	RefuseTrafficWithoutServiceName bool                            `yaml:"refuse_traffic_without_service_name"` // For WithRefuseTrafficWithoutServiceName
	EnableContextTimeout            bool                            `yaml:"enable_context_timeout"`              // For WithEnableContextTimeout
	// advanced
	ServerBasicInfo *ServerBasicInfo `yaml:"server_basic_info"` // For WithServerBasicInfo
	Proxy           *NetAddr         `yaml:"proxy"`             // For WithProxy
	ReusePort       bool             `yaml:"reuse_port"`        // For WithReusePort
	// stream
	CompatibleMiddlewareForUnary bool `yaml:"compatible_middleware_for_unary"` // For WithCompatibleMiddlewareForUnary
}

// NetAddr defines the network and address of the service.
// Network name of the network (for example, "tcp", "udp").
// Address string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80").
type NetAddr struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type GRPCKeepaliveParams struct {
	MaxConnectionIdle     TimeInterval `yaml:"max_connection_idle"`
	MaxConnectionAge      TimeInterval `yaml:"max_connection_age"`
	MaxConnectionAgeGrace TimeInterval `yaml:"max_connection_age_grace"`
	Time                  TimeInterval `yaml:"time"`
	Timeout               TimeInterval `yaml:"timeout"`
}

type GRPCKeepaliveEnforcementPolicy struct {
	MinTime             TimeInterval `yaml:"min_time"`
	PermitWithoutStream bool         `yaml:"permit_without_stream"`
}

type ContextBackup struct {
	Enable bool `yaml:"enable"`
	Async  bool `yaml:"async"`
}

// TimeInterval common time interval struct.
// Unit ns, us, ms, s, m, h.
// Value is the timeout value.
type TimeInterval struct {
	Unit  string `yaml:"unit"`
	Value int    `yaml:"value"`
}

func (t *TimeInterval) Transform() time.Duration {
	var unit time.Duration
	switch t.Unit {
	case "ns":
		unit = time.Nanosecond
	case "us":
		unit = time.Microsecond
	case "ms":
		unit = time.Millisecond
	case "s":
		unit = time.Second
	case "m":
		unit = time.Minute
	case "h":
		unit = time.Hour
	}
	return time.Duration(t.Value) * unit
}

type ServerBasicInfo struct {
	ServiceName string            `yaml:"service_name"`
	Method      string            `yaml:"method"`
	Tags        map[string]string `yaml:"tags"`
}
