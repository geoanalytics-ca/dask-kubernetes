// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubernetesdaskorgv1 "github.com/JordanGunn/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DaskJobStatusApplyConfiguration represents an declarative configuration of the DaskJobStatus type for use
// with apply.
type DaskJobStatusApplyConfiguration struct {
	ClusterName      *string                        `json:"clusterName,omitempty"`
	EndTime          *v1.Time                       `json:"endTime,omitempty"`
	JobRunnerPodName *string                        `json:"jobRunnerPodName,omitempty"`
	JobStatus        *kubernetesdaskorgv1.JobStatus `json:"jobStatus,omitempty"`
	StartTime        *v1.Time                       `json:"startTime,omitempty"`
}

// DaskJobStatusApplyConfiguration constructs an declarative configuration of the DaskJobStatus type for use with
// apply.
func DaskJobStatus() *DaskJobStatusApplyConfiguration {
	return &DaskJobStatusApplyConfiguration{}
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *DaskJobStatusApplyConfiguration) WithClusterName(value string) *DaskJobStatusApplyConfiguration {
	b.ClusterName = &value
	return b
}

// WithEndTime sets the EndTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndTime field is set to the value of the last call.
func (b *DaskJobStatusApplyConfiguration) WithEndTime(value v1.Time) *DaskJobStatusApplyConfiguration {
	b.EndTime = &value
	return b
}

// WithJobRunnerPodName sets the JobRunnerPodName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JobRunnerPodName field is set to the value of the last call.
func (b *DaskJobStatusApplyConfiguration) WithJobRunnerPodName(value string) *DaskJobStatusApplyConfiguration {
	b.JobRunnerPodName = &value
	return b
}

// WithJobStatus sets the JobStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JobStatus field is set to the value of the last call.
func (b *DaskJobStatusApplyConfiguration) WithJobStatus(value kubernetesdaskorgv1.JobStatus) *DaskJobStatusApplyConfiguration {
	b.JobStatus = &value
	return b
}

// WithStartTime sets the StartTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTime field is set to the value of the last call.
func (b *DaskJobStatusApplyConfiguration) WithStartTime(value v1.Time) *DaskJobStatusApplyConfiguration {
	b.StartTime = &value
	return b
}
