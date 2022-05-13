PRAGMA foreign_keys = ON;
CREATE TABLE hexrepo(id TEXT UNIQUE);
CREATE TABLE hexdata(hexid TEXT, key TEXT, value TEXT, FOREIGN KEY(hexid) REFERENCES hexrepo(id));
CREATE TABLE hexmap(x INTEGER, y INTEGER, z INTEGER, hexid TEXT, FOREIGN KEY(hexid) REFERENCES hexrepo(id));
CREATE UNIQUE INDEX hmi ON hexmap(x, y);
CREATE TABLE mapdata(x INTEGER, y INTEGER, key TEXT, value TEXT, FOREIGN KEY(x,y) REFERENCES hexmap(x,y));
CREATE UNIQUE INDEX mdi ON mapdata(x, y, key);