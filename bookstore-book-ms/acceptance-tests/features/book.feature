Feature: Flow of book CRUD operations with database

  Background:
    Given a default book
    And the book is created in book-db

  Scenario: The book is created
    Then the book id read in book-db is "default-id"
    And the book title read in book-db is "default-title"
    And the book author-id read in book-db is "default-author-id"

  Scenario: The book is obtained
    When the book is obtained in book-db
    Then the book id read in book-db is "default-id"
    And the book title read in book-db is "default-title"
    And the book author-id read in book-db is "default-author-id"

  Scenario: The book is updated
    Given the book id is "da374ae1-cfda-4b47-b45d-36064ee3c161"
    And the book title is "Go in Action"
    And the book author-id is "3e21ae38-b5f6-11ed-afa1-0242ac120002"
    When the book is updated in book-db
    Then the book id read in book-db is "da374ae1-cfda-4b47-b45d-36064ee3c161"
    And the book title read in book-db is "Go in Action"
    And the book author-id read in book-db is "3e21ae38-b5f6-11ed-afa1-0242ac120002"

  Scenario: The book is deleted
    Then the book id read in book-db is "default-id"
    When the book is deleted in book-db
    Then the book id "default-id" isn't exists in book-db
