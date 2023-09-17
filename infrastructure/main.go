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
				Annotations: pulumi.StringMap{
					"run.googleapis.com/client-name":    pulumi.String("cloud-console"),
					"run.googleapis.com/ingress":        pulumi.String("all"),
					"run.googleapis.com/ingress-status": pulumi.String("all"),
					"run.googleapis.com/operation-id":   pulumi.String("48d2d5a7-52c8-4969-9b2c-bb79c85a831e"),
					"serving.knative.dev/creator":       pulumi.String("2010southafrica@gmail.com"),
					"serving.knative.dev/lastModifier":  pulumi.String("2010southafrica@gmail.com"),
				},
				Labels: pulumi.StringMap{
					"cloud.googleapis.com/location": pulumi.String(region),
					"gcb-trigger-id":                pulumi.String("c5a72a2d-1d2a-4d1c-8efa-df0b8880bdad"),
					"gcb-trigger-region":            pulumi.String("global"),
					"managed-by":                    pulumi.String("gcp-cloud-build-deploy-cloud-run"),
				},
				Namespace: pulumi.String(project),
			},
			Name:    pulumi.String(serviceName),
			Project: pulumi.String(project),
			Template: &cloudrun.ServiceTemplateArgs{
				Metadata: &cloudrun.ServiceTemplateMetadataArgs{
					Annotations: pulumi.StringMap{
						"autoscaling.knative.dev/maxScale":         pulumi.String("2"),
						"run.googleapis.com/client-name":           pulumi.String("cloud-console"),
						"run.googleapis.com/execution-environment": pulumi.String("gen1"),
						"run.googleapis.com/startup-cpu-boost":     pulumi.String("true"),
					},
					Labels: pulumi.StringMap{
						"client.knative.dev/nonce":            pulumi.String("8784e239-ddaf-4820-af94-4e1c6ad2ff2c"),
						"run.googleapis.com/startupProbeType": pulumi.String("Default"),
					},
				},
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
