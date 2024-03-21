package autoyaml


// Possibly move var flag *string globally here and check if it is set or not

// Could take flag as argument i.e -b(or blank as -b is default), -a for which version to return
func GenerateDeployment(flag *string) Deployment {
    if *flag == "-a" {
        // TODO filled version all variables
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
                        ServiceAccount: "default #Change or Remove",
                        Containers: Containers{
                            Name: "test-app-container",
                            Image: "<Account-Name>/<Repository>:<x,y,z>",
                            ImagePullPolicy: "Always",
                            VolumeMounts: VolumeMount{
                                Name: "test-volume",
                                MountPath: "/mnt/data/",
                            },
                        },
                        NodeName: "node-01 #Change or Remove",
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
                    ServiceAccount: "default #Change or Remove",
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
