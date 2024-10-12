from decimal import Decimal
from datetime import date
from typing import List, Dict


class User:
    def __init__(self, user_id: int, first_name: str, last_name: str, email: str):
        self.id = user_id
        self.first_name = first_name
        self.last_name = last_name
        self.email = email

    def get_user_info(self) -> str:
        pass

    def login(self):
        pass

    def register(self):
        pass

    def verified(self) -> bool:
        pass


class Librarian(User):
    def add_book(self, book: 'Book'):
        pass

    def remove_book(self, book: 'Book'):
        pass

    def update_book_info(self, book: 'Book'):
        pass

    def add_genre(self, genre: 'Genre'):
        pass


class Reader(User):
    def __init__(self, user_id: int, first_name: str, last_name: str, email: str, rented_books: List['Rental'],
                 fines: List['Fine']):
        super().__init__(user_id, first_name, last_name, email)
        self.rented_books = rented_books
        self.fines = fines

    def rent_book(self, book_id: int) -> 'Rental':
        pass

    def return_book(self, rental: 'Rental'):
        pass

    def pay_fine(self, fine: 'Fine'):
        pass


class Book:
    def __init__(self, book_id: int, title: str, author: str, publication_date: date, pages: int,
                 genres: List['Genre']):
        self.id = book_id
        self.title = title
        self.author = author
        self.publication_date = publication_date
        self.pages = pages
        self.genres = genres

    def get_book_info(self) -> str:
        pass

    def add_genre(self, genre: 'Genre'):
        pass

    def update_book_info(self, data: Dict):
        pass

    def remove_genre(self, genre: 'Genre'):
        pass


class Genre:
    def __init__(self, genre_id: int, name: str):
        self.id = genre_id
        self.name = name

    def update_genre_info(self, data: Dict):
        pass


class Rental:
    def __init__(self, rental_id: int, book: 'Book', reader: 'Reader', rental_date: date, due_date: date,
                 is_returned: bool):
        self.id = rental_id
        self.book = book
        self.reader = reader
        self.rental_date = rental_date
        self.due_date = due_date
        self.is_returned = is_returned

    def calculate_fine(self) -> 'Fine':
        pass

    def mark_as_returned(self):
        pass


class Fine:
    def __init__(self, fine_id: int, amount: Decimal, created_date: date, is_paid: bool):
        self.id = fine_id
        self.amount = amount
        self.created_date = created_date
        self.is_paid = is_paid

    def pay_fine(self):
        pass
