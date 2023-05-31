package loadbalancer

const (
	// for server
	ServerHeaderKey = "upstream"
	RateHeaderKey   = "rate"

	// for stat
	NameRequestTotal          = "grpc_server_request_total"
	HelpRequestTotal          = "Total number of RPCs handled on the server."
	NameRequestLatency        = "grpc_server_request_latency_seconds"
	HelpRequestLatency        = "RPC latency distribution in seconds."
	NameServerlessTaskTotal   = "serverless_task_total"
	HelpServerlessTaskTotal   = "Number of tasks in queue and being processed."
	NameServerlessTaskRunning = "serverless_task_running"
	HelpServerlessTaskRunning = "Number of tasks being processed."
)
