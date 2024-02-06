CREATE TABLE tasks (
	id INTEGER PRIMARY KEY,
	created_at DATETIME,
	title VARCHAR(256) NOT NULL,
	completed BOOLEAN
);

INSERT INTO tasks
	(title, completed)
VALUES
	('Tarefa Teste', false),
	('Testezinho 2', true)
;