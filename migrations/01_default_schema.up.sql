CREATE TABLE clients (
  id       SERIAL PRIMARY KEY,
  name     VARCHAR(40) NOT NULL,
  address_1 VARCHAR(40) NOT NULL,
  address_2 VARCHAR(40) NOT NULL,
  city     VARCHAR(40) NOT NULL,
  state    VARCHAR(40) NOT NULL,
  zip      VARCHAR(40) NOT NULL
);

CREATE TABLE contacts (
  id SERIAL PRIMARY KEY,
  client_id INTEGER NOT NULL ,
  first_name VARCHAR(40) NOT NULL ,
  last_name VARCHAR(40) NOT NULL ,
  email VARCHAR(100) NOT NULL ,
  is_primary BOOLEAN NOT NULL DEFAULT FALSE ,
  CONSTRAINT fk_contact_client
     FOREIGN KEY (client_id)
     REFERENCES clients (id)
);

CREATE TABLE time_entries (
  id SERIAL PRIMARY KEY,
  client_id INTEGER NOT NULL ,
  start_time TIMESTAMP WITH TIME ZONE NOT NULL ,
  end_time TIMESTAMP WITH TIME ZONE NOT NULL ,
  rate NUMERIC(6, 2) NOT NULL ,
  CONSTRAINT fk_time_entry_client
     FOREIGN KEY (client_id)
     REFERENCES clients (id)
);

CREATE TABLE rates (
  id SERIAL PRIMARY KEY,
  client_id INTEGER NOT NULL ,
  rate NUMERIC(6, 2) NOT NULL ,
  CONSTRAINT fk_rate_client
     FOREIGN KEY (client_id)
     REFERENCES clients (id)
);

CREATE TABLE invoices (
  id SERIAL PRIMARY KEY,
  client_id INTEGER NOT NULL ,
  start_date DATE NOT NULL ,
  end_date DATE NOT NULL ,
  hours NUMERIC(6, 2) NOT NULL ,
  rate NUMERIC(6, 2) NOT NULL ,
  CONSTRAINT fk_invoice_client
     FOREIGN KEY (client_id)
     REFERENCES clients (id)
);