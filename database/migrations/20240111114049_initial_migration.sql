-- +goose Up
-- +goose StatementBegin
CREATE TABLE measurement_unit (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    reference_unit_id INT,
    reference_unit_value DECIMAL(10, 2),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (reference_unit_id) REFERENCES measurement_unit (id),
    CONSTRAINT measurement_unit_uuid_unique UNIQUE (uuid)
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    measurement_unit_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    barcode VARCHAR(255) NOT NULL,
    measurement_value DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (measurement_unit_id) REFERENCES measurement_unit (id),
    CONSTRAINT product_barcode_unique UNIQUE (barcode),
    CONSTRAINT product_uuid_unique UNIQUE (uuid)
);

CREATE TABLE recipe (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT recipe_uuid_unique UNIQUE (uuid)
);

CREATE TABLE recipe_product (
    id SERIAL PRIMARY KEY,
    recipe_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (recipe_id) REFERENCES recipe (id),
    FOREIGN KEY (product_id) REFERENCES product (id),
    CONSTRAINT recipe_product_unique UNIQUE (recipe_id, product_id)
);

CREATE TABLE recipe_selling_history (
    id SERIAL PRIMARY KEY,
    recipe_id INT NOT NULL,
    cost_price DECIMAL(10, 2) NOT NULL,
    selling_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (recipe_id) REFERENCES recipe (id),
    CONSTRAINT recipe_selling_history_unique UNIQUE (recipe_id)
);

CREATE TABLE stock_operation (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT stock_operation_uuid_unique UNIQUE (uuid)
);

CREATE TABLE stock_history (
    id SERIAL PRIMARY KEY,
    stock_operation_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (stock_operation_id) REFERENCES stock_operation (id),
    FOREIGN KEY (product_id) REFERENCES product (id),
    CONSTRAINT stock_history_unique UNIQUE (stock_operation_id, product_id)
);

CREATE TABLE changelog (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    object_name VARCHAR(255) NOT NULL,
    object_id INT NOT NULL,
    parent_object_name VARCHAR(255),
    parent_object_id INT,
    operation VARCHAR(255) NOT NULL,
    field_name VARCHAR(255) NOT NULL,
    old_value VARCHAR(255),
    new_value VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT changelog_uuid_unique UNIQUE (uuid)
);

CREATE INDEX changelog_object_name_object_id_index ON changelog (object_name, object_id);
CREATE INDEX changelog_parent_object_name_parent_object_id_index ON changelog (parent_object_name, parent_object_id);
CREATE INDEX changelog_operation_index ON changelog (operation);
CREATE INDEX changelog_field_name_index ON changelog (field_name);
CREATE INDEX changelog_created_at_index ON changelog (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE changelog;
DROP TABLE stock_history;
DROP TABLE stock_operation;
DROP TABLE recipe_selling_history;
DROP TABLE recipe_product;
DROP TABLE recipe;
DROP TABLE product;
DROP TABLE measurement_unit;
-- +goose StatementEnd
