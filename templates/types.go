

/*  This file contains all structs holding the .yaml templates
    for the auto-generation. */


package autoyaml

type YamlKind struct {
    Deployment Deployment
    ServiceAccount ServiceAccount
    Role Role
    RoleBinding RoleBinding
    ClusterRole ClusterRole
    ClusterRoleBinding ClusterRoleBinding
    Pod Pod
    // TODO
}

// -------- KIND --------
type Deployment struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Metadata ObjectMeta `yaml:"metadata"`
    Spec DeploymentSpec `yaml:"spec"`
}

type ServiceAccount struct {
    APIVersion string `yaml:"apiVersion"`
    AutomatedServiceAccountToken bool `yaml:"automationServiceAccountToken"`
    ImagePullSecrets LocalObjectReference `yaml:"imagePullSecrets"`
    Kind string `yaml:"kind"`
    Metadata ObjectMeta `yaml:"metadata"`
    Secrets ObjectReference `yaml:"secrets"`
}

type Role struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Metadata ObjectMeta `yaml:"metadata"`
    Rules PolicyRule `yaml:"rules"`
}

type RoleBinding struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"string"`
    Metadata ObjectMeta `yaml:"metadata"`
    RoleRef RoleRef `yaml:"roleRef"`
    Subjects Subject `yaml:"subjects"`
}

type ClusterRole struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Metadata ObjectMeta `yaml:"metadata"`
    Rules PolicyRule `yaml:"rules"`
}

type ClusterRoleBinding struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Metadata ObjectMeta `yaml:"metadata"`
    RoleRef RoleRef `yaml:"roleRef"`
    Subjects Subject `yaml:"subjects"`
}

type Pod struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Metadata ObjectMeta `yaml:"metadata"`
    Spec PodSpec `yaml:"spec"`
}

// -------- First Degree Properties --------
type ObjectMeta struct {
    Name string `yaml:"name"`
    Namespace string `yaml:"namespace"`
    ClusterName string `yaml:"clusterName"`
    Labels MetadataLabels `yaml:"labels"`
}

type DeploymentSpec struct {
    Replicas int `yaml:"replicas"`
    Selector LabelSelector `yaml:"selector"`
    Template PodTemplateSpec `yaml:"template"`
}

type LocalObjectReference struct {
    Name string `yaml:"name"`
}

type ObjectReference struct {
    APIVersion string `yaml:"apiVersion"`
    FieldPath string `yaml:"fieldPath"`
    Kind string `yaml:"kind"`
    Name string `yaml:"name"`
    Namespace string `yaml:"namespace"`
    ResourceVersion string `yaml:"resourceVersion"`
    UID string `yaml:"uid"`
}

type PolicyRule struct {
    APIGroups []string `yaml:"apiGroups"`
    NonResourceURLs []string `yaml:"nonResourceURLs"`
    ResourceNames []string `yaml:"resourceNames"`
    Resources []string `yaml:"resources"`
    Verbs []string `yaml:"verbs"`
}

type RoleRef struct {
    APIGroup string `yaml:"apiGroup"`
    Kind string `yaml:"kind"`
    Name string `yaml:"name"`
}

type Subject struct {
    APIGroup string `yaml:"apiGroup"`
    Kind string `yaml:"kind"`
    Name string `yaml:"name"`
    Namespace string `yaml:"namespace"`
}

// -------- Second Degree Properties --------
type MetadataLabels struct {
    App string `yaml:"app"`
}

type LabelSelector struct {
    MatchLabels MatchLabels `yaml:"matchLabels"`
}

type PodTemplateSpec struct {
    Metadata ObjectMeta `yaml:"metadata"`
    Spec PodSpec `yaml:"spec"`
}

// -------- Third Degree Properties --------
type PodSpec struct {
    Containers Containers `yaml:"containers"`
    NodeName string `yaml:"nodeName"`
    RestartPolicy string `yaml:"restartPolicy"`
    SchedulerName string `yaml:"schedulerName"`
    ServiceAccount string `yaml:"serviceaccount"`
    Volumes Volume `yaml:"volumes"`
}

// -------- Forth Degree Properties --------
type Containers struct {
    Args []string `yaml:"args"`
    Command []string `yaml:"command"`
    Image string `yaml:"image"`
    ImagePullPolicy string `yaml:"imagePullPolicy"` // Always, Never, IfNotPresent
    Name string `yaml:"name"`
    LivenessProbe Probe `yaml:"livenessProbe"`
    Ports ContainerPort `yaml:"ports"`
    ReadinessProbe Probe `yaml:"readinessProbe"`
    Resources ResourceRequirements `yaml:"resources"`
    VolumeMounts VolumeMount `yaml:"volumeMounts"`
}

type Volume struct {
    Name string `yaml:"name"`
    PersistentVolumeClaim PersistentVolumeClaim `yaml:"persistentVolumeClaim"`
}

// -------- Fifth Degree Properties --------
type PersistentVolumeClaim struct {
    ClaimName string `yaml:"claimName"`
    ReadOnly bool `yaml:"readOnly"`
}

type Probe struct {
    Exec ExecAction `yaml:"exec"`
    failureThreshold int `yaml:"failureThreshold"` // Defaults to 3, min. 1
    HTTPGet HTTPGetAction `yaml:"httpGet"`
    InitialDelaySeconds int `yaml:"initialDelaySeconds"`
    periodSeconds int `yaml:"periodSeconds"` // Defaults to 10, min. 1
    successThreshold int `yaml:"successThreshold"` // Defaults to 1, min. 1 (Must be 1 for liveness and startup)
    TimeoutSeconds int `yaml:"timeoutSeconds"` // Defaults to 1sec, min. 1
}

type ContainerPort struct {
    ContainerPort int `yaml:"containerPort"`
    HostIP string `yaml:"hostIP"`
    HostPort int `yaml:"hostPort"`
    Name string `yaml:"name"`
    Protocol string `yaml:"protocol"`
}

type ResourceRequirements struct {
    Limits Limits `yaml:"limits"`
    Requests Requests `yaml:"requests"`
}

type VolumeMount struct {
    MountPath string `yaml:"mountPath"`
    Name string `yaml:"name"`
    ReadOnly bool `yaml:"readOnly"`
    SubPath string `yaml:"subPath"`
}

// -------- Sixth Degree Properties --------
type ExecAction struct {
    Command []string `yaml:"command"`
}

type HTTPGetAction struct {
    Host string `yaml:"host"`
    HTTPHeader HTTPHeader `yaml:"httpHeader"`
    Path string `yaml:"path"`
    Port int `yaml:"port"`
    Scheme string `yaml:"scheme"`
}

// -------- Seventh Degree Properties --------
type HTTPHeader struct {
    Name string `yaml:"name"`
    Value string `yaml:"value"`
}

// -------- Property Objects --------
type MatchLabels struct { // ?????
    Key string
    Value int
}

//      Limits and Requests take a byte value as plain int
//      E,P,T,G,M,k can also be used as suffixes
//      Ei,Pi,Ti,Gi,Mi,Ki can also be used as suffixes
type Limits struct {
    CPU string
    Memory string
}

type Requests struct {
    CPU string
    Memory string
}



