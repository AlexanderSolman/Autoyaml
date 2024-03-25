package autoyaml

import "fmt"

func GenerateDeployment(flag *string) Deployment {
    // flag is always defaulted as empty string 
    fmt.Println("From generate:", *flag) // DEBUG
    if *flag == "-a" {
        fmt.Printf("If flag is -a this is awkward: %s", *flag)
        return Deployment{
            APIVersion: "apps/v1",
            Kind: "Deployment",
            Metadata: ObjectMeta{
                Name: "test-app",
                Namespace: "default",
                Labels: MetadataLabels{
                    App: "test-app",
                },
            },
            Spec: DeploymentSpec{
                Replicas: 1,
                Template: PodTemplateSpec{
                    Metadata: ObjectMeta{
                        Name: "test-app",
                        Namespace: "default",
                    },
                    Spec: PodSpec{
                        SchedulerName: "kube-scheduler",
                        ServiceAccount: "default",
                        Containers: Containers{
                            Name: "test-app-container",
                            Image: "<Account-Name>/<Repository>:<x,y,z>",
                            ImagePullPolicy: "Always",
                            VolumeMounts: VolumeMount{
                                Name: "test-volume",
                                MountPath: "/mnt/data/",
                            },
                        },
                        NodeName: "node-01",
                        RestartPolicy: "OnFailure",
                        Volumes: Volume{
                            Name: "test-volume",
                            PersistentVolumeClaim: PersistentVolumeClaim{
                                ClaimName: "test-pvc",
                                ReadOnly: true,
                            },
                        },
                    },
                },
            },
        }
    }
    fmt.Println("Correct?")
    return Deployment{
        APIVersion: "apps/v1",
        Kind: "Deployment",
        Metadata: ObjectMeta{
            Name: "test-app",
            Namespace: "default",
            Labels: MetadataLabels{
                App: "test-app",
            },
        },
        Spec: DeploymentSpec{
            Replicas: 1,
            Template: PodTemplateSpec{
                Metadata: ObjectMeta{
                    Name: "test-app",
                    Namespace: "default",
                },
                Spec: PodSpec{
                    SchedulerName: "kube-scheduler",
                    ServiceAccount: "default",
                    Containers: Containers{
                        Name: "test-app-container",
                        Image: "<Account-Name>/<Repository>:<x,y,z>",
                        ImagePullPolicy: "Always",
                        VolumeMounts: VolumeMount{
                            Name: "test-volume",
                            MountPath: "/mnt/data/",
                        },
                    },
                    RestartPolicy: "OnFailure",
                    Volumes: Volume{
                        Name: "test-volume",
                        PersistentVolumeClaim: PersistentVolumeClaim{
                            ClaimName: "test-pvc",
                        },
                    },
                },
            },
        },
    }
}

func GenerateServiceAccount(flag *string) ServiceAccount {

    return ServiceAccount{

    }
}

func GenerateRole(flag *string) Role {

    return Role{

    }
}

func GenerateRoleBinding(flag *string) RoleBinding {

    return RoleBinding{

    }
}

func GenerateClusterRole(flag *string) ClusterRole {

    return ClusterRole{

    }
}

func GenerateClusterRoleBinding(flag *string) ClusterRoleBinding {
    
    return ClusterRoleBinding{

    }
}

func GeneratePod(flag *string) Pod {

    return Pod{

    }
}
