Feature: User Update
  As an administrator
  I want to update user information
  So that user data remains current

  Scenario: Successfully update user profile
    Given I am an authenticated administrator
    And a user exists with email "john@email.com"
    When I update the user with the following details:
      | first_name | last_name | username |
      | Johnny     | Doe       | johnny   |
    Then the user should be updated successfully
    And the user's updated_at time should be current

  Scenario: Update user email to existing email
    Given I am an authenticated administrator
    And a user exists with email "john@email.com"
    And a user exists with email "jane@email.com"
    When I update the user "john@email.com" email to "jane@email.com"
    Then I should see an error "email already exists"