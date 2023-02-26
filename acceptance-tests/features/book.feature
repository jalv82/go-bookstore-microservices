Feature: Flow of book register in database

  Background:
    Given a default book

  Scenario: The book register in database is successful
    Given the book title is "Go in Action"
    And the book author-id is "3e21ae38-b5f6-11ed-afa1-0242ac120002"
#    When the book is registered in author-db
#    Then the book id read in book-db is "default-id"
#    And the book title read in book-db is "Go in Action"
#    And the author-id of the book author read in book-db is "3e21ae38-b5f6-11ed-afa1-0242ac120002"