package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Chaos defines the chaos monkey configuration
type ChaosMonkeySpec struct {
	Monkeys []Monkey `json:"monkeys,omitempty"`
}

// Chaos monkey configuration
type Monkey struct {
	Name          string           `json:"name,omitempty"`
	RateSeconds   *int64           `json:"rateSeconds,omitempty"`
	PeriodSeconds *int64           `json:"periodSeconds,omitempty"`
	Jitter        *float64         `json:"jitter,omitempty"`
	Selector      *MonkeySelector  `json:"selector,omitempty"`
	Crash         *CrashMonkey     `json:"crash,omitempty"`
	Partition     *PartitionMonkey `json:"partition,omitempty"`
	Stress        *StressMonkey    `json:"stress,omitempty"`
}

// Monkey selector
type MonkeySelector struct {
	*metav1.LabelSelector `json:",inline"`
	*GroupSelector        `json:",inline"`
	*PodSelector          `json:",inline"`
}

// Group selector
type GroupSelector struct {
	MatchGroups []string `json:"groups,omitempty"`
}

// Pod selector
type PodSelector struct {
	MatchPods []string `json:"pods,omitempty"`
}

// Chaos monkey that crashes nodes
type CrashMonkey struct {
	CrashStrategy CrashStrategy `json:"crashStrategy,omitempty"`
}

// Crash strategy
type CrashStrategy struct {
	Type CrashStrategyType `json:"type,omitempty"`
}

type CrashStrategyType string

const (
	CrashContainer CrashStrategyType = "Container"
	CrashPod       CrashStrategyType = "Pod"
)

// Chaos monkey that partitions nodes
type PartitionMonkey struct {
	PartitionStrategy PartitionStrategy `json:"partitionStrategy,omitempty"`
}

// Partition strategy.
type PartitionStrategy struct {
	Type PartitionStrategyType `json:"type,omitempty"`
}

type PartitionStrategyType string

const (
	PartitionIsolate PartitionStrategyType = "Isolate"
	PartitionBridge  PartitionStrategyType = "Bridge"
)

// Chaos monkey that stresses nodes
type StressMonkey struct {
	StressStrategy StressStrategy `json:"type,omitempty"`
	IO             *StressIO      `json:"io,omitempty"`
	CPU            *StressCPU     `json:"cpu,omitempty"`
	Memory         *StressMemory  `json:"memory,omitempty"`
	HDD            *StressHDD     `json:"hdd,omitempty"`
	Network        *StressNetwork `json:"network,omitempty"`
}

// Stress monkey strategy
type StressStrategy struct {
	Type StressStrategyType `json:"type,omitempty"`
}

type StressStrategyType string

const (
	StressRandom StressStrategyType = "Random"
	StressAll    StressStrategyType = "All"
)

type StressConfig struct {
	Workers *int `json:"workers,omitempty"`
}

// Configuration for stressing node I/O
type StressIO struct {
	StressConfig `json:",inline"`
}

// Configuration for stressing CPU
type StressCPU struct {
	StressConfig `json:",inline"`
}

// Configuration for stressing memory allocation
type StressMemory struct {
	StressConfig `json:",inline"`
}

// Configuration for stressing hard drive usage
type StressHDD struct {
	StressConfig `json:",inline"`
}

// Configuration for stressing network usage
type StressNetwork struct {
	LatencyMilliseconds *int64               `json:"latencyMilliseconds,omitempty"`
	Jitter              *float64             `json:"jitterMilliseconds,omitempty"`
	Correlation         *float64             `json:"correlation,omitempty"`
	Distribution        *LatencyDistribution `json:"distribution,omitempty"`
}

type LatencyDistribution string

const (
	LatencyDistributionNormal       LatencyDistribution = "normal"
	LatencyDistributionPareto       LatencyDistribution = "pareto"
	LatencyDistributionParetoNormal LatencyDistribution = "paretonormal"
)

// ChaosMonkeyStatus defines the observed state of ChaosMonkey
type ChaosMonkeyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosMonkey is the Schema for the chaosmonkeys API
// +k8s:openapi-gen=true
type ChaosMonkey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChaosMonkeySpec   `json:"spec,omitempty"`
	Status ChaosMonkeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosMonkeyList contains a list of ChaosMonkey
type ChaosMonkeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChaosMonkey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChaosMonkey{}, &ChaosMonkeyList{})
}
