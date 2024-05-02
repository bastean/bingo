Feature: Create a new Binary

  Scenario: Create a valid Binary
    Given I am on / page
    Then I fill the Filename with bingo
    * I fill the Description with Custom executable
    * I select the Platform in Linux 64-bit
    * I click the Build button
    And I see Download started! notification

  Scenario: Create an invalid Binary
    Given I am on / page
    Then I click the Build button
    But I see Filename must have a length between 1 and 50 characters notification
    * I see Description must have a length between 1 and 100 characters notification
    * I see Please select a target platform notification
