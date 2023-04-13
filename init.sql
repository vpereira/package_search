CREATE TABLE packages (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    release TEXT NOT NULL,
    location TEXT NOT NULL,
    summary TEXT NOT NULL
);

INSERT INTO packages (name, version, release, location, summary)
VALUES ('package1', '1.0.0', '1', '/usr/local', 'Package 1 summary'),
       ('package2', '1.1.0', '2', '/usr/local', 'Package 2 summary'),
       ('package3', '2.0.0', '3', '/usr/local', 'Package 3 summary'),
       ('package4', '2.1.0', '4', '/usr/local', 'Package 4 summary'),
       ('package5', '3.0.0', '5', '/usr/local', 'Package 5 summary'),
       ('package6', '3.1.0', '6', '/usr/local', 'Package 6 summary'),
       ('package7', '4.0.0', '7', '/usr/local', 'Package 7 summary'),
       ('package8', '4.1.0', '8', '/usr/local', 'Package 8 summary'),
       ('package9', '5.0.0', '9', '/usr/local', 'Package 9 summary'),
       ('package10', '5.1.0', '10', '/usr/local', 'Package 10 summary');

