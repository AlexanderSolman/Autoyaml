

package autoyaml

import (
)
// Could take flag as argument i.e -b, -a for which version to return
func generateDeployment(flag *string) Deployment {
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
                    //NodeName: "node-01 #Change or Remove",
                    RestartPolicy: "OnFailure",
                    Volumes: Volume{
                        Name: "test-volume",
                        PersistentVolumeClaim: PersistentVolumeClaim{
                            ClaimName: "test-pvc",
                            //ReadOnly: true,
                        },
                    },
                },
            },
        },
    }
}
