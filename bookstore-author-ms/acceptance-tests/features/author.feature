Feature: Flow of author CRUD operations with database

  Background:
    Given a default author
    And the author is created in author-db

  Scenario: The author is created
    Then the author id read in author-db is "default-id"
    And the author name read in author-db is "default-name"
    And the author book-id read in author-db is "default-book-id"

  Scenario: The author is obtained
    When the author is obtained in author-db
    Then the author id read in author-db is "default-id"
    And the author name read in author-db is "default-name"
    And the author book-id read in author-db is "default-book-id"

  Scenario: The author is updated
    Given the author id is "9f9c19fc-68e6-428b-9483-bdfc3eb50046"
    And the author name is "William"
    And the author book-id is "2d12d0a0-b555-11ed-afa1-0242ac120002"
    When the author is updated in author-db
    Then the author id read in author-db is "9f9c19fc-68e6-428b-9483-bdfc3eb50046"
    And the author name read in author-db is "William"
    And the author book-id read in author-db is "2d12d0a0-b555-11ed-afa1-0242ac120002"

  Scenario: The author is deleted
    Then the author id read in author-db is "default-id"
    When the author is deleted in author-db
    Then the author id "default-id" isn't exists in author-db
