CREATE DATABASE `employees` DEFAULT CHARACTER SET utf8mb4;

USE `employees`;
-- Create employees table
CREATE TABLE staff (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    position VARCHAR(100),
    department VARCHAR(100),
    salary DECIMAL(10, 2),
    hire_date DATE
);

-- Insert employee data
INSERT INTO staff (first_name, last_name, position, department, salary, hire_date) VALUES 
('Taro', 'Yamada', 'Manager', 'Sales', 60000.00, '2020-01-15'),
('Hanako', 'Suzuki', 'Engineer', 'Development', 50000.00, '2020-02-20'),
('Jiro', 'Sato', 'Sales Representative', 'Sales', 45000.00, '2020-03-10'),
('Saburo', 'Tanaka', 'Assistant', 'Administration', 40000.00, '2020-04-05'),
('Miki', 'Ito', 'Designer', 'Design', 55000.00, '2020-05-12'),
('Kenta', 'Kimura', 'Engineer', 'Development', 52000.00, '2020-06-18'),
('Mai', 'Yamamoto', 'Sales Representative', 'Sales', 47000.00, '2020-07-22'),
('Natsumi', 'Nakamura', 'Assistant', 'Administration', 42000.00, '2020-08-30'),
('Satoshi', 'Inoue', 'Engineer', 'Development', 53000.00, '2020-09-25'),
('Miho', 'Kobayashi', 'Designer', 'Design', 56000.00, '2020-10-10');
