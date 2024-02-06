## API

| Ação | Rota | Função |
| -- | -- | -- |
| GET | /tasks | Retornar todas as notas |
| GET | /tasks/id/{:id} | Retornar uma nota de acordo com o __id__ |
| POST | /tasks | Criar uma nota com os dados fornecidos no __body__ |
| PUT | /tasks/{:id} | Editar as informações de uma nota fornecidas pelo __body__ de acordo com o __id__ |
| DELETE | /tasks/{:id} | Deletar uma nota de acordo com o __id__ |