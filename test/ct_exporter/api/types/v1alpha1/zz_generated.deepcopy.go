// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployModel) DeepCopyInto(out *DeployModel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployModel.
func (in *DeployModel) DeepCopy() *DeployModel {
	if in == nil {
		return nil
	}
	out := new(DeployModel)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploySpec) DeepCopyInto(out *DeploySpec) {
	*out = *in
	out.Model = in.Model
	in.Trigger.DeepCopyInto(&out.Trigger)
	in.HardExampleMining.DeepCopyInto(&out.HardExampleMining)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploySpec.
func (in *DeploySpec) DeepCopy() *DeploySpec {
	if in == nil {
		return nil
	}
	out := new(DeploySpec)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvalSpec) DeepCopyInto(out *EvalSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvalSpec.
func (in *EvalSpec) DeepCopy() *EvalSpec {
	if in == nil {
		return nil
	}
	out := new(EvalSpec)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardExampleMining) DeepCopyInto(out *HardExampleMining) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParaSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardExampleMining.
func (in *HardExampleMining) DeepCopy() *HardExampleMining {
	if in == nil {
		return nil
	}
	out := new(HardExampleMining)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILDataset) DeepCopyInto(out *ILDataset) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILDataset.
func (in *ILDataset) DeepCopy() *ILDataset {
	if in == nil {
		return nil
	}
	out := new(ILDataset)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILJobCondition) DeepCopyInto(out *ILJobCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILJobCondition.
func (in *ILJobCondition) DeepCopy() *ILJobCondition {
	if in == nil {
		return nil
	}
	out := new(ILJobCondition)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILJobSpec) DeepCopyInto(out *ILJobSpec) {
	*out = *in
	out.Dataset = in.Dataset
	out.InitialModel = in.InitialModel
	in.TrainSpec.DeepCopyInto(&out.TrainSpec)
	in.EvalSpec.DeepCopyInto(&out.EvalSpec)
	in.DeploySpec.DeepCopyInto(&out.DeploySpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILJobSpec.
func (in *ILJobSpec) DeepCopy() *ILJobSpec {
	if in == nil {
		return nil
	}
	out := new(ILJobSpec)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILJobStatus) DeepCopyInto(out *ILJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ILJobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILJobStatus.
func (in *ILJobStatus) DeepCopy() *ILJobStatus {
	if in == nil {
		return nil
	}
	out := new(ILJobStatus)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncrementalLearningJob) DeepCopyInto(out *IncrementalLearningJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncrementalLearningJob.
func (in *IncrementalLearningJob) DeepCopy() *IncrementalLearningJob {
	if in == nil {
		return nil
	}
	out := new(IncrementalLearningJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IncrementalLearningJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncrementalLearningJobList) DeepCopyInto(out *IncrementalLearningJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IncrementalLearningJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncrementalLearningJobList.
func (in *IncrementalLearningJobList) DeepCopy() *IncrementalLearningJobList {
	if in == nil {
		return nil
	}
	out := new(IncrementalLearningJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IncrementalLearningJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialModel) DeepCopyInto(out *InitialModel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialModel.
func (in *InitialModel) DeepCopy() *InitialModel {
	if in == nil {
		return nil
	}
	out := new(InitialModel)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParaSpec) DeepCopyInto(out *ParaSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParaSpec.
func (in *ParaSpec) DeepCopy() *ParaSpec {
	if in == nil {
		return nil
	}
	out := new(ParaSpec)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timer) DeepCopyInto(out *Timer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timer.
func (in *Timer) DeepCopy() *Timer {
	if in == nil {
		return nil
	}
	out := new(Timer)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainSpec) DeepCopyInto(out *TrainSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Trigger.DeepCopyInto(&out.Trigger)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainSpec.
func (in *TrainSpec) DeepCopy() *TrainSpec {
	if in == nil {
		return nil
	}
	out := new(TrainSpec)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Timer != nil {
		in, out := &in.Timer, &out.Timer
		*out = new(Timer)
		**out = **in
	}
	out.Condition = in.Condition
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dataset) DeepCopyInto(out *Dataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dataset.
func (in *Dataset) DeepCopy() *Dataset {
	if in == nil {
		return nil
	}
	out := new(Dataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dataset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetList) DeepCopyInto(out *DatasetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dataset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetList.
func (in *DatasetList) DeepCopy() *DatasetList {
	if in == nil {
		return nil
	}
	out := new(DatasetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetSpec) DeepCopyInto(out *DatasetSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetSpec.
func (in *DatasetSpec) DeepCopy() *DatasetSpec {
	if in == nil {
		return nil
	}
	out := new(DatasetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetStatus) DeepCopyInto(out *DatasetStatus) {
	*out = *in
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetStatus.
func (in *DatasetStatus) DeepCopy() *DatasetStatus {
	if in == nil {
		return nil
	}
	out := new(DatasetStatus)
	in.DeepCopyInto(out)
	return out
}
