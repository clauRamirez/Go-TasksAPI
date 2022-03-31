DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    Id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    is_done BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO tasks(name, description, is_done) VALUES('learn Go basics', 'learn basics of the Go programming language', true);
INSERT INTO tasks(name, description) VALUES('write api', 'create a simple api with crud operations with Go');
INSERT INTO tasks(name, description) VALUES('push github', 'push project to public github repo');
INSERT INTO tasks(name, description) VALUES('learn fyne', 'learn the fyne toolkit');