Feature: Flow of author register in database

  Background:
    Given a default author

  Scenario: The author register in database is successful
    Given the author name is "William"
    And the author book-id is "2d12d0a0-b555-11ed-afa1-0242ac120002"
#    When the author is registered in author-db
#    Then the author id read in author-db is "default-id"
#    And the author name read in author-db is "William"
#    And the book-id of the author book read in author-db is "2d12d0a0-b555-11ed-afa1-0242ac120002"