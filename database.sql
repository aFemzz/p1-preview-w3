-- Create the author table
CREATE TABLE author (
    author_id INT AUTO_INCREMENT PRIMARY KEY,
    author_name VARCHAR(255) UNIQUE,
    author_email VARCHAR(255)
);

-- Create the book table
CREATE TABLE book (
    book_id INT AUTO_INCREMENT PRIMARY KEY,
    book_title VARCHAR(255) UNIQUE,
    book_type VARCHAR(255),
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES author(author_id)
);

-- Create the order table
CREATE TABLE bookstore_order (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(255),
    customer_email VARCHAR(255),
    order_date DATE,
    price DECIMAL(10,2),
    book_id INT,
    FOREIGN KEY (book_id) REFERENCES book(book_id)
);