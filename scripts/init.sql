CREATE DATABASE IF NOT EXISTS webapp;

USE webapp;

CREATE TABLE IF NOT EXISTS Account(account_id INT AUTO_INCREMENT PRIMARY KEY, document_number VARCHAR(255) UNIQUE NOT NULL);

-- Create the Operation_type table and insert the data
CREATE TABLE IF NOT EXISTS Operation_type ( operation_type_id INT AUTO_INCREMENT PRIMARY KEY, description VARCHAR(255));

INSERT INTO Operation_type (operation_type_id, description) VALUES (1, 'Normal Purchase'), (2, 'Installment Purchase'), (3, 'Withdrawal'), (4, 'Credit Voucher');

-- Create the Transaction table
CREATE TABLE IF NOT EXISTS Transaction (
                                           transaction_id INT AUTO_INCREMENT PRIMARY KEY,
                                           account_id INT NOT NULL,
                                           operation_type_id INT NOT NULL,
                                           amount FLOAT NOT NULL,
                                           balance FLOAT NOT NULL,
                                           event_date DATETIME NOT NULL,
                                           FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (operation_type_id) REFERENCES Operation_type(operation_type_id)
    );
