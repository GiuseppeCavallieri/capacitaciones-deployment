-- Veterinaria Citiaps - PostgreSQL Schema & Seed Data

CREATE TABLE IF NOT EXISTS duenos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    edad INT NOT NULL,
    sexo VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS perros (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    raza VARCHAR(255) NOT NULL,
    color VARCHAR(100) NOT NULL,
    edad INT NOT NULL,
    id_dueno INT NOT NULL REFERENCES duenos(id)
);

CREATE TABLE IF NOT EXISTS vacunas (
    id SERIAL PRIMARY KEY,
    fecha VARCHAR(50) NOT NULL,
    nombrevacuna VARCHAR(255) NOT NULL,
    id_perro INT NOT NULL REFERENCES perros(id)
);

-- Seed duenos
INSERT INTO duenos (id, nombre, edad, sexo) VALUES
    (1, 'Giuseppe Asd', 22, 'Masculino'),
    (2, 'Humano2', 25, 'Masculino'),
    (3, 'Humano3', 30, 'Femenino'),
    (4, 'Humano4', 45, 'Masculino'),
    (5, 'Humano5', 28, 'Femenino');

-- Seed perros
INSERT INTO perros (id, nombre, raza, color, edad, id_dueno) VALUES
    (101, 'Hachi', 'Pug', 'Beige', 3, 1),
    (102, 'Charawi', 'Pastor', 'Negro', 5, 2),
    (103, 'Perro3', 'Beagle', 'Blanco', 2, 3),
    (104, 'Perro4', 'Poodle', 'Gris', 4, 4),
    (105, 'Perro5', 'Labrador', 'Dorado', 1, 5);

-- Seed vacunas
INSERT INTO vacunas (id, fecha, nombrevacuna, id_perro) VALUES
    (501, '2026-03-03', 'Vacuna1', 101),
    (502, '2026-03-03', 'Vacuna2', 102),
    (503, '2026-03-03', 'Vacuna3', 103),
    (504, '2026-03-03', 'Vacuna4', 104),
    (505, '2026-03-03', 'Vacuna5', 105);

-- Advance sequences past the seeded IDs
SELECT setval('duenos_id_seq', (SELECT MAX(id) FROM duenos));
SELECT setval('perros_id_seq', (SELECT MAX(id) FROM perros));
SELECT setval('vacunas_id_seq', (SELECT MAX(id) FROM vacunas));
