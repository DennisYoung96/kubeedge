package constants

// Service level constants
const (
	ResourceNode = "node"

	// GroupResource Group
	GroupResource = "resource"

	// NvidiaGPUStatusAnnotationKey Nvidia Constants
	// NvidiaGPUStatusAnnotationKey is the key of the node annotation for GPU status
	NvidiaGPUStatusAnnotationKey = "huawei.com/gpu-status"
	// NvidiaGPUScalarResourceName is the device plugin resource name used for special handling
	NvidiaGPUScalarResourceName = "nvidia.com/gpu"
	vcoreResourceName = "h3c.com/vcuda-core"
	vmemoryResourceName = "h3c.com/vcuda-memory"
)
