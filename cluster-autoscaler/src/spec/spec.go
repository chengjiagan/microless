package spec

const (
	Cpu = "cpu" // in vCPU
	Mem = "mem" // in GB
)

type Spec map[string]float64

func NewSpec(cpu, mem float64) Spec {
	s := make(Spec)
	s[Cpu] = cpu
	s[Mem] = mem
	return s
}

// return true if s1 is less than s2 in both cpu and mem
func (s1 Spec) LessThan(s2 Spec) bool {
	return s1[Cpu] < s2[Cpu] && s1[Mem] < s2[Mem]
}
