package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
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
		secretName := conf.Get("secret")
		// TODO: Add DomainMapping

		serviceAccount, err := serviceaccount.NewAccount(ctx, serviceName, &serviceaccount.AccountArgs{
			AccountId:   pulumi.String(serviceName),
			DisplayName: pulumi.String(serviceName),
			Project:     pulumi.String(project),
		})
		if err != nil {
			return err
		}

		_, err = serviceaccount.NewIAMBinding(ctx, "gha-server", &serviceaccount.IAMBindingArgs{
			Members: pulumi.StringArray{
				pulumi.String("serviceAccount:" + serviceName + "@" + project + ".iam.gserviceaccount.com"),
			},
			Role:             pulumi.String("roles/iam.serviceAccountUser"),
			ServiceAccountId: serviceAccount.Name,
		})
		if err != nil {
			return err
		}
		_, err = secretmanager.NewSecretIamBinding(ctx, "gha-secret-accessor", &secretmanager.SecretIamBindingArgs{
			SecretId: pulumi.String(secretName),
			Role:     pulumi.String("roles/secretmanager.secretAccessor"),
			Members: pulumi.StringArray{
				pulumi.String("serviceAccount:" + serviceName + "@" + project + ".iam.gserviceaccount.com"),
			},
		})
		if err != nil {
			return err
		}

		secret, err := secretmanager.LookupSecret(ctx, &secretmanager.LookupSecretArgs{
			SecretId: secretName,
		}, nil)
		if err != nil {
			return err
		}

		_, err = cloudrun.NewService(ctx, serviceName, &cloudrun.ServiceArgs{
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
											Name: pulumi.String(secret.SecretId),
										},
									},
								},
							},
							Image: pulumi.String(imageName),
							Ports: cloudrun.ServiceTemplateSpecContainerPortArray{
								&cloudrun.ServiceTemplateSpecContainerPortArgs{
									ContainerPort: pulumi.Int(containerPort),
								},
							},
							Resources: &cloudrun.ServiceTemplateSpecContainerResourcesArgs{
								Limits: pulumi.StringMap{
									"cpu":    pulumi.String(cpu),
									"memory": pulumi.String(memory),
								},
							},
						},
					},
					ServiceAccountName: serviceAccount.Email,
					TimeoutSeconds:     pulumi.Int(300),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
