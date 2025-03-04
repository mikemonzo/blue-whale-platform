package steps

import (
	"context"
	"fmt"
	"time"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service/impl"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repository"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/persistence/memory"
)

type UserCreationContext struct {
	userService  service.UserService
	mailService  service.MailService
	userRepo     repository.UserRepository
	currentUser  *model.User
	createdUser  *model.User
	existingUser *model.User
	err          error
}

func NewUserCreationContext(userService service.UserService, mailService service.MailService, userRepo repository.UserRepository) *UserCreationContext {
	return &UserCreationContext{
		userService: userService,
		mailService: mailService,
		userRepo:    userRepo,
	}
}

func (ctx *UserCreationContext) InitializeScenario(sc *godog.ScenarioContext) {
	// Add BeforeScenario hook to reset state
	sc.BeforeScenario(func(*godog.Scenario) {
		// Reset the in-memory repository before each scenario
		ctx.userRepo = memory.NewInMemoryUserRepository()
		ctx.userService = impl.NewUserService(ctx.userRepo, ctx.mailService)
		ctx.createdUser = nil
		ctx.existingUser = nil
		ctx.err = nil
	})

	// Register steps
	sc.Step(`^a welcome email should be sent to "([^"]*)"$`, ctx.aWelcomeEmailShouldBeSentTo)
	sc.Step(`^I am an authenticated administrator$`, ctx.iAmAnAuthenticatedAdministrator)
	sc.Step(`^I create a new user with the following details:$`, ctx.iCreateANewUserWithTheFollowingDetails)
	sc.Step(`^the user should be created successfully$`, ctx.theUserShouldBeCreatedSuccessfully)
	sc.Step(`^the user should be inactive by default$`, ctx.theUserShouldBeInactiveByDefault)
	// Add more steps here
	sc.Step(`^a user exists with the email "([^"]*)"$`, ctx.aUserExistsWithEmail)
	sc.Step(`^I create a user with email "([^"]*)"$`, ctx.iCreateAUserWithEmail)
	sc.Step(`^I should see an error "([^"]*)"$`, ctx.iShouldSeeAnError)
}

func (ctx *UserCreationContext) iAmAnAuthenticatedAdministrator() error {
	// Simulamos un usuario administrador autenticado
	ctx.currentUser = &model.User{
		ID:       "1",
		Username: "admin",
		Email:    "admin@example.com",
		Status:   model.UserStatusActive,
	}
	return nil
}

func (ctx *UserCreationContext) iCreateANewUserWithTheFollowingDetails(table *godog.Table) error {
	var email, username, firstName, lastName, password string

	// Get data from the table
	for _, row := range table.Rows[1:] {
		email = row.Cells[0].Value
		username = row.Cells[1].Value
		firstName = row.Cells[2].Value
		lastName = row.Cells[3].Value
		password = row.Cells[4].Value
	}

	// Create the user using service
	user, err := ctx.userService.Create(context.Background(), email, username, firstName, lastName, password)
	if err != nil {
		ctx.err = err
		return err
	}

	ctx.createdUser = user
	return nil
}

func (ctx *UserCreationContext) theUserShouldBeCreatedSuccessfully() error {
	if ctx.err != nil {
		return ctx.err
	}
	if ctx.createdUser == nil {
		return fmt.Errorf("user was not created")
	}
	return nil
}

func (ctx *UserCreationContext) theUserShouldBeInactiveByDefault() error {
	if ctx.createdUser.Status != model.UserStatusInactive {
		return fmt.Errorf("expected user to be inactive, but got %s", ctx.createdUser.Status)
	}
	return nil
}

func (ctx *UserCreationContext) aWelcomeEmailShouldBeSentTo(email string) error {
	// Check if email is the same as the user's email
	if ctx.createdUser.Email != email {
		return fmt.Errorf("expected email %s, but got %s", email, ctx.createdUser.Email)
	}
	return nil
}

func (ctx *UserCreationContext) aUserExistsWithEmail(email string) error {
	// Create user directly through repository to avoid duplicate email check
	user := &model.User{
		ID:        uuid.New().String(),
		Email:     email,
		Username:  "existinguser",
		FirstName: "Existing",
		LastName:  "User",
		Password:  "password123",
		Status:    model.UserStatusInactive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := ctx.userRepo.Create(context.Background(), user); err != nil {
		return fmt.Errorf("failed to create existing user: %w", err)
	}

	ctx.existingUser = user
	return nil
}

func (ctx *UserCreationContext) iCreateAUserWithEmail(email string) error {
	// Try to create a user with the same email
	_, err := ctx.userService.Create(context.Background(),
		email,
		"newuser",
		"New",
		"User",
		"password123")
	ctx.err = err
	return nil
}

func (ctx *UserCreationContext) iShouldSeeAnError(errMsg string) error {
	if ctx.err == nil {
		return fmt.Errorf("expected error, but got nil")
	}
	if ctx.err.Error() != errMsg {
		return fmt.Errorf("expected error %s, but got %s", errMsg, ctx.err.Error())
	}
	return nil
}
