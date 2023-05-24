package loadbalancer

const (
	OverloadHeaderKey = "microless-server-overload"
	UpdateInterval    = 100 // in milliseconds

	// for statistics
	NameRequestTotal        = "grpc_server_request_total"
	HelpRequestTotal        = "Total number of RPCs handled on the server."
	NameRequestLatency      = "grpc_server_request_latency_seconds"
	HelpRequestLatency      = "RPC latency distribution in seconds."
	NameServerlessTaskCount = "serverless_task_count"
	HelpServerlessTaskCount = "Number of tasks in queue and being processed."
)
