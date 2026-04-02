CREATE DATABASE p1ppteam3;
USE p1ppteam3;

--  Products
CREATE TABLE products (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    sku VARCHAR(50) UNIQUE, -- Index unik untuk pencarian cepat
    name VARCHAR(100) NOT NULL,
    type ENUM('finished', 'raw', 'semi-finished') NOT NULL, -- Menggunakan ENUM agar data konsisten
    unit VARCHAR(20) NOT NULL,
    standard_cost DECIMAL(15,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--  Bill of Material (BOM)
CREATE TABLE bill_of_material (
    bill_id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    component_id INT,
    qty_required DECIMAL(10,3) NOT NULL, -- 3 desimal untuk akurasi bahan (gram ke kg)
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE,
    FOREIGN KEY (component_id) REFERENCES products(product_id)
);

--  Production Orders
CREATE TABLE production_orders (
    production_order_id INT AUTO_INCREMENT PRIMARY KEY,
    order_code VARCHAR(50) UNIQUE,
    product_id INT,
    qty_plan DECIMAL(10,2) NOT NULL,
    qty_actual DECIMAL(10,2) DEFAULT 0.00,
    status ENUM('planned', 'running', 'completed', 'cancelled') DEFAULT 'planned',
    start_date DATETIME,
    end_date DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    CONSTRAINT chk_qty_plan CHECK (qty_plan > 0)
);

--  Production Summaries
CREATE TABLE production_summaries (
    production_order_id INT PRIMARY KEY,
    total_output INT DEFAULT 0,
    total_defect INT DEFAULT 0,
    efficiency DECIMAL(5,2) DEFAULT 0.00,
    runtime_minutes INT DEFAULT 0,
    downtime_minutes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (production_order_id) REFERENCES production_orders(production_order_id) ON DELETE CASCADE
);

--  Machines & Operators
CREATE TABLE machines (
    machine_id INT AUTO_INCREMENT PRIMARY KEY,
    machine_code VARCHAR(50) UNIQUE,
    name VARCHAR(100),
    status ENUM('active', 'maintenance', 'broken') DEFAULT 'active'
);

CREATE TABLE operators (
    operator_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(255) NOT NULL,
    status ENUM('active', 'inactive') DEFAULT 'active'
);

--  Logs (Production & Machine)
CREATE TABLE production_logs (
    log_id INT AUTO_INCREMENT PRIMARY KEY,
    production_order_id INT,
    machine_id INT,
    operator_id INT,
    output_qty INT DEFAULT 0,
    defect_qty INT DEFAULT 0,
    log_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (production_order_id) REFERENCES production_orders(production_order_id),
    FOREIGN KEY (machine_id) REFERENCES machines(machine_id),
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id)
);

CREATE TABLE machine_usage_logs (
    usage_id INT AUTO_INCREMENT PRIMARY KEY,
    production_order_id INT,
    machine_id INT,
    operator_id INT,
    start_time DATETIME,
    end_time DATETIME,
    FOREIGN KEY (production_order_id) REFERENCES production_orders(production_order_id),
    FOREIGN KEY (machine_id) REFERENCES machines(machine_id)
);

--  Inventory Transactions
CREATE TABLE inventory_transactions (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    production_order_id INT,
    product_id INT, -- Ditambahkan agar tahu barang apa yang bergerak
    transaction_type ENUM('IN', 'OUT') NOT NULL,
    qty DECIMAL(10,2) NOT NULL,
    transaction_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (production_order_id) REFERENCES production_orders(production_order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);
	

-- === Products ===
INSERT INTO products (sku, name, type, unit, standard_cost) VALUES
('FG-001', 'Roti Tawar Gandum', 'finished', 'pcs', 12000),
('FG-002', 'Roti Manis Cokelat', 'finished', 'pcs', 5000),
('FG-003', 'Croissant Butter', 'finished', 'pcs', 15000),
('FG-004', 'Donat Gula', 'finished', 'pcs', 4000),
('FG-005', 'Baguette French', 'finished', 'pcs', 18000),
('RM-001', 'Tepung Terigu Pro Tinggi', 'raw', 'kg', 11000),
('RM-002', 'Gula Pasir', 'raw', 'kg', 14000),
('RM-003', 'Mentega/Butter', 'raw', 'kg', 85000),
('RM-004', 'Ragi Instan', 'raw', 'pack', 5000),
('RM-005', 'Cokelat Batang', 'raw', 'kg', 60000);

-- === BOM ===
INSERT INTO bill_of_material (product_id, component_id, qty_required) VALUES
(1, 6, 0.500), (1, 9, 0.010), -- Roti Tawar: Tepung & Ragi
(2, 6, 0.300), (2, 10, 0.050), -- Roti Manis: Tepung & Cokelat
(3, 8, 0.200), (3, 6, 0.400), -- Croissant: Butter & Tepung
(4, 6, 0.200), (4, 7, 0.020), -- Donat: Tepung & Gula
(5, 6, 0.600), (5, 9, 0.020); -- Baguette: Tepung & Ragi

-- === Machines ===
INSERT INTO machines (machine_code, name, status) VALUES
('M-01', 'Mixer Heavy Duty', 'active'), ('M-02', 'Oven Rotary Large', 'active'),
('M-03', 'Proofer Room', 'active'), ('M-04', 'Bread Slicer', 'active'),
('M-05', 'Packaging Auto', 'active'), ('M-06', 'Fryer Donut', 'active'),
('M-07', 'Cooling Rack Conveyor', 'active'), ('M-08', 'Dough Divider', 'active'),
('M-09', 'Mixer Small Batch', 'maintenance'), ('M-10', 'Oven Deck 02', 'active');

-- === Operator ===
INSERT INTO operators (name, email, password, status) VALUES
('Andi', 'andi@food.com', 'pass1', 'active'), ('Siti', 'siti@food.com', 'pass2', 'active'),
('Budi', 'budi@food.com', 'pass3', 'active'), ('Dewi', 'dewi@food.com', 'pass4', 'active'),
('Eko', 'eko@food.com', 'pass5', 'active'), ('Rina', 'rina@food.com', 'pass6', 'active'),
('Tono', 'tono@food.com', 'pass7', 'active'), ('Lani', 'lani@food.com', 'pass8', 'active'),
('Hadi', 'hadi@food.com', 'pass9', 'active'), ('Yati', 'yati@food.com', 'pass10', 'active');

-- === Product Orders ===
INSERT INTO production_orders (order_code, product_id, qty_plan, status) VALUES
('PO-001', 1, 100, 'completed'), ('PO-002', 2, 200, 'completed'),
('PO-003', 3, 50, 'running'), ('PO-004', 4, 150, 'running'),
('PO-005', 5, 80, 'planned'), ('PO-006', 1, 120, 'planned'),
('PO-007', 2, 300, 'planned'), ('PO-008', 3, 60, 'planned'),
('PO-009', 4, 200, 'planned'), ('PO-010', 5, 100, 'planned');

-- === Production Logs ===
INSERT INTO production_logs (production_order_id, machine_id, operator_id, output_qty, defect_qty) VALUES
(1, 1, 1, 50, 1), (1, 2, 2, 49, 0), (2, 1, 3, 100, 2), (2, 2, 4, 98, 1),
(3, 1, 5, 25, 0), (3, 3, 6, 25, 0), (4, 6, 7, 75, 5), (4, 5, 8, 70, 2),
(1, 4, 9, 49, 0), (2, 5, 10, 97, 1);

-- === Machine Logs ===
INSERT INTO machine_usage_logs (production_order_id, machine_id, operator_id, start_time, end_time) VALUES
(1, 1, 1, '2024-05-01 08:00', '2024-05-01 09:00'), (1, 2, 2, '2024-05-01 09:30', '2024-05-01 11:00'),
(2, 1, 3, '2024-05-01 08:00', '2024-05-01 10:00'), (2, 2, 4, '2024-05-01 10:30', '2024-05-01 12:30'),
(3, 1, 5, '2024-05-02 08:00', '2024-05-02 09:00'), (4, 6, 7, '2024-05-02 08:30', '2024-05-02 10:30'),
(1, 4, 9, '2024-05-01 13:00', '2024-05-01 14:00'), (2, 5, 10, '2024-05-01 14:00', '2024-05-01 15:00'),
(3, 3, 6, '2024-05-02 09:30', '2024-05-02 11:00'), (4, 5, 8, '2024-05-02 12:00', '2024-05-02 13:00');

INSERT INTO inventory_transactions (production_order_id, product_id, transaction_type, qty) VALUES
(1, 1, 'IN', 98), (2, 2, 'IN', 195), (1, 6, 'OUT', 50), (1, 9, 'OUT', 1),
(2, 6, 'OUT', 60), (2, 10, 'OUT', 10), (3, 8, 'OUT', 10), (3, 6, 'OUT', 20),
(4, 6, 'OUT', 30), (4, 7, 'OUT', 3);

INSERT INTO production_summaries (production_order_id, total_output, total_defect, efficiency, runtime_minutes) VALUES
(1, 98, 1, 98.00, 120), (2, 195, 3, 97.50, 240), (3, 50, 0, 100.00, 90),
(4, 145, 7, 96.60, 180), (5, 0, 0, 0.00, 0), (6, 0, 0, 0.00, 0),
(7, 0, 0, 0.00, 0), (8, 0, 0, 0.00, 0), (9, 0, 0, 0.00, 0), (10, 0, 0, 0.00, 0);

select * from production_summaries;
select * from operators;
select * from production_logs;
select * from products;
