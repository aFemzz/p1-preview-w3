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

-- Inserting Data

-- Insert books authored by John Doe
INSERT INTO book (book_title, book_type, author_id) VALUES
('Book Title 1', 'Physical', 1),
('Book Title 2', 'E-book', 1);

-- Insert books authored by Jane Smith
INSERT INTO book (book_title, book_type, author_id) VALUES
('Book Title 3', 'Physical', 2);

-- Imsert Author credential
INSERT INTO author (author_name, author_email) VALUES
('John Doe', 'john.doe@example.com'),
('Jane Smith', 'jane.smith@example.com');


-- Insert sales for the first author
INSERT INTO bookstore_order (customer_name, customer_email, order_date, price, book_id) VALUES
('John Smith', 'john.smith@example.com', '2024-05-01', 20.50, 1),  -- Assuming book_id 1 corresponds to John Doe's book
('John Smith', 'john.smith@example.com', '2024-05-05', 15.75, 2);  -- Assuming book_id 2 corresponds to John Doe's book

-- Insert sales for the second author
INSERT INTO bookstore_order (customer_name, customer_email, order_date, price, book_id) VALUES
('Alice Johnson', 'alice.johnson@example.com', '2024-05-02', 12.99, 3);  -- Assuming book_id 3 corresponds to Jane Smith's book

-- Insert additional data to sales
INSERT INTO bookstore_order (customer_name, customer_email, order_date, price, book_id) VALUES
('Alice Johnson', 'alice.johnson@example.com', '2024-05-06', 25.99, 3),  -- Assuming book_id 3 corresponds to a book
('Bob Williams', 'bob.williams@example.com', '2024-05-07', 18.50, 1),     -- Assuming book_id 1 corresponds to a book
('Alice Johnson', 'alice.johnson@example.com', '2024-05-08', 22.75, 2),  -- Assuming book_id 2 corresponds to a book
('John Smith', 'john.smith@example.com', '2024-05-09', 30.00, 4),         -- Assuming book_id 4 corresponds to a book
('Emily Davis', 'emily.davis@example.com', '2024-05-10', 12.99, 5);        -- Assuming book_id 5 corresponds to a book
