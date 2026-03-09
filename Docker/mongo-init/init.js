db = db.getSiblingDB("veterinaria_citiaps");

db.Duenos.insertMany([
  { _id: 1, nombre: "Giuseppe Asd", edad: 22, sexo: "Masculino" },
  { _id: 2, nombre: "Humano2", edad: 25, sexo: "Masculino" },
  { _id: 3, nombre: "Humano3", edad: 30, sexo: "Femenino" },
  { _id: 4, nombre: "Humano4", edad: 45, sexo: "Masculino" },
  { _id: 5, nombre: "Humano5", edad: 28, sexo: "Femenino" },
]);

db.Perros.insertMany([
  {
    _id: 101,
    nombre: "Hachi",
    raza: "Pug",
    color: "Beige",
    edad: 3,
    id_dueno: 1,
  },
  {
    _id: 102,
    nombre: "Charawi",
    raza: "Pastor",
    color: "Negro",
    edad: 5,
    id_dueno: 2,
  },
  {
    _id: 103,
    nombre: "Perro3",
    raza: "Beagle",
    color: "Blanco",
    edad: 2,
    id_dueno: 3,
  },
  {
    _id: 104,
    nombre: "Perro4",
    raza: "Poodle",
    color: "Gris",
    edad: 4,
    id_dueno: 4,
  },
  {
    _id: 105,
    nombre: "Perro5",
    raza: "Labrador",
    color: "Dorado",
    edad: 1,
    id_dueno: 5,
  },
]);

db.Vacunas.insertMany([
  { _id: 501, fecha: "2026-03-03", nombrevacuna: "Vacuna1", id_perro: 101 },
  { _id: 502, fecha: "2026-03-03", nombrevacuna: "Vacuna2", id_perro: 102 },
  { _id: 503, fecha: "2026-03-03", nombrevacuna: "Vacuna3", id_perro: 103 },
  { _id: 504, fecha: "2026-03-03", nombrevacuna: "Vacuna4", id_perro: 104 },
  { _id: 505, fecha: "2026-03-03", nombrevacuna: "Vacuna5", id_perro: 105 },
]);
