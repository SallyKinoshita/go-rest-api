package config

import (
	"context"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type AWS struct {
	Region  string `envconfig:"AWS_REGION" required:"true" default:"ap-northeast-1"`
	Cognito Cognito
}

type CustomEndpointResolver struct {
	Endpoint string
}

func (c *CustomEndpointResolver) ResolveEndpoint(service, region string) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:               c.Endpoint,
		HostnameImmutable: true,
	}, nil
}

func NewAWS(ctx context.Context, aws AWS, environment Environment) (*aws.Config, error) {
	if environment.Environment == Local {
		// ローカル環境用のエンドポイントリゾルバーを作成
		customResolver := &CustomEndpointResolver{
			Endpoint: aws.Cognito.Endpoint, // 環境変数や設定に基づくカスタムエンドポイント
		}

		cfg, err := awsconfig.LoadDefaultConfig(
			ctx,
			awsconfig.WithRegion(aws.Region),
			awsconfig.WithEndpointResolver(customResolver),
		)
		if err != nil {
			return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
		}

		return &cfg, nil
	}

	// 通常のAWS設定
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(aws.Region))
	if err != nil {
		return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	return &cfg, nil
}

type Cognito struct {
	UserPoolID   string `envconfig:"AWS_COGNITO_USER_POOL_ID" required:"true"`
	ClientID     string `envconfig:"AWS_COGNITO_CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"AWS_COGNITO_CLIENT_SECRET" required:"true"`
	Endpoint     string `envconfig:"AWS_COGNITO_ENDPOINT"` // カスタムエンドポイントを環境変数で指定
}

func NewCognito(awsCfg *aws.Config) *cognitoidentityprovider.Client {
	return cognitoidentityprovider.NewFromConfig(*awsCfg)
}
