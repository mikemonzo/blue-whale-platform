package features

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/features/steps"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service/impl"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service/mocks"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/persistence/memory"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			// Initialize repositories and services
			userRepo := memory.NewInMemoryUserRepository()
			mailService := mocks.NewMockMailService()
			userService := impl.NewUserService(userRepo, mailService)

			// Initialize contexts for different features
			userCreationContext := steps.NewUserCreationContext(userService, mailService, userRepo)
			userUpdateContext := steps.NewUserUpdateContext(userService, mailService, userRepo)

			// Initialize scenarios for all features
			userCreationContext.InitializeScenario(sc)
			userUpdateContext.InitializeScenario(sc)
		},
		Options: &godog.Options{
			Format: "pretty",
			Paths: []string{
				"user_creation.feature",
				"user_update.feature",
			},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
