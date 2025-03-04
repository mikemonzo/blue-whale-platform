Feature: User Creation
    As an administrator
    I want to create a new user
    So that they can access the system

    Scenario: Successfully create a new user
        Given I am an authenticated administrator
        When I create a new user with the following details:
            | email             | username  | first_name | last_name | password     |
            | john@email.com    | johndoe   | John       | Doe       | SecurePass1! |
        Then the user should be created successfully
        And the user should be inactive by default
        And  a welcome email should be sent to "john@email.com"

    Scenario: Create user with existing email
        Given I am an authenticated administrator
        And a user exists with the email "john@email.com"
        When I create a user with email "john@email.com"
        Then I should see an error "user with email john@email.com already exists"