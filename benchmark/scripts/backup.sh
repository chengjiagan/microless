# kubectl -n monitoring port-forward svc/prometheus-stack-kube-prom-prometheus 9090:9090&
# pf=$!
# sleep 10
# prometheus="localhost:9090"

# fetch_prom="../fetch_prometheus/fetch_prometheus.py"
# for n in ${thread[@]}; do
#     input="./${timestamp}/load_t${n}.csv"
#     output="./${timestamp}/prometheus_t${n}.xlsx"
#     $fetch_prom -a $prometheus -i $input -o $output -n social-network-kn
# done
# kill $pf

# kubectl -n jaeger port-forward svc/jaeger 16685:16685&
# pf=$!
# sleep 10
# jaeger="localhost:16685"

# export PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python
# fetch_jaeger="../fetch_jaeger/fetch_jaeger.py"
# for n in ${thread[@]}; do
#     input="./${timestamp}/load_t${n}.csv"
#     output="./${timestamp}/jaeger_t${n}.pb"
#     $fetch_jaeger -a $jaeger -i $input -o $output
# done
# kill $pf