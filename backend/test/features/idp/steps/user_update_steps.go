package steps

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service/impl"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repository"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/persistence/memory"
)

type UserUpdateContext struct {
	userService service.UserService
	mailService service.MailService
	userRepo    repository.UserRepository
	currentUser *model.User
	updatedUser *model.User
	err         error
}

func NewUserUpdateContext(userService service.UserService, mailService service.MailService, userRepo repository.UserRepository) *UserUpdateContext {
	return &UserUpdateContext{
		userService: userService,
		mailService: mailService,
		userRepo:    userRepo,
	}
}

func (ctx *UserUpdateContext) InitializeScenario(sc *godog.ScenarioContext) {
	// Add BeforeScenario hook to reset state
	sc.Before(func(testCtx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Create new instances for each scenario
		ctx.userRepo = memory.NewInMemoryUserRepository()
		ctx.userService = impl.NewUserService(ctx.userRepo, ctx.mailService)
		ctx.currentUser = nil
		ctx.updatedUser = nil
		ctx.err = nil
		return testCtx, nil
	})

	sc.Step(`^a user exists with email "([^"]*)"$`, ctx.aUserExistsWithEmail)
	sc.Step(`^I am an authenticated administrator$`, ctx.iAmAnAuthenticatedAdministrator)
	sc.Step(`^I should see an error "([^"]*)"$`, ctx.iShouldSeeAnError)
	sc.Step(`^I update the user "([^"]*)" email to "([^"]*)"$`, ctx.iUpdateTheUserEmailTo)
	sc.Step(`^I update the user with the following details:$`, ctx.iUpdateTheUserWithTheFollowingDetails)
	sc.Step(`^the user should be updated successfully$`, ctx.theUserShouldBeUpdatedSuccessfully)
	sc.Step(`^the user\'s updated_at time should be current$`, ctx.theUsersUpdatedAtTimeShouldBeCurrent)
}

func (ctx *UserUpdateContext) aUserExistsWithEmail(email string) error {
	user := &model.User{
		ID:        uuid.New().String(),
		Email:     email,
		Username:  "existinguser",
		FirstName: "Existing",
		LastName:  "User",
		Password:  "password123",
		Status:    model.UserStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := ctx.userRepo.Create(context.Background(), user); err != nil {
		return fmt.Errorf("failed to create existing user: %w", err)
	}

	ctx.currentUser = user
	log.Printf("Created user: %+v\n", user)
	return nil
}

func (ctx *UserUpdateContext) iAmAnAuthenticatedAdministrator() error {
	ctx.currentUser = &model.User{
		ID:       "admin-123",
		Username: "admin",
		Email:    "admin@example.com",
		Status:   model.UserStatusActive,
	}
	log.Printf("Authenticated as administrator: %+v\n", ctx.currentUser)
	return nil
}

func (ctx *UserUpdateContext) iShouldSeeAnError(expectedError string) error {
	if ctx.err == nil {
		return fmt.Errorf("expected error '%s' but got nil", expectedError)
	}
	if ctx.err.Error() != expectedError {
		return fmt.Errorf("expected error '%s' but got '%s'", expectedError, ctx.err.Error())
	}
	log.Printf("Saw expected error: %s\n", expectedError)
	return nil
}

func (ctx *UserUpdateContext) iUpdateTheUserEmailTo(oldEmail, newEmail string) error {
	// Get the existing user
	existingUser, err := ctx.userRepo.GetUserByEmail(context.Background(), oldEmail)
	if err != nil {
		ctx.err = err
		log.Printf("Error getting user by email: %v\n", err)
		return nil
	}
	if existingUser == nil {
		ctx.err = fmt.Errorf("user not found")
		log.Printf("User not found with email: %s\n", oldEmail)
		return nil
	}

	// Create updated user with new email
	updatedUser := *existingUser
	updatedUser.Email = newEmail

	// Try to update
	err = ctx.userService.Update(context.Background(), &updatedUser)
	ctx.err = err // Store error for later assertion
	if err == nil {
		ctx.updatedUser = &updatedUser
		log.Printf("Updated user: %+v\n", updatedUser)
	} else {
		log.Printf("Error updating user: %v\n", err)
	}
	return nil
}

func (ctx *UserUpdateContext) iUpdateTheUserWithTheFollowingDetails(table *godog.Table) error {
	if ctx.currentUser == nil {
		return fmt.Errorf("no user exists to update")
	}

	var firstName, lastName, username string

	// Get data from the table
	for _, row := range table.Rows[1:] {
		firstName = row.Cells[0].Value
		lastName = row.Cells[1].Value
		username = row.Cells[2].Value
	}

	updatedUser := &model.User{
		ID:        ctx.currentUser.ID,
		Email:     ctx.currentUser.Email,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Status:    ctx.currentUser.Status,
		CreatedAt: ctx.currentUser.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if err := ctx.userService.Update(context.Background(), updatedUser); err != nil {
		ctx.err = err
		log.Printf("Failed to update user: %v\n", err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	ctx.updatedUser = updatedUser
	log.Printf("Updated user with details: %+v\n", updatedUser)
	return nil
}

func (ctx *UserUpdateContext) theUserShouldBeUpdatedSuccessfully() error {
	if ctx.err != nil {
		return fmt.Errorf("unexpected error: %w", ctx.err)
	}
	if ctx.updatedUser == nil {
		return fmt.Errorf("expected updated user but got nil")
	}
	log.Printf("User updated successfully: %+v\n", ctx.updatedUser)
	return nil
}

func (ctx *UserUpdateContext) theUsersUpdatedAtTimeShouldBeCurrent() error {
	if ctx.updatedUser == nil {
		return fmt.Errorf("expected updated user but got nil")
	}

	// Check if UpdatedAt is within the last second
	if time.Since(ctx.updatedUser.UpdatedAt) > time.Second {
		return fmt.Errorf("expected UpdatedAt to be current but got %v", ctx.updatedUser.UpdatedAt)
	}
	log.Printf("User's updated_at time is current: %v\n", ctx.updatedUser.UpdatedAt)
	return nil
}
