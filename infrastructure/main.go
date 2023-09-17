package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		gcpConf := config.New(ctx, "gcp")
		project := gcpConf.Get("project")
		region := gcpConf.Get("region")

		conf := config.New(ctx, "run")
		serviceName := conf.Get("serviceName")
		imageName := conf.Get("imageName")
		concurrency := conf.GetInt("concurrency")
		containerPort := conf.GetInt("containerPort")
		cpu := conf.Get("cpu")
		memory := conf.Get("memory")
		// TODO: Add DomainMapping
		//domain := conf.Get("domain")

		_, err := cloudrun.NewService(ctx, serviceName, &cloudrun.ServiceArgs{
			Location: pulumi.String(region),
			Metadata: &cloudrun.ServiceMetadataArgs{
				Namespace: pulumi.String(project),
			},
			Name:    pulumi.String(serviceName),
			Project: pulumi.String(project),
			Template: &cloudrun.ServiceTemplateArgs{
				Spec: &cloudrun.ServiceTemplateSpecArgs{
					ContainerConcurrency: pulumi.Int(concurrency),
					Containers: cloudrun.ServiceTemplateSpecContainerArray{
						&cloudrun.ServiceTemplateSpecContainerArgs{
							Envs: cloudrun.ServiceTemplateSpecContainerEnvArray{
								&cloudrun.ServiceTemplateSpecContainerEnvArgs{
									Name: pulumi.String("GITHUB_TOKEN"),
									ValueFrom: &cloudrun.ServiceTemplateSpecContainerEnvValueFromArgs{
										SecretKeyRef: &cloudrun.ServiceTemplateSpecContainerEnvValueFromSecretKeyRefArgs{
											Key:  pulumi.String("1"),
											Name: pulumi.String("gha-token"),
										},
									},
								},
							},
							Image: pulumi.String(imageName),
							Ports: cloudrun.ServiceTemplateSpecContainerPortArray{
								&cloudrun.ServiceTemplateSpecContainerPortArgs{
									ContainerPort: pulumi.Int(containerPort),
									Name:          pulumi.String("http1"),
								},
							},
							Resources: &cloudrun.ServiceTemplateSpecContainerResourcesArgs{
								Limits: pulumi.StringMap{
									"cpu":    pulumi.String(cpu),
									"memory": pulumi.String(memory),
								},
							},
							StartupProbe: &cloudrun.ServiceTemplateSpecContainerStartupProbeArgs{
								FailureThreshold: pulumi.Int(1),
								PeriodSeconds:    pulumi.Int(240),
								TcpSocket: &cloudrun.ServiceTemplateSpecContainerStartupProbeTcpSocketArgs{
									Port: pulumi.Int(containerPort),
								},
								TimeoutSeconds: pulumi.Int(240),
							},
						},
					},
					// TODO: Generate ServiceAccount with Pulumi
					ServiceAccountName: pulumi.String("33511728547-compute@developer.gserviceaccount.com"),
					TimeoutSeconds:     pulumi.Int(300),
				},
			},
			Traffics: cloudrun.ServiceTrafficArray{
				&cloudrun.ServiceTrafficArgs{
					LatestRevision: pulumi.Bool(true),
					Percent:        pulumi.Int(100),
				},
			},
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		return nil
	})
}
