package cmd

import (
	"context"
	"fmt"

	//"cloud.google.com/go/secretmanager"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	//"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// Variable to store the retrieved secret string
var sourceSecretString string

// secretsCmd represents the secrets command
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Retrieve and create cloud secrets",
	Run: func(cmd *cobra.Command, args []string) {
		cloudProvider, _ := cmd.Flags().GetString("source-cloud-provider")

		// Load AWS and GCP configurations based on profile names
		sourceAWSProfile, _ := cmd.Flags().GetString("source-aws-profile")
		destAWSProfile, _ := cmd.Flags().GetString("dest-aws-profile")
		sourceGCPProfile, _ := cmd.Flags().GetString("source-gcp-profile")
		destGCPProfile, _ := cmd.Flags().GetString("dest-gcp-profile")

		sourceAWSConfig, err := loadAWSConfig(sourceAWSProfile)
		if err != nil {
			fmt.Printf("Error loading source AWS configuration: %v\n", err)
			return
		}

		destAWSConfig, err := loadAWSConfig(destAWSProfile)
		if err != nil {
			fmt.Printf("Error loading destination AWS configuration: %v\n", err)
			return
		}

		sourceGCPConfig, err := loadGCPConfig(sourceGCPProfile)
		if err != nil {
			fmt.Printf("Error loading source GCP configuration: %v\n", err)
			return
		}

		destGCPConfig, err := loadGCPConfig(destGCPProfile)
		if err != nil {
			fmt.Printf("Error loading destination GCP configuration: %v\n", err)
			return
		}

		switch cloudProvider {
		case "aws":
			awsSecretName, _ := cmd.Flags().GetString("source-secret-name")
			secretString, err := getAWSSecretString(awsSecretName, sourceAWSConfig) // Pass sourceAWSConfig here
			if err != nil {
				fmt.Printf("Error retrieving AWS secret: %v\n", err)
				return
			}
			fmt.Printf("AWS Secret String: %s\n", secretString)
			sourceSecretString = secretString

		case "gcp":
			gcpProjectID, _ := cmd.Flags().GetString("gcp-project-id")
			gcpSecretName, _ := cmd.Flags().GetString("source-secret-name")
			secretString, err := getGCPSecretString(gcpSecretName, gcpProjectID, sourceGCPConfig) // Pass secretName, projectID, and sourceGCPConfig
			if err != nil {
				fmt.Printf("Error retrieving GCP secret: %v\n", err)
				return
			}
			fmt.Printf("GCP Secret String: %s\n", secretString)
			sourceSecretString = secretString

		default:
			fmt.Println("Unsupported source cloud provider choice.")
			return
		}

		destCloudProvider, _ := cmd.Flags().GetString("dest-cloud-provider")
		if destCloudProvider != "" {
			switch destCloudProvider {
			case "aws":
				destAWSSecretName, _ := cmd.Flags().GetString("dest-secret-name")
				destAWSSecretString := sourceSecretString                                     // Use the retrieved secret string
				err := createAWSSecret(destAWSSecretName, destAWSSecretString, destAWSConfig) // Pass destAWSConfig here
				if err != nil {
					fmt.Printf("Error creating AWS secret: %v\n", err)
					return
				}
				fmt.Println("AWS secret created successfully.")

			case "gcp":
				destGCPSecretName, _ := cmd.Flags().GetString("dest-secret-name")
				destGCPSecretString := sourceSecretString // Use the retrieved secret string
				destGCPProjectID, _ := cmd.Flags().GetString("gcp-project-id")
				err := createGCPSecret(destGCPSecretName, destGCPSecretString, destGCPProjectID, destGCPConfig) // Pass destGCPConfig here
				if err != nil {
					fmt.Printf("Error creating GCP secret: %v\n", err)
					return
				}
				fmt.Println("GCP secret created successfully.")

			default:
				fmt.Println("Unsupported destination cloud provider choice.")
			}
		} else {
			fmt.Println("No secret creation requested.")
		}
	},
}

func init() {
	migrateCmd.AddCommand(secretsCmd)

	secretsCmd.Flags().String("source-cloud-provider", "aws", "Source cloud provider (aws or gcp)")
	secretsCmd.Flags().String("dest-cloud-provider", "", "Destination cloud provider (aws or gcp)")

	// Flags for AWS secrets
	secretsCmd.Flags().String("source-secret-name", "", "Source secret name")
	secretsCmd.Flags().String("source-aws-profile", "", "AWS profile name for the source")
	secretsCmd.Flags().String("dest-aws-profile", "", "AWS profile name for the destination")

	// Flags for GCP secrets
	secretsCmd.Flags().String("gcp-project-id", "", "GCP project ID")
	secretsCmd.Flags().String("source-gcp-profile", "", "GCP profile name for the source")
	secretsCmd.Flags().String("dest-gcp-profile", "", "GCP profile name for the destination")

	// Flags for secret creation
	secretsCmd.Flags().String("dest-secret-name", "", "Destination secret name")

	// Flags for secret string
	secretsCmd.Flags().String("secret-string", "", "Secret string")
}

// Functions to load AWS and GCP configurations based on profile names
func loadAWSConfig(profileName string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profileName),
	)
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}

func loadGCPConfig(profileName string) (*secretmanager.Client, error) {
	ctx := context.Background()

	// Load the specified GCP profile if provided
	if profileName != "" {
		client, err := secretmanager.NewClient(ctx, option.WithCredentialsFile(profileName))
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	// Load the default GCP client if no profile is specified
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Function to create an AWS secret
func createAWSSecret(secretName, secretString string, awsConfig aws.Config) error {
	svc := secretsmanager.NewFromConfig(awsConfig)

	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(secretName),
		SecretString: aws.String(secretString),
	}

	_, err := svc.CreateSecret(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

// Function to create a GCP secret
func createGCPSecret(secretName, secretString, projectID string, gcpClient *secretmanager.Client) error {
	ctx := context.Background()

	parent := fmt.Sprintf("projects/%s", projectID)
	secret := fmt.Sprintf("%s/secrets/%s", parent, secretName)

	payload := []byte(secretString)
	_, err := gcpClient.AddSecretVersion(ctx, &secretmanagerpb.AddSecretVersionRequest{
		Parent:  secret,
		Payload: &secretmanagerpb.SecretPayload{Data: payload},
	})
	if err != nil {
		return err
	}

	return nil
}

func getAWSSecretString(secretName string, awsConfig aws.Config) (string, error) {
	svc := secretsmanager.NewFromConfig(awsConfig)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", err
	}

	if result.SecretString != nil {
		secretString := *result.SecretString
		return secretString, nil
	}

	return "", nil
}

func getGCPSecretString(secretName, projectID string, gcpClient *secretmanager.Client) (string, error) {
	ctx := context.Background()

	// Construct the secret name
	secretPath := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName)

	// Access the secret version
	version, err := gcpClient.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretPath,
	})
	if err != nil {
		return "", err
	}

	// Convert the payload to a string
	secretString := string(version.Payload.Data)

	return secretString, nil
}
