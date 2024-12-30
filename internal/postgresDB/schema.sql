CREATE TABLE equipment(
    inventory_number SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    manufacturer VARCHAR(255) NOT NULL,
    start_date DATE,
    lifeTime INTEGER
);

CREATE TABLE material (
    material_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255),
    unit_price NUMERIC(10, 2),
    unit_of_measurement VARCHAR(50),
    alternative VARCHAR(255)
);

CREATE TABLE product_specification (
    specification_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    production_duration INTEGER,
    equipment_id INTEGER NOT NULL,
    material_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (equipment_id) REFERENCES equipment (inventory_number) ON DELETE CASCADE,
    FOREIGN KEY (material_id) REFERENCES material (material_id) ON DELETE CASCADE
);